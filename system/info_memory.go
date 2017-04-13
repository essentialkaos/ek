// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MemInfo contains info about system memory
type MemInfo struct {
	MemTotal   uint64 `json:"total"`       // Total usable ram (i.e. physical ram minus a few reserved bits and the kernel binary code)
	MemFree    uint64 `json:"free"`        // The sum of MemFree - (Buffers + Cached)
	MemUsed    uint64 `json:"used"`        // MemTotal - MemFree
	Buffers    uint64 `json:"buffers"`     // Relatively temporary storage for raw disk blocks shouldn't get tremendously large (20MB or so)
	Cached     uint64 `json:"cached"`      // In-memory cache for files read from the disk (the pagecache).  Doesn't include SwapCached
	Active     uint64 `json:"active"`      // Memory that has been used more recently and usually not reclaimed unless absolutely necessary
	Inactive   uint64 `json:"inactive"`    // Memory which has been less recently used.  It is more eligible to be reclaimed for other purposes
	SwapTotal  uint64 `json:"swap_total"`  // Total amount of swap space available
	SwapFree   uint64 `json:"swap_free"`   // Memory which has been evicted from RAM, and is temporarily on the disk still also is in the swapfile
	SwapUsed   uint64 `json:"swap_used"`   // SwapTotal - SwapFree
	SwapCached uint64 `json:"spaw_cached"` // Memory that once was swapped out, is swapped back in but
	Dirty      uint64 `json:"dirty"`       // Memory which is waiting to get written back to the disk
	Slab       uint64 `json:"slab"`        // In-kernel data structures cache
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with memory info in procfs
var procMemInfoFile = "/proc/meminfo"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMemInfo return memory info
func GetMemInfo() (*MemInfo, error) {
	content, err := readFileContent(procMemInfoFile)

	if err != nil {
		return nil, err
	}

	mem := &MemInfo{}

	for _, line := range content {
		lineSlice := splitLine(line)

		if len(lineSlice) < 2 {
			continue
		}

		switch lineSlice[0] {
		case "MemTotal:":
			mem.MemTotal, err = parseSize(lineSlice[1])
		case "MemFree:":
			mem.MemFree, err = parseSize(lineSlice[1])
		case "Buffers:":
			mem.Buffers, err = parseSize(lineSlice[1])
		case "Cached:":
			mem.Cached, err = parseSize(lineSlice[1])
		case "SwapCached:":
			mem.SwapCached, err = parseSize(lineSlice[1])
		case "Active:":
			mem.Active, err = parseSize(lineSlice[1])
		case "Inactive:":
			mem.Inactive, err = parseSize(lineSlice[1])
		case "SwapTotal:":
			mem.SwapTotal, err = parseSize(lineSlice[1])
		case "SwapFree:":
			mem.SwapFree, err = parseSize(lineSlice[1])
		case "Dirty:":
			mem.Dirty, err = parseSize(lineSlice[1])
		case "Slab:":
			mem.Slab, err = parseSize(lineSlice[1])
		default:
			continue
		}

		if err != nil {
			return nil, err
		}
	}

	mem.MemFree += mem.Cached + mem.Buffers
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	return mem, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseSize convert string with size in kb to uint64 bytes
func parseSize(s string) (uint64, error) {
	size, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, err
	}

	return size * 1024, nil
}
