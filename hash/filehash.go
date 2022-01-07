// Package hash contains different hash algorithms and utilities
package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// FileHash generates an SHA-256 hash for a given file
func FileHash(file string) string {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return ""
	}

	defer fd.Close()

	hasher := sha256.New()

	io.Copy(hasher, fd)

	return fmt.Sprintf("%064x", hasher.Sum(nil))
}
