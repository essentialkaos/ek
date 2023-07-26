// Package panel provides methods for rendering panels with text
package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/sliceutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Option uint8

const (
	// WRAP is panel rendering option for automatic text wrapping
	WRAP Option = iota + 1

	// INDENT_OUTER is panel rendering option for indent using new lines
	// before and after panel
	INDENT_OUTER

	// INDENT_INNER is panel rendering option for indent using new lines
	// before and after panel data
	INDENT_INNER

	// BOTTOM_LINE is panel rendering option for drawing bottom line of panel
	BOTTOM_LINE

	// LABEL_POWERLINE is panel rendering option for using powerline symbols
	LABEL_POWERLINE
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrorColorTag is fmtc color tag used for error messages
	ErrorColorTag = "{r}"

	// WarnColorTag is fmtc color tag used for warning messages
	WarnColorTag = "{y}"

	// InfoColorTag is fmtc color tag used for info messages
	InfoColorTag = "{c-}"
)

// Width is panel width (≥ 40) if option WRAP is set
var Width = 88

// ////////////////////////////////////////////////////////////////////////////////// //

// minWidth is minimal panel width
var minWidth = 38

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrorPanel shows panel with error message
func ErrorPanel(title, message string, options ...Option) {
	Panel("ERROR", ErrorColorTag, title, message, options...)
}

// WarnPanel shows panel with warning message
func WarnPanel(title, message string, options ...Option) {
	Panel("WARNING", WarnColorTag, title, message, options...)
}

// InfoPanel shows panel with warning message
func InfoPanel(title, message string, options ...Option) {
	Panel("INFO", InfoColorTag, title, message, options...)
}

// Panel show panel with given label, title, and message
func Panel(label, colorTag, title, message string, options ...Option) {
	var buf *bytes.Buffer

	if sliceutil.Contains(options, INDENT_OUTER) {
		fmtc.NewLine()
	}

	if sliceutil.Contains(options, LABEL_POWERLINE) {
		fmtc.Printf(colorTag+"{@*} %s {!}"+colorTag+"{!} "+colorTag+"%s{!}\n", label, title)
	} else {
		fmtc.Printf(colorTag+"{@*} %s {!} "+colorTag+"%s{!}\n", label, title)
	}

	if sliceutil.Contains(options, INDENT_INNER) {
		fmtc.Println(colorTag + "┃{!}")
	}

	switch {
	case sliceutil.Contains(options, WRAP):
		buf = bytes.NewBufferString(
			fmtutil.Wrap(fmtc.Sprint(message), "", mathutil.Max(minWidth, Width-2)) + "\n",
		)
	default:
		buf = bytes.NewBufferString(message)
		buf.WriteRune('\n')
	}

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		fmtc.Print(colorTag + "┃{!} " + line)
	}

	if sliceutil.Contains(options, INDENT_INNER) {
		fmtc.Println(colorTag + "┃{!}")
	}

	if sliceutil.Contains(options, BOTTOM_LINE) {
		fmtc.Println(colorTag + "┖" + strings.Repeat("─", mathutil.Max(minWidth, Width-1)))
	}

	if sliceutil.Contains(options, INDENT_OUTER) {
		fmtc.NewLine()
	}
}
