// Package ansi provides methods for working with ANSI/VT100 control sequences
package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// HasCodes returns true if given string contains ANSI/VT100 control sequences
func HasCodes(s string) bool {
	for _, r := range s {
		if r == 0x1B {
			return true
		}
	}

	return false
}

// RemoveCodes removes all ANSI/VT100 control sequences from given string
func RemoveCodes(s string) string {
	if !HasCodes(s) || s == "" {
		return s
	}

	var b strings.Builder
	var skip bool

	for _, r := range s {
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

		b.WriteRune(r)
	}

	return b.String()
}

// ////////////////////////////////////////////////////////////////////////////////// //
