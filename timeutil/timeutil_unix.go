//go:build linux || freebsd || darwin

package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// localZoneFile is path to link to current tz info file
var localZoneFile = "/etc/localtime"

// ////////////////////////////////////////////////////////////////////////////////// //

// LocalTimezone returns name of local timezone
func LocalTimezone() string {
	tzName := os.Getenv("TZ")

	if tzName != "" {
		return tzName
	}

	zoneFile, err := os.Readlink(localZoneFile)

	if err != nil {
		return "Local"
	}

	index := strings.Index(zoneFile, "zoneinfo/")

	if index == -1 {
		return "Local"
	}

	return zoneFile[index+9:]
}
