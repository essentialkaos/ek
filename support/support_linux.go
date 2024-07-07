package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/system"
	"github.com/essentialkaos/ek/v13/system/container"
)

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

	if i.OS != nil {
		i.System.Name = i.OS.Name
	}

	i.System.ContainerEngine = container.GetEngine()
}

// appendOSInfo appends OS info
func (i *Info) appendOSInfo() {
	osInfo, err := system.GetOSInfo()

	if err != nil {
		return
	}

	i.OS = &OSInfo{
		Name:        osInfo.Name,
		PrettyName:  osInfo.PrettyName,
		Version:     osInfo.Version,
		ID:          osInfo.ID,
		IDLike:      osInfo.IDLike,
		VersionID:   osInfo.VersionID,
		VersionCode: osInfo.VersionCodename,
		PlatformID:  osInfo.PlatformID,
		CPE:         osInfo.CPEName,

		coloredName:       osInfo.ColoredName(),
		coloredPrettyName: osInfo.ColoredPrettyName(),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //
