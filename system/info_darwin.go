// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"

	"github.com/essentialkaos/ek/v13/strutil"
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

	arch = getMacOSArch(arch)

	return &SystemInfo{
		Hostname: hostname,
		ID:       getSystemID(),
		OS:       os,
		Kernel:   kernel,
		Arch:     arch,
		ArchName: getArchName(arch),
		ArchBits: 64,
	}, nil
}

// GetOSInfo returns info about OS
func GetOSInfo() (*OSInfo, error) {
	versionData, err := exec.Command("sw_vers").Output()

	if err != nil {
		return nil, err
	}

	info := &OSInfo{}

	for _, line := range strings.Split(string(versionData), "\n") {
		name := strutil.ReadField(line, 0, false, ':')
		value := strutil.ReadField(line, 1, false, ':')
		value = strings.Trim(value, " \r\t\n")

		switch name {
		case "ProductName":
			info.Name = value
		case "ProductVersion":
			info.Version = value
		case "BuildVersion":
			info.Build = value
		}
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getMacOSArch returns info about arch
func getMacOSArch(archInfo string) string {
	switch {
	case strings.Contains(archInfo, "X86_64"):
		return "x86_64"
	case strings.Contains(archInfo, "ARM64"):
		return "arm64"
	}

	return "unknown"
}

// getArchName returns name for given arch
func getArchName(arch string) string {
	if arch == "x86_64" {
		return "amd64"
	}

	return arch
}

// getSystemID returns unique system ID
func getSystemID() string {
	ioData, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()

	if err != nil {
		return ""
	}

	for _, line := range strings.Split(string(ioData), "\n") {
		if !strings.Contains(line, `"IOPlatformUUID"`) {
			continue
		}

		value := strutil.ReadField(line, 1, false, '=')
		return strings.Trim(value, " \r\t\n\"")
	}

	return ""
}
