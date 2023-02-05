package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/sys/unix"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime returns uptime in seconds from 1/1/1970
func GetUptime() (uint64, error) {
	tv, err := unix.SysctlTimeval("kern.boottime")

	if err != nil {
		return 0, err
	}

	return uint64(tv.Sec), nil
}
