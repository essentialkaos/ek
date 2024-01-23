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

// BetweenU returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenU(val, min, max uint) uint {
	return Between(val, min, max)
}

// MinU returns a smaller value
//
// Deprecated: Use method Min instead
func MinU(val1, val2 uint) uint {
	return Min(val1, val2)
}

// MaxU returns a greater value
//
// Deprecated: Use method Max instead
func MaxU(val1, val2 uint) uint {
	return Max(val1, val2)
}

// Between8 returns value between min and max values
//
// Deprecated: Use method Between instead
func Between8(val, min, max int8) int8 {
	return Between(val, min, max)
}

// Min8 returns a smaller value
//
// Deprecated: Use method Min instead
func Min8(val1, val2 int8) int8 {
	return Min(val1, val2)
}

// Max8 returns a greater value
//
// Deprecated: Use method Max instead
func Max8(val1, val2 int8) int8 {
	return Max(val1, val2)
}

// Between16 returns value between min and max values
//
// Deprecated: Use method Between instead
func Between16(val, min, max int16) int16 {
	return Between(val, min, max)
}

// Min16 returns a smaller value
//
// Deprecated: Use method Min instead
func Min16(val1, val2 int16) int16 {
	return Min(val1, val2)
}

// Max16 returns a greater value
//
// Deprecated: Use method Max instead
func Max16(val1, val2 int16) int16 {
	return Max(val1, val2)
}

// Between32 returns value between min and max values
//
// Deprecated: Use method Between instead
func Between32(val, min, max int32) int32 {
	return Between(val, min, max)
}

// Min32 returns a smaller value
//
// Deprecated: Use method Min instead
func Min32(val1, val2 int32) int32 {
	return Min(val1, val2)
}

// Max32 returns a greater value
//
// Deprecated: Use method Max instead
func Max32(val1, val2 int32) int32 {
	return Max(val1, val2)
}

// Between64 returns value between min and max values
//
// Deprecated: Use method Between instead
func Between64(val, min, max int64) int64 {
	return Between(val, min, max)
}

// Min64 returns a smaller value
//
// Deprecated: Use method Min instead
func Min64(val1, val2 int64) int64 {
	return Min(val1, val2)
}

// Max64 returns a greater value
//
// Deprecated: Use method Max instead
func Max64(val1, val2 int64) int64 {
	return Max(val1, val2)
}

// BetweenU8 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenU8(val, min, max uint8) uint8 {
	return Between(val, min, max)
}

// MinU8 returns a smaller value
//
// Deprecated: Use method Min instead
func MinU8(val1, val2 uint8) uint8 {
	return Min(val1, val2)
}

// MaxU8 returns a greater value
//
// Deprecated: Use method Max instead
func MaxU8(val1, val2 uint8) uint8 {
	return Max(val1, val2)
}

// BetweenU16 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenU16(val, min, max uint16) uint16 {
	return Between(val, min, max)
}

// MinU16 returns a smaller value
//
// Deprecated: Use method Min instead
func MinU16(val1, val2 uint16) uint16 {
	return Min(val1, val2)
}

// MaxU16 returns a greater value
//
// Deprecated: Use method Max instead
func MaxU16(val1, val2 uint16) uint16 {
	return Max(val1, val2)
}

// BetweenU32 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenU32(val, min, max uint32) uint32 {
	return Between(val, min, max)
}

// MinU32 returns a smaller value
//
// Deprecated: Use method Min instead
func MinU32(val1, val2 uint32) uint32 {
	return Min(val1, val2)
}

// MaxU32 returns a greater value
//
// Deprecated: Use method Max instead
func MaxU32(val1, val2 uint32) uint32 {
	return Max(val1, val2)
}

// BetweenU64 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenU64(val, min, max uint64) uint64 {
	return Between(val, min, max)
}

// MinU64 returns a smaller value
func MinU64(val1, val2 uint64) uint64 {
	return Min(val1, val2)
}

// MaxU64 returns a greater value
//
// Deprecated: Use method Max instead
func MaxU64(val1, val2 uint64) uint64 {
	return Max(val1, val2)
}

// BetweenF returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenF(val, min, max float64) float64 {
	return Between(val, min, max)
}

// BetweenF32 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenF32(val, min, max float32) float32 {
	return Between(val, min, max)
}

// BetweenF64 returns value between min and max values
//
// Deprecated: Use method Between instead
func BetweenF64(val, min, max float64) float64 {
	return Between(val, min, max)
}

// Abs8 returns absolute value
//
// Deprecated: Use method Abs instead
func Abs8(val int8) int8 {
	return Abs(val)
}

// Abs16 returns absolute value
//
// Deprecated: Use method Abs instead
func Abs16(val int16) int16 {
	return Abs(val)
}

// Abs32 returns absolute value
//
// Deprecated: Use method Abs instead
func Abs32(val int32) int32 {
	return Abs(val)
}

// Abs64 returns absolute value
//
// Deprecated: Use method Abs instead
func Abs64(val int64) int64 {
	return Abs(val)
}

// AbsF returns absolute value
//
// Deprecated: Use method Abs instead
func AbsF(val float64) float64 {
	return Abs(val)
}

// AbsF32 returns absolute value
//
// Deprecated: Use method Abs instead
func AbsF32(val float32) float32 {
	return Abs(val)
}

// AbsF64 returns absolute value
//
// Deprecated: Use method Abs instead
func AbsF64(val float64) float64 {
	return Abs(val)
}

// ////////////////////////////////////////////////////////////////////////////////// //
