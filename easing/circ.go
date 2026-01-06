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

// CircIn accelerating from zero velocity
// https://easings.net/#easeInCirc
func CircIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d

	return -c*(math.Sqrt(1-t*t)-1) + b
}

// CircOut decelerating to zero velocity
// https://easings.net/#easeOutCirc
func CircOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d
	t--

	return c*math.Sqrt(1-t*t) + b
}

// CircInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutCirc
func CircInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return -c/2*(math.Sqrt(1-t*t)-1) + b
	}

	t -= 2

	return c/2*(math.Sqrt(1-t*t)+1) + b
}
