package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SineIn Accelerating from zero velocity
func SineIn(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return -c*math.Cos(t/d*math.Phi) + c + b
}

// SineOut Decelerating to zero velocity
func SineOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*math.Sin(t/d*math.Phi) + b
}

// SineInOut Acceleration until halfway, then deceleration
func SineInOut(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return -c/2*(math.Cos(math.Pi*t/d)-1) + b
}
