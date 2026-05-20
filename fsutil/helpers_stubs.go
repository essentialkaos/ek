//go:build !linux && !darwin && !freebsd

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ CopyFile copies file using bufio
func CopyFile(from, to string, perms ...os.FileMode) error {
	panic("UNSUPPORTED")
}

// ❗ CopyAttr copies attributes (mode, ownership, timestamps) from one object
// (file or directory) to another
func CopyAttr(from, to string) error {
	panic("UNSUPPORTED")
}

// ❗ MoveFile moves file
func MoveFile(from, to string, perms ...os.FileMode) error {
	panic("UNSUPPORTED")
}

// ❗ CopyDir copies directory content recursively to target directory
func CopyDir(from, to string) error {
	panic("UNSUPPORTED")
}

// ❗ TouchFile creates empty file
func TouchFile(path string, perm os.FileMode) error {
	panic("UNSUPPORTED")
}

// ❗ CountLines returns the number of newline-delimited lines in the given file
func CountLines(file string) (int, error) {
	panic("UNSUPPORTED")
}
