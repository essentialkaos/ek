// +build windows

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	STATE_RUNNING   = "R"
	STATE_SLEEPING  = "S"
	STATE_DISK_WAIT = "D"
	STATE_ZOMBIE    = "Z"
	STATE_STOPPED   = "T"
	STATE_DEAD      = "X"
	STATE_WAKEKILL  = "K"
	STATE_WAKING    = "W"
	STATE_PARKED    = "P"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcSample contains value for usage calculation
type ProcSample uint

// ProcInfo contains partial info from /proc/[PID]/stat
type ProcInfo struct {
	PID        int    `json:"pid"`         // The process ID
	Comm       string `json:"comm"`        // The filename of the executable, in parentheses
	State      string `json:"state"`       // Process state
	PPID       int    `json:"ppid"`        // The PID of the parent of this process
	Session    int    `json:"session"`     // The session ID of the process
	TTYNR      int    `json:"tty_nr"`      // The controlling terminal of the process
	TPGid      int    `json:"tpgid"`       // The ID of the foreground process group of the controlling terminal of the process
	UTime      uint64 `json:"utime"`       // Amount of time that this process has been scheduled in user mode, measured in clock ticks
	STime      uint64 `json:"stime"`       // Amount of time that this process has been scheduled in kernel mode, measured in clock ticks
	CUTime     uint64 `json:"cutime"`      // Amount of time that this process's waited-for children have been scheduled in user mode, measured in clock ticks
	CSTime     uint64 `json:"cstime"`      // Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks
	Priority   int    `json:"priority"`    // Priority
	Nice       int    `json:"nice"`        // The nice value
	NumThreads int    `json:"num_threads"` // Number of threads in this process
}

// MemInfo contains process memory usage stats
type MemInfo struct {
	VmPeak uint64 `json:"peak"` // Peak virtual memory size
	VmSize uint64 `json:"size"` // Virtual memory size
	VmLck  uint64 `json:"lck"`  // Locked memory size
	VmPin  uint64 `json:"pin"`  // Pinned memory size (since Linux 3.2)
	VmHWM  uint64 `json:"hwm"`  // Peak resident set size ("high water mark")
	VmRSS  uint64 `json:"rss"`  // Resident set size
	VmData uint64 `json:"data"` // Size of data
	VmStk  uint64 `json:"stk"`  // Size of stack
	VmExe  uint64 `json:"exe"`  // Size of text segments
	VmLib  uint64 `json:"lib"`  // Shared library code size
	VmPTE  uint64 `json:"pte"`  // Page table entries size (since Linux 2.6.10)
	VmSwap uint64 `json:"swap"` // Swap size
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ToSample converts ProcInfo to ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	return 0
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInfo returns process info from procfs
func GetInfo(pid int) (*ProcInfo, error) {
	return nil, nil
}

// GetSample returns ProcSample for CPU usage calculation
func GetSample(pid int) (ProcSample, error) {
	return 0, nil
}

// CalculateCPUUsage calculate CPU usage
func CalculateCPUUsage(s1, s2 ProcSample, duration time.Duration) float64 {
	return 0.0
}

// GetMemInfo returns info about process memory usage
func GetMemInfo(pid int) (*MemInfo, error) {
	return nil, nil
}
