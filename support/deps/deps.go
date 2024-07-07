// Package pkgs provides methods for collecting information about used dependencies
package deps

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek.v13/support"

	"github.com/essentialkaos/depsy"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Extract extracts dependencies info from gomod data
func Extract(gomod []byte) []support.Dep {
	var result []support.Dep

	for _, dep := range depsy.Extract(gomod, false) {
		result = append(result, support.Dep{
			Version: dep.Version,
			Path:    dep.PrettyPath(),
			Extra:   dep.Extra,
		})
	}

	return result
}
