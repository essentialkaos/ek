// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"strconv"
	"strings"
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
	data, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/status")

	if err != nil {
		return nil, err
	}

	info := &MemInfo{}

	for _, line := range strings.Split(string(data), "\n") {
		lineSlice := splitLine(line)

		if len(lineSlice) < 2 {
			continue
		}

		switch lineSlice[0] {
		case "VmPeak:":
			info.VmPeak, err = parseSize(lineSlice[1])
		case "VmSize:":
			info.VmSize, err = parseSize(lineSlice[1])
		case "VmLck:":
			info.VmLck, err = parseSize(lineSlice[1])
		case "VmPin:":
			info.VmPin, err = parseSize(lineSlice[1])
		case "VmHWM:":
			info.VmHWM, err = parseSize(lineSlice[1])
		case "VmRSS:":
			info.VmRSS, err = parseSize(lineSlice[1])
		case "VmData:":
			info.VmData, err = parseSize(lineSlice[1])
		case "VmStk:":
			info.VmStk, err = parseSize(lineSlice[1])
		case "VmExe:":
			info.VmExe, err = parseSize(lineSlice[1])
		case "VmLib:":
			info.VmLib, err = parseSize(lineSlice[1])
		case "VmPTE:":
			info.VmPTE, err = parseSize(lineSlice[1])
		case "VmSwap:":
			info.VmSwap, err = parseSize(lineSlice[1])
		default:
			continue
		}

		if err != nil {
			return nil, err
		}
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// splitLine split line to slice by whitespace symbol
func splitLine(line string) []string {
	if line == "" {
		return nil
	}

	var (
		result []string
		buffer string
		space  bool
	)

	for _, r := range line {
		if r == ' ' || r == '\t' {
			space = true
			continue
		}

		if space == true {
			if buffer != "" {
				result = append(result, buffer)
			}

			buffer, space = "", false
		}

		buffer += string(r)
	}

	if buffer != "" {
		result = append(result, buffer)
	}

	return result
}

// parseSize convert string with size in kb to uint64 bytes
func parseSize(s string) (uint64, error) {
	size, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, err
	}

	return size * 1024, nil
}
