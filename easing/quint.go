package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// QuintIn accelerating from zero velocity
func QuintIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d
	return c*t*t*t*t*t + b
}

// QuintOut decelerating to zero velocity
func QuintOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d
	t--

	return c*(t*t*t*t*t+1) + b
}

// QuintInOut acceleration until halfway, then deceleration
func QuintInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	t /= d / 2

	if t < 1 {
		return c/2*t*t*t*t*t + b
	}

	t -= 2

	return c/2*(t*t*t*t*t+2) + b
}
