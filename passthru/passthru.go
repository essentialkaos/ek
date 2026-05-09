// Package passthru provides Reader and Writer with information about the amount
// of data being passed.
package passthru

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

// Reader is a pass-through wrapper around io.Reader that tracks the number
// of bytes read
type Reader struct {
	Calculator *Calculator

	UpdateN        func(n int)
	Update         func(r *Reader)
	UpdateInterval time.Duration

	r          io.Reader
	current    int64
	total      int64
	lastUpdate time.Time
}

// Writer is a pass-through wrapper around io.Writer that tracks the number
// of bytes written
type Writer struct {
	Calculator *Calculator

	UpdateN        func(n int)
	Update         func(w *Writer)
	UpdateInterval time.Duration

	w          io.Writer
	current    int64
	total      int64
	lastUpdate time.Time
}

// Calculator computes exponentially weighted moving average speed and estimated
// time remaining
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
	// ErrNilReader is returned when Read is called on a nil [Reader]
	ErrNilReader = errors.New("reader is nil")

	// ErrNilWriter is returned when Write is called on a nil [Writer]
	ErrNilWriter = errors.New("writer is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader wraps the given io.Reader and tracks up to total bytes
func NewReader(reader io.Reader, total int64) *Reader {
	return &Reader{r: reader, total: total}
}

// NewWriter wraps the given io.Writer and tracks up to total bytes
func NewWriter(writer io.Writer, total int64) *Writer {
	return &Writer{w: writer, total: total}
}

// NewCalculator returns a Calculator for the given total size using an EWMA window
// of winSizeSec seconds
func NewCalculator(total int64, winSizeSec float64) *Calculator {
	return &Calculator{
		total: float64(total),
		decay: 2 / (winSizeSec + 1),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read implements [io.Reader] and atomically accumulates the number of bytes read
func (r *Reader) Read(p []byte) (int, error) {
	if r == nil || r.r == nil {
		return 0, ErrNilReader
	}

	n, err := r.r.Read(p)

	atomic.AddInt64(&r.current, int64(n))

	if n > 0 && (r.Update != nil || r.UpdateN != nil) {
		now := time.Now()

		if r.UpdateInterval == 0 || (r.UpdateInterval > 0 && now.Sub(r.lastUpdate) > r.UpdateInterval) {
			if r.UpdateN != nil {
				r.UpdateN(n)
			}

			if r.Update != nil {
				r.Update(r)
			}
		}

		r.lastUpdate = now
	}

	return n, err
}

// Current returns the total number of bytes read so far
func (r *Reader) Current() int64 {
	if r == nil {
		return 0
	}

	return atomic.LoadInt64(&r.current)
}

// Total returns the expected total number of bytes to be read
func (r *Reader) Total() int64 {
	if r == nil {
		return 0
	}

	return atomic.LoadInt64(&r.total)
}

// Progress returns the percentage of data read relative to the total
func (r *Reader) Progress() float64 {
	if r == nil {
		return 0
	}

	total := r.Total()

	if total == 0 {
		return 0
	}

	return (float64(r.Current()) / float64(total)) * 100.0
}

// SetTotal updates the expected total number of bytes to be read
func (r *Reader) SetTotal(total int64) {
	if r == nil {
		return
	}

	atomic.StoreInt64(&r.total, total)
}

// Speed returns the current read speed in bytes per second and the estimated
// time remaining. A default [Calculator] with a 10-second window is created
// automatically if one was not set.
func (r *Reader) Speed() (float64, time.Duration) {
	if r == nil {
		return 0, 0
	}

	if r.Calculator == nil {
		r.Calculator = NewCalculator(r.Total(), 10.0)
	}

	return r.Calculator.Calculate(atomic.LoadInt64(&r.current))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Write implements [io.Writer] and atomically accumulates the number of bytes written
func (w *Writer) Write(p []byte) (int, error) {
	if w == nil || w.w == nil {
		return 0, ErrNilWriter
	}

	n, err := w.w.Write(p)

	atomic.AddInt64(&w.current, int64(n))

	if n > 0 && (w.Update != nil || w.UpdateN != nil) {
		now := time.Now()

		if w.UpdateInterval == 0 || (w.UpdateInterval > 0 && now.Sub(w.lastUpdate) > w.UpdateInterval) {
			if w.UpdateN != nil {
				w.UpdateN(n)
			}

			if w.Update != nil {
				w.Update(w)
			}
		}

		w.lastUpdate = now
	}

	return n, err
}

// Current returns the total number of bytes written so far
func (w *Writer) Current() int64 {
	if w == nil {
		return 0
	}

	return atomic.LoadInt64(&w.current)
}

// Total returns the expected total number of bytes to be written
func (w *Writer) Total() int64 {
	if w == nil {
		return 0
	}

	return atomic.LoadInt64(&w.total)
}

// Progress returns the percentage of data written relative to the total
func (w *Writer) Progress() float64 {
	if w == nil {
		return 0
	}

	total := w.Total()

	if total == 0 {
		return 0
	}

	return (float64(w.Current()) / float64(total)) * 100.0
}

// SetTotal updates the expected total number of bytes to be written
func (w *Writer) SetTotal(total int64) {
	if w == nil {
		return
	}

	atomic.StoreInt64(&w.total, total)
}

// Speed returns the current write speed in bytes per second and the estimated
// time remaining. A default [Calculator] with a 10-second window is created automatically
// if one was not set.
func (w *Writer) Speed() (float64, time.Duration) {
	if w == nil {
		return 0, 0
	}

	if w.Calculator == nil {
		w.Calculator = NewCalculator(w.Total(), 10.0)
	}

	return w.Calculator.Calculate(atomic.LoadInt64(&w.current))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Calculate updates the EWMA speed estimate using the current byte count and returns
// speed and remaining time. Results are throttled: cached values are returned if called
// within 500ms of the last update.
func (c *Calculator) Calculate(current int64) (float64, time.Duration) {
	if c == nil {
		return 0, 0
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cur := float64(current)
	now := time.Now()

	if c.lastUpdate.IsZero() {
		c.prev, c.lastUpdate = cur, now
		return 0, 0
	}

	if !c.lastUpdate.IsZero() && now.Sub(c.lastUpdate) < time.Second/2 {
		return c.speed, c.remaining
	}

	speed := math.Abs(cur-c.prev) / time.Since(c.lastUpdate).Seconds()
	c.speed = (speed * c.decay) + (c.speed * (1.0 - c.decay))

	if c.total != 0 && c.prev != 0 && c.speed > 0 {
		c.remaining = time.Duration((c.total-cur)/c.speed) * time.Second
	}

	if c.remaining > time.Hour {
		c.remaining = time.Hour
	}

	c.prev, c.lastUpdate = cur, now

	return c.speed, c.remaining
}

// SetTotal updates the total number of bytes used for remaining-time estimation
func (c *Calculator) SetTotal(total int64) {
	if c == nil {
		return
	}

	c.mu.Lock()
	c.total = float64(total)
	c.mu.Unlock()
}
