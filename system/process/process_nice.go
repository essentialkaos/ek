//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"

	"github.com/essentialkaos/ek/v13/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	PRIO_CLASS_NONE        = 0
	PRIO_CLASS_REAL_TIME   = 1
	PRIO_CLASS_BEST_EFFORT = 2
	PRIO_CLASS_IDLE        = 3
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUPriority returns process CPU scheduling priority (PR, NI, error)
func GetCPUPriority(pid int) (int, int, error) {
	pr, err := syscall.Getpriority(syscall.PRIO_PROCESS, pid)

	if err != nil {
		return 0, 0, err
	}

	ni := 20 - pr

	return 20 + ni, ni, nil
}

// SetCPUPriority sets process CPU scheduling priority
func SetCPUPriority(pid, niceness int) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, pid, niceness)
}

// GetIOPriority returns process IO scheduling priority (class, classdata, error)
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

// SetIOPriority sets process IO scheduling priority
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
