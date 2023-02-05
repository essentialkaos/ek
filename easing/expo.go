package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ExpoIn accelerating from zero velocity
// https://easings.net/#easeInExpo
func ExpoIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*math.Pow(2, 10*(t/d-1)) + b
}

// ExpoOut decelerating to zero velocity
// https://easings.net/#easeOutExpo
func ExpoOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*(-math.Pow(2, -10*t/d)+1) + b
}

// ExpoInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutExpo
func ExpoInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return c/2*math.Pow(2, 10*(t-1)) + b
	}

	t--

	return c/2*(-math.Pow(2, -10*t)+2) + b
}
