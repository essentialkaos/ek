package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsWorks return if process pid is not -1
func IsWorks(name string) bool {
	pid := Get(name)

	if pid == -1 {
		return false
	}

	err := exec.Command("/usr/bin/procstat", strconv.Itoa(pid)).Run()

	return err == nil
}
