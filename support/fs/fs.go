//go:build !windows

// Package pkgs provides methods for collecting information about filesystem
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/system"

	"github.com/essentialkaos/ek/v13/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect returns usage statistics for all currently mounted filesystems.
// Returns nil if filesystem information cannot be retrieved from the OS.
func Collect() []support.FSInfo {
	fsInfo, err := system.GetFSUsage()

	if err != nil {
		return nil
	}

	info := make([]support.FSInfo, len(fsInfo))

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
