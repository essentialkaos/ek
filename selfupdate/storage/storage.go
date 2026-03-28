package storage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
	// ErrNilStorage indicates that a nil storage implementation was provided
	ErrNilStorage = errors.New("storage is nil")

	// ErrEmptyName indicates that the application name is empty
	ErrEmptyName = errors.New("app name is empty")

	// ErrEmptyVersion indicates that the application version is empty
	ErrEmptyVersion = errors.New("app version is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Storage describes a backend capable of checking for application updates
type Storage interface {
	// Check queries the storage for updates and returns update info and whether update
	// is available
	Check(app, version string) (selfupdate.Update, bool, error)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// OSName returns a normalized name of the current operating system
func OSName() string {
	if runtime.GOOS == "darwin" {
		return "macos"
	}

	return runtime.GOOS
}

// ArchName returns a normalized name of the current CPU architecture
func ArchName() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86_64"
	case "386":
		return "i386"
	}

	return runtime.GOARCH
}
