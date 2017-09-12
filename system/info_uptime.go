// +build !windows

package system

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

	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with uptime info in procfs
var procUptimeFile = "/proc/uptime"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime return system uptime in seconds
func GetUptime() (uint64, error) {
	fd, err := os.OpenFile(procUptimeFile, os.O_RDONLY, 0)

	if err != nil {
		return 0, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, err := r.ReadString('\n')

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
