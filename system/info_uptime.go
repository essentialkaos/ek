//go:build linux
// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"

	"pkg.re/essentialkaos/ek.v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with uptime info in procfs
var procUptimeFile = "/proc/uptime"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime returns system uptime in seconds
func GetUptime() (uint64, error) {
	fd, err := os.OpenFile(procUptimeFile, os.O_RDONLY, 0)

	if err != nil {
		return 0, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		return 0, err
	}

	return parseUptime(text)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseUptime parses uptime data
func parseUptime(text string) (uint64, error) {
	uptimeStr := strutil.ReadField(text, 0, true)

	if uptimeStr == "" {
		return 0, errors.New("Can't parse file " + procUptimeFile)
	}

	uptimeInt, err := strconv.ParseFloat(uptimeStr, 64)

	if err != nil {
		return 0, errors.New("Can't parse file " + procUptimeFile)
	}

	return uint64(uptimeInt), nil
}
