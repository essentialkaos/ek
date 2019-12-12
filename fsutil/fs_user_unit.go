// +build unit

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"pkg.re/essentialkaos/ek.v11/system"
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
