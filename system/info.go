// +build !windows

// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"

	"pkg.re/essentialkaos/ek.v9/errutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// OS names
const (
	LINUX_ARCH   = "Arch"
	LINUX_CENTOS = "CentOS"
	LINUX_DEBIAN = "Debian"
	LINUX_FEDORA = "Dedora"
	LINUX_GENTOO = "Gentoo"
	LINUX_RHEL   = "RHEL"
	LINUX_SUSE   = "SuSe"
	LINUX_UBUNTU = "Ubuntu"
	DARWIN_OSX   = "OSX"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SystemInfo contains info about a system (hostname, OS, arch...)
type SystemInfo struct {
	Hostname     string `json:"hostname"`     // Hostname
	OS           string `json:"os"`           // OS name
	Distribution string `json:"distribution"` // OS distribution
	Version      string `json:"version"`      // OS version
	Kernel       string `json:"kernel"`       // Kernel version
	Arch         string `json:"arch"`         // System architecture (i386/i686/x86_64/etc...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseSize(v string, errs *errutil.Errors) uint64 {
	size, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return size * 1024
}

func parseUint(v string, errs *errutil.Errors) uint64 {
	value, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return value
}

func parseFloat(v string, errs *errutil.Errors) float64 {
	value, err := strconv.ParseFloat(v, 64)

	if err != nil {
		errs.Add(err)
		return 0.0
	}

	return value
}

func parseInt(v string, errs *errutil.Errors) int {
	value, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return int(value)
}

// readField read field from data
func readField(data string, index int) string {
	if data == "" {
		return ""
	}

	curIndex, startPointer := -1, -1

	for i, r := range data {
		if r == ' ' || r == '\t' {
			if curIndex == index {
				return data[startPointer:i]
			}

			startPointer = -1
			continue
		}

		if startPointer == -1 {
			startPointer = i
			curIndex++
		}
	}

	if index > curIndex {
		return ""
	}

	return data[startPointer:]
}
