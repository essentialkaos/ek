package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// CubicIn Accelerating from zero velocity
func CubicIn(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d

	return c*t*t*t + b
}

// CubicOut Decelerating to zero velocity
func CubicOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d
	t--

	return c*(t*t*t+1) + b
}

// CubicInOut Acceleration until halfway, then deceleration
func CubicInOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	t /= d / 2

	if t < 1 {
		return c/2*t*t*t + b
	}

	t -= 2

	return c/2*(t*t*t+2) + b
}
