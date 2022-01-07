// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo returns system info
func GetSystemInfo() (*SystemInfo, error) {
	hostname, err := syscall.Sysctl("kern.hostname")

	if err != nil || hostname == "" {
		return nil, errors.New("Can't read hostname info")
	}

	os, err := syscall.Sysctl("kern.ostype")

	if err != nil || os == "" {
		return nil, errors.New("Can't read os info")
	}

	kernel, err := syscall.Sysctl("kern.osrelease")

	if err != nil || kernel == "" {
		return nil, errors.New("Can't read kernel info")
	}

	arch, err := syscall.Sysctl("kern.version")

	if err != nil || arch == "" {
		return nil, errors.New("Can't read arch info")
	}

	return &SystemInfo{
		Hostname:     hostname,
		OS:           os,
		Distribution: DARWIN_OSX,
		Version:      getMacOSVersion(),
		Kernel:       kernel,
		Arch:         getMacOSArch(arch),
		ArchBits:     64,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMacOSArch(archInfo string) string {
	switch {
	case strings.Contains(archInfo, "X86_64"):
		return "x86_64"
	case strings.Contains(archInfo, "ARM64"):
		return "arm64"
	}

	return "unknown"
}

func getMacOSVersion() string {
	cmd := exec.Command("sw_vers", "-productVersion")

	versionData, err := cmd.Output()

	if err != nil {
		return ""
	}

	return strings.Trim(string(versionData), "\r\n")
}
