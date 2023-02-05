// Package mathutil provides some additional math methods
package mathutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Between returns value between min and max values
func Between(val, min, max int) int {
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
func Min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max returns a greater value
func Max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenU returns value between min and max values
func BetweenU(val, min, max uint) uint {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// MinU returns a smaller value
func MinU(val1, val2 uint) uint {
	if val1 < val2 {
		return val1
	}

	return val2
}

// MaxU returns a greater value
func MaxU(val1, val2 uint) uint {
	if val1 > val2 {
		return val1
	}

	return val2
}

// Between8 returns value between min and max values
func Between8(val, min, max int8) int8 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Min8 returns a smaller value
func Min8(val1, val2 int8) int8 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max8 returns a greater value
func Max8(val1, val2 int8) int8 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// Between16 returns value between min and max values
func Between16(val, min, max int16) int16 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Min16 returns a smaller value
func Min16(val1, val2 int16) int16 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max16 returns a greater value
func Max16(val1, val2 int16) int16 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// Between32 returns value between min and max values
func Between32(val, min, max int32) int32 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Min32 returns a smaller value
func Min32(val1, val2 int32) int32 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max32 returns a greater value
func Max32(val1, val2 int32) int32 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// Between64 returns value between min and max values
func Between64(val, min, max int64) int64 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Min64 returns a smaller value
func Min64(val1, val2 int64) int64 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// Max64 returns a greater value
func Max64(val1, val2 int64) int64 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenU8 returns value between min and max values
func BetweenU8(val, min, max uint8) uint8 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// MinU8 returns a smaller value
func MinU8(val1, val2 uint8) uint8 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// MaxU8 returns a greater value
func MaxU8(val1, val2 uint8) uint8 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenU16 returns value between min and max values
func BetweenU16(val, min, max uint16) uint16 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// MinU16 returns a smaller value
func MinU16(val1, val2 uint16) uint16 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// MaxU16 returns a greater value
func MaxU16(val1, val2 uint16) uint16 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenU32 returns value between min and max values
func BetweenU32(val, min, max uint32) uint32 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// MinU32 returns a smaller value
func MinU32(val1, val2 uint32) uint32 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// MaxU32 returns a greater value
func MaxU32(val1, val2 uint32) uint32 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenU64 returns value between min and max values
func BetweenU64(val, min, max uint64) uint64 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// MinU64 returns a smaller value
func MinU64(val1, val2 uint64) uint64 {
	if val1 < val2 {
		return val1
	}

	return val2
}

// MaxU64 returns a greater value
func MaxU64(val1, val2 uint64) uint64 {
	if val1 > val2 {
		return val1
	}

	return val2
}

// BetweenF returns value between min and max values
func BetweenF(val, min, max float64) float64 {
	return BetweenF64(val, min, max)
}

// BetweenF32 returns value between min and max values
func BetweenF32(val, min, max float32) float32 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// BetweenF64 returns value between min and max values
func BetweenF64(val, min, max float64) float64 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

// Abs returns absolute value
func Abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

// Abs8 returns absolute value
func Abs8(val int8) int8 {
	if val < 0 {
		return val * -1
	}

	return val
}

// Abs16 returns absolute value
func Abs16(val int16) int16 {
	if val < 0 {
		return val * -1
	}

	return val
}

// Abs32 returns absolute value
func Abs32(val int32) int32 {
	if val < 0 {
		return val * -1
	}

	return val
}

// Abs64 returns absolute value
func Abs64(val int64) int64 {
	if val < 0 {
		return val * -1
	}

	return val
}

// AbsF returns absolute value
func AbsF(val float64) float64 {
	return AbsF64(val)
}

// AbsF32 returns absolute value
func AbsF32(val float32) float32 {
	if val < 0 {
		return val * -1
	}

	return val
}

// AbsF64 returns absolute value
func AbsF64(val float64) float64 {
	if val < 0 {
		return val * -1
	}

	return val
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
