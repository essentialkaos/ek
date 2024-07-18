package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"runtime"

	"github.com/essentialkaos/ek/v13/fmtc"

	"golang.org/x/sys/windows"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// appendSystemInfo appends system info
func (i *Info) appendSystemInfo() {
	i.System = &SystemInfo{
		Name:   "Windows",
		Arch:   formatArchName(runtime.GOARCH),
		Kernel: "Windows NT",
	}
}

// appendOSInfo appends OS info
func (i *Info) appendOSInfo() {
	major, minor, build := windows.RtlGetNtVersionNumbers()
	i.OS = &OSInfo{
		Name:    "Windows",
		ID:      "win",
		Version: fmt.Sprintf("%d.%d.%d", major, minor, build),
	}

	switch build {
	case 22000, 22621, 22631, 26100:
		i.OS.Name = "Windows 11"
	case 10240, 10586, 15063, 16299, 17134, 18362,
		18363, 19041, 19042, 19043, 19044, 19045:
		i.OS.Name = "Windows 10"
	case 20348:
		i.OS.Name = "Windows Server 2022"
	case 17763:
		i.OS.Name = "Windows Server 2019"
	case 14393:
		i.OS.Name = "Windows Server 2016"
	case 9600:
		i.OS.Name = "Windows 8.1 / Windows Server 2012 R2"
	case 9200:
		i.OS.Name = "Windows 8 / Windows Server 2012"
	}

	switch build {
	// Win 11
	case 22000:
		i.OS.VersionID = "24H2"
	case 22621:
		i.OS.VersionID = "23H2"
	case 22631:
		i.OS.VersionID = "22H2"

	// Win 10
	case 19045:
		i.OS.VersionID = "22H2"
	case 19044:
		i.OS.VersionID = "21H2"
	case 19043:
		i.OS.VersionID = "21H1"
	case 19042:
		i.OS.VersionID = "20H2"
	}

	if i.OS.VersionID != "" {
		i.OS.PrettyName = i.OS.Name + " " + i.OS.VersionID
	} else {
		i.OS.PrettyName = i.OS.Name
	}

	i.OS.coloredName = fmtc.Sprintf("{c}%s{!}", i.OS.Name)
	i.OS.coloredPrettyName = fmtc.Sprintf("{c}%s{!}", i.OS.PrettyName)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// formatArchName formats arch name
func formatArchName(goos string) string {
	if goos == "amd64" {
		return "x86_64"
	}

	return goos
}
