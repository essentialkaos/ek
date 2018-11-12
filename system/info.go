// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// OS names
const (
	LINUX_ARCH   = "Arch"
	LINUX_CENTOS = "CentOS"
	LINUX_DEBIAN = "Debian"
	LINUX_FEDORA = "Dedora"
	LINUX_GENTOO = "Gentoo"
	LINUX_RHEL   = "RHEL"
	LINUX_SUSE   = "SuSe"
	LINUX_UBUNTU = "Ubuntu"
	DARWIN_OSX   = "OSX"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// LoadAvg contains information about average system load
type LoadAvg struct {
	Min1  float64 `json:"min1"`  // LA in last 1 minute
	Min5  float64 `json:"min5"`  // LA in last 5 minutes
	Min15 float64 `json:"min15"` // LA in last 15 minutes
	RProc int     `json:"rproc"` // Number of currently runnable kernel scheduling entities
	TProc int     `json:"tproc"` // Number of kernel scheduling entities that currently exist on the system
}

// MemUsage contains info about system memory usage
type MemUsage struct {
	MemTotal     uint64 `json:"total"`        // Total usable ram (i.e. physical ram minus a few reserved bits and the kernel binary code)
	MemFree      uint64 `json:"free"`         // The sum of MemFree - (Buffers + Cached)
	MemUsed      uint64 `json:"used"`         // MemTotal - MemFree
	Buffers      uint64 `json:"buffers"`      // Relatively temporary storage for raw disk blocks shouldn't get tremendously large (20MB or so)
	Cached       uint64 `json:"cached"`       // In-memory cache for files read from the disk (the pagecache).  Doesn't include SwapCached
	Active       uint64 `json:"active"`       // Memory that has been used more recently and usually not reclaimed unless absolutely necessary
	Inactive     uint64 `json:"inactive"`     // Memory which has been less recently used.  It is more eligible to be reclaimed for other purposes
	SwapTotal    uint64 `json:"swap_total"`   // Total amount of swap space available
	SwapFree     uint64 `json:"swap_free"`    // Memory which has been evicted from RAM, and is temporarily on the disk still also is in the swapfile
	SwapUsed     uint64 `json:"swap_used"`    // SwapTotal - SwapFree
	SwapCached   uint64 `json:"swap_cached"`  // Memory that once was swapped out, is swapped back in but
	Dirty        uint64 `json:"dirty"`        // Memory which is waiting to get written back to the disk
	Shmem        uint64 `json:"shmem"`        // Total used shared memory
	Slab         uint64 `json:"slab"`         // In-kernel data structures cache
	SReclaimable uint64 `json:"sreclaimable"` // The part of the Slab that might be reclaimed (such as caches)
}

// CPUUsage contains info about CPU usage
type CPUUsage struct {
	User    float64 `json:"user"`    // Normal processes executing in user mode
	System  float64 `json:"system"`  // Processes executing in kernel mode
	Nice    float64 `json:"nice"`    // Niced processes executing in user mode
	Idle    float64 `json:"idle"`    // Twiddling thumbs
	Wait    float64 `json:"wait"`    // Waiting for I/O to complete
	Average float64 `json:"average"` // Average CPU usage
	Count   int     `json:"count"`   // Number of CPU cores
}

// CPUInfo contains info about CPU
type CPUInfo struct {
	Vendor    string    `json:"vendor"`     // Processor vandor name
	Model     string    `json:"model"`      // Common name of the processor
	Cores     int       `json:"cores"`      // Number of cores
	Siblings  int       `json:"siblings"`   // Total number of sibling CPUs on the same physical CPU
	CacheSize uint64    `json:"cache_size"` // Amount of level 2 memory cache available to the processor
	Speed     []float64 `json:"speed"`      // Speed in megahertz for the processor
}

// CPUStats contains basic CPU stats
type CPUStats struct {
	User   uint64 `json:"user"`
	Nice   uint64 `json:"nice"`
	System uint64 `json:"system"`
	Idle   uint64 `json:"idle"`
	Wait   uint64 `json:"wait"`
	IRQ    uint64 `json:"irq"`
	SRQ    uint64 `json:"srq"`
	Steal  uint64 `json:"steal"`
	Total  uint64 `json:"total"`
	Count  int    `json:"count"`
}

// FSUsage contains info about FS usage
type FSUsage struct {
	Type    string   `json:"type"`    // FS type (ext4/ntfs/etc...)
	Device  string   `json:"device"`  // Device spec
	Used    uint64   `json:"used"`    // Used space
	Free    uint64   `json:"free"`    // Free space
	Total   uint64   `json:"total"`   // Total space
	IOStats *IOStats `json:"iostats"` // IO statistics
}

// IOStats contains information about I/O
type IOStats struct {
	ReadComplete  uint64 `json:"read_complete"`  // Reads completed successfully
	ReadMerged    uint64 `json:"read_merged"`    // Reads merged
	ReadSectors   uint64 `json:"read_sectors"`   // Sectors read
	ReadMs        uint64 `json:"read_ms"`        // Time spent reading (ms)
	WriteComplete uint64 `json:"write_complete"` // Writes completed
	WriteMerged   uint64 `json:"write_merged"`   // Writes merged
	WriteSectors  uint64 `json:"write_sectors"`  // Sectors written
	WriteMs       uint64 `json:"write_ms"`       // Time spent writing (ms)
	IOPending     uint64 `json:"io_pending"`     // I/Os currently in progress
	IOMs          uint64 `json:"io_ms"`          // Time spent doing I/Os (ms)
	IOQueueMs     uint64 `json:"io_queue_ms"`    // Weighted time spent doing I/Os (ms)
}

// InterfaceStats contains stats about network interfaces usage
type InterfaceStats struct {
	ReceivedBytes      uint64 `json:"received_bytes"`
	ReceivedPackets    uint64 `json:"received_packets"`
	TransmittedBytes   uint64 `json:"transmitted_bytes"`
	TransmittedPackets uint64 `json:"transmitted_packets"`
}

// SystemInfo contains info about a system (hostname, OS, arch...)
type SystemInfo struct {
	Hostname     string `json:"hostname"`     // Hostname
	OS           string `json:"os"`           // OS name
	Distribution string `json:"distribution"` // OS distribution
	Version      string `json:"version"`      // OS version
	Kernel       string `json:"kernel"`       // Kernel version
	Arch         string `json:"arch"`         // System architecture (i386/i686/x86_64/etc...)
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
