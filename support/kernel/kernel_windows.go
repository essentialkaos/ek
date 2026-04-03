// Package kernel provides methods for collecting information from OS kernel
package kernel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "github.com/essentialkaos/ek/v13/support"

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect returns kernel parameters matching the given names or prefix patterns.
// Patterns ending with "*" are treated as prefix globs (e.g. "vm.*" matches all vm
// parameters). Returns nil if no params match or if the kernel cannot be queried.
func Collect(params ...string) []support.KernelParam {
	return nil
}
