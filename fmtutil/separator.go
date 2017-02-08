package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"pkg.re/essentialkaos/ek.v6/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SEPARATOR = "----------------------------------------------------------------------------------------"

// ////////////////////////////////////////////////////////////////////////////////// //

// SeparatorColorTag is fmtc color tag used for separator (light grey by default)
var SeparatorColorTag string = "{s}"

// SeparatorTitleColorTag is fmtc color tag used for separator title (light grey by default)
var SeparatorTitleColorTag = "{s}"

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator print separator to output
func Separator(tiny bool, args ...string) {
	sep := SeparatorColorTag + _SEPARATOR + "{!}"

	if len(args) != 0 {
		name := args[0]
		sep = SeparatorColorTag + "-- {!}" + SeparatorTitleColorTag + name + "{!} " + SeparatorColorTag + _SEPARATOR[:(84-len(name))] + "{!}"
	}

	if !tiny {
		sep = "\n" + sep + "\n"
	}

	fmtc.Println(sep)
}
