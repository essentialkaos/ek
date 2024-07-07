// Package panel provides methods for rendering panels with text
package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"

	"github.com/essentialkaos/ek.v13/fmtc"
	"github.com/essentialkaos/ek.v13/fmtutil"
	"github.com/essentialkaos/ek.v13/mathutil"
	"github.com/essentialkaos/ek.v13/sliceutil"
	"github.com/essentialkaos/ek.v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Option uint8

type Options []Option

const (
	// WRAP is panel rendering option for automatic text wrapping
	WRAP Option = iota + 1

	// INDENT_OUTER is panel rendering option for indent using new lines
	// before and after panel
	INDENT_OUTER

	// INDENT_INNER is panel rendering option for indent using new lines
	// before and after panel data
	INDENT_INNER

	// TOP_LINE is panel rendering option for drawing top line of panel
	TOP_LINE

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

// Indent is indent from the left side of terminal
var Indent = 0

// DefaultOptions is the default options used for rendering the panel if
// no options are passed
var DefaultOptions Options

// ////////////////////////////////////////////////////////////////////////////////// //

// minWidth is the minimal panel width
var minWidth = 38

// maxWidth is the maximum panel width
var maxWidth = 256

// maxIndent is the maximum indent value
var maxIndent = 24

// ////////////////////////////////////////////////////////////////////////////////// //

// Error shows panel with error message
func Error(title, message string, options ...Option) {
	Panel("ERROR", ErrorColorTag, title, message, options...)
}

// Warn shows panel with warning message
func Warn(title, message string, options ...Option) {
	Panel("WARNING", WarnColorTag, title, message, options...)
}

// Info shows panel with warning message
func Info(title, message string, options ...Option) {
	Panel("INFO", InfoColorTag, title, message, options...)
}

// Panel shows panel with given label, title, and message
func Panel(label, colorTag, title, message string, options ...Option) {
	label = strutil.Q(label, "•••")

	if len(options) == 0 {
		options = DefaultOptions
	}

	colorTag = strutil.B(fmtc.IsTag(colorTag), colorTag, "")

	renderPanel(label, colorTag, title, message, options)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if options slice contains given option
func (o Options) Has(option Option) bool {
	if len(o) == 0 {
		return false
	}

	return sliceutil.Contains(o, option)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// renderPanel renders panel
func renderPanel(label, colorTag, title, message string, options Options) {
	var buf *bytes.Buffer

	width := mathutil.Between(Width, minWidth, maxWidth)
	indent := strings.Repeat(" ", mathutil.Between(Indent, 0, maxIndent))

	if options.Has(INDENT_OUTER) {
		fmtc.NewLine()
	}

	labelFormat := "{@*} %s {!}"

	if fmtc.DisableColors {
		labelFormat = "[%s]"
	}

	if options.Has(LABEL_POWERLINE) && !fmtc.DisableColors {
		fmtc.Printf(colorTag+indent+labelFormat+colorTag+"{!} "+colorTag+"%s{!}", label, title)
	} else {
		fmtc.Printf(colorTag+indent+labelFormat+colorTag+" %s{!}", label, title)
	}

	if !options.Has(TOP_LINE) {
		fmtc.NewLine()
	} else {
		lineSize := width - (strutil.LenVisual(label+title) + 4)

		if options.Has(LABEL_POWERLINE) && !fmtc.DisableColors {
			lineSize--
		}

		if lineSize > 0 {
			fmtc.Printf(colorTag+" %s{!}\n", strings.Repeat("─", lineSize))
		} else {
			fmtc.NewLine()
		}
	}

	if options.Has(INDENT_INNER) || options.Has(BOTTOM_LINE) {
		fmtc.Println(colorTag + indent + "┃{!}")
	}

	switch {
	case options.Has(WRAP):
		buf = bytes.NewBufferString(
			fmtutil.Wrap(fmtc.Sprint(message), "", width-2) + "\n",
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

		fmtc.Print(colorTag + indent + "┃{!} " + line)
	}

	if options.Has(INDENT_INNER) {
		fmtc.Println(colorTag + indent + "┃{!}")
	}

	if options.Has(BOTTOM_LINE) {
		fmtc.Println(colorTag + indent + "┖" + strings.Repeat("─", width-1))
	}

	if options.Has(INDENT_OUTER) {
		fmtc.NewLine()
	}
}
