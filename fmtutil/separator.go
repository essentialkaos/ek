package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/terminal/tty"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SeparatorColorTag is fmtc color tag used for separator (light gray by default)
var SeparatorColorTag = "{s}"

// SeparatorTitleColorTag is fmtc color tag used for separator title (light gray by default)
var SeparatorTitleColorTag = "{s}"

// SeparatorFullscreen allow enabling full-screen separator
var SeparatorFullscreen = false

// SeparatorSymbol used for separator generation
var SeparatorSymbol = "-"

// SeparatorSize contains size of separator
var SeparatorSize = 88

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator renders separator to console
func Separator(tiny bool, args ...string) {
	var separator string
	var size int

	if SeparatorFullscreen {
		size = between(tty.GetWidth(), 16, 999999)
	} else {
		size = between(SeparatorSize, 80, 999999)
	}

	if len(args) != 0 {
		name := args[0]
		suffixSize := between((size-4)-len(name), 0, 999999)

		separator += SeparatorColorTag
		separator += strings.Repeat(SeparatorSymbol, 2) + "{!} "
		separator += SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + strings.Repeat(SeparatorSymbol, suffixSize) + "{!}"
	} else {
		separator = SeparatorColorTag + getSeparator(size) + "{!}"
	}

	if !tiny {
		separator = "\n" + separator + "\n"
	}

	fmtc.Println(separator)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getSeparator(size int) string {
	return strings.Repeat(SeparatorSymbol, size)
}

func between(val, min, max int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}
