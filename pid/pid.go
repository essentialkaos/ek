// Package pid provides methods for working with PID files
package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyName is returned if pid file name is empty
	ErrEmptyName = errors.New("PID file name can't be blank")

	// ErrInvalidPID returns if PID read from file is invalid
	ErrInvalidPID = errors.New("PID is invalid (PID ≤ 0)")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is the path to the directory where PID files are stored
var Dir = "/var/run"

// ////////////////////////////////////////////////////////////////////////////////// //

// Create writes the current process PID to a file named after the given service name.
// The file is created in [Dir] and the ".pid" extension is appended if absent.
func Create(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	err := fsutil.ValidatePerms("DRW", Dir)

	if err != nil {
		return err
	}

	pidFile := getPIDFilePath(name)

	if fsutil.IsExist(pidFile) {
		os.Remove(pidFile)
	}

	return os.WriteFile(
		pidFile,
		[]byte(fmt.Sprintf("%d\n", os.Getpid())),
		0644,
	)
}

// Remove deletes the PID file associated with the given service name from [Dir]
func Remove(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	err := fsutil.ValidatePerms("DRW", Dir)

	if err != nil {
		return err
	}

	pidFile := getPIDFilePath(name)

	return os.Remove(pidFile)
}

// Get returns the PID stored in the named PID file inside [Dir]
func Get(name string) (int, error) {
	if name == "" {
		return 0, ErrEmptyName
	}

	err := fsutil.ValidatePerms("DR", Dir)

	if err != nil {
		return 0, err
	}

	pidFile := getPIDFilePath(name)

	return Read(pidFile)
}

// Read reads and parses a PID from the given absolute PID file path
func Read(pidFile string) (int, error) {
	data, err := os.ReadFile(pidFile)

	if err != nil {
		return 0, fmt.Errorf("can't read PID file: %w", err)
	}

	pid, err := strconv.Atoi(strings.TrimRight(string(data), "\n"))

	if err != nil {
		return 0, fmt.Errorf("can't parse PID: %w", err)
	}

	if pid <= 0 {
		return 0, ErrInvalidPID
	}

	return pid, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getPIDFilePath returns path to PID file
func getPIDFilePath(name string) string {
	return filepath.Join(Dir, normalizePIDFilename(name))
}

// normalizePIDFilename ensures that the PID file name ends with ".pid"
func normalizePIDFilename(name string) string {
	if !strings.HasSuffix(name, ".pid") {
		return name + ".pid"
	}

	return name
}
