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
	"pkg.re/essentialkaos/ek.v9/strutil"
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

		switch strutil.ReadField(text, 0, true) {
		case "VmPeak:":
			info.VmPeak = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmSize:":
			info.VmSize = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmLck:":
			info.VmLck = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmPin:":
			info.VmPin = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmHWM:":
			info.VmHWM = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmRSS:":
			info.VmRSS = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmData:":
			info.VmData = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmStk:":
			info.VmStk = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmExe:":
			info.VmExe = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmLib:":
			info.VmLib = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmPTE:":
			info.VmPTE = parseSize(strutil.ReadField(text, 1, true), errs)
		case "VmSwap:":
			info.VmSwap = parseSize(strutil.ReadField(text, 1, true), errs)
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
