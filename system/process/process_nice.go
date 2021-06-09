// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SetCPUPriority sets CPU scheduling priority
func SetCPUPriority(pid, niceness int) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, pid, niceness)
}

// GetCPUPriority returns CPU scheduling priority (PR, NI, error)
func GetCPUPriority(pid int) (int, int, error) {
	pr, err := syscall.Getpriority(syscall.PRIO_PROCESS, pid)

	if err != nil {
		return 0, 0, err
	}

	ni := 20 - pr

	return 20 + ni, ni, nil
}
