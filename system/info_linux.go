// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

// getDistributionInfo try to find current OS distribution and version
func getDistributionInfo() (string, string) {
	var distribution string
	var version string

	switch {
	case isFileExist("/etc/arch-release"):
		distribution = LINUX_ARCH
		version = getRawRelease("/etc/arch-release")

	case isFileExist("/etc/centos-release"):
		distribution = LINUX_CENTOS
		version = getReleasePart("/etc/centos-release")

	case isFileExist("/etc/fedora-release"):
		distribution = LINUX_FEDORA
		version = getReleasePart("/etc/fedora-release")

	case isFileExist("/etc/gentoo-release"):
		distribution = LINUX_GENTOO
		version = getReleasePart("/etc/gentoo-release")

	case isFileExist("/etc/redhat-release"):
		distribution = LINUX_RHEL
		version = getReleasePart("/etc/redhat-release")

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

// getReleasePart extract release part from version info
func getReleasePart(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	return findOSVersionNumber(string(data))
}

// findOSVersionNumber try to find OS version number
func findOSVersionNumber(data string) string {
WORDLOOP:
	for _, word := range strings.Split(data, " ") {
		for _, r := range word {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
				continue
			default:
				continue WORDLOOP
			}
		}

		return word
	}

	return ""
}

// getRawRelease extract raw release data from given file
func getRawRelease(file string) string {
	data, err := ioutil.ReadFile(file)

	if err != nil || len(data) == 0 {
		return ""
	}

	return string(data)
}

// getUbuntuRelease extract info about Ubuntu release from given release file
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

// getUbuntuRelease extract info about SuSe release from given release file
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

// byteSliceToString convert byte slice to string
func byteSliceToString(s [65]int8) string {
	result := ""

	for _, r := range s {
		if r == 0 {
			break
		}

		result += string(r)
	}

	return result
}

// isFileExist check if file exist
func isFileExist(path string) bool {
	if path == "" {
		return false
	}

	return syscall.Access(path, syscall.F_OK) == nil
}
