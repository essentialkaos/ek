// +build !windows

// Package ek is set of auxiliary packages
package ek

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/crypto/bcrypt"

	"pkg.re/essentialkaos/go-linenoise.v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// VERSION is current ek package version
const VERSION = "12.13.0"

// ////////////////////////////////////////////////////////////////////////////////// //

// worthless is used as dependency fix
func worthless() {
	linenoise.Clear()
	bcrypt.Cost(nil)
}
