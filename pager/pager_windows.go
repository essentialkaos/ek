// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "errors"

// ////////////////////////////////////////////////////////////////////////////////// //

// DEFAULT is default pager command
const DEFAULT = "more"

// ////////////////////////////////////////////////////////////////////////////////// //

var ErrAlreadySet = errors.New("Pager already set")

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

// ❗ In most cases, you should use Setup and Complete because you can handle an
// error from Setup.
func Redirect(pager ...string) func() {
	return nil
}
