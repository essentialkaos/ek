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

// ElasticIn Accelerating from zero velocity
func ElasticIn(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	s := math.SqrtPi
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

// ElasticOut Decelerating to zero velocity
func ElasticOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	s := math.SqrtPi
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

// ElasticInOut Acceleration until halfway, then deceleration
func ElasticInOut(t, b, c, d float64) float64 {
	if t > d {
		t = d
	}

	s := math.SqrtPi
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
