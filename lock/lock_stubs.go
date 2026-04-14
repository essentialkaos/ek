//go:build !linux && !darwin

// Package lock provides methods for working with lock files
package lock

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Dir is the directory where lock files are created and looked up
var Dir = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Create creates a new lock file with the given name in [Dir]
func Create(name string) error {
	panic("UNSUPPORTED")
}

// ❗ Remove deletes the lock file with the given name from [Dir]
func Remove(name string) error {
	panic("UNSUPPORTED")
}

// ❗ Has reports whether a lock file with the given name currently exists
func Has(name string) bool {
	panic("UNSUPPORTED")
}

// ❗ Wait blocks until the named lock file is removed or the deadline is exceeded.
// Returns true if the lock was released, false if the deadline was reached.
func Wait(name string, deadline time.Time) bool {
	panic("UNSUPPORTED")
}

// ❗ IsExpired reports whether the named lock file has existed longer than TTL
func IsExpired(name string, ttl time.Duration) bool {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //
