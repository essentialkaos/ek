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

// MountInfo contains information about mounts
// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
type MountInfo struct {
	MountID        uint16   `json:"mount_id"`        // unique identifier of the mount (may be reused after umount)
	ParentID       uint16   `json:"parent_id"`       // ID of parent (or of self for the top of the mount tree)
	StDevMajor     uint16   `json:"stdev_major"`     // major value of st_dev for files on filesystem
	StDevMinor     uint16   `json:"stdev_minor"`     // minor value of st_dev for files on filesystem
	Root           string   `json:"root"`            // root of the mount within the filesystem
	MountPoint     string   `json:"mount_point"`     // mount point relative to the process's root
	MountOptions   []string `json:"mount_options"`   // per mount options
	OptionalFields []string `json:"optional_fields"` // zero or more fields of the form "tag[:value]"
	FSType         string   `json:"fs_type"`         // name of filesystem of the form "type[.subtype]"
	MountSource    string   `json:"mount_source"`    // filesystem specific information or "none"
	SuperOptions   []string `json:"super_options"`   // per super block options
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

// GetMountInfo returns info about process mounts
func GetMountInfo(pid int) ([]*MountInfo, error) {
	return nil, nil
}

// GetCPUPriority returns process CPU scheduling priority (PR, NI, error)
func GetCPUPriority(pid int) (int, int, error) {
	return 0, 0, nil
}

// SetCPUPriority sets process CPU scheduling priority
func SetCPUPriority(pid, niceness int) error {
	return nil
}

// GetIOPriority returns process IO scheduling priority (class, classdata, error)
func GetIOPriority(pid int) (int, int, error) {
	return 0, 0, nil
}

// SetIOPriority sets process IO scheduling priority
func SetIOPriority(pid, class, classdata int) error {
	return nil
}
