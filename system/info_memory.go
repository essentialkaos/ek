// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
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

// defaultMemProps map with names of required stats
var defaultMemProps map[string]bool

// Path to file with memory info in procfs
var procMemInfoFile = "/proc/meminfo"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMemInfo return memory info
func GetMemInfo() (*MemInfo, error) {
	if defaultMemProps == nil {
		defaultMemProps = map[string]bool{
			"MemTotal":   true,
			"MemFree":    true,
			"Buffers":    true,
			"Cached":     true,
			"SwapCached": true,
			"Active":     true,
			"Inactive":   true,
			"SwapTotal":  true,
			"SwapFree":   true,
			"Dirty":      true,
			"Slab":       true,
		}
	}

	content, err := readFileContent(procMemInfoFile)

	if err != nil {
		return nil, err
	}

	mem := &MemInfo{}

	for _, line := range content {
		if line == "" {
			continue
		}

		lineSlice := strings.Split(line, ":")

		if len(lineSlice) != 2 {
			return nil, errors.New("Can't parse file " + procMemInfoFile)
		}

		if !defaultMemProps[lineSlice[0]] {
			continue
		}

		strValue := strings.TrimRight(lineSlice[1], " kB")
		strValue = strings.Replace(strValue, " ", "", -1)
		uintValue, err := strconv.ParseUint(strValue, 10, 64)

		if err != nil {
			return nil, err
		}

		switch lineSlice[0] {
		case "MemTotal":
			mem.MemTotal = uintValue * 1024
		case "MemFree":
			mem.MemFree = uintValue * 1024
		case "Buffers":
			mem.Buffers = uintValue * 1024
		case "Cached":
			mem.Cached = uintValue * 1024
		case "SwapCached":
			mem.SwapCached = uintValue * 1024
		case "Active":
			mem.Active = uintValue * 1024
		case "Inactive":
			mem.Inactive = uintValue * 1024
		case "SwapTotal":
			mem.SwapTotal = uintValue * 1024
		case "SwapFree":
			mem.SwapFree = uintValue * 1024
		case "Dirty":
			mem.Dirty = uintValue * 1024
		case "Slab":
			mem.Slab = uintValue * 1024
		}
	}

	mem.MemFree += mem.Cached + mem.Buffers
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	return mem, nil
}
