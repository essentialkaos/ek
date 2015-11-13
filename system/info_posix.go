// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	result := &SystemInfo{}

	info := &syscall.Utsname{}
	err := syscall.Uname(info)

	if err != nil {
		return result, err
	}

	result.Hostname = byteSliceToString(info.Nodename)
	result.OS = byteSliceToString(info.Sysname)
	result.Kernel = byteSliceToString(info.Release)
	result.Arch = byteSliceToString(info.Machine)

	return result, nil
}
