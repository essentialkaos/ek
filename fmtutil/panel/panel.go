// Package panel provides methods for rendering panels with text in terminal
package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"slices"
	"strings"

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/fmtutil"
	"github.com/essentialkaos/ek/v14/mathutil"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	// WRAP enables automatic text wrapping based on panel width
	WRAP Option = iota + 1

	// INDENT_OUTER adds empty lines before and after the panel
	INDENT_OUTER

	// INDENT_INNER adds empty lines before and after the panel content
	INDENT_INNER

	// TOP_LINE draws a horizontal line after the panel label row
	TOP_LINE

	// BOTTOM_LINE draws a horizontal line at the bottom of the panel
	BOTTOM_LINE

	// LABEL_POWERLINE uses powerline symbols for the label rendering
	LABEL_POWERLINE
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Option represents a single panel rendering option
type Option uint8

// Options is a slice of panel rendering options
type Options []Option

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrorColorTag is the fmtc color tag used for error panels
	ErrorColorTag = "{r}"

	// WarnColorTag is the fmtc color tag used for warning panels
	WarnColorTag = "{y}"

	// InfoColorTag is the fmtc color tag used for info panels
	InfoColorTag = "{c-}"
)

// Width is the panel content width in characters (≥ 40), applied when WRAP is set
var Width = 88

// Indent is the number of spaces prepended to each panel line as left-side padding
var Indent = 0

// DefaultOptions holds the options applied when no options are passed to [Panel]
var DefaultOptions Options

// ////////////////////////////////////////////////////////////////////////////////// //

// minWidth is the minimum allowed panel width
var minWidth = 38

// maxWidth is the maximum allowed panel width
var maxWidth = 256

// maxIndent is the maximum allowed indent size
var maxIndent = 24

// ////////////////////////////////////////////////////////////////////////////////// //

// Error renders a panel with the ERROR label using [ErrorColorTag]
func Error(title, message string, options ...Option) {
	Panel("ERROR", ErrorColorTag, title, message, options...)
}

// Warn renders a panel with the WARNING label using [WarnColorTag]
func Warn(title, message string, options ...Option) {
	Panel("WARNING", WarnColorTag, title, message, options...)
}

// Info renders a panel with the INFO label using [InfoColorTag]
func Info(title, message string, options ...Option) {
	Panel("INFO", InfoColorTag, title, message, options...)
}

// Panel renders a panel with the given label, color tag, title, and message
func Panel(label, colorTag, title, message string, options ...Option) {
	label = strutil.Q(label, "•••")

	if len(options) == 0 {
		options = DefaultOptions
	}

	colorTag = strutil.B(fmtc.IsTag(colorTag), colorTag, "")

	renderPanel(label, colorTag, title, message, options)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if the options slice contains the given option
func (o Options) Has(option Option) bool {
	if len(o) == 0 {
		return false
	}

	return slices.Contains(o, option)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// renderPanel writes the formatted panel to the terminal using fmtc primitives
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
		fmtc.Printf(colorTag+indent+labelFormat+colorTag+"{!}", label)
	} else {
		fmtc.Printf(colorTag+indent+labelFormat+"{!}", label)
	}

	if title != "" {
		fmtc.Printf(" "+colorTag+"%s{!}", title)
	}

	if !options.Has(TOP_LINE) {
		fmtc.NewLine()
	} else {
		lineSize := width - (strutil.LenVisual(label+title) + 4)

		if options.Has(LABEL_POWERLINE) && !fmtc.DisableColors {
			lineSize--
		}

		if lineSize > 0 {
			fmtc.Printfn(colorTag+" %s{!}", strings.Repeat("─", lineSize))
		} else {
			fmtc.NewLine()
		}
	}

	if options.Has(INDENT_INNER) {
		fmtc.Println(colorTag + indent + "┃{!}")
	}

	if options.Has(WRAP) {
		buf = bytes.NewBufferString(
			fmtutil.Wrap(fmtc.Sprint(message), "", width-2) + "\n",
		)
	} else {
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
