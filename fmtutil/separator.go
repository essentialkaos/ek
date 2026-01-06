package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
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

// SeparatorTitleAlign aligning of separator title (l/left, c/center/, r/right)
var SeparatorTitleAlign = "left"

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator renders separator to console
func Separator(tiny bool, args ...string) {
	var separator string
	var size int

	if SeparatorFullscreen {
		size = mathutil.Between(tty.GetWidth(), 16, 999999)
	} else {
		size = mathutil.Between(SeparatorSize, 80, 999999)
	}

	if len(args) != 0 {
		separator = SeparatorColorTag + getAligned(size, args[0]) + "{!}"
	} else {
		separator = SeparatorColorTag + strings.Repeat(SeparatorSymbol, size) + "{!}"
	}

	if !tiny {
		separator = "\n" + separator + "\n"
	}

	fmtc.Println(separator)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getAligned(size int, name string) string {
	var separator string

	lineSize := mathutil.Between((size-2)-strutil.LenVisual(name), 4, 999999)

	switch strings.ToLower(SeparatorTitleAlign) {
	case "c", "center":
		lineSize /= 2
		separator = SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize) + "{!} "
		separator += strutil.B(size%((lineSize*2)+strutil.LenVisual(name)+2) == 1, " ", "")
		separator += SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize) + "{!}"
	case "r", "right":
		separator = SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize-2) + "{!} "
		separator += SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + strings.Repeat(SeparatorSymbol, 2) + "{!}"
	default:
		separator = SeparatorColorTag + strings.Repeat(SeparatorSymbol, 2) + "{!} "
		separator += SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize-2) + "{!}"
	}

	return separator
}
