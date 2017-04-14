// +build !windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with uptime info in procfs
var procUptimeFile = "/proc/uptime"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime return system uptime in seconds
func GetUptime() (uint64, error) {
	content, err := readFileContent(procUptimeFile)

	if err != nil {
		return 0, err
	}

	ca := strings.Split(content[0], " ")

	if len(ca) != 2 {
		return 0, errors.New("Can't parse file " + procUptimeFile)
	}

	up, _ := strconv.ParseFloat(ca[0], 64)

	return uint64(up), nil
}
