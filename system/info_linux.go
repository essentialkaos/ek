//go:build linux
// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
	"syscall"

	"github.com/essentialkaos/ek/v12/strutil"
	"github.com/essentialkaos/ek/v12/system/container"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var osReleaseFile = "/etc/os-release"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo returns system info
func GetSystemInfo() (*SystemInfo, error) {
	info := &syscall.Utsname{}
	err := syscall.Uname(info)

	if err != nil {
		return nil, err
	}

	arch := byteSliceToString(info.Machine)

	return &SystemInfo{
		Hostname:        byteSliceToString(info.Nodename),
		OS:              byteSliceToString(info.Sysname),
		Kernel:          byteSliceToString(info.Release),
		Arch:            arch,
		ArchName:        getArchName(arch),
		ArchBits:        getCPUArchBits(),
		ContainerEngine: container.GetEngine(),
	}, nil
}

// GetOSInfo returns info about OS
func GetOSInfo() (*OSInfo, error) {
	return ParseOSInfo(osReleaseFile)
}

// ParseOSInfo parses data in given os-release file
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

		name := strutil.ReadField(line, 0, false, "=")
		value := strings.Trim(strutil.ReadField(line, 1, false, "="), "\"")

		applyOSInfo(info, name, value)
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

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
	}

	switch {
	case strings.HasSuffix(name, "SUPPORT_PRODUCT"):
		info.SupportProduct = value
	case strings.HasSuffix(name, "SUPPORT_PRODUCT_VERSION"):
		info.SupportProductVersion = value
	}
}

// getArchName returns name for given arch
func getArchName(arch string) string {
	switch arch {
	case "i386":
		return "386"
	case "i586":
		return "586"
	case "i686":
		return "686"
	case "x86_64":
		return "amd64"
	}

	return arch
}
