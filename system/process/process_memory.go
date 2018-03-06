// +build linux

package process

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
	"strconv"

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

	for s.Scan() {
		text := s.Text()

		if len(text) < 2 || text[:2] != "Vm" {
			continue
		}

		switch strutil.ReadField(text, 0, true) {
		case "VmPeak:":
			info.VmPeak, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmSize:":
			info.VmSize, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmLck:":
			info.VmLck, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmPin:":
			info.VmPin, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmHWM:":
			info.VmHWM, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmRSS:":
			info.VmRSS, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmData:":
			info.VmData, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmStk:":
			info.VmStk, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmExe:":
			info.VmExe, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmLib:":
			info.VmLib, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmPTE:":
			info.VmPTE, err = parseSize(strutil.ReadField(text, 1, true))
		case "VmSwap:":
			info.VmSwap, err = parseSize(strutil.ReadField(text, 1, true))
		}

		if err != nil {
			return nil, errors.New("Can't parse status file for given process")
		}
	}

	if info.VmPeak+info.VmSize == 0 {
		return nil, errors.New("Can't parse status file for given process")
	}

	return info, nil
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
