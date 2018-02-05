package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var _disableCopyFileChecks bool // Flags for testing purposes only
var _disableMoveFileChecks bool // Flags for testing purposes only

// ////////////////////////////////////////////////////////////////////////////////// //

// CopyFile simple file copying with bufio
func CopyFile(from, to string, perms ...os.FileMode) error {
	var targetExist = IsExist(to)

	if !_disableCopyFileChecks {
		dir := path.Dir(to)

		switch {
		case from == "":
			return errors.New("Source file can't be blank")
		case to == "":
			return errors.New("Target file can't be blank")

		case !IsExist(from):
			return errors.New("File " + from + " does not exists")
		case !IsRegular(from):
			return errors.New("File " + from + " is not a regular file")
		case !IsReadable(from):
			return errors.New("File " + from + " is not readable")

		case !targetExist && !IsExist(dir):
			return errors.New("Directory " + dir + " does not exists")
		case !targetExist && !IsWritable(dir):
			return errors.New("Directory " + dir + " is not writable")

		case targetExist && !IsWritable(to):
			return errors.New("Target file " + to + " is not writable")
		case targetExist && !IsRegular(to):
			return errors.New("Target is not a file")
		}
	}

	if IsExist(to) {
		targetExist = true
	}

	var perm os.FileMode = 0644

	if len(perms) != 0 {
		perm = perms[0]
	}

	ffd, err := os.OpenFile(from, os.O_RDONLY, 0)

	if err != nil {
		return err
	}

	defer ffd.Close()

	tfd, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)

	if err != nil {
		return err
	}

	defer tfd.Close()

	reader := bufio.NewReader(ffd)

	_, err = io.Copy(tfd, reader)

	if err != nil {
		return err
	}

	if targetExist {
		return os.Chmod(to, perm)
	}

	return nil
}

// MoveFile move file
func MoveFile(from, to string, perms ...os.FileMode) error {
	if !_disableMoveFileChecks {
		targetExist := IsExist(to)
		dir := path.Dir(to)

		switch {
		case from == "":
			return errors.New("Source file can't be blank")
		case to == "":
			return errors.New("Target file can't be blank")

		case !IsExist(from):
			return errors.New("File " + from + " does not exists")
		case !IsRegular(from):
			return errors.New("File " + from + " is not a regular file")
		case !IsReadable(from):
			return errors.New("File " + from + " is not readable")

		case !targetExist && !IsExist(dir):
			return errors.New("Directory " + dir + " does not exists")
		case !targetExist && !IsWritable(dir):
			return errors.New("Directory " + dir + " is not writable")
		}
	}

	err := os.Rename(from, to)

	if err != nil {
		return err
	}

	if len(perms) == 0 {
		return nil
	}

	return os.Chmod(to, perms[0])
}
