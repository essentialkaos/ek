// Package mathutil provides some additional math methods
package mathutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type NumericNeg interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Between returns value between min and max values
func Between[N Numeric](val, min, max N) N {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Min returns a smaller value
func Min[N Numeric](val1, val2 N) N {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max returns a greater value
func Max[N Numeric](val1, val2 N) N {
	if val1 > val2 {
		return val1
	}

	return val2
}

// Abs returns absolute value
func Abs[N NumericNeg](val N) N {
	if val < 0 {
		return val * -1
	}

	return val
}

// Perc calculates percentage
func Perc[N Numeric](val1, val2 N) float64 {
	if val2 == 0 {
		return 0
	}

	return float64(val1) / float64(val2) * 100.0
}

// Round returns rounded value
func Round(v float64, p int) float64 {
	pow := math.Pow(10, float64(p))
	digit := pow * v
	_, div := math.Modf(digit)

	if div >= 0.5 {
		return math.Ceil(digit) / pow
	}

	return math.Floor(digit) / pow
}

// ////////////////////////////////////////////////////////////////////////////////// //
