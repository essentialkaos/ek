package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
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
	return nil
}

// ❗ CopyAttr copies attributes (mode, ownership, timestamps) from one object
// (file or directory) to another
func CopyAttr(from, to string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ MoveFile moves file
func MoveFile(from, to string, perms ...os.FileMode) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ CopyDir copies directory content recursively to target directory
func CopyDir(from, to string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ TouchFile creates empty file
func TouchFile(path string, perm os.FileMode) error {
	panic("UNSUPPORTED")
	return nil
}
