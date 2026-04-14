// Package ansi provides methods for working with ANSI/VT100 control sequences
package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"slices"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if given string contains ANSI/VT100 control sequences
func Has(s string) bool {
	for _, r := range s {
		if r == 0x1B {
			return true
		}
	}

	return false
}

// HasBytes returns true if given byte slice contains ANSI/VT100 control sequences
func HasBytes(b []byte) bool {
	return slices.Contains(b, 0x1B)
}

// Remove returns string without all ANSI/VT100 control sequences
func Remove(s string) string {
	if s == "" || !Has(s) {
		return s
	}

	return string(removeCodesBytes([]byte(s)))
}

// RemoveBytes returns a byte slice with all ANSI/VT100 SGR control sequences removed.
//
// Aliasing: if the input contains no escape sequences, the original slice is
// returned as-is without any allocation. Mutating the result in that case will
// silently mutate the source buffer. If an independent copy is required
// regardless of content, use [bytes.Clone]:
//
//	safe := bytes.Clone(RemoveBytes(b))
func RemoveBytes(b []byte) []byte {
	if len(b) == 0 || !HasBytes(b) {
		return b
	}

	return removeCodesBytes(b)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// removeCodesBytes removes all ANSI/VT100 control sequences from given byte slice
func removeCodesBytes(b []byte) []byte {
	var skip bool

	result := make([]byte, 0, len(b))

	for _, r := range b {
		if r == 0x1B {
			skip = true
			continue
		}

		if skip {
			if r != 0x6D {
				continue
			}

			skip = false
			continue
		}

		result = append(result, r)
	}

	return result
}
