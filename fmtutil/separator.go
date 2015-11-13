package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _SEPARATOR = "----------------------------------------------------------------------------------------"

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
		fmtc.Printf("{s}%s{!}\n", sep)
	case false:
		fmtc.Printf("\n{s}%s{!}\n\n", sep)
	}

}
