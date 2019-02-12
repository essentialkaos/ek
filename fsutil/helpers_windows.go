package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CopyFile simple file copying with bufio
func CopyFile(from, to string, perms ...os.FileMode) error {
	return nil
}

// MoveFile move file
func MoveFile(from, to string, perms ...os.FileMode) error {
	return nil
}

// CopyDir copy directory content recursively to target directory
func CopyDir(from, to string) error {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
