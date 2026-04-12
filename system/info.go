// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
	"strings"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// User contains information about a system user account
type User struct {
	UID      int      `json:"uid"`
	GID      int      `json:"gid"`
	Name     string   `json:"name"`
	Groups   []*Group `json:"groups"`
	Comment  string   `json:"comment"`
	Shell    string   `json:"shell"`
	HomeDir  string   `json:"home_dir"`
	RealUID  int      `json:"real_uid"`  // UID of the original user before sudo elevation
	RealGID  int      `json:"real_gid"`  // GID of the original user before sudo elevation
	RealName string   `json:"real_name"` // Name of the original user before sudo elevation
}

// Group contains information about a system group
type Group struct {
	Name string `json:"name"`
	GID  int    `json:"gid"`
}

// SessionInfo contains information about an active login session
type SessionInfo struct {
	Username         string    `json:"username"`
	Host             string    `json:"host"`
	LoginTime        time.Time `json:"login_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// LoadAvg contains average system load over 1, 5, and 15 minute intervals
type LoadAvg struct {
	Min1  float64 `json:"min1"`  // LA in last 1 minute
	Min5  float64 `json:"min5"`  // LA in last 5 minutes
	Min15 float64 `json:"min15"` // LA in last 15 minutes
	RProc int     `json:"rproc"` // Number of currently runnable kernel scheduling entities
	TProc int     `json:"tproc"` // Number of kernel scheduling entities that currently exist on the system
}

// MemUsage contains information about physical and swap memory usage
type MemUsage struct {
	MemTotal     uint64 `json:"total"`                  // Total usable ram (i.e. physical ram minus a few reserved bits and the kernel binary code)
	MemFree      uint64 `json:"free"`                   // The sum of MemFree - (Buffers + Cached)
	MemUsed      uint64 `json:"used"`                   // MemTotal - MemFree
	MemAvailable uint64 `json:"available"`              // Available memory
	SwapTotal    uint64 `json:"swap_total"`             // Total amount of swap space available
	SwapFree     uint64 `json:"swap_free"`              // Memory which has been evicted from RAM, and is temporarily on the disk still also is in the swapfile
	SwapUsed     uint64 `json:"swap_used"`              // SwapTotal - SwapFree
	Active       uint64 `json:"active"`                 // Memory that has been used more recently and usually not reclaimed unless absolutely necessary
	Inactive     uint64 `json:"inactive"`               // Memory which has been less recently used.  It is more eligible to be reclaimed for other purposes
	Buffers      uint64 `json:"buffers,omitempty"`      // Relatively temporary storage for raw disk blocks shouldn't get tremendously large (20MB or so)
	Cached       uint64 `json:"cached,omitempty"`       // In-memory cache for files read from the disk (the pagecache).  Doesn't include SwapCached
	SwapCached   uint64 `json:"swap_cached,omitempty"`  // Memory that once was swapped out, is swapped back in but
	Dirty        uint64 `json:"dirty,omitempty"`        // Memory which is waiting to get written back to the disk
	Shmem        uint64 `json:"shmem,omitempty"`        // Total used shared memory
	Slab         uint64 `json:"slab,omitempty"`         // In-kernel data structures cache
	SReclaimable uint64 `json:"sreclaimable,omitempty"` // Reclaimable portion of Slab (e.g. caches)
}

// CPUUsage contains percentage-based CPU usage breakdown for a measured interval
type CPUUsage struct {
	User    float64 `json:"user"`    // Time spent running user-space processes
	System  float64 `json:"system"`  // Time spent running kernel-space processes
	Nice    float64 `json:"nice"`    // Time spent running niced user-space processes
	Idle    float64 `json:"idle"`    // Time spent idle
	Wait    float64 `json:"wait"`    // Time spent waiting for I/O to complete
	Average float64 `json:"average"` // Overall average CPU usage across all states
	Count   int     `json:"count"`   // Number of logical CPU cores
}

// CPUInfo contains static information about a physical CPU package
type CPUInfo struct {
	Vendor    string    `json:"vendor"`          // Processor vendor identifier (e.g. GenuineIntel)
	Model     string    `json:"model"`           // Full model name of the processor
	Cores     int       `json:"cores"`           // Number of physical cores
	Siblings  int       `json:"siblings"`        // Total logical CPUs on the same physical package
	CacheSize uint64    `json:"cache_size"`      // L2 cache size in bytes
	Speed     []float64 `json:"speed,omitempty"` // Per-core speed in MHz
}

// CPUStats contains raw cumulative CPU time counters read from /proc/stat
type CPUStats struct {
	User   uint64 `json:"user"`   // Time in user mode
	Nice   uint64 `json:"nice"`   // Time in user mode with low priority
	System uint64 `json:"system"` // Time in system (kernel) mode
	Idle   uint64 `json:"idle"`   // Time spent idle
	Wait   uint64 `json:"wait"`   // Time waiting for I/O
	IRQ    uint64 `json:"irq"`    // Time servicing hardware interrupts
	SRQ    uint64 `json:"srq"`    // Time servicing software interrupts
	Steal  uint64 `json:"steal"`  // Time stolen by a hypervisor for other guests
	Total  uint64 `json:"total"`  // Sum of all CPU time fields
	Count  int    `json:"count"`  // Number of logical CPU cores seen in /proc/stat
}

// CPUCount contains the number of CPUs in each availability state
type CPUCount struct {
	Possible uint32 `json:"possible"` // CPUs that can ever be online on this system
	Present  uint32 `json:"present"`  // CPUs currently present (plugged in)
	Online   uint32 `json:"online"`   // CPUs currently online and schedulable
	Offline  uint32 `json:"offline"`  // CPUs present but currently offline
}

// FSUsage contains info about FS usage
type FSUsage struct {
	Type    string   `json:"type"`    // Filesystem type (ext4, xfs, tmpfs, etc.)
	Device  string   `json:"device"`  // Block device or remote spec
	Used    uint64   `json:"used"`    // Used space in bytes
	Free    uint64   `json:"free"`    // Available space in bytes
	Total   uint64   `json:"total"`   // Total capacity in bytes
	IOStats *IOStats `json:"iostats"` // I/O statistics for the backing device
}

// IOStats contains raw I/O counters for a block device from /proc/diskstats
type IOStats struct {
	ReadComplete  uint64 `json:"read_complete"`  // Total reads completed successfully
	ReadMerged    uint64 `json:"read_merged"`    // Adjacent reads merged by the I/O scheduler
	ReadSectors   uint64 `json:"read_sectors"`   // Total sectors read
	ReadMs        uint64 `json:"read_ms"`        // Total time spent reading in milliseconds
	WriteComplete uint64 `json:"write_complete"` // Total writes completed successfully
	WriteMerged   uint64 `json:"write_merged"`   // Adjacent writes merged by the I/O scheduler
	WriteSectors  uint64 `json:"write_sectors"`  // Total sectors written
	WriteMs       uint64 `json:"write_ms"`       // Total time spent writing in milliseconds
	IOPending     uint64 `json:"io_pending"`     // Number of I/Os currently in flight
	IOMs          uint64 `json:"io_ms"`          // Total time spent on I/Os in milliseconds
	IOQueueMs     uint64 `json:"io_queue_ms"`    // Weighted time spent on I/Os (reflects queue depth)
}

// InterfaceStats contains cumulative traffic counters for a network interface
type InterfaceStats struct {
	ReceivedBytes      uint64 `json:"received_bytes"`
	ReceivedPackets    uint64 `json:"received_packets"`
	TransmittedBytes   uint64 `json:"transmitted_bytes"`
	TransmittedPackets uint64 `json:"transmitted_packets"`
}

// SystemInfo contains general information about the host system
type SystemInfo struct {
	Hostname        string `json:"hostname"`         // System hostname
	ID              string `json:"id"`               // Unique machine ID
	OS              string `json:"os"`               // OS name
	Kernel          string `json:"kernel"`           // Kernel version string
	Arch            string `json:"arch"`             // Raw architecture identifier (x86_64, aarch64, etc.)
	ArchName        string `json:"arch_name"`        // Normalised architecture name (amd64, 386, etc.)
	ContainerEngine string `json:"container_engine"` // Detected container engine
	ArchBits        int    `json:"arch_bits"`        // Pointer width of the architecture (32 or 64)
}

// OSInfo contains information parsed from /etc/os-release
type OSInfo struct {
	Name                  string `json:"name"`
	PrettyName            string `json:"pretty_name"`
	Version               string `json:"version"`
	Build                 string `json:"build"`
	VersionID             string `json:"version_id"`
	VersionCodename       string `json:"version_codename"`
	ID                    string `json:"id"`
	IDLike                string `json:"id_like"`
	PlatformID            string `json:"platform_id"`
	Variant               string `json:"variant"`
	VariantID             string `json:"variant_id"`
	CPEName               string `json:"cpe_name"`
	HomeURL               string `json:"home_url"`
	BugReportURL          string `json:"bugreport_url"`
	DocumentationURL      string `json:"documentation_url"`
	Logo                  string `json:"logo"`
	ANSIColor             string `json:"ansi_color"`
	SupportURL            string `json:"support_url"`
	SupportProduct        string `json:"support_product"`
	SupportProductVersion string `json:"support_product_version"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ColoredPrettyName returns the pretty OS name wrapped with its ANSI color code
func (i *OSInfo) ColoredPrettyName() string {
	if !isValidANSIColor(i.ANSIColor) {
		return i.PrettyName
	}

	return "\033[" + i.ANSIColor + "m" + i.PrettyName + "\033[0m"
}

// ColoredName returns the OS name wrapped with its ANSI color code
func (i *OSInfo) ColoredName() string {
	if !isValidANSIColor(i.ANSIColor) {
		return i.Name
	}

	return "\033[" + i.ANSIColor + "m" + i.Name + "\033[0m"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// MemUsedPerc returns used physical memory as a percentage of total RAM
func (m *MemUsage) MemUsedPerc() float64 {
	if m == nil {
		return 0
	}

	return (float64(m.MemUsed) / float64(m.MemTotal)) * 100.0
}

// SwapUsedPerc returns used swap space as a percentage of total swap
func (m *MemUsage) SwapUsedPerc() float64 {
	if m == nil || m.SwapTotal == 0 {
		return 0
	}

	return (float64(m.SwapUsed) / float64(m.SwapTotal)) * 100.0
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseSize parse size in kB
func parseSize(v string) (uint64, error) {
	size, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		return 0, err
	}

	return size * 1024, nil
}

// getArchName returns name for given arch
func getArchName(arch string) string {
	switch arch {
	case "i386":
		return "386"
	case "i586":
		return "586"
	case "i686":
		return "686"
	case "x86_64":
		return "amd64"
	}

	return arch
}

// isValidANSIColor validates ansi color code
func isValidANSIColor(color string) bool {
	return color != "" && strings.Trim(color, "0123456789;") == ""
}
