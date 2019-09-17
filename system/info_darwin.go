// +build darwin

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

	archSlice := strings.Split(arch, "/")

	if len(archSlice) != 2 {
		return nil, errors.New("Can't read arch info")
	}

	cleanArch := strings.ToLower(strings.Replace(archSlice[len(archSlice)-1], "RELEASE_", "", -1))

	return &SystemInfo{
		Hostname:     hostname,
		OS:           os,
		Distribution: DARWIN_OSX,
		Version:      getOSXVersion(),
		Kernel:       kernel,
		Arch:         cleanArch,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getOSXVersion() string {
	cmd := exec.Command("sw_vers", "-productVersion")

	versionData, err := cmd.Output()

	if err != nil {
		return ""
	}

	return string(versionData)
}
