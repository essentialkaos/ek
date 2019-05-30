package pid

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// IsWorks returns true if process with PID from PID file is works
func IsWorks(name string) bool {
	pid := Get(name)

	if pid == -1 {
		return false
	}

	return IsProcessWorks(pid)
}

// IsProcessWorks returns true if process with given PID is works
func IsProcessWorks(pid int) bool {
	pr, err := os.FindProcess(pid)

	if pr == nil || err != nil {
		return false
	}

	return true
}
