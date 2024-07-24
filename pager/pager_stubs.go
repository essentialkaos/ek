//go:build !linux && !darwin && !freebsd
// +build !linux,!darwin,!freebsd

// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "errors"

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrAlreadySet = errors.New("Pager already set")
	ErrNoPager    = errors.New("There is no pager on the system")
	ErrStdinPipe  = errors.New("Can't get pager stdin")
)

// AllowEnv is a flag that allows to user to define pager binary using PAGER environment
// variable
var AllowEnv bool

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Setup set up pager for work. After calling this method, any data sent to Stdout and
// Stderr (using fmt, fmtc, or terminal packages) will go to the pager.
func Setup(pager ...string) error {
	return nil
}

// ❗ Complete finishes output redirect to pager
func Complete() {
	return
}
