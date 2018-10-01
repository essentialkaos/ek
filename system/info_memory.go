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

	"pkg.re/essentialkaos/ek.v10/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with memory info in procfs
var procMemInfoFile = "/proc/meminfo"

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

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

	for s.Scan() {
		text := s.Text()

		switch strutil.ReadField(text, 0, true) {
		case "MemTotal:":
			mem.MemTotal, err = parseSize(strutil.ReadField(text, 1, true))
		case "MemFree:":
			mem.MemFree, err = parseSize(strutil.ReadField(text, 1, true))
		case "Buffers:":
			mem.Buffers, err = parseSize(strutil.ReadField(text, 1, true))
		case "Cached:":
			mem.Cached, err = parseSize(strutil.ReadField(text, 1, true))
		case "SwapCached:":
			mem.SwapCached, err = parseSize(strutil.ReadField(text, 1, true))
		case "Active:":
			mem.Active, err = parseSize(strutil.ReadField(text, 1, true))
		case "Inactive:":
			mem.Inactive, err = parseSize(strutil.ReadField(text, 1, true))
		case "SwapTotal:":
			mem.SwapTotal, err = parseSize(strutil.ReadField(text, 1, true))
		case "SwapFree:":
			mem.SwapFree, err = parseSize(strutil.ReadField(text, 1, true))
		case "Dirty:":
			mem.Dirty, err = parseSize(strutil.ReadField(text, 1, true))
		case "Shmem:":
			mem.Shmem, err = parseSize(strutil.ReadField(text, 1, true))
		case "Slab:":
			mem.Slab, err = parseSize(strutil.ReadField(text, 1, true))
		case "SReclaimable:":
			mem.SReclaimable, err = parseSize(strutil.ReadField(text, 1, true))
		}

		if err != nil {
			return nil, errors.New("Can't parse file " + procMemInfoFile)
		}
	}

	if mem.MemTotal == 0 {
		return nil, errors.New("Can't parse file " + procMemInfoFile)
	}

	mem.MemFree += (mem.Cached + mem.Buffers + mem.SReclaimable) - mem.Shmem
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	return mem, nil
}

// codebeat:enable[LOC,ABC]
