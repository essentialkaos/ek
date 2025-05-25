// Package hash contains different hash algorithms and utilities
//
// Deprecated: Use package hashutil instead
package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha256"

	"github.com/essentialkaos/ek/v13/hashutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// FileHash generates an SHA-256 hash for a given file
//
// Deprecated: Use package hashutil instead
func FileHash(file string) string {
	return hashutil.File(file, sha256.New())
}
