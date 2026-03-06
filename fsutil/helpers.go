//go:build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	chmodFunc    = os.Chmod
	chownFunc    = os.Chown
	chtimesFunc  = os.Chtimes
	openFileFunc = os.OpenFile
	mkDirFunc    = os.Mkdir
	ioCopyFunc   = io.Copy

	modeFunc  = GetMode
	ownerFunc = GetOwner
	timesFunc = GetTimes
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CopyFile copies file using bufio with given permissions
func CopyFile(from, to string, perms ...os.FileMode) error {
	targetExist := IsExist(to)

	if targetExist && IsDir(to) {
		to = path.Join(to, path.Base(from))
		targetExist = false
	}

	dir := path.Dir(to)

	switch {
	case from == "":
		return fmt.Errorf("can't copy file: source file can't be blank")
	case to == "":
		return fmt.Errorf("can't copy file: target file can't be blank")

	case !IsExist(from):
		return fmt.Errorf("can't copy file: file %q does not exists", from)
	case !IsRegular(from):
		return fmt.Errorf("can't copy file: file %q is not a regular file", from)
	case !IsReadable(from):
		return fmt.Errorf("can't copy file: file %q is not readable", from)

	case !targetExist && !IsExist(dir):
		return fmt.Errorf("can't copy file: directory %q does not exists", dir)
	case !targetExist && !IsWritable(dir):
		return fmt.Errorf("can't copy file: directory %q is not writable", dir)

	case targetExist && !IsWritable(to):
		return fmt.Errorf("can't copy file: target file %q is not writable", to)
	}

	var targetPerms os.FileMode

	if len(perms) == 0 {
		targetPerms = GetMode(from)
	} else {
		targetPerms = perms[0]
	}

	return copyFile(from, to, targetPerms)
}

// CopyAttr copies attributes (mode, ownership, timestamps) from one object
// (file or directory) to another
func CopyAttr(from, to string) error {
	switch {
	case from == "":
		return fmt.Errorf("can't copy attributes: source object can't be blank")
	case to == "":
		return fmt.Errorf("can't copy attributes: target object can't be blank")

	case !IsExist(from):
		return fmt.Errorf("can't copy attributes: %q does not exists", from)
	case !IsReadable(from):
		return fmt.Errorf("can't copy attributes: %q is not readable", from)
	case !IsExist(to):
		return fmt.Errorf("can't copy attributes: %q does not exists", to)
	case !IsWritable(to):
		return fmt.Errorf("can't copy attributes: %q is not writable", to)
	}

	return copyAttributes(from, to)
}

// MoveFile moves file
func MoveFile(from, to string, perms ...os.FileMode) error {
	targetExist := IsExist(to)

	if targetExist && IsDir(to) {
		to = path.Join(to, path.Base(from))
		targetExist = false
	}

	dir := path.Dir(to)

	switch {
	case from == "":
		return fmt.Errorf("can't move file: source file can't be blank")
	case to == "":
		return fmt.Errorf("can't move file: target file can't be blank")

	case !IsExist(from):
		return fmt.Errorf("can't move file: file %q does not exists", from)
	case !IsRegular(from):
		return fmt.Errorf("can't move file: file %q is not a regular file", from)
	case !IsReadable(from):
		return fmt.Errorf("can't move file: file %q is not readable", from)

	case !targetExist && !IsExist(dir):
		return fmt.Errorf("can't move file: directory %q does not exists", dir)
	case !targetExist && !IsWritable(dir):
		return fmt.Errorf("can't move file: directory %q is not writable", dir)
	}

	var targetPerms os.FileMode

	if len(perms) == 0 {
		targetPerms = GetMode(from)
	} else {
		targetPerms = perms[0]
	}

	return moveFile(from, to, targetPerms)
}

// CopyDir copies directory content recursively to target directory
func CopyDir(from, to string) error {
	switch {
	case from == "":
		return fmt.Errorf("can't copy directory: source directory can't be blank")
	case to == "":
		return fmt.Errorf("can't copy directory: target directory can't be blank")

	case !IsExist(from):
		return fmt.Errorf("can't copy directory: directory %q does not exists", from)
	case !IsDir(from):
		return fmt.Errorf("can't copy directory: target %q is not a directory", from)
	case !IsReadable(from):
		return fmt.Errorf("can't copy directory: directory %q is not readable", from)

	case IsExist(to) && !IsDir(to):
		return fmt.Errorf("can't copy directory: target %q is not a directory", to)
	case IsExist(to) && !IsWritable(to):
		return fmt.Errorf("can't copy directory: directory %q is not writable", to)
	}

	if !IsExist(to) {
		err := mkDirFunc(to, GetMode(from))

		if err != nil {
			return err
		}
	}

	return copyDir(from, to)
}

// TouchFile creates empty file
func TouchFile(path string, perm os.FileMode) error {
	fd, err := openFileFunc(path, os.O_CREATE, perm)

	if err != nil {
		return err
	}

	return fd.Close()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// copyFile copies file using bufio with given permissions
func copyFile(from, to string, perms os.FileMode) error {
	from, to = path.Clean(from), path.Clean(to)

	ffd, err := openFileFunc(from, os.O_RDONLY, 0)

	if err != nil {
		return err
	}

	defer ffd.Close()

	tfd, err := openFileFunc(to, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perms)

	if err != nil {
		return err
	}

	defer tfd.Close()

	r := bufio.NewReader(ffd)
	w := bufio.NewWriter(tfd)

	_, err = ioCopyFunc(w, r)

	if err != nil {
		return err
	}

	err = w.Flush()

	if err != nil {
		return err
	}

	return chmodFunc(to, perms)
}

// copyAttributes copies attributes (mode, ownership, timestamps) from one object
// (file or directory) to another
func copyAttributes(from, to string) error {
	from, to = path.Clean(from), path.Clean(to)

	fMode := modeFunc(from)

	if fMode == 0 {
		return fmt.Errorf("error while reading source object mode")
	}

	tMode := modeFunc(to)

	if tMode == 0 {
		return fmt.Errorf("error while reading target object mode")
	}

	fUID, fGID, err := ownerFunc(from)

	if err != nil {
		return err
	}

	tUID, tGID, err := ownerFunc(to)

	if err != nil {
		return err
	}

	fAtime, fMtime, _, err := timesFunc(from)

	if err != nil {
		return err
	}

	tAtime, tMtime, _, err := timesFunc(to)

	if err != nil {
		return err
	}

	if fMode != tMode {
		err = chmodFunc(to, fMode)

		if err != nil {
			return err
		}
	}

	if fUID != tUID || fGID != tGID {
		err = chownFunc(to, fUID, fGID)

		if err != nil {
			return err
		}
	}

	if fAtime != tAtime || fMtime != tMtime {
		err = chtimesFunc(to, fAtime, fMtime)

		if err != nil {
			return err
		}
	}

	return nil
}

// moveFile moves file with given permissions
func moveFile(from, to string, perms os.FileMode) error {
	from, to = path.Clean(from), path.Clean(to)

	err := os.Rename(from, to)

	if err != nil {
		return err
	}

	if perms == 0 {
		return nil
	}

	return chmodFunc(to, perms)
}

// copyDir copies directory content recursively to target directory
func copyDir(from, to string) error {
	from, to = path.Clean(from), path.Clean(to)

	var err error

	for _, target := range List(from, false) {
		fp := from + "/" + target
		tp := to + "/" + target

		if IsDir(fp) {
			err = mkDirFunc(tp, GetMode(fp))

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
