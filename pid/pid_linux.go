package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"syscall"

	"pkg.re/essentialkaos/ek.v11/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// procfsDir is path to procfs directory
var procfsDir = "/proc"

// ////////////////////////////////////////////////////////////////////////////////// //

// IsWorks returns true if process with PID from PID file is works
func IsWorks(name string) bool {
	pid := Get(name)

	if pid == -1 {
		return false
	}

	if !fsutil.IsExist(fmt.Sprintf("%s/%d", procfsDir, pid)) {
		return false
	}

	_, _, initCDate, _ := fsutil.GetTimestamps(fmt.Sprintf("%s/%d", procfsDir, 1))
	_, _, procCDate, _ := fsutil.GetTimestamps(fmt.Sprintf("%s/%d", procfsDir, pid))

	if initCDate != -1 && procCDate != -1 && initCDate > procCDate {
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
