package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime returns system uptime in seconds
func GetUptime() (uint64, error) {
	info := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(info)

	return uint64(info.Uptime), err
}
