package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Linear No easing, no acceleration
func Linear(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*t/d + b
}
