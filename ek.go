//go:build !windows
// +build !windows

// Package ek is a set of auxiliary packages
package ek

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/essentialkaos/go-linenoise/v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// VERSION is current ek package version
const VERSION = "12.95.0"

// ////////////////////////////////////////////////////////////////////////////////// //

// worthless is used as dependency fix
func worthless() {
	linenoise.Clear()
	bcrypt.Cost(nil)
}
