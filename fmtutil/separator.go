package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"pkg.re/essentialkaos/ek.v3/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SEPARATOR = "----------------------------------------------------------------------------------------"

// ////////////////////////////////////////////////////////////////////////////////// //

// SeparatorColorTag is color tag used for separator output (light grey by default)
var SeparatorColorTag string = "{s}"

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator print separator to output
func Separator(tiny bool, args ...string) {
	var sep = _SEPARATOR

	if len(args) != 0 {
		name := args[0]

		sep = "-- "
		sep += name
		sep += " "
		sep += _SEPARATOR[0:(84 - len(name))]
	}

	switch tiny {
	case true:
		fmtc.Printf(SeparatorColorTag+"%s{!}\n", sep)
	case false:
		fmtc.Printf("\n"+SeparatorColorTag+"%s{!}\n\n", sep)
	}
}
