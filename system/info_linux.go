//go:build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"os"
	"strings"
	"syscall"

	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/system/container"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// osReleaseFile is the path to a file with information about operating system
var osReleaseFile = "/etc/os-release"

// machineIDFile is the path to a file with unique system ID
var machineIDFile = "/etc/machine-id"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo returns general information about the host system
func GetSystemInfo() (*SystemInfo, error) {
	info := &syscall.Utsname{}
	err := syscall.Uname(info)

	if err != nil {
		return nil, err
	}

	arch := byteSliceToString(info.Machine)

	return &SystemInfo{
		Hostname:        byteSliceToString(info.Nodename),
		ID:              getSystemID(),
		OS:              byteSliceToString(info.Sysname),
		Kernel:          byteSliceToString(info.Release),
		Arch:            arch,
		ArchName:        getArchName(arch),
		ArchBits:        getCPUArchBits(),
		ContainerEngine: container.GetEngine(),
	}, nil
}

// GetOSInfo returns information parsed from the default os-release file
func GetOSInfo() (*OSInfo, error) {
	return ParseOSInfo(osReleaseFile)
}

// ParseOSInfo parses OS release information from the given os-release file path
func ParseOSInfo(file string) (*OSInfo, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	info := &OSInfo{}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		name := strutil.ReadField(line, 0, false, '=')
		value := strings.Trim(strutil.ReadField(line, 1, false, '='), "\"")

		applyOSInfo(info, name, value)
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getFileScanner opens file and creates scanner for reading text files line by line
func getFileScanner(file string) (*bufio.Scanner, func() error, error) {
	fd, err := os.Open(file)

	if err != nil {
		return nil, nil, err
	}

	s := bufio.NewScanner(fd)

	return s, fd.Close, nil
}

// applyOSInfo applies record from os-release
func applyOSInfo(info *OSInfo, name, value string) {
	switch name {
	case "NAME":
		info.Name = value
	case "VERSION":
		info.Version = value
	case "ID":
		info.ID = value
	case "ID_LIKE":
		info.IDLike = value
	case "PRETTY_NAME":
		info.PrettyName = value
	case "VERSION_ID":
		info.VersionID = value
	case "PLATFORM_ID":
		info.PlatformID = value
	case "VARIANT":
		info.Variant = value
	case "VARIANT_ID":
		info.VariantID = value
	case "CPE_NAME":
		info.CPEName = value
	case "VERSION_CODENAME":
		info.VersionCodename = value
	case "HOME_URL":
		info.HomeURL = value
	case "BUG_REPORT_URL":
		info.BugReportURL = value
	case "SUPPORT_URL":
		info.SupportURL = value
	case "DOCUMENTATION_URL":
		info.DocumentationURL = value
	case "LOGO":
		info.Logo = value
	case "ANSI_COLOR":
		info.ANSIColor = value
	}

	switch {
	case strings.HasSuffix(name, "SUPPORT_PRODUCT"):
		info.SupportProduct = value
	case strings.HasSuffix(name, "SUPPORT_PRODUCT_VERSION"):
		info.SupportProductVersion = value
	}
}

// getSystemID returns unique system ID
func getSystemID() string {
	id, err := os.ReadFile(machineIDFile)

	if err != nil {
		return ""
	}

	return strings.TrimRight(string(id), "\n")
}
