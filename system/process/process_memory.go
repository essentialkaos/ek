// +build linux

package process

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
	"strconv"

	"pkg.re/essentialkaos/ek.v9/errutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MemInfo contains process memory usage stats
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

// GetMemInfo return info about process memory usage
func GetMemInfo(pid int) (*MemInfo, error) {
	fd, err := os.OpenFile("/proc/"+strconv.Itoa(pid)+"/status", os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	info := &MemInfo{}
	errs := errutil.NewErrors()

	for s.Scan() {
		text := s.Text()

		if len(text) < 2 || text[:2] != "Vm" {
			continue
		}

		switch readField(text, 0) {
		case "VmPeak:":
			info.VmPeak = parseSize(readField(text, 1), errs)
		case "VmSize:":
			info.VmSize = parseSize(readField(text, 1), errs)
		case "VmLck:":
			info.VmLck = parseSize(readField(text, 1), errs)
		case "VmPin:":
			info.VmPin = parseSize(readField(text, 1), errs)
		case "VmHWM:":
			info.VmHWM = parseSize(readField(text, 1), errs)
		case "VmRSS:":
			info.VmRSS = parseSize(readField(text, 1), errs)
		case "VmData:":
			info.VmData = parseSize(readField(text, 1), errs)
		case "VmStk:":
			info.VmStk = parseSize(readField(text, 1), errs)
		case "VmExe:":
			info.VmExe = parseSize(readField(text, 1), errs)
		case "VmLib:":
			info.VmLib = parseSize(readField(text, 1), errs)
		case "VmPTE:":
			info.VmPTE = parseSize(readField(text, 1), errs)
		case "VmSwap:":
			info.VmSwap = parseSize(readField(text, 1), errs)
		}

		if errs.HasErrors() {
			return nil, errs.Last()
		}
	}

	if info.VmPeak+info.VmSize == 0 {
		return nil, errors.New("Can't parse status file for given process")
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
