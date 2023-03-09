// Package container provides methods for checking container engine info
package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DOCKER = "docker" // Docker (Moby)
	PODMAN = "podman" // Podman
	LXC    = "lxc"    // LXC
)

// ////////////////////////////////////////////////////////////////////////////////// //

// mountsFile is path to mounts file for init process
var mountsFile = "/proc/1/mounts"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetEngine returns container engine name if used
func GetEngine() string {
	mountsData, err := os.ReadFile(mountsFile)

	if err != nil {
		return ""
	}

	return guessEngine(string(mountsData))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// guessEngine tries to guess container engine based on information from /proc/1/mounts
func guessEngine(mountsData string) string {
	switch {
	case strings.Contains(mountsData, "lxcfs "):
		return LXC
	case strings.Contains(mountsData, "workdir=/var/lib/containers"):
		return PODMAN
	case strings.Contains(mountsData, "workdir=/var/lib/docker"):
		return DOCKER
	}

	return ""
}
