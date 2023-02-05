//go:build unit
// +build unit

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"github.com/essentialkaos/ek/v12/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var useFakeUser bool
var getUserError bool

// ////////////////////////////////////////////////////////////////////////////////// //

func getCurrentUser() (*system.User, error) {
	if useFakeUser {
		return &system.User{
			Name:    "test",
			UID:     65534,
			GID:     65534,
			RealUID: 65534,
			RealGID: 65534,
			HomeDir: "/unknown",
		}, nil
	}

	if getUserError {
		return nil, errors.New("Error")
	}

	return system.CurrentUser()
}
