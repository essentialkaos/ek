// Package pkgs provides methods for collecting information about filesystem
package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v14/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect returns usage statistics for all currently mounted filesystems.
// Returns nil if filesystem information cannot be retrieved from the OS.
func Collect() []support.FSInfo {
	return nil
}
