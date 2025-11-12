//go:build !windows

// Package lock provides methods for working with lock files
package lock

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"time"

	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Create creates new lock file
func Create(name string) error {
	err := fsutil.ValidatePerms("DW", Dir)

	if err != nil {
		return err
	}

	return fsutil.TouchFile(getLockPath(name), 0644)
}

// Remove deletes lock file
func Remove(name string) error {
	err := fsutil.ValidatePerms("DW", Dir)

	if err != nil {
		return err
	}

	return os.Remove(getLockPath(name))
}

// Has returns true if lock file exists
func Has(name string) bool {
	return fsutil.IsExist(getLockPath(name))
}

// Wait waits until lock file being deleted
func Wait(name string, deadline time.Time) bool {
	if !Has(name) {
		return true
	}

	for range time.NewTicker(time.Second / 4).C {
		if !deadline.IsZero() && time.Now().After(deadline) {
			return false
		}

		if !Has(name) {
			break
		}
	}

	return true
}

// IsExpired returns true if lock file reached TTL
func IsExpired(name string, ttl time.Duration) bool {
	ct, err := fsutil.GetCTime(getLockPath(name))

	if err != nil {
		return false
	}

	return time.Since(ct) > ttl
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getLockPath returns path to lock file
func getLockPath(name string) string {
	return Dir + "/" + name + ".lock"
}
