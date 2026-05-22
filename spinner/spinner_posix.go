//go:build linux || darwin

package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v14/mathutil"
	"github.com/essentialkaos/ek/v14/terminal/tty"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getMaxDescSize returns the maximum character width available for the description text
func getMaxDescSize() int {
	w := tty.GetWidth()
	return mathutil.B(w < 20, 9999, w-14)
}
