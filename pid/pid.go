// Package pid provides methods for working with PID files
package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v10/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is path to directory with PID files
var Dir = "/var/run"

// ////////////////////////////////////////////////////////////////////////////////// //

// Create creates file with process PID file
func Create(name string) error {
	err := checkPIDDir(Dir, true)

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

	return ioutil.WriteFile(
		pidFile,
		[]byte(fmt.Sprintf("%d\n", os.Getpid())),
		0644,
	)
}

// Remove remove file with process PID file
func Remove(name string) error {
	err := checkPIDDir(Dir, true)

	if err != nil {
		return err
	}

	pidFile := Dir + "/" + normalizePIDFilename(name)

	return os.Remove(pidFile)
}

// Get return PID from PID file
func Get(name string) int {
	err := checkPIDDir(Dir, false)

	if err != nil {
		return -1
	}

	pidFile := Dir + "/" + normalizePIDFilename(name)
	data, err := ioutil.ReadFile(pidFile)

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

// checkPIDDir checks directory path and return error if directory not ok
func checkPIDDir(path string, requireModify bool) error {
	switch {
	case !fsutil.IsExist(path):
		return errors.New("Directory " + path + " does not exist")

	case !fsutil.IsDir(path):
		return errors.New(path + " is not directory")

	case !fsutil.IsWritable(path) && requireModify:
		return errors.New("Directory " + path + " is not writable")

	case !fsutil.IsReadable(path):
		return errors.New("Directory " + path + " is not readable")
	}

	return nil
}

// normalizePIDFilename returns PID file name with extension
func normalizePIDFilename(name string) string {
	if !strings.Contains(name, ".pid") {
		return name + ".pid"
	}

	return name
}
