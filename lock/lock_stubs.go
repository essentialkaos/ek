//go:build !linux && !darwin

// Package lock provides methods for working with lock files
package lock

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is a path to directory with lock files
var Dir = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Create creates new lock file
func Create(name string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ Remove deletes lock file
func Remove(name string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ Has returns true if lock file exists
func Has(name string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ Expired returns true if lock file reached TTL
func Expired(name string, ttl time.Duration) bool {
	panic("UNSUPPORTED")
	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //
