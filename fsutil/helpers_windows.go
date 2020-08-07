package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
