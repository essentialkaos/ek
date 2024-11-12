package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "time"

// ////////////////////////////////////////////////////////////////////////////////// //

// limiter is request limiter
type limiter struct {
	lastCall time.Time
	delay    time.Duration
}

// ////////////////////////////////////////////////////////////////////////////////// //

// createLimiter creates new limiter
func createLimiter(rps float64) *limiter {
	if rps <= 0 {
		return nil
	}

	return &limiter{
		delay: time.Duration(float64(time.Second) / rps),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Wait blocks current goroutine execution until next time slot become available
func (l *limiter) Wait() {
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
