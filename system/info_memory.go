// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"

	"pkg.re/essentialkaos/ek.v9/errutil"
	"pkg.re/essentialkaos/ek.v9/strutil"
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

		switch strutil.ReadField(text, 0, true) {
		case "MemTotal:":
			mem.MemTotal = parseSize(strutil.ReadField(text, 1, true), errs)
		case "MemFree:":
			mem.MemFree = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Buffers:":
			mem.Buffers = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Cached:":
			mem.Cached = parseSize(strutil.ReadField(text, 1, true), errs)
		case "SwapCached:":
			mem.SwapCached = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Active:":
			mem.Active = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Inactive:":
			mem.Inactive = parseSize(strutil.ReadField(text, 1, true), errs)
		case "SwapTotal:":
			mem.SwapTotal = parseSize(strutil.ReadField(text, 1, true), errs)
		case "SwapFree:":
			mem.SwapFree = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Dirty:":
			mem.Dirty = parseSize(strutil.ReadField(text, 1, true), errs)
		case "Slab:":
			mem.Slab = parseSize(strutil.ReadField(text, 1, true), errs)
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
