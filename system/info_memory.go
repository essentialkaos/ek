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
