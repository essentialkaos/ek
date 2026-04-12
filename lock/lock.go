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
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	if !isValidLockName(name) {
		return fmt.Errorf("invalid lock name %q", name)
	}

	return fsutil.TouchFile(getLockPath(name), 0644)
}

// Remove deletes the lock file with the given name from [Dir]
func Remove(name string) error {
	err := fsutil.ValidatePerms("DW", Dir)

	if err != nil {
		return err
	}

	if !isValidLockName(name) {
		return fmt.Errorf("invalid lock name %q", name)
	}

	return os.Remove(getLockPath(name))
}

// Has reports whether a lock file with the given name currently exists
func Has(name string) bool {
	if !isValidLockName(name) {
		return false
	}

	return fsutil.IsExist(getLockPath(name))
}

// Wait blocks until the named lock file is removed or the deadline is exceeded.
// Returns true if the lock was released, false if the deadline was reached.
// Pass a zero time.Time{} deadline to wait indefinitely.
func Wait(name string, deadline time.Time) bool {
	if !Has(name) {
		return true
	}

	ticker := time.NewTicker(time.Second / 4)
	defer ticker.Stop()

	for range ticker.C {
		if !Has(name) {
			break
		}

		if !deadline.IsZero() && time.Now().After(deadline) {
			return false
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

// isValidLockName returns false if given lock name is invalid
func isValidLockName(name string) bool {
	switch {
	case name == "..", name == ".", strings.ContainsAny(name, "/\\"):
		return false
	}

	return true
}

// getLockPath returns the full filesystem path for the named lock file
func getLockPath(name string) string {
	return filepath.Join(Dir, filepath.Base(name)+".lock")
}
