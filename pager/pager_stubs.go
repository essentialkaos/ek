//go:build !linux && !darwin && !freebsd

// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "errors"

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ❗ ErrAlreadySet is returned by [Setup] if the pager has already been initialized
	ErrAlreadySet = errors.New("pager already set")

	// ❗ ErrNoPager is returned by [Setup] if no supported pager binary (less, more)
	// was found on the system and no explicit pager was provided
	ErrNoPager = errors.New("no pager found on the system")

	// ❗ ErrStdinPipe is returned by [Setup] if the pager process stdin pipe could not
	// be obtained as an *os.File, which is required for stdout redirection
	ErrStdinPipe = errors.New("can't get pager stdin pipe")

	// ❗ ErrPagerError is returned by [Complete] if the pager process exited with a
	// non-zero exit code
	ErrPagerError = errors.New("pager exited with an error")
)

// ❗ AllowEnv allows the user to define the pager binary via the PAGER environment
// variable
var AllowEnv bool

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Setup redirects os.Stdout and os.Stderr through the given pager process.
// If no pager is provided, less or more is located automatically.
func Setup(pager ...string) error {
	return nil
}

// ❗ Complete closes the pager stdin pipe, waits for the pager process to exit,
// and restores [os.Stdout] and [os.Stderr] to their original values
func Complete() {
	return
}
