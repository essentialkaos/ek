package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _HALF_PI = math.Pi / 2

// ////////////////////////////////////////////////////////////////////////////////// //

// SineIn accelerating from zero velocity
// https://easings.net/#easeInSine
func SineIn(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	return -c*math.Cos(t/d*_HALF_PI) + c + b
}

// SineOut decelerating to zero velocity
// https://easings.net/#easeOutSine
func SineOut(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	return c*math.Sin(t/d*_HALF_PI) + b
}

// SineInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutSine
func SineInOut(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	return -c/2*(math.Cos(math.Pi*t/d)-1) + b
}
