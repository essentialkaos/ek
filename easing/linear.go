package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Linear No easing, no acceleration
func Linear(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*t/d + b
}
