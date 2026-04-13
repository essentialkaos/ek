//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMemInfo returns memory usage statistics for the given process
func GetMemInfo(pid int) (*MemInfo, error) {
	statusFile := path.Join(procFS, strconv.Itoa(pid), "status")
	s, closeFunc, err := getFileScanner(statusFile)

	if err != nil {
		return nil, err
	}

	defer closeFunc()

	info := &MemInfo{}

	for s.Scan() {
		text := s.Text()

		if !strings.HasPrefix(text, "Vm") {
			continue
		}

		field := strutil.ReadField(text, 0, true)

		switch field {
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
			return nil, fmt.Errorf("can't parse field %q from stats file %s", field, statusFile)
		}
	}

	if info.VmPeak+info.VmSize == 0 {
		return nil, fmt.Errorf("invalid status data: VmPeak+VmSize is equals to 0")
	}

	return info, s.Err()
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
