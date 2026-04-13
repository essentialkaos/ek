package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"

	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with memory info in procfs
var procMemInfoFile = "/proc/meminfo"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMemUsage returns current physical and swap memory usage
func GetMemUsage() (*MemUsage, error) {
	s, closer, err := getFileScanner(procMemInfoFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseMemUsage(s)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseMemUsage parses memory usage info
func parseMemUsage(s *bufio.Scanner) (*MemUsage, error) {
	var err error

	mem := &MemUsage{}

	for s.Scan() {
		text := s.Text()
		fieldName := strutil.ReadField(text, 0, true)

		switch fieldName {
		case "MemTotal:":
			mem.MemTotal, err = parseSize(strutil.ReadField(text, 1, true))
		case "MemFree:":
			mem.MemFree, err = parseSize(strutil.ReadField(text, 1, true))
		case "MemAvailable":
			mem.MemAvailable, err = parseSize(strutil.ReadField(text, 1, true))
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
			return nil, fmt.Errorf(
				"can't parse field %q in procfs file %s",
				fieldName, procMemInfoFile,
			)
		}
	}

	mem.MemFree += (mem.Cached + mem.Buffers + mem.SReclaimable) - mem.Shmem
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	return mem, nil
}
