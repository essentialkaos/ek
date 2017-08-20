// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"

	"pkg.re/essentialkaos/ek.v9/errutil"
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
	fd, err := os.OpenFile(procMemInfoFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	mem := &MemInfo{}
	errs := errutil.NewErrors()

	for s.Scan() {
		text := s.Text()

		switch readField(text, 0) {
		case "MemTotal:":
			mem.MemTotal = parseSize(readField(text, 1), errs)
		case "MemFree:":
			mem.MemFree = parseSize(readField(text, 1), errs)
		case "Buffers:":
			mem.Buffers = parseSize(readField(text, 1), errs)
		case "Cached:":
			mem.Cached = parseSize(readField(text, 1), errs)
		case "SwapCached:":
			mem.SwapCached = parseSize(readField(text, 1), errs)
		case "Active:":
			mem.Active = parseSize(readField(text, 1), errs)
		case "Inactive:":
			mem.Inactive = parseSize(readField(text, 1), errs)
		case "SwapTotal:":
			mem.SwapTotal = parseSize(readField(text, 1), errs)
		case "SwapFree:":
			mem.SwapFree = parseSize(readField(text, 1), errs)
		case "Dirty:":
			mem.Dirty = parseSize(readField(text, 1), errs)
		case "Slab:":
			mem.Slab = parseSize(readField(text, 1), errs)
		}

		if errs.HasErrors() {
			return nil, errs.Last()
		}
	}

	if mem.MemTotal == 0 {
		return nil, errors.New("Can't parse file " + procMemInfoFile)
	}

	mem.MemFree += mem.Cached + mem.Buffers
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	return mem, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
