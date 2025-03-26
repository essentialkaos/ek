// Package container provides methods for checking container engine info
package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DOCKER       = "docker"       // Docker (Moby)
	PODMAN       = "podman"       // Podman
	LXC          = "lxc"          // LXC
	CONTAINERD   = "containerd"   // containerd
	YANDEX       = "yandex"       // Yandex Serverless
	DOCKER_RUNSC = "docker+runsc" // Docker (Moby) + runsc (gVisor)
)

// ////////////////////////////////////////////////////////////////////////////////// //

// mountsFile is path to mounts file for init process
var mountsFile = "/proc/1/mounts"

// dockerEnv is path to env file inside a Docker container
var dockerEnv = "/.dockerenv"

// ////////////////////////////////////////////////////////////////////////////////// //

// engineChecked set to true if engine was checked
var engineChecked bool

// engineCache cached engine info
var engineCache string

// isK8s is set to true if we inside K8s pod container
var isK8s bool

// ////////////////////////////////////////////////////////////////////////////////// //

// GetEngine returns container engine name if used
func GetEngine() string {
	if engineChecked {
		return engineCache
	}

	mountsData, err := os.ReadFile(mountsFile)

	if err != nil {
		return ""
	}

	engineChecked = true
	engineCache = guessEngine(string(mountsData))

	return engineCache
}

// IsContainer returns true if we are inside container
func IsContainer() bool {
	return GetEngine() != ""
}

// IsK8s returns true if we are inside K8s pod container
func IsK8s() bool {
	return GetEngine() != "" && isK8s
}

// ////////////////////////////////////////////////////////////////////////////////// //

// guessEngine tries to guess container engine based on information from /proc/1/mounts
func guessEngine(mountsData string) string {
	_, err := os.Stat(dockerEnv)

	isK8s = strings.Contains(mountsData, "kubernetes.io")

	switch {
	case strings.Contains(mountsData, "overlay-container /function/code/rootfs"):
		return YANDEX
	case strings.Contains(mountsData, "io.containerd"):
		return CONTAINERD
	case strings.Contains(mountsData, "lxcfs "):
		return LXC
	case strings.Contains(mountsData, "workdir=/var/lib/containers"):
		return PODMAN
	case err == nil &&
		strings.Contains(mountsData, "none /etc/hostname 9p") &&
		strings.Contains(mountsData, "none /etc/hosts 9p"):
		return DOCKER_RUNSC
	case strings.Contains(mountsData, "workdir=/var/lib/docker"):
		return DOCKER
	}

	return ""
}
