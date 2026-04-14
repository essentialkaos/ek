package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Linear no easing and no acceleration
func Linear(t, b, c, d float64) float64 {
	switch {
	case t <= 0:
		return b
	case t >= d:
		return b + c
	}

	return c*t/d + b
}
