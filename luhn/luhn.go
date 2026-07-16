// Package luhn provides methods to calculate and validate Luhn checksum
package luhn

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"

	"github.com/essentialkaos/ek/v14/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyData is returned if input has no data
	ErrEmptyData = errors.New("input is empty")

	// ErrInvalidData is returned if input data is invalid
	ErrInvalidData = errors.New("input is invalid")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsValid returns true if given value contain valid Luhn checksum
func IsValid(v string) bool {
	switch {
	case v == "" || len(v) < 2:
		return false
	case !mathutil.IsInt(v):
		return false
	}

	var sum int
	var double bool

	for i := len(v) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(v[i]))

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}

// Calculate calculates Luhn checksum for given value
func Calculate(v string) (string, error) {
	switch {
	case v == "":
		return "", ErrEmptyData
	case len(v) < 2 || !mathutil.IsInt(v):
		return "", ErrInvalidData
	}

	if IsValid(v) {
		return v[len(v)-1:], nil
	}

	var sum int

	for i := len(v) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(v[i]))

		if (len(v)-1-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	checksum := (10 - (sum % 10)) % 10

	return strconv.Itoa(checksum), nil
}

// Normalize checks if given value contains valid Luhn checksum and if not calculates
// it and append to the given value
func Normalize(v string) (string, error) {
	switch {
	case v == "":
		return "", ErrEmptyData
	case len(v) < 2 || !mathutil.IsInt(v):
		return "", ErrInvalidData
	}

	if IsValid(v) {
		return v, nil
	}

	checksum, _ := Calculate(v)

	return v + checksum, nil
}
