//go:build !windows

// Package ek is a set of auxiliary packages
package ek

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/essentialkaos/depsy"
	"github.com/essentialkaos/go-linenoise/v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// worthless is used as dependency fix
func worthless() {
	linenoise.Clear()
	depsy.Extract(nil, false)
	bcrypt.Cost(nil)
}
