package crypto

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

	_, err = io.Copy(hasher, fd)

	return fmt.Sprintf("%064x", hasher.Sum(nil))
}
