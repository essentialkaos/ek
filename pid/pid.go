// Package pid provides methods for working with PID files
package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is a path to directory with PID files
var Dir = "/var/run"

// ////////////////////////////////////////////////////////////////////////////////// //

// Create creates file with process PID file
func Create(name string) error {
	err := fsutil.ValidatePerms("DRW", Dir)

	if err != nil {
		return err
	}

	if name == "" {
		return errors.New("PID file name can't be blank")
	}

	pidFile := Dir + "/" + normalizePIDFilename(name)

	if fsutil.IsExist(pidFile) {
		os.Remove(pidFile)
	}

	return os.WriteFile(
		pidFile,
		[]byte(fmt.Sprintf("%d\n", os.Getpid())),
		0644,
	)
}

// Remove removes file with process PID file
func Remove(name string) error {
	err := fsutil.ValidatePerms("DRW", Dir)

	if err != nil {
		return err
	}

	pidFile := Dir + "/" + normalizePIDFilename(name)

	return os.Remove(pidFile)
}

// Get returns PID from PID file
func Get(name string) int {
	err := fsutil.ValidatePerms("DR", Dir)

	if err != nil {
		return -1
	}

	pidFile := Dir + "/" + normalizePIDFilename(name)

	return Read(pidFile)
}

// Read just reads PID from PID file
func Read(pidFile string) int {
	data, err := os.ReadFile(pidFile)

	if err != nil {
		return -1
	}

	pid, err := strconv.Atoi(strings.TrimRight(string(data), "\n"))

	if err != nil {
		return -1
	}

	return pid
}

// ////////////////////////////////////////////////////////////////////////////////// //

// normalizePIDFilename returns PID file name with extension
func normalizePIDFilename(name string) string {
	if !strings.Contains(name, ".pid") {
		return name + ".pid"
	}

	return name
}
