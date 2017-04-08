package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"

	"pkg.re/essentialkaos/ek.v7/fmtc"
	"pkg.re/essentialkaos/ek.v7/terminal/window"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SEPARATOR = "----------------------------------------------------------------------------------------"

// ////////////////////////////////////////////////////////////////////////////////// //

// SeparatorColorTag is fmtc color tag used for separator (light grey by default)
var SeparatorColorTag string = "{s}"

// SeparatorTitleColorTag is fmtc color tag used for separator title (light grey by default)
var SeparatorTitleColorTag = "{s}"

// FullscreenSeparator allow to enable full screen separator
var FullscreenSeparator = false

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator print separator to output
func Separator(tiny bool, args ...string) {
	var separator string

	if len(args) != 0 {
		name := args[0]
		sep := getSeparator()
		rem := (len(sep) - 4) - len(name)
		separator = SeparatorColorTag + "--{!} " + SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + sep[:rem] + "{!}"
	} else {
		separator = SeparatorColorTag + getSeparator() + "{!}"
	}

	if !tiny {
		separator = "\n" + separator + "\n"
	}

	fmtc.Println(separator)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getSeparator() string {
	if !FullscreenSeparator {
		return _SEPARATOR
	}

	width := window.GetWidth()

	if width == -1 {
		return _SEPARATOR
	}

	return strings.Repeat("-", width)
}
