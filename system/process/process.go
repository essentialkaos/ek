// Package process provides methods for gathering information about active processes
package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Process state flags
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

const (
	PRIO_CLASS_NONE        = 0
	PRIO_CLASS_REAL_TIME   = 1
	PRIO_CLASS_BEST_EFFORT = 2
	PRIO_CLASS_IDLE        = 3
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcessInfo contains basic info about process
type ProcessInfo struct {
	Command  string         // Full command
	User     string         // Username
	PID      int            // PID
	Parent   int            // Parent process PID
	Children []*ProcessInfo // Slice with child processes
	IsThread bool           // True if process is thread
}

// ////////////////////////////////////////////////////////////////////////////////// //

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

// ProcSample contains value for usage calculation
type ProcSample uint

// ////////////////////////////////////////////////////////////////////////////////// //

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

// ////////////////////////////////////////////////////////////////////////////////// //

// MountInfo contains information about mount
// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
type MountInfo struct {
	ID             uint16   `json:"mount_id"`        // Unique identifier of the mount (may be reused after umount)
	ParentID       uint16   `json:"parent_id"`       // ID of parent (or of self for the top of the mount tree)
	StDevMajor     uint16   `json:"stdev_major"`     // Major value of st_dev for files on filesystem
	StDevMinor     uint16   `json:"stdev_minor"`     // Minor value of st_dev for files on filesystem
	Root           string   `json:"root"`            // Root of the mount within the filesystem
	MountPoint     string   `json:"mount_point"`     // Mount point relative to the process's root
	MountOptions   []string `json:"mount_options"`   // Per mount options
	OptionalFields []string `json:"optional_fields"` // Zero or more fields of the form "tag[:value]"
	FSType         string   `json:"fs_type"`         // Name of filesystem of the form "type[.subtype]"
	MountSource    string   `json:"mount_source"`    // Filesystem specific information or "none"
	SuperOptions   []string `json:"super_options"`   // Per super block options
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ToSample converts ProcInfo to ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	return ProcSample(pi.UTime + pi.STime + pi.CUTime + pi.CSTime)
}

// ////////////////////////////////////////////////////////////////////////////////// //
