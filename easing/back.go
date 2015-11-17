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

// BackIn Accelerating from zero velocity
func BackIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	s := math.SqrtPi
	t /= d

	return c*t*t*((s+1)*t-s) + b
}

// BackOut Decelerating to zero velocity
func BackOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	s := math.SqrtPi
	t = t/d - 1

	return c*(t*t*((s+1)*t+s)+1) + b
}

// BackInOut Acceleration until halfway, then deceleration
func BackInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	s := math.SqrtPi * 1.525
	t /= d / 2

	if t < 1 {
		return c/2*(t*t*((s+1)*t-s)) + b
	}

	t -= 2

	return c/2*(t*t*((s+1)*t+s)+2) + b
}
