package pid

import (
	"os"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
	// On Unix systems, FindProcess always succeeds and returns a Process
	// for the given pid, regardless of whether the process exists.
	pr, _ := os.FindProcess(pid)
	return pr.Signal(syscall.Signal(0)) == nil
}
