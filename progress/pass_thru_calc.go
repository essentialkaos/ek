package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _MAX_REMAINING = 5999 * time.Second

// ////////////////////////////////////////////////////////////////////////////////// //

// PassThruCalc is pass thru calculator struct
type PassThruCalc struct {
	total      float64
	prev       float64
	speed      float64
	decay      float64
	remaining  time.Duration
	lastUpdate time.Time
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewPassThruCalc creates new pass thru calculator
func NewPassThruCalc(total int64, winSizeSec float64) *PassThruCalc {
	return &PassThruCalc{
		total: float64(total),
		decay: 2 / (winSizeSec + 1),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Calculate calculates number of objects per seconds and remaining time
func (p *PassThruCalc) Calculate(v int64) (float64, time.Duration) {
	c := float64(v)
	now := time.Now()

	if !p.lastUpdate.IsZero() && now.Sub(p.lastUpdate) < time.Second/2 {
		return p.speed, p.remaining
	}

	speed := math.Abs(c-p.prev) / time.Since(p.lastUpdate).Seconds()

	p.speed = (speed * p.decay) + (p.speed * (1.0 - p.decay))

	if p.prev != 0 && p.speed > 0 {
		p.remaining = time.Duration((p.total-c)/p.speed) * time.Second
	}

	if p.remaining > _MAX_REMAINING {
		p.remaining = _MAX_REMAINING
	}

	p.prev = c
	p.lastUpdate = now

	return p.speed, p.remaining
}

// SetTotal sets total number of objects
func (p *PassThruCalc) SetTotal(v int64) {
	p.total = float64(v)
}
