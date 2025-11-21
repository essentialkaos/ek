//go:build !unit

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getCurrentUser returns the current user information
func getCurrentUser() (*system.User, error) {
	return system.CurrentUser()
}
