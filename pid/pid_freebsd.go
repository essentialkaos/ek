package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strconv"
)

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
	return exec.Command("/usr/bin/procstat", strconv.Itoa(pid)).Run() == nil
}
