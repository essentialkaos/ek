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

const DOUBLE_PI = math.Pi * 2

// ////////////////////////////////////////////////////////////////////////////////// //

// ElasticIn accelerating from zero velocity
// https://easings.net/#easeInElastic
func ElasticIn(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	p := d * 0.3

	t /= d

	var s float64

	if c < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DOUBLE_PI * math.Asin(c/c)
	}

	t--

	return -(c * math.Pow(2, 10*t) * math.Sin((t*d-s)*DOUBLE_PI/p)) + b
}

// ElasticOut decelerating to zero velocity
// https://easings.net/#easeOutElastic
func ElasticOut(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	p := d * 0.3

	t /= d

	var s float64

	if c < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DOUBLE_PI * math.Asin(c/c)
	}

	return c*math.Pow(2, -10*t)*math.Sin((t*d-s)*DOUBLE_PI/p) + c + b
}

// ElasticInOut acceleration until halfway, then deceleration
// https://easings.net/#easeInOutElastic
func ElasticInOut(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	p := d * 0.45

	t /= d / 2

	var s float64

	if c < math.Abs(c) {
		s = p / 4
	} else {
		s = p / DOUBLE_PI * math.Asin(c/c)
	}

	t--

	if t < 0 {
		return -0.5*(c*math.Pow(2, 10*t)*math.Sin((t*d-s)*DOUBLE_PI/p)) + b
	}

	return c*math.Pow(2, -10*t)*math.Sin((t*d-s)*DOUBLE_PI/p)*0.5 + c + b
}
