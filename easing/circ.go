package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CircIn Accelerating from zero velocity
func CircIn(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d

	return -c*(math.Sqrt(1-t*t)-1) + b
}

// CircOut Decelerating to zero velocity
func CircOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d
	t--

	return c*math.Sqrt(1-t*t) + b
}

// CircInOut Acceleration until halfway, then deceleration
func CircInOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d / 2

	if t < 1 {
		return -c/2*(math.Sqrt(1-t*t)-1) + b
	}

	t -= 2

	return c/2*(math.Sqrt(1-t*t)+1) + b
}
