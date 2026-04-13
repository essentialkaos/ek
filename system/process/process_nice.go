//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"

	"github.com/essentialkaos/ek/v14/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUPriority returns the CPU scheduling priority of the given process.
// It returns the PR (priority) and NI (nice) values as reported by the kernel.
func GetCPUPriority(pid int) (int, int, error) {
	pr, err := syscall.Getpriority(syscall.PRIO_PROCESS, pid)

	if err != nil {
		return 0, 0, err
	}

	ni := 20 - pr

	return 20 + ni, ni, nil
}

// SetCPUPriority sets the CPU scheduling nice value for the given process
func SetCPUPriority(pid, niceness int) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, pid, niceness)
}

// GetIOPriority returns the I/O scheduling class and class-data for the given process
func GetIOPriority(pid int) (int, int, error) {
	v, _, errNo := syscall.Syscall(
		syscall.SYS_IOPRIO_GET, uintptr(1), uintptr(pid), uintptr(0),
	)

	if errNo != 0 {
		return 0, 0, errNo
	}

	prio := int(v)
	class := prio >> 13
	classdata := prio & ((1 << 13) - 1)

	return class, classdata, nil
}

// SetIOPriority sets the I/O scheduling class and priority for the given process.
// class is clamped to [0, 3] and classdata to [0, 7].
func SetIOPriority(pid, class, classdata int) error {
	class = mathutil.Between(class, 0, 3)
	classdata = mathutil.Between(classdata, 0, 7)
	prio := (class << 13) | classdata

	_, _, errNo := syscall.Syscall(
		syscall.SYS_IOPRIO_SET, uintptr(1), uintptr(pid), uintptr(prio),
	)

	if errNo != 0 {
		return errNo
	}

	return nil
}
