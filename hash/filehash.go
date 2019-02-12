// Package hash contains different hash algorithms and utilities
package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// FileHash generate SHA-256 hash for file
func FileHash(file string) string {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0644)

	if err != nil {
		return ""
	}

	defer fd.Close()

	hasher := sha256.New()

	io.Copy(hasher, fd)

	return fmt.Sprintf("%064x", hasher.Sum(nil))
}
