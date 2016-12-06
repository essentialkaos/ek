// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	info := &syscall.Utsname{}
	err := syscall.Uname(info)

	if err != nil {
		return nil, err
	}

	dist, version := getDistributionInfo()

	return &SystemInfo{
		Hostname:     byteSliceToString(info.Nodename),
		OS:           byteSliceToString(info.Sysname),
		Distribution: dist,
		Version:      version,
		Kernel:       byteSliceToString(info.Release),
		Arch:         byteSliceToString(info.Machine),
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getDistributionInfo() (string, string) {
	var distribution string
	var version string

	switch {
	case isFileExist("/etc/arch-release"):
		distribution = LINUX_ARCH
		version = getRawRelease("/etc/arch-release")

	case isFileExist("/etc/centos-release"):
		distribution = LINUX_CENTOS
		version = getReleasePart("/etc/centos-release", 2)

	case isFileExist("/etc/fedora-release"):
		distribution = LINUX_FEDORA
		version = getReleasePart("/etc/fedora-release", 2)

	case isFileExist("/etc/gentoo-release"):
		distribution = LINUX_GENTOO
		version = getReleasePart("/etc/gentoo-release", 4)

	case isFileExist("/etc/redhat-release"):
		distribution = LINUX_RHEL
		version = getReleasePart("/etc/redhat-release", 6)

	case isFileExist("/etc/SuSE-release"):
		distribution = LINUX_SUSE
		version = getSuseVersion("/etc/SuSE-release")

	case isFileExist("/etc/lsb-release"):
		distribution = LINUX_UBUNTU
		version = getUbuntuRelease("/etc/lsb-release")

	case isFileExist("/etc/debian_version"):
		distribution = LINUX_DEBIAN
		version = getRawRelease("/etc/debian_version")
	}

	return distribution, version
}

func getReleasePart(file string, field int) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	versionSlice := strings.Split(string(data), " ")

	if len(versionSlice) < field+1 {
		return ""
	}

	return versionSlice[field]
}

func getRawRelease(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	return string(data)
}

func getUbuntuRelease(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) < 4 {
		return ""
	}

	versionSlice := strings.Split(lines[3], " ")

	if len(versionSlice) < 3 {
		return ""
	}

	return versionSlice[1]
}

func getSuseVersion(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) < 3 {
		return ""
	}

	versionSlice := strings.Split(lines[1], " ")
	patchSlice := strings.Split(lines[2], " ")

	if len(versionSlice) != 3 || len(patchSlice) != 3 {
		return ""
	}

	return versionSlice[2] + "." + patchSlice[2]
}
