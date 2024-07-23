// Package passthru provides Reader and Writer with information about the amount
// of data being passed.
package passthru

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Reader is pass-thru Reader
type Reader struct {
	Calculator *Calculator
	Update     func(n int)

	r       io.Reader
	current int64
	total   int64
}

// Writer is pass-thru Writer
type Writer struct {
	Calculator *Calculator
	Update     func(n int)

	w       io.Writer
	current int64
	total   int64
}

// Calculator calculates pass-thru speed and remaining time
type Calculator struct {
	total      float64
	prev       float64
	speed      float64
	decay      float64
	remaining  time.Duration
	lastUpdate time.Time

	mu sync.Mutex
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNilReader = errors.New("Reader is nil")
	ErrNilWriter = errors.New("Writer is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader creates new passthru reader
func NewReader(reader io.Reader, total int64) *Reader {
	return &Reader{r: reader, total: total}
}

// NewWriter creates new passthru writer
func NewWriter(writer io.Writer, total int64) *Writer {
	return &Writer{w: writer, total: total}
}

// NewCalculator creates new Calculator struct
func NewCalculator(total int64, winSizeSec float64) *Calculator {
	return &Calculator{
		total: float64(total),
		decay: 2 / (winSizeSec + 1),
		mu:    sync.Mutex{},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read implements the standard Read interface
func (r *Reader) Read(p []byte) (int, error) {
	if r == nil || r.r == nil {
		return -1, ErrNilReader
	}

	n, err := r.r.Read(p)

	atomic.AddInt64(&r.current, int64(n))

	if r.Update != nil && n > 0 {
		r.Update(n)
	}

	return n, err
}

// Current returns read amount of data
func (r *Reader) Current() int64 {
	if r == nil {
		return 0
	}

	return atomic.LoadInt64(&r.current)
}

// Total returns total amount of data
func (r *Reader) Total() int64 {
	if r == nil {
		return 0
	}

	return atomic.LoadInt64(&r.total)
}

// Progress returns percentage of data read
func (r *Reader) Progress() float64 {
	if r == nil {
		return 0
	}

	return (float64(r.Current()) / float64(r.Total())) * 100.0
}

// SetTotal sets total amount of data
func (r *Reader) SetTotal(total int64) {
	if r == nil {
		return
	}

	atomic.StoreInt64(&r.total, total)
}

// Speed calculates passthru speed and time remaining to process all data.
// If Calculator was not set on the first call, a default calculator with
// a 10 second window is created.
func (r *Reader) Speed() (float64, time.Duration) {
	if r == nil || r.Total() <= 0 {
		return 0, 0
	}

	if r.Calculator == nil {
		r.Calculator = NewCalculator(r.Total(), 10.0)
	}

	return r.Calculator.Calculate(atomic.LoadInt64(&r.current))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Write implements the standard Write interface
func (w *Writer) Write(p []byte) (int, error) {
	if w == nil || w.w == nil {
		return -1, ErrNilWriter
	}

	n, err := w.w.Write(p)

	atomic.AddInt64(&w.current, int64(n))

	if w.Update != nil && n > 0 {
		w.Update(n)
	}

	return n, err
}

// Current returns written amount of data
func (w *Writer) Current() int64 {
	if w == nil {
		return 0
	}

	return atomic.LoadInt64(&w.current)
}

// Total returns total amount of data
func (w *Writer) Total() int64 {
	if w == nil {
		return 0
	}

	return atomic.LoadInt64(&w.total)
}

// Progress returns percentage of data written
func (w *Writer) Progress() float64 {
	if w == nil {
		return 0
	}

	return (float64(w.Current()) / float64(w.Total())) * 100.0
}

// SetTotal sets total amount of data
func (w *Writer) SetTotal(total int64) {
	if w == nil {
		return
	}

	atomic.StoreInt64(&w.total, total)
}

// Speed calculates passthru speed and time remaining to process all data.
// If Calculator was not set on the first call, a default calculator with
// a 10 second window is created.
func (w *Writer) Speed() (float64, time.Duration) {
	if w == nil || w.Total() <= 0 {
		return 0, 0
	}

	if w.Calculator == nil {
		w.Calculator = NewCalculator(w.Total(), 10.0)
	}

	return w.Calculator.Calculate(atomic.LoadInt64(&w.current))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Calculate calculates speed and remaining time
func (c *Calculator) Calculate(current int64) (float64, time.Duration) {
	if c == nil {
		return 0, 0
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cur := float64(current)
	now := time.Now()

	if c.total == 0 {
		return 0, 0
	}

	if !c.lastUpdate.IsZero() && now.Sub(c.lastUpdate) < time.Second/2 {
		return c.speed, c.remaining
	}

	speed := math.Abs(cur-c.prev) / time.Since(c.lastUpdate).Seconds()
	c.speed = (speed * c.decay) + (c.speed * (1.0 - c.decay))

	if c.prev != 0 && c.speed > 0 {
		c.remaining = time.Duration((c.total-cur)/c.speed) * time.Second
	}

	if c.remaining > time.Hour {
		c.remaining = time.Hour
	}

	c.prev = cur
	c.lastUpdate = now

	return c.speed, c.remaining
}

// SetTotal sets total number of objects
func (c *Calculator) SetTotal(total int64) {
	if c == nil {
		return
	}

	c.mu.Lock()
	c.total = float64(total)
	c.mu.Unlock()
}
