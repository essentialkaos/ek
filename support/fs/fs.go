//go:build !windows
// +build !windows

// Package pkgs provides methods for collecting information about filesystem
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v12/system"

	"github.com/essentialkaos/ek/v12/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects info about filesystem
func Collect() []support.FSInfo {
	fsInfo, err := system.GetFSUsage()

	if err != nil {
		return nil
	}

	var info []support.FSInfo

	for mPath, mInfo := range fsInfo {
		info = append(info, support.FSInfo{
			Path:   mPath,
			Device: mInfo.Device,
			Type:   mInfo.Type,
			Used:   mInfo.Used,
			Free:   mInfo.Free,
		})
	}

	return info
}
