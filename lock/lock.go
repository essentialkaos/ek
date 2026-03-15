//go:build !windows

// Package lock provides methods for working with lock files
package lock

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path/filepath"
	"time"

	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Create creates a new lock file with the given name in [Dir]
func Create(name string) error {
	err := fsutil.ValidatePerms("DW", Dir)

	if err != nil {
		return err
	}

	return fsutil.TouchFile(getLockPath(name), 0644)
}

// Remove deletes the lock file with the given name from [Dir]
func Remove(name string) error {
	err := fsutil.ValidatePerms("DW", Dir)

	if err != nil {
		return err
	}

	return os.Remove(getLockPath(name))
}

// Has reports whether a lock file with the given name currently exists
func Has(name string) bool {
	return fsutil.IsExist(getLockPath(name))
}

// Wait blocks until the named lock file is removed or the deadline is exceeded.
// Returns true if the lock was released, false if the deadline was reached.
func Wait(name string, deadline time.Time) bool {
	if !Has(name) {
		return true
	}

	ticker := time.NewTicker(time.Second / 4)
	defer ticker.Stop()

	for range ticker.C {
		if !deadline.IsZero() && time.Now().After(deadline) {
			return false
		}

		if !Has(name) {
			break
		}
	}

	return true
}

// IsExpired reports whether the named lock file has existed longer than TTL
func IsExpired(name string, ttl time.Duration) bool {
	ct, err := fsutil.GetCTime(getLockPath(name))

	if err != nil {
		return false
	}

	return time.Since(ct) > ttl
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLockPath returns the full filesystem path for the named lock file
func getLockPath(name string) string {
	return filepath.Join(Dir, name+".lock")
}
