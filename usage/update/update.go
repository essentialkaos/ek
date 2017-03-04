// Package update provides methods for checking application updates
package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"pkg.re/essentialkaos/ek.v6/version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// isNewerVersion return true if latest version is greater than current
func isNewerVersion(current, latest string) bool {
	v1, err := version.Parse(current)

	if err != nil {
		return false
	}

	v2, err := version.Parse(latest)

	if err != nil {
		return false
	}

	return v2.Greater(v1)
}
