package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type PassThruCalc struct {
	current    float64
	total      float64
	winSize    time.Duration
	lastUpdate time.Time
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewPassThruCalc creates new pass thru calculator
func NewPassThruCalc(total int64, winSize time.Duration) *PassThruCalc {
	return &PassThruCalc{
		total:   float64(total),
		winSize: winSize,
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Calculate calculates number of objects per seconds and remaining time
func (p *PassThruCalc) Calculate(v int64) (float64, time.Duration) {
	if p.lastUpdate.IsZero() {
		p.lastUpdate = time.Now()
		p.current = float64(v)
		return 0, 0
	}

	c := float64(v)

	if c > p.total {
		return 0, 0
	}

	if c-p.current <= 0 {
		return 0, 99 * time.Minute
	}

	t := time.Since(p.lastUpdate)
	n := (c - p.current) / t.Seconds()
	r := time.Duration((p.total-c)/n) * time.Second

	if t >= p.winSize {
		p.current = c
		p.lastUpdate = time.Now()
	}

	return n, r
}

// SetTotal sets total number of objects
func (p *PassThruCalc) SetTotal(v int64) {
	p.total = float64(v)
}
