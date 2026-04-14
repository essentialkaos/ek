//go:build !linux && !darwin

// Package directio provides methods for reading/writing files with direct io
package directio

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

const (
	BLOCK_SIZE = 0 // ❗ Minimal block size
	ALIGN_SIZE = 0 // ❗ Align size
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ReadFile reads the file at the given path using Direct IO, bypassing the OS page
// cache
func ReadFile(file string) ([]byte, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ WriteFile writes data to the file at the given path using Direct IO, bypassing
// the OS page cache. The file is created with the given permissions if it does
// not exist.
func WriteFile(file string, data []byte, perms os.FileMode) error {
	panic("UNSUPPORTED")
	return nil
}
