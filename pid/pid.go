// Package pid provides methods for working with PID files
package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
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

	"pkg.re/essentialkaos/ek.v8/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is path to directory with PID files
var Dir = "/var/run"

// ////////////////////////////////////////////////////////////////////////////////// //

// Create create file with process PID file
func Create(name string) error {
	err := checkPidDir(Dir)

	if err != nil {
		return err
	}

	if name == "" {
		return errors.New("Pid file name can't be blank")
	}

	pidFile := Dir + "/" + normalizePidFilename(name)

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
	err := checkPidDir(Dir)

	if err != nil {
		return err
	}

	pidFile := Dir + "/" + normalizePidFilename(name)

	return os.Remove(pidFile)
}

// Get return PID from PID file
func Get(name string) int {
	err := checkPidDir(Dir)

	if err != nil {
		return -1
	}

	pidFile := Dir + "/" + normalizePidFilename(name)
	data, err := ioutil.ReadFile(pidFile)

	if err != nil {
		return -1
	}

	pid, err := strconv.Atoi(strings.TrimRight(string(data[:]), "\n"))

	if err != nil {
		return -1
	}

	return pid
}

// ////////////////////////////////////////////////////////////////////////////////// //

// checkPidDir check directory path and return error if directory not ok
func checkPidDir(path string) error {
	switch {
	case fsutil.IsExist(path) == false:
		return errors.New("Directory " + path + " does not exist")
	case fsutil.IsDir(path) == false:
		return errors.New(path + " is not directory")
	case fsutil.IsWritable(path) == false:
		return errors.New("Directory " + path + " is not writable")
	case fsutil.IsReadable(path) == false:
		return errors.New("Directory " + path + " is not readable")
	}

	return nil
}

// normalizePidFilename return PID file name with extension
func normalizePidFilename(name string) string {
	if !strings.Contains(name, ".pid") {
		return name + ".pid"
	}

	return name
}
