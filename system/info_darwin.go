// +build darwin

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	hostname, err := syscall.Sysctl("kern.hostname")

	if err != nil || hostname == "" {
		return &SystemInfo{}, errors.New("Can't read hostname info")
	}

	os, err := syscall.Sysctl("kern.ostype")

	if err != nil || os == "" {
		return &SystemInfo{}, errors.New("Can't read os info")
	}

	kernel, err := syscall.Sysctl("kern.osrelease")

	if err != nil || kernel == "" {
		return &SystemInfo{}, errors.New("Can't read kernel info")
	}

	arch, err := syscall.Sysctl("kern.version")

	if err != nil || arch == "" {
		return &SystemInfo{}, errors.New("Can't read arch info")
	}

	archSlice := strings.Split(arch, "/")

	if len(archSlice) != 2 {
		return &SystemInfo{}, errors.New("Can't read arch info")
	}

	return &SystemInfo{
		hostname,
		os,
		kernel,
		strings.ToLower(strings.Replace(archSlice[len(archSlice)-1], "RELEASE_", "", -1)),
	}, nil
}
