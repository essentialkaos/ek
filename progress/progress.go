// Package progress provides methods and structs for creating terminal progress bar
package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/passthru"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MIN_WIDTH is minimal progress bar width
const MIN_WIDTH = 80

// MIN_REFRESH_RATE is minimal refresh rate (1 ms)
const MIN_REFRESH_RATE = time.Duration(time.Millisecond)

// PROGRESS_BAR_SYMBOL is symbol for creating progress bar
const PROGRESS_BAR_SYMBOL = "—"

// ////////////////////////////////////////////////////////////////////////////////// //

// Bar is progress bar struct
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

	mu *sync.RWMutex
}

// Settings contains progress bar settings
type Settings struct {
	RefreshRate time.Duration

	NameColorTag      string
	BarFgColorTag     string
	BarBgColorTag     string
	PercentColorTag   string
	ProgressColorTag  string
	SpeedColorTag     string
	RemainingColorTag string

	ShowSpeed      bool
	ShowName       bool
	ShowPercentage bool
	ShowProgress   bool
	ShowRemaining  bool

	Width    int
	NameSize int

	WindowSizeSec int64 // Window size for passtru reader

	IsSize bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultSettings is default progress bar settings
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
	WindowSizeSec:     15.0,
}

var (
	ErrBarIsNil = fmt.Errorf("Progress bar struct is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new progress bar struct
func New(total int64, name string) *Bar {
	return &Bar{
		settings: DefaultSettings,
		name:     name,
		total:    total,
		mu:       &sync.RWMutex{},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Start starts progress processing
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

// Finish finishes progress processing
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

// UpdateSettings updates progress settings
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

// SetName sets progress bar name
func (b *Bar) SetName(name string) {
	if b == nil {
		return
	}

	b.mu.Lock()
	b.name = name
	b.mu.Unlock()
}

// Name returns progress bar name
func (b *Bar) Name() string {
	if b == nil {
		return ""
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.name
}

// SetTotal sets total progress bar value
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

// Total returns total progress bar value
func (b *Bar) Total() int64 {
	if b == nil {
		return 0
	}

	return atomic.LoadInt64(&b.total)
}

// SetCurrent sets current progress bar value
func (b *Bar) SetCurrent(v int64) {
	if b == nil {
		return
	}

	atomic.StoreInt64(&b.current, v)
}

// Current returns current progress bar value
func (b *Bar) Current() int64 {
	if b == nil {
		return 0
	}

	return atomic.LoadInt64(&b.current)
}

func (b *Bar) Add(v int) {
	if b == nil {
		return
	}

	atomic.AddInt64(&b.current, int64(v))
}

// Add64 adds given value ti
func (b *Bar) Add64(v int64) {
	if b == nil {
		return
	}

	atomic.AddInt64(&b.current, v)
}

// IsFinished returns true if progress processing is finished
func (b *Bar) IsFinished() bool {
	if b == nil {
		return false
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.finished
}

// IsStarted returns true if progress processing is started
func (b *Bar) IsStarted() bool {
	if b == nil {
		return false
	}

	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.started
}

// Reader creates and returns pass thru-proxy reader
func (b *Bar) Reader(r io.Reader) io.Reader {
	if b == nil {
		return nil
	}

	pr := passthru.NewReader(r, b.total)
	pr.Update = b.Add
	b.reader = pr

	return pr
}

// Writer creates and returns pass-thru proxy reader
func (b *Bar) Writer(w io.Writer) io.Writer {
	if b == nil {
		return nil
	}

	pw := passthru.NewWriter(w, b.total)
	pw.Update = b.Add
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
	defer b.mu.RUnlock()

	result := b.renderElements(isFinished)

	// render text only if changed or if finished
	if b.buffer != result || isFinished {
		fmtc.TPrint(result)
	}

	if b.total > 0 {
		b.buffer = result
	}

	if isFinished {
		fmtc.NewLine()
		b.finished = true
		close(b.finishChan)
		b.finishGroup.Done()
	}
}

// renderElements returns text with all progress bar graphics and text
func (b *Bar) renderElements(isFinished bool) string {
	var size, totalSize int
	var name, percentage, bar, progress, speed, remaining string
	var statSpeed float64
	var statRemaining time.Duration

	if b.passThruCalc != nil && (b.settings.ShowSpeed || b.settings.ShowRemaining) {
		if isFinished {
			statRemaining = time.Since(b.startTime)
			statSpeed = float64(b.current) / statRemaining.Seconds()
		} else {
			statSpeed, statRemaining = b.passThruCalc.Calculate(b.current)
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
	size := mathutil.Max(5, mathutil.Max(MIN_WIDTH, b.settings.Width)-dataSize)

	if b.total <= 0 {
		return b.renderPlaceholder(size)
	}

	if b.current >= b.total {
		switch fmtc.DisableColors || b.settings.BarFgColorTag == "" {
		case true:
			return strings.Repeat(PROGRESS_BAR_SYMBOL, size)
		case false:
			return b.settings.BarFgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, size) + "{!}"
		}
	}

	cur := int((float64(b.current) / float64(b.total)) * float64(size))

	if fmtc.DisableColors || b.settings.BarFgColorTag == "" {
		return strings.Repeat(PROGRESS_BAR_SYMBOL, cur) + strings.Repeat(" ", size-cur)
	}

	return b.settings.BarFgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, cur) + b.settings.BarBgColorTag + strings.Repeat(PROGRESS_BAR_SYMBOL, size-cur) + "{!}"
}

// renderPlaceholder returns placeholder bar graphics
func (b *Bar) renderPlaceholder(size int) string {
	var result string

	disableColors := fmtc.DisableColors || b.settings.BarFgColorTag == ""

	for i := 0; i < size; i++ {
		if disableColors {
			if i%3 == b.phCounter {
				result += PROGRESS_BAR_SYMBOL
			} else {
				result += " "
			}
		} else {
			if i%3 == b.phCounter {
				result += b.settings.BarFgColorTag
			} else {
				result += b.settings.BarBgColorTag
			}

			result += PROGRESS_BAR_SYMBOL
		}
	}

	b.phCounter++

	if b.phCounter == 3 {
		b.phCounter = 0
	}

	if disableColors {
		return result
	}

	return result + "{!}"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate validates settings struct
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
