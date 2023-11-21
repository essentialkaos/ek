//go:build !windows
// +build !windows

// Package tty provides methods for working with TTY
package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var stdin = os.Stdin
var stdout = os.Stdout

// ////////////////////////////////////////////////////////////////////////////////// //

// IsTTY returns true if current output device is TTY
func IsTTY() bool {
	si, _ := stdin.Stat()
	so, _ := stdout.Stat()

	if si.Mode()&os.ModeCharDevice != 0 &&
		so.Mode()&os.ModeCharDevice == 0 &&
		!IsFakeTTY() {
		return false
	}

	return true
}

// IsFakeTTY returns true is fake TTY is used
func IsFakeTTY() bool {
	return os.Getenv("FAKETTY") != ""
}
