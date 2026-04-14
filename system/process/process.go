// Package process provides methods for gathering information about active processes
package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// // Process state flags reported in /proc/[pid]/stat field 3
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

// ProcessInfo contains basic information about a process read from procfs
type ProcessInfo struct {
	Command  string         // Full command line with arguments
	User     string         // Name of the user owning the process
	PID      int            // Process ID
	Parent   int            // Parent process ID
	Children []*ProcessInfo // Child processes and threads
	IsThread bool           // True if this entry represents a kernel thread
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcInfo contains a partial view of /proc/[pid]/stat fields
type ProcInfo struct {
	PID        int    `json:"pid"`         // Process ID
	Comm       string `json:"comm"`        // Executable filename (without path, without parentheses)
	State      string `json:"state"`       // Single-character process state
	PPID       int    `json:"ppid"`        // Parent process ID
	Session    int    `json:"session"`     // Session ID
	TTYNR      int    `json:"tty_nr"`      // Controlling terminal device number
	TPGid      int    `json:"tpgid"`       // Foreground process group ID of the controlling terminal
	UTime      uint64 `json:"utime"`       // User-mode CPU time, in clock ticks
	STime      uint64 `json:"stime"`       // Kernel-mode CPU time, in clock ticks
	CUTime     uint64 `json:"cutime"`      // Waited-for children user-mode CPU time, in clock ticks
	CSTime     uint64 `json:"cstime"`      // Waited-for children kernel-mode CPU time, in clock ticks
	Priority   int    `json:"priority"`    // Kernel scheduling priority value
	Nice       int    `json:"nice"`        // Nice value in the range [-20, 19]
	NumThreads int    `json:"num_threads"` // Number of threads in the process
}

// ProcSample is a single CPU-time snapshot used to compute usage between two samples
type ProcSample uint

// ////////////////////////////////////////////////////////////////////////////////// //

// MemInfo contains virtual and physical memory usage counters from /proc/[pid]/status.
// All size values are in bytes.
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

// MountInfo contains a single parsed entry from /proc/[pid]/mountinfo. See
// https://www.kernel.org/doc/Documentation/filesystems/proc.txt for field semantics.
type MountInfo struct {
	ID             uint16   `json:"mount_id"`        // Unique mount identifier (may be reused after umount)
	ParentID       uint16   `json:"parent_id"`       // ID of the parent mount (self for the tree root)
	StDevMajor     uint16   `json:"stdev_major"`     // Major component of st_dev for this filesystem
	StDevMinor     uint16   `json:"stdev_minor"`     // Minor component of st_dev for this filesystem
	Root           string   `json:"root"`            // Pathname of the directory in the filesystem that forms the root of this mount
	MountPoint     string   `json:"mount_point"`     // Pathname of the mount point relative to the process root
	MountOptions   []string `json:"mount_options"`   // Per-mount options (e.g. "rw", "relatime")
	OptionalFields []string `json:"optional_fields"` // Zero or more tagged fields of the form "tag[:value]"
	FSType         string   `json:"fs_type"`         // Filesystem type, optionally with a subtype ("type[.subtype]")
	MountSource    string   `json:"mount_source"`    // Filesystem-specific source information, or "none"
	SuperOptions   []string `json:"super_options"`   // Per-superblock options (e.g. "rw", "errors=remount-ro")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ToSample converts ProcInfo to a ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	return ProcSample(pi.UTime + pi.STime + pi.CUTime + pi.CSTime)
}

// ////////////////////////////////////////////////////////////////////////////////// //
