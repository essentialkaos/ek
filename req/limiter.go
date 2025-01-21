package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "time"

// ////////////////////////////////////////////////////////////////////////////////// //

// Limiter is request limiter
type Limiter struct {
	lastCall time.Time
	delay    time.Duration
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewLimiter creates a new limiter. If rps is less than or equal to 0, it returns nil.
func NewLimiter(rps float64) *Limiter {
	if rps <= 0 {
		return nil
	}

	return &Limiter{
		delay: time.Duration(float64(time.Second) / rps),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Wait blocks current goroutine execution until next time slot become available
func (l *Limiter) Wait() {
	if l == nil {
		return
	}

	if l.lastCall.IsZero() {
		l.lastCall = time.Now()
		return
	}

	w := time.Since(l.lastCall)

	if w < l.delay {
		time.Sleep(l.delay - w)
	}

	l.lastCall = time.Now()
}
