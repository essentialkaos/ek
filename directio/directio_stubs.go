//go:build !linux || !darwin
// +build !linux !darwin

// Package directio provides methods for reading/writing files with direct io
package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

// ❗ ReadFile read file with Direct IO without buffering data in page cache
func ReadFile(file string) ([]byte, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ WriteFile write file with Direct IO without buffering data in page cache
func WriteFile(file string, data []byte, perms os.FileMode) error {
	panic("UNSUPPORTED")
	return nil
}
