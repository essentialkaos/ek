// Package hashutil contains various helper functions for working with hashes
package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"hash"
	"io"
	"os"
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// File calculates hash of the file using given hasher
func File(file string, hasher hash.Hash) string {
	if hasher == nil {
		return ""
	}

	fd, err := os.Open(file)

	if err != nil {
		return ""
	}

	defer fd.Close()

	io.Copy(hasher, fd)

	return getSum(hasher)
}

// Bytes calculates data hash using given hasher
func Bytes(data []byte, hasher hash.Hash) string {
	if len(data) == 0 || hasher == nil {
		return ""
	}

	hasher.Write(data)

	return getSum(hasher)
}

// String calculates string hash using given hasher
func String(data string, hasher hash.Hash) string {
	if len(data) == 0 || hasher == nil {
		return ""
	}

	fmt.Fprint(hasher, data)

	return getSum(hasher)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getSum returns string representation of hash
func getSum(hasher hash.Hash) string {
	return fmt.Sprintf("%0"+strconv.Itoa(hasher.Size()/2)+"x", hasher.Sum(nil))
}
