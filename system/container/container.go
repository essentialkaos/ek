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
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DOCKER = "docker" // Docker
	PODMAN = "podman" // Podman
)

// ////////////////////////////////////////////////////////////////////////////////// //

var dockerEnv = "/.dockerenv"
var podmanEnv = "/run/.containerenv"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetEngine returns container engine name if used
func GetEngine() string {
	switch {
	case isFileExist(dockerEnv):
		return DOCKER
	case isFileExist(podmanEnv):
		return PODMAN
	}

	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// isFileExist returns true if given file exist
func isFileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
