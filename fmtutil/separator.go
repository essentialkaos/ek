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

// SeparatorColorTag is the fmtc color tag applied to separator dashes
var SeparatorColorTag = "{s}"

// SeparatorTitleColorTag is the fmtc color tag applied to the separator title
var SeparatorTitleColorTag = "{s}"

// SeparatorFullscreen makes Separator span the full terminal width when true
var SeparatorFullscreen = false

// SeparatorSymbol is the character repeated to form the separator line
var SeparatorSymbol = "-"

// SeparatorSize is the fixed width of the separator when SeparatorFullscreen is false
var SeparatorSize = 88

// SeparatorTitleAlign controls title placement
var SeparatorTitleAlign = LEFT

// ////////////////////////////////////////////////////////////////////////////////// //

// Separator prints a horizontal separator to stdout.
// Pass tiny=true for a single line; tiny=false adds a blank line above and below.
// An optional first element of args is rendered as a title inside the separator.
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

// getAligned returns a full-width separator string of size with name positioned
// according to SeparatorTitleAlign, including fmtc color tags
func getAligned(size int, name string) string {
	var separator string

	lineSize := mathutil.Between((size-2)-strutil.LenVisual(name), 4, 999999)

	switch SeparatorTitleAlign {
	case CENTER:
		lineSize /= 2
		separator = SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize) + "{!} "
		separator += strutil.B(size%((lineSize*2)+strutil.LenVisual(name)+2) == 1, " ", "")
		separator += SeparatorTitleColorTag + name + "{!} "
		separator += SeparatorColorTag + strings.Repeat(SeparatorSymbol, lineSize) + "{!}"
	case RIGHT:
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
