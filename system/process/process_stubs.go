//go:build !linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ToSample converts ProcInfo to ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetInfo returns process info from procfs
func GetInfo(pid int) (*ProcInfo, error) {
	panic("UNSUPPORTED")
}

// ❗ GetSample returns ProcSample for CPU usage calculation
func GetSample(pid int) (ProcSample, error) {
	panic("UNSUPPORTED")
}

// ❗ CalculateCPUUsage calculate CPU usage
func CalculateCPUUsage(s1, s2 ProcSample, duration time.Duration) float64 {
	panic("UNSUPPORTED")
}

// ❗ GetMemInfo returns info about process memory usage
func GetMemInfo(pid int) (*MemInfo, error) {
	panic("UNSUPPORTED")
}

// ❗ GetMountInfo returns info about process mounts
func GetMountInfo(pid int) ([]*MountInfo, error) {
	panic("UNSUPPORTED")
}

// ❗ GetCPUPriority returns process CPU scheduling priority (PR, NI, error)
func GetCPUPriority(pid int) (int, int, error) {
	panic("UNSUPPORTED")
}

// ❗ SetCPUPriority sets process CPU scheduling priority
func SetCPUPriority(pid, niceness int) error {
	panic("UNSUPPORTED")
}

// ❗ GetIOPriority returns process IO scheduling priority (class, classdata, error)
func GetIOPriority(pid int) (int, int, error) {
	panic("UNSUPPORTED")
}

// ❗ SetIOPriority sets process IO scheduling priority
func SetIOPriority(pid, class, classdata int) error {
	panic("UNSUPPORTED")
}
