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
	"errors"
	"io/ioutil"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

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

// SystemInfo contains info about system (hostname, OS, arch...)
type SystemInfo struct {
	Hostname     string `json:"hostname"`     // Hostname
	OS           string `json:"os"`           // OS name
	Distribution string `json:"distribution"` // OS distribution
	Version      string `json:"version"`      // OS version
	Kernel       string `json:"kernel"`       // Kernel version
	Arch         string `json:"arch"`         // System architecture (i386/i686/x86_64/etc...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readFileContent(file string) ([]string, error) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	if string(content) == "" {
		return nil, errors.New("File " + file + " is empty")
	}

	return strings.Split(string(content), "\n"), nil
}

func cleanSlice(s []string) []string {
	var result []string

	for _, item := range s {
		if item != "" {
			result = append(result, item)
		}
	}

	return result
}
