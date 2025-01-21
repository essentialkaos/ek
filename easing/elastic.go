package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ElasticIn accelerating from zero velocity
// https://easings.net/#easeInElastic
func ElasticIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	var s float64

	p := d * 0.3
	a := c

	if t == 0 {
		return b
	}

	t /= d

	if t == 1 {
		return b + c
	}

	if a < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DoublePi * math.Asin(c/a)
	}

	t--

	return -(a * math.Pow(2, 10*t) * math.Sin((t*d-s)*DoublePi/p)) + b
}

// ElasticOut decelerating to zero velocity
// https://easings.net/#easeOutElastic
func ElasticOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	var s float64

	p := d * 0.3
	a := c

	if t == 0 {
		return b
	}

	t /= d

	if t == 1 {
		return b + c
	}

	if a < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DoublePi * math.Asin(c/a)
	}

	return a*math.Pow(2, -10*t)*math.Sin((t*d-s)*DoublePi/p) + c + b
}

// ElasticInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutElastic
func ElasticInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	var s float64

	p := d * (0.3 * 1.5)
	a := c

	if t == 0 {
		return b
	}

	t /= d / 2

	if t == 2 {
		return b + c
	}

	if a < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DoublePi * math.Asin(c/a)
	}

	t--

	if t < 0 {
		return -0.5*(a*math.Pow(2, 10*t)*math.Sin((t*d-s)*DoublePi/p)) + b
	}

	return a*math.Pow(2, -10*t)*math.Sin((t*d-s)*DoublePi/p)*0.5 + c + b
}
