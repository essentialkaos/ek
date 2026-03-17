// Package mathutil provides some additional math methods
package mathutil

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

// Integer is a type constraint that matches all signed and unsigned integer types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float is a type constraint that matches float32 and float64
type Float interface {
	~float32 | ~float64
}

// Numeric is a type constraint that matches all integer, unsigned integer, uintptr,
// and float types
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// NumericNeg is a type constraint that matches numeric types capable of holding
// negative values
type NumericNeg interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// ////////////////////////////////////////////////////////////////////////////////// //

// B returns positive if cond is true, otherwise returns negative
func B[N Numeric](cond bool, positive, negative N) N {
	if cond {
		return positive
	}

	return negative
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsInt returns true if every character in s is a valid integer symbol (digits and
// leading minus).
//
// Note: this does not structurally validate the value; use [strconv.Atoi] for
// full parsing.
func IsInt(s string) bool {
	if s == "" {
		return false
	}

	var i int

	if s[0] == '-' {
		i = 1
	}

	if i == len(s) {
		return false
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}

// IsFloat returns true if every character in s is a valid float symbol (digits,
// leading minus, and decimal point).
//
// Note: this does not structurally validate the value; use [strconv.ParseFloat] for
// full parsing.
func IsFloat(s string) bool {
	if s == "" {
		return false
	}

	var i int

	if s[i] == '-' {
		i++
	}

	if i == len(s) {
		return false
	}

	var hasDot, hasDigit bool

	for ; i < len(s); i++ {
		c := s[i]

		switch {
		case c >= '0' && c <= '9':
			hasDigit = true

		case c == '.':
			if hasDot {
				return false
			}
			hasDot = true

		default:
			return false
		}
	}

	return hasDigit
}

// IsNumber returns true if s passes either [IsInt] or [IsFloat].
func IsNumber(s string) bool {
	return IsFloat(s)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Between returns val clamped to the range [min, max]
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

// Abs returns the absolute value of val
func Abs[N NumericNeg](val N) N {
	if val < 0 {
		return val * -1
	}

	return val
}

// Perc returns the percentage that current represents of total.
// Returns 0 if total is 0.
func Perc[N Numeric](current, total N) float64 {
	if current == 0 || total == 0 {
		return 0
	}

	return (float64(current) / float64(total)) * 100.0
}

// FromPerc returns the value corresponding to perc percent of total.
// Returns 0 if perc is <= 0 or total is 0.
func FromPerc(perc float64, total float64) float64 {
	if perc <= 0 || total == 0 {
		return 0
	}

	return (total / 100.0) * perc
}

// Round returns v rounded to p decimal places using half-up rounding.
// p should be in range [0, 15] for float64 precision.
func Round(v float64, p int) float64 {
	pow := math.Pow(10, float64(Between(p, 0, 15)))
	digit := pow * v

	_, frac := math.Modf(digit)

	if frac >= 0.5 {
		return math.Ceil(digit) / pow
	}

	return math.Floor(digit) / pow
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Min returns a smaller value
//
// Deprecated: Use built-in function `min` instead
func Min[N Numeric](val1, val2 N) N {
	return min(val1, val2)
}

// Max returns a greater value
//
// Deprecated: Use built-in function `max` instead
func Max[N Numeric](val1, val2 N) N {
	return max(val1, val2)
}

// ////////////////////////////////////////////////////////////////////////////////// //
