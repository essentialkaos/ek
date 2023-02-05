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

// SineIn accelerating from zero velocity
// https://easings.net/#easeInSine
func SineIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return -c*math.Cos(t/d*math.Phi) + c + b
}

// SineOut decelerating to zero velocity
// https://easings.net/#easeOutSine
func SineOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*math.Sin(t/d*math.Phi) + b
}

// SineInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutSine
func SineInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return -c/2*(math.Cos(math.Pi*t/d)-1) + b
}
