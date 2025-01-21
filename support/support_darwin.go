package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "github.com/essentialkaos/ek/v13/system"

// ////////////////////////////////////////////////////////////////////////////////// //

// appendSystemInfo appends system info
func (i *Info) appendSystemInfo() {
	systemInfo, err := system.GetSystemInfo()

	if err != nil {
		return
	}

	i.System = &SystemInfo{
		Name:   systemInfo.OS,
		Arch:   systemInfo.Arch,
		Kernel: systemInfo.Kernel,
	}
}

// appendOSInfo appends OS info
func (i *Info) appendOSInfo() {
	osInfo, err := system.GetOSInfo()

	if err != nil {
		return
	}

	i.OS = &OSInfo{
		Name:    osInfo.Name,
		Version: osInfo.VersionID,
		Build:   osInfo.Build,
	}
}
