// Package ansi provides methods for working with ANSI/VT100 control sequences
package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
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

// HasCodes returns true if given byte slice contains ANSI/VT100 control sequences
func HasCodesBytes(b []byte) bool {
	for _, r := range b {
		if r == 0x1B {
			return true
		}
	}

	return false
}

// RemoveCodesBytes returns string without all ANSI/VT100 control sequences
func RemoveCodes(s string) string {
	if s == "" || !HasCodes(s) {
		return s
	}

	return string(RemoveCodesBytes([]byte(s)))
}

// RemoveCodesBytes returns byte slice without all ANSI/VT100 control sequences
func RemoveCodesBytes(b []byte) []byte {
	if len(b) == 0 || !HasCodesBytes(b) {
		return b
	}

	var buf bytes.Buffer
	var skip bool

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

		buf.WriteByte(r)
	}

	return buf.Bytes()
}

// ////////////////////////////////////////////////////////////////////////////////// //
