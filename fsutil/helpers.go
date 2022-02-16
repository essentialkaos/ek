//go:build !windows
// +build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

var _disableCopyFileChecks bool // Flag for testing purposes only
var _disableMoveFileChecks bool // Flag for testing purposes only
var _disableCopyDirChecks bool  // Flag for testing purposes only

// ////////////////////////////////////////////////////////////////////////////////// //

// CopyFile copies file using bufio
func CopyFile(from, to string, perms ...os.FileMode) error {
	targetExist := IsExist(to)

	if targetExist && IsDir(to) {
		to = path.Join(to, path.Base(from))
		targetExist = false
	}

	dir := path.Dir(to)

	if !_disableCopyFileChecks {
		switch {
		case from == "":
			return errors.New("Can't copy file: Source file can't be blank")
		case to == "":
			return errors.New("Can't copy file: Target file can't be blank")

		case !IsExist(from):
			return errors.New("Can't copy file: File " + from + " does not exists")
		case !IsRegular(from):
			return errors.New("Can't copy file: File " + from + " is not a regular file")
		case !IsReadable(from):
			return errors.New("Can't copy file: File " + from + " is not readable")

		case !targetExist && !IsExist(dir):
			return errors.New("Can't copy file: Directory " + dir + " does not exists")
		case !targetExist && !IsWritable(dir):
			return errors.New("Can't copy file: Directory " + dir + " is not writable")

		case targetExist && !IsWritable(to):
			return errors.New("Can't copy file: Target file " + to + " is not writable")
		}
	}

	return copyFile(from, to, perms)
}

// MoveFile moves file
func MoveFile(from, to string, perms ...os.FileMode) error {
	targetExist := IsExist(to)

	if targetExist && IsDir(to) {
		to = path.Join(to, path.Base(from))
		targetExist = false
	}

	dir := path.Dir(to)

	if !_disableMoveFileChecks {
		switch {
		case from == "":
			return errors.New("Can't move file: Source file can't be blank")
		case to == "":
			return errors.New("Can't move file: Target file can't be blank")

		case !IsExist(from):
			return errors.New("Can't move file: File " + from + " does not exists")
		case !IsRegular(from):
			return errors.New("Can't move file: File " + from + " is not a regular file")
		case !IsReadable(from):
			return errors.New("Can't move file: File " + from + " is not readable")

		case !targetExist && !IsExist(dir):
			return errors.New("Can't move file: Directory " + dir + " does not exists")
		case !targetExist && !IsWritable(dir):
			return errors.New("Can't move file: Directory " + dir + " is not writable")
		}
	}

	return moveFile(from, to, perms)
}

// CopyDir copies directory content recursively to target directory
func CopyDir(from, to string) error {
	if !_disableCopyDirChecks {
		switch {
		case from == "":
			return errors.New("Can't copy directory: Source directory can't be blank")
		case to == "":
			return errors.New("Can't copy directory: Target directory can't be blank")

		case !IsExist(from):
			return errors.New("Can't copy directory: Directory " + from + " does not exists")
		case !IsDir(from):
			return errors.New("Can't copy directory: Target " + from + " is not a directory")
		case !IsReadable(from):
			return errors.New("Can't copy directory: Directory " + from + " is not readable")
		case IsExist(to) && !IsDir(to):
			return errors.New("Can't copy directory: Target " + to + " is not a directory")
		case IsExist(to) && !IsWritable(to):
			return errors.New("Can't copy directory: Directory " + to + " is not writable")
		}
	}

	if !IsExist(to) {
		err := os.Mkdir(to, GetMode(from))

		if err != nil {
			return err
		}
	}

	return copyDir(from, to)
}

// TouchFile creates empty file
func TouchFile(path string, perm os.FileMode) error {
	fd, err := os.OpenFile(path, os.O_CREATE, perm)

	if err != nil {
		return err
	}

	return fd.Close()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func copyFile(from, to string, perms []os.FileMode) error {
	from, to = path.Clean(from), path.Clean(to)

	var targetExist bool
	var perm os.FileMode

	if IsExist(to) {
		targetExist = true
	}

	if len(perms) == 0 {
		perm = GetMode(from)
	} else {
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

func moveFile(from, to string, perms []os.FileMode) error {
	from, to = path.Clean(from), path.Clean(to)

	err := os.Rename(from, to)

	if err != nil {
		return err
	}

	if len(perms) == 0 {
		return nil
	}

	return os.Chmod(to, perms[0])
}

func copyDir(from, to string) error {
	from, to = path.Clean(from), path.Clean(to)

	var err error

	for _, target := range List(from, false) {
		fp := from + "/" + target
		tp := to + "/" + target

		if IsDir(fp) {
			err = os.Mkdir(tp, GetMode(fp))

			if err != nil {
				return err
			}

			err = copyDir(fp, tp)

			if err != nil {
				return err
			}
		} else {
			err = CopyFile(fp, tp)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
