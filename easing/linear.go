package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Linear no easing and no acceleration
func Linear(t, b, c, d float64) float64 {
	if t > d {
		return c
	}

	return c*t/d + b
}
