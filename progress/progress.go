// Package progress provides methods and structs for creating terminal progress bar
package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/fmtutil"
	"github.com/essentialkaos/ek/v14/passthru"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MIN_WIDTH is the minimum allowed width of the progress bar in characters
const MIN_WIDTH = 80

// MIN_REFRESH_RATE is the minimum allowed render refresh rate (1 ms)
const MIN_REFRESH_RATE = time.Duration(time.Millisecond)

// PROGRESS_BAR_SYMBOL is the character used to draw the progress bar fill
const PROGRESS_BAR_SYMBOL = "—"

// ////////////////////////////////////////////////////////////////////////////////// //

// Bar represents a terminal progress bar with configurable rendering and
// thread-safe state
type Bar struct {
	settings Settings

	startTime time.Time
	started   bool
	finished  bool

	finishChan  chan bool
	finishGroup sync.WaitGroup

	current int64
	total   int64
	name    string

	buffer string

	ticker       *time.Ticker
	passThruCalc *passthru.Calculator
	phCounter    int

	reader *passthru.Reader
	writer *passthru.Writer

	mu sync.RWMutex
}

// Settings holds all configuration options for a progress bar instance
type Settings struct {
	// RefreshRate is the interval between consecutive render frames; must
	// be >= [MIN_REFRESH_RATE]
	RefreshRate time.Duration

	// NameColorTag is the fmtc color tag applied to the bar name
	NameColorTag string

	// BarFgColorTag is the fmtc color tag for the filled portion of the bar
	BarFgColorTag string

	// BarBgColorTag is the fmtc color tag for the unfilled portion of the bar
	BarBgColorTag string

	// PercentColorTag is the fmtc color tag for the percentage value
	PercentColorTag string

	// ProgressColorTag is the fmtc color tag for the current/total progress text
	ProgressColorTag string

	// SpeedColorTag is the fmtc color tag for the transfer speed text
	SpeedColorTag string

	// RemainingColorTag is the fmtc color tag for the estimated remaining time
	RemainingColorTag string

	// ShowSpeed enables rendering of the current transfer speed
	ShowSpeed bool

	// ShowName enables rendering of the bar name
	ShowName bool

	// ShowPercentage enables rendering of the completion percentage
	ShowPercentage bool

	// ShowProgress enables rendering of the current/total progress values
	ShowProgress bool

	// ShowRemaining enables rendering of the estimated time remaining
	ShowRemaining bool

	// Width is the total character width of the rendered bar line
	Width int

	// NameSize is the fixed column width reserved for the name; shorter names
	// are right-padded. Zero disables fixed-width padding.
	NameSize int

	// WindowSizeSec is the sliding window duration in seconds used by the passthru
	// speed calculator
	WindowSizeSec int64 // Window size for passtru reader

	// IsSize controls value formatting: true renders values as byte sizes (KB/MB/GB),
	// false renders them as plain numbers (K/M/B)
	IsSize bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultSettings is the ready-to-use default configuration applied to every new Bar
var DefaultSettings = Settings{
	RefreshRate:       100 * time.Millisecond,
	NameColorTag:      "{b}",
	BarFgColorTag:     "{r}",
	BarBgColorTag:     "{s-}",
	PercentColorTag:   "{m}",
	SpeedColorTag:     "{r}",
	ProgressColorTag:  "{g}",
	RemainingColorTag: "{c}",
	ShowName:          true,
	ShowPercentage:    true,
	ShowProgress:      true,
	ShowSpeed:         true,
	ShowRemaining:     true,
	IsSize:            true,
	Width:             88,
	WindowSizeSec:     15,
}

var (
	// ErrBarIsNil is returned by Bar methods when called on a nil receiver
	ErrBarIsNil = errors.New("progress bar struct is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates a new Bar with the given total value and display name, using
// [DefaultSettings]
func New(total int64, name string) *Bar {
	return &Bar{
		settings: DefaultSettings,
		name:     name,
		total:    total,
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Start initialises internal state and launches the async rendering goroutine.
// It is a no-op if the bar is already running.
func (b *Bar) Start() error {
	if b == nil {
		return ErrBarIsNil
	}

	if b.IsStarted() && !b.IsFinished() {
		return nil
	}

	b.phCounter = 0
	b.current = 0
	b.started = true
	b.finished = false
	b.startTime = time.Now()
	b.finishChan = make(chan bool)
	b.finishGroup = sync.WaitGroup{}
	b.ticker = time.NewTicker(b.settings.RefreshRate)

	if b.total > 0 {
		b.passThruCalc = passthru.NewCalculator(
			b.total, math.Max(1.0, float64(b.settings.WindowSizeSec)),
		)
	}

	go b.renderer()

	return nil
}

// Finish signals the rendering goroutine to stop and blocks until the final frame
// is drawn
func (b *Bar) Finish() error {
	if b == nil {
		return ErrBarIsNil
	}

	b.mu.RLock()

	if b.finished || !b.started {
		b.mu.RUnlock()
		return nil
	}

	b.mu.RUnlock()

	b.finishGroup.Add(1)
	b.finishChan <- true
	b.finishGroup.Wait()

	return nil
}

// UpdateSettings validates and applies new settings to the bar.
// RefreshRate changes take effect only after the next [Start] call.
func (b *Bar) UpdateSettings(s Settings) error {
	if b == nil {
		return ErrBarIsNil
	}

	err := s.Validate()

	if err != nil {
		return err
	}

	b.mu.Lock()
	b.settings = s
	b.mu.Unlock()

	return nil
}

// SetName updates the display name of the bar
func (b *Bar) SetName(name string) {
	if b == nil {
		return
	}

	b.mu.Lock()
	b.name = name
	b.mu.Unlock()
}

// Name returns the current display name of the bar
func (b *Bar) Name() string {
	if b == nil {
		return ""
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.name
}

// SetTotal updates the total value and reconfigures the passthru speed
// calculator accordingly
func (b *Bar) SetTotal(v int64) {
	if b == nil {
		return
	}

	b.mu.Lock()

	if b.passThruCalc == nil {
		b.passThruCalc = passthru.NewCalculator(
			v, math.Max(1.0, float64(b.settings.WindowSizeSec)),
		)
	} else {
		b.passThruCalc.SetTotal(v)
	}

	b.mu.Unlock()

	atomic.StoreInt64(&b.total, v)
}

// Total returns the current total value of the bar
func (b *Bar) Total() int64 {
	if b == nil {
		return 0
	}

	return atomic.LoadInt64(&b.total)
}

// SetCurrent atomically sets the current progress value
func (b *Bar) SetCurrent(v int64) {
	if b == nil {
		return
	}

	atomic.StoreInt64(&b.current, v)
}

// Current atomically returns the current progress value
func (b *Bar) Current() int64 {
	if b == nil {
		return 0
	}

	return atomic.LoadInt64(&b.current)
}

// Add atomically increments the current progress value by v
func (b *Bar) Add(v int) {
	if b == nil {
		return
	}

	atomic.AddInt64(&b.current, int64(v))
}

// Add64 atomically increments the current progress value by v
func (b *Bar) Add64(v int64) {
	if b == nil {
		return
	}

	atomic.AddInt64(&b.current, v)
}

// IsFinished returns true if the bar has completed rendering after a
// Finish call
func (b *Bar) IsFinished() bool {
	if b == nil {
		return false
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.finished
}

// IsStarted returns true if the bar has been started and not yet reset
func (b *Bar) IsStarted() bool {
	if b == nil {
		return false
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.started
}

// Reader wraps r in a passthru reader that automatically advances the bar
// as data is read
func (b *Bar) Reader(r io.Reader) io.Reader {
	if b == nil {
		return nil
	}

	pr := passthru.NewReader(r, b.total)
	pr.UpdateN = b.Add
	b.reader = pr

	return pr
}

// Writer wraps w in a passthru writer that automatically advances the bar
// as data is written
func (b *Bar) Writer(w io.Writer) io.Writer {
	if b == nil {
		return nil
	}

	pw := passthru.NewWriter(w, b.total)
	pw.UpdateN = b.Add
	b.writer = pw

	return pw
}

// ////////////////////////////////////////////////////////////////////////////////// //

// renderer is rendering loop func
func (b *Bar) renderer() {
	for {
		select {
		case <-b.finishChan:
			b.ticker.Stop()
			b.render(true)
			return
		case <-b.ticker.C:
			b.render(false)
		}
	}
}

// render renders current progress bar state
func (b *Bar) render(isFinished bool) {
	b.mu.RLock()
	result := b.renderElements(isFinished)
	changed := b.buffer != result
	b.mu.RUnlock()

	if changed || isFinished {
		fmtc.TPrint(result)
	}

	if isFinished {
		fmtc.NewLine()

		b.mu.Lock()
		b.buffer = result
		b.finished = true
		b.mu.Unlock()

		close(b.finishChan)
		b.finishGroup.Done()

		return
	}

	if changed && b.Total() > 0 {
		b.mu.Lock()
		b.buffer = result
		b.mu.Unlock()
	}
}

// renderElements returns text with all progress bar graphics and text
func (b *Bar) renderElements(isFinished bool) string {
	var size, totalSize int
	var name, percentage, bar, progress, speed, remaining string
	var statSpeed float64
	var statRemaining time.Duration

	if b.passThruCalc != nil && (b.settings.ShowSpeed || b.settings.ShowRemaining) {
		current := atomic.LoadInt64(&b.current)

		if isFinished {
			statRemaining = time.Since(b.startTime)
			statSpeed = float64(current) / statRemaining.Seconds()
		} else {
			statSpeed, statRemaining = b.passThruCalc.Calculate(current)
		}
	}

	if b.settings.ShowName && b.name != "" {
		name, size = b.renderName()
		totalSize += size + 1
	}

	if b.total > 0 {
		if b.settings.ShowPercentage {
			percentage, size = b.renderPercentage()
			totalSize += size + 1
		}

		if b.settings.ShowProgress {
			progress, size = b.renderProgress()
			totalSize += size + 3
		}

		if b.settings.ShowSpeed {
			speed, size = b.renderSpeed(statSpeed)
			totalSize += size + 3
		}

		if b.settings.ShowRemaining {
			remaining, size = b.renderRemaining(statRemaining)
			totalSize += size + 3
		}
	}

	bar = b.renderBar(totalSize)

	var result string

	if b.settings.ShowName && name != "" {
		result += name + " "
	}

	result += bar + " "

	if b.total > 0 {
		if b.settings.ShowPercentage {
			result += percentage
		}

		if b.settings.ShowProgress {
			result += " {s-}•{!} " + progress
		}

		if b.settings.ShowSpeed {
			result += " {s-}•{!} " + speed
		}

		if b.settings.ShowRemaining {
			result += " {s-}•{!} " + remaining
		}
	}

	return result
}

// renderName returns name text
func (b *Bar) renderName() (string, int) {
	var result string

	if b.settings.NameSize > 0 && len(b.name) < b.settings.NameSize {
		result = fmt.Sprintf("%"+strconv.Itoa(b.settings.NameSize)+"s", b.name)
	} else {
		result = b.name
	}

	if fmtc.DisableColors || b.settings.NameColorTag == "" {
		return result, len(result)
	}

	return b.settings.NameColorTag + result + "{!}", len(result)
}

// renderPercentage returns parcentage text
func (b *Bar) renderPercentage() (string, int) {
	var perc float64
	var result string

	switch {
	case b.total <= 0:
		perc = 0.0
	case b.current > b.total:
		perc = 100.0
	default:
		perc = (float64(b.current) / float64(b.total)) * 100.0
	}

	if perc == 100.0 {
		result = "100%"
	} else {
		result = fmt.Sprintf("%5.1f%%", perc)
	}

	if fmtc.DisableColors || b.settings.PercentColorTag == "" {
		return result, len(result)
	}

	return b.settings.PercentColorTag + result + "{!}", len(result)
}

// renderProgress returns progress text
func (b *Bar) renderProgress() (string, int) {
	var result, curText, totText, label string
	var size int

	if b.settings.IsSize {
		curText, totText, label = getPrettyCTSize(b.current, b.total)
	} else {
		curText, totText, label = getPrettyCTNum(b.current, b.total)
	}

	size = (len(totText) * 2) + len(label) + 1

	if label == "" {
		result = fmt.Sprintf("%"+strconv.Itoa(size)+"s", curText+"/"+totText)
	} else {
		result = fmt.Sprintf("%"+strconv.Itoa(size)+"s", curText+"/"+totText+label)
	}

	if fmtc.DisableColors || b.settings.ProgressColorTag == "" {
		return result, size
	}

	return b.settings.ProgressColorTag + result + "{!}", size
}

// renderSpeed returns speed text
func (b *Bar) renderSpeed(speed float64) (string, int) {
	var result string

	if b.settings.IsSize {
		result = fmt.Sprintf("%9s/s", fmtutil.PrettySize(speed, " "))
	} else {
		result = formatSpeedNum(speed)
	}

	if fmtc.DisableColors || b.settings.SpeedColorTag == "" {
		return result, len(result)
	}

	return b.settings.SpeedColorTag + result + "{!}", len(result)
}

// renderRemaining returns remaining text
func (b *Bar) renderRemaining(remaining time.Duration) (string, int) {
	var result string
	var min, sec int

	d := int(remaining.Seconds())

	if d >= 60 {
		min = d / 60
		sec = d % 60
	} else {
		sec = d
	}

	result = fmt.Sprintf("%2d:%02d", min, sec)

	if fmtc.DisableColors || b.settings.RemainingColorTag == "" {
		return result, len(result)
	}

	return b.settings.RemainingColorTag + result + "{!}", len(result)
}

// renderBar returns bar graphics
func (b *Bar) renderBar(dataSize int) string {
	size := max(5, max(MIN_WIDTH, b.settings.Width)-dataSize)

	if b.total <= 0 {
		return b.renderPlaceholder(size)
	}

	current := atomic.LoadInt64(&b.current)
	total := atomic.LoadInt64(&b.total)

	if current >= total {
		if fmtc.DisableColors || b.settings.BarFgColorTag == "" {
			return strings.Repeat(PROGRESS_BAR_SYMBOL, size)
		}

		return b.settings.BarFgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, size) + "{!}"
	}

	cur := int((float64(current) / float64(total)) * float64(size))

	if fmtc.DisableColors || b.settings.BarFgColorTag == "" {
		return strings.Repeat(PROGRESS_BAR_SYMBOL, cur) + strings.Repeat(" ", size-cur)
	}

	return b.settings.BarFgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, cur) + b.settings.BarBgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, size-cur) + "{!}"
}

// renderPlaceholder returns placeholder bar graphics
func (b *Bar) renderPlaceholder(size int) string {
	var result strings.Builder

	result.Grow(size * 4)

	disableColors := fmtc.DisableColors || b.settings.BarFgColorTag == ""

	for i := range size {
		if disableColors {
			if i%3 == b.phCounter {
				result.WriteString(PROGRESS_BAR_SYMBOL)
			} else {
				result.WriteRune(' ')
			}
		} else {
			if i%3 == b.phCounter {
				result.WriteString(b.settings.BarFgColorTag)
			} else {
				result.WriteString(b.settings.BarBgColorTag)
			}

			result.WriteString(PROGRESS_BAR_SYMBOL)
		}
	}

	b.phCounter++

	if b.phCounter == 3 {
		b.phCounter = 0
	}

	if disableColors {
		return result.String()
	}

	return result.String() + "{!}"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate checks that all color tags are valid fmtc tags and that RefreshRate is
// within bounds
func (s Settings) Validate() error {
	switch {
	case !fmtc.IsTag(s.NameColorTag):
		return fmt.Errorf("NameColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.BarFgColorTag):
		return fmt.Errorf("BarFgColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.BarBgColorTag):
		return fmt.Errorf("BarBgColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.PercentColorTag):
		return fmt.Errorf("PercentColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.ProgressColorTag):
		return fmt.Errorf("ProgressColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.SpeedColorTag):
		return fmt.Errorf("SpeedColorTag value is not a valid color tag")

	case !fmtc.IsTag(s.RemainingColorTag):
		return fmt.Errorf("RemainingColorTag value is not a valid color tag")

	case s.RefreshRate != 0 && s.RefreshRate < MIN_REFRESH_RATE:
		return fmt.Errorf("RefreshRate too small (less than 1ms)")
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getPrettyCTSize returns formatted current/total size text
func getPrettyCTSize(current, total int64) (string, string, string) {
	var mod float64
	var label string

	switch {
	case total > 1024*1024*1024:
		mod = 1024 * 1024 * 1024
		label = " GB"
	case total > 1024*1024:
		mod = 1024 * 1024
		label = " MB"
	case total > 1024:
		mod = 1024
		label = " KB"
	default:
		mod = 1
		label = " B"
	}

	curText := fmt.Sprintf("%.1f", float64(current)/mod)
	totText := fmt.Sprintf("%.1f", float64(total)/mod)

	return curText, totText, label
}

// getPrettyCTNum returns formatted current/total number text
func getPrettyCTNum(current, total int64) (string, string, string) {
	var mod float64
	var label, curText, totText string

	switch {
	case total > 1000*1000*1000:
		mod = 1000 * 1000 * 1000
		label = "B"
	case total > 1000*1000:
		mod = 1000 * 1000
		label = "M"
	case total > 1000:
		mod = 1000
		label = "K"
	default:
		mod = 1
	}

	if total > 1000 {
		curText = fmt.Sprintf("%.1f", float64(current)/mod)
		totText = fmt.Sprintf("%.1f", float64(total)/mod)
	} else {
		curText = fmt.Sprintf("%d", int(float64(current)/mod))
		totText = fmt.Sprintf("%d", int(float64(total)/mod))
	}

	return curText, totText, label
}

// formatSpeedNum formats speed number
func formatSpeedNum(s float64) string {
	var mod float64
	var label string

	switch {
	case s > 1000.0*1000.0*1000.0:
		mod = 1000.0 * 1000.0 * 1000.0
		label = "B"
	case s > 1000.0*1000.0:
		mod = 1000.0 * 1000.0
		label = "M"
	case s > 1000.0:
		mod = 1000.0
		label = "K"
	default:
		mod = 1
	}

	return fmt.Sprintf("%6g%s/s", fmtutil.Float(s/mod), label)
}
