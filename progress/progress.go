// Package progress provides methods and structs for creating terminal progress bar
package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"pkg.re/essentialkaos/ek.v12/fmtc"
	"pkg.re/essentialkaos/ek.v12/fmtutil"
	"pkg.re/essentialkaos/ek.v12/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MIN_WIDTH is minimal progress bar width
const MIN_WIDTH = 80

// PROGRESS_BAR_SYMBOL is symbol for creating progress bar
const PROGRESS_BAR_SYMBOL = "—"

// ////////////////////////////////////////////////////////////////////////////////// //

// Bar is progress bar struct
type Bar struct {
	settings Settings

	startTime time.Time
	started   bool
	finished  bool

	finishChan chan bool

	current int64
	total   int64
	name    string

	buffer string

	ticker       *time.Ticker
	passThruCalc *PassThruCalc
	phCounter    int

	reader *passThruReader
	writer *passThruWriter

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

	IsSize bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

type passThruReader struct {
	io.Reader
	bar *Bar
}

type passThruWriter struct {
	io.Writer
	bar *Bar
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
}

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
func (b *Bar) Start() {
	if b.IsStarted() && !b.IsFinished() {
		return
	}

	b.phCounter = 0
	b.current = 0
	b.started = true
	b.finished = false
	b.startTime = time.Now()
	b.ticker = time.NewTicker(b.settings.RefreshRate)
	b.finishChan = make(chan bool)

	if b.total > 0 {
		b.passThruCalc = NewPassThruCalc(b.total, 10.0)
	}

	go b.renderer()
}

// Finish finishes progress processing
func (b *Bar) Finish() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.finished || !b.started {
		return
	}

	fmtc.TPrintf(b.renderElements())
	fmtc.NewLine()

	b.finishChan <- true
}

// UpdateSettings updates progress settings
func (b *Bar) UpdateSettings(s Settings) {
	b.mu.Lock()
	b.settings = s
	b.mu.Unlock()
}

// SetName sets progress bar name
func (b *Bar) SetName(name string) {
	b.mu.Lock()
	b.name = name
	b.mu.Unlock()
}

// Name returns progress bar name
func (b *Bar) Name() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.name
}

// SetTotal sets total progress bar value
func (b *Bar) SetTotal(v int64) {
	b.mu.Lock()

	if b.passThruCalc == nil {
		b.passThruCalc = NewPassThruCalc(v, 10.0)
	} else {
		b.passThruCalc.SetTotal(v)
	}

	b.mu.Unlock()

	atomic.StoreInt64(&b.total, v)
}

// Total returns total progress bar value
func (b *Bar) Total() int64 {
	return atomic.LoadInt64(&b.total)
}

// SetCurrent sets current progress bar value
func (b *Bar) SetCurrent(v int64) {
	atomic.StoreInt64(&b.current, v)
}

// Current returns current progress bar value
func (b *Bar) Current() int64 {
	return atomic.LoadInt64(&b.current)
}

func (b *Bar) Add(v int) {
	atomic.AddInt64(&b.current, int64(v))
}

// Add64 adds given value ti
func (b *Bar) Add64(v int64) {
	atomic.AddInt64(&b.current, v)
}

// IsFinished returns true if progress proccesing is finished
func (b *Bar) IsFinished() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.finished
}

// IsStarted returns true if progress proccesing is started
func (b *Bar) IsStarted() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.started
}

// Reader creates and returns pass thru proxy reader
func (b *Bar) Reader(r io.Reader) io.Reader {
	if b.reader != nil {
		b.reader.Reader = r
	} else {
		b.reader = &passThruReader{
			Reader: r,
			bar:    b,
		}
	}

	return b.reader
}

// Writer creates and returns pass thru proxy reader
func (b *Bar) Writer(w io.Writer) io.Writer {
	if b.writer != nil {
		b.writer.Writer = w
	} else {
		b.writer = &passThruWriter{
			Writer: w,
			bar:    b,
		}
	}

	return b.writer
}

// ////////////////////////////////////////////////////////////////////////////////// //

// renderer is rendering loop func
func (b *Bar) renderer() {
	for {
		select {
		case <-b.finishChan:
			b.finished = true
			b.ticker.Stop()
			b.render()
			return
		case <-b.ticker.C:
			b.render()
		}
	}
}

// render renders current progress bar state
func (b *Bar) render() {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if !b.finished && b.total > 0 && b.current >= b.total {
		b.finished = true
		b.ticker.Stop()
	}

	result := b.renderElements()

	// render text only if changed
	if b.buffer != result {
		fmtc.TPrintf(result)
	}

	if b.total > 0 {
		b.buffer = result
	}

	if b.finished {
		fmtc.NewLine()
	}
}

// renderElements returns text with all progress bar graphics and text
func (b *Bar) renderElements() string {
	var size, totalSize int
	var name, percentage, bar, progress, speed, remaining string
	var statSpeed float64
	var statRemaining time.Duration

	if b.passThruCalc != nil && (b.settings.ShowSpeed || b.settings.ShowRemaining) {
		if b.finished {
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
		result = "100%%"
	} else {
		result = fmt.Sprintf("%5.1f", perc) + "%%"
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

// Read reads data and updates progress bar
func (r *passThruReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)

	if n > 0 {
		r.bar.Add(n)
	}

	return n, err
}

// Write writes data and updates progress bar
func (w *passThruWriter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)

	if n > 0 {
		w.bar.Add(n)
	}

	return n, err
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
