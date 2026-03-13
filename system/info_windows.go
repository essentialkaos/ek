package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/sys/windows"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSystemInfo returns system info
func GetSystemInfo() (*SystemInfo, error) {
	return &SystemInfo{
		OS:       "Windows",
		ID:       "win",
		Arch:     getArchName(runtime.GOARCH),
		Kernel:   "Windows NT",
		ArchBits: getArchBits(runtime.GOARCH),
	}, nil
}

// GetOSInfo returns info about OS
func GetOSInfo() (*OSInfo, error) {
	versionInfo := windows.RtlGetVersion()

	if versionInfo == nil {
		return nil, fmt.Errorf("can't get windows version info")
	}

	version := fmt.Sprintf(
		"%d.%d.%d",
		versionInfo.MajorVersion, versionInfo.MinorVersion, versionInfo.BuildNumber,
	)

	info := &OSInfo{
		Name:       "Windows",
		ID:         "win",
		Version:    version,
		ANSIColor:  "36",
		PlatformID: "platform:win",
	}

	switch versionInfo.BuildNumber {
	case 22000, 22621, 22631, 26100, 26200, 28000, 26300:
		info.Name = "Windows 11"
		info.PlatformID += "11"
	case 10240, 10586, 15063, 16299, 17134, 18362,
		18363, 19041, 19042, 19043, 19044, 19045:
		info.Name = "Windows 10"
		info.PlatformID += "10"
	case 20348:
		info.Name = "Windows Server 2022"
		info.PlatformID += "10"
	case 17763:
		info.Name = "Windows Server 2019"
		info.PlatformID += "10"
	case 14393:
		info.Name = "Windows Server 2016"
		info.PlatformID += "10"
	case 9600:
		info.Name = "Windows 8.1"
		info.PlatformID += "8"
	case 9200:
		info.Name = "Windows 8"
		info.PlatformID += "8"
	}

	if versionInfo.ProductType != 0x0000001 { // 0x0000001 == NT_WORKSTATION
		switch versionInfo.BuildNumber {
		case 26100:
			info.Name = "Windows Server 2025"
		case 9600:
			info.Name = "Windows Server 2012 R2"
		case 9200:
			info.Name = "Windows Server 2012"
		}
	}

	switch versionInfo.BuildNumber {
	// Win 11
	case 26300:
		info.VersionID = "26H2"
	case 28000:
		info.VersionID = "26H1"
	case 26200:
		info.VersionID = "25H2"
	case 26100:
		info.VersionID = "24H2"
	case 22621:
		info.VersionID = "23H2"
	case 22631:
		info.VersionID = "22H2"
	case 22000:
		info.VersionID = "21H2"

	// Win 10
	case 19045:
		info.VersionID = "22H2"
	case 19044:
		info.VersionID = "21H2"
	case 19043:
		info.VersionID = "21H1"
	case 19042:
		info.VersionID = "20H2"
	}

	if info.VersionID != "" {
		info.PrettyName = info.Name + " " + info.VersionID
	} else {
		info.PrettyName = info.Name
	}

	return info, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getArchBits returns system arch bits (32 or 64)
func getArchBits(arch string) int {
	if strings.Contains(arch, "64") && !strings.Contains(arch, "32") {
		return 64
	}

	return 32
}
