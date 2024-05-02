//go:build !windows
// +build !windows

// Package tty provides methods for working with TTY
package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var stdout, _ = os.Stdout.Stat()

// ////////////////////////////////////////////////////////////////////////////////// //

// IsTTY returns true if current output device is TTY
func IsTTY() bool {
	if stdout.Mode()&os.ModeCharDevice == 0 && !IsFakeTTY() {
		return false
	}

	return true
}

// IsTMUX returns true if we are currently working in tmux
func IsTMUX() (bool, error) {
	if os.Getenv("TMUX") != "" {
		return true, nil
	}

	return isTmuxAncestor()
}

// IsFakeTTY returns true is fake TTY is used
func IsFakeTTY() bool {
	return os.Getenv("FAKETTY") != ""
}

// IsSystemd returns true if process started by systemd
func IsSystemd() bool {
	return os.Getppid() == 1
}

// ////////////////////////////////////////////////////////////////////////////////// //
