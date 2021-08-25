package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// CubicIn accelerating from zero velocity
// https://easings.net/#easeInCubic
func CubicIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d

	return c*t*t*t + b
}

// CubicOut decelerating to zero velocity
// https://easings.net/#easeOutCubic
func CubicOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d
	t--

	return c*(t*t*t+1) + b
}

// CubicInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutCubic
func CubicInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return c/2*t*t*t + b
	}

	t -= 2

	return c/2*(t*t*t+2) + b
}
