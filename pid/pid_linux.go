package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"pkg.re/essentialkaos/ek.v7/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsWorks return if process with pid from pid file is works
func IsWorks(name string) bool {
	pid := Get(name)

	if pid == -1 {
		return false
	}

	return fsutil.IsExist(fmt.Sprintf("/proc/%d", pid))
}
