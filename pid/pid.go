// Package pid provides methods for working with pid files
package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v1/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is path to directory with pid files
var Dir = "/var/run"

// ////////////////////////////////////////////////////////////////////////////////// //

// Create create file with process pid file
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
		err = os.Remove(pidFile)

		if err != nil {
			return err
		}
	}

	fd, err := os.OpenFile(pidFile, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	_, err = fd.WriteString(fmt.Sprintf("%d\n", os.Getpid()))

	return err
}

// Remove remove file with process pid file
func Remove(name string) error {
	err := checkPidDir(Dir)

	if err != nil {
		return err
	}

	pidFile := Dir + "/" + normalizePidFilename(name)

	os.Remove(pidFile)

	return nil
}

// Get return pid from pid file
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

func checkPidDir(path string) error {
	switch {
	case fsutil.IsExist(path) == false:
		return errors.New("Directory " + path + " is not exist")
	case fsutil.IsDir(path) == false:
		return errors.New(path + " is not directory")
	case fsutil.IsWritable(path) == false:
		return errors.New("Directory " + path + " is not writable")
	case fsutil.IsReadable(path) == false:
		return errors.New("Directory " + path + " is not readable")
	}

	return nil
}

func normalizePidFilename(name string) string {
	if !strings.Contains(name, ".pid") {
		return name + ".pid"
	}

	return name
}
