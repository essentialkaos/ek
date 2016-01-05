package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

// CopyFile simple file copying with bufio
func CopyFile(from, to string, perms ...os.FileMode) error {
	dir := path.Dir(to)

	switch {
	case IsExist(from) == false:
		return errors.New("File " + from + " is not exists")
	case IsRegular(from) == false:
		return errors.New("File " + from + " is not a regular file")
	case IsReadable(from) == false:
		return errors.New("File " + from + " is not readable")
	case IsExist(dir) == false:
		return errors.New("Directory " + from + " is not exists")
	case IsWritable(dir) == false:
		return errors.New("Directory " + from + " is not writable")
	}

	if IsExist(to) {
		os.Remove(to)
	}

	var perm os.FileMode = 0644

	if len(perms) != 0 {
		perm = perms[0]
	}

	ffd, err := os.OpenFile(from, os.O_RDONLY, 0644)

	if err != nil {
		return err
	}

	defer ffd.Close()

	tfd, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY, perm)

	if err != nil {
		return err
	}

	defer tfd.Close()

	reader := bufio.NewReader(ffd)

	_, err = io.Copy(tfd, reader)

	return err
}
