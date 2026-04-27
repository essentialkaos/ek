//go:build linux || darwin

// Package tty provides methods for working with TTY
package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var stdout os.FileInfo
var isFakeTTY bool
var isSystemd bool
var isTMUX bool

var initOnce sync.Once

// ////////////////////////////////////////////////////////////////////////////////// //

// IsTTY returns true if the current output device is a TTY
func IsTTY() bool {
	checkSystem()

	if stdout != nil && stdout.Mode()&os.ModeCharDevice == 0 && !isFakeTTY {
		return false
	}

	return true
}

// IsTMUX returns true if the process is running inside a tmux session
func IsTMUX() (bool, error) {
	checkSystem()

	if isTMUX {
		return true, nil
	}

	return isTmuxAncestor()
}

// IsFakeTTY returns true if a fake TTY is in use via the FAKETTY environment variable
func IsFakeTTY() bool {
	checkSystem()
	return isFakeTTY
}

// IsSystemd returns true if the current process was started by systemd
func IsSystemd() bool {
	checkSystem()
	return isSystemd
}

// ////////////////////////////////////////////////////////////////////////////////// //

// checkSystem checks system for tmux/faketty
func checkSystem() {
	initOnce.Do(func() {
		stat, err := os.Stdout.Stat()

		if err == nil {
			stdout = stat
		}

		isTMUX = os.Getenv("TMUX") != ""
		isFakeTTY = os.Getenv("FAKETTY") != ""
		isSystemd = os.Getppid() == 1
	})
}
