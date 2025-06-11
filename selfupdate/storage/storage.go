package storage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"runtime"

	"github.com/essentialkaos/ek/v13/selfupdate"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilStorage is returned if storage is nil
	ErrNilStorage = errors.New("Storage is nil")

	// ErrEmptyName is returned if app name is empty
	ErrEmptyName = errors.New("App name is empty")

	// ErrEmptyVersion is returned if app version is empty
	ErrEmptyVersion = errors.New("App name is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Storage is an interface for update storage
type Storage interface {
	// Check checks the storage to see if there's an update for the app
	Check(app, version string) (selfupdate.Update, bool, error)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// OSName returns OS name
func OSName() string {
	if runtime.GOOS == "darwin" {
		return "macos"
	}

	return runtime.GOOS
}

// ArchName returns arch name
func ArchName() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86_64"
	case "386":
		return "i386"
	}

	return runtime.GOARCH
}
