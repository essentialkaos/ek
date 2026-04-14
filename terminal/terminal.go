// Package terminal provides methods for printing messages to terminal
package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrorColorTag is the fmtc color tag applied to error messages
	ErrorColorTag = "{r}"

	// WarnColorTag is the fmtc color tag applied to warning messages
	WarnColorTag = "{y}"

	// InfoColorTag is fmtc color tag used for info messages
	InfoColorTag = "{c-}"
)

var (
	// ErrorPrefix is the string prepended to every error message
	ErrorPrefix = ""

	// WarnPrefix is the string prepended to every warning message
	WarnPrefix = ""

	// InfoPrefix is the string prepended to every info message
	InfoPrefix = ""
)

// ////////////////////////////////////////////////////////////////////////////////// //

// PrintActionMessage prints a bold action description followed by a space,
// intended to be paired with a subsequent PrintActionStatus call
func PrintActionMessage(message string) {
	fmtc.Printf("{*}%s:{!} ", message)
}

// PrintActionStatus prints a colored status label (OK, ERROR, WARNING, or UNKNOWN)
// corresponding to the given numeric status code
func PrintActionStatus(status int) {
	switch status {
	case 0:
		fmtc.Println("{g}OK{!}")
	case 1:
		fmtc.Println("{r}ERROR{!}")
	case 2:
		fmtc.Println("{y}WARNING{!}")
	default:
		fmtc.Println("{s}UNKNOWN{!}")
	}
}

// Error prints a formatted error message in red to stderr
func Error(message any, args ...any) {
	fmtc.Fprintf(
		os.Stderr, ErrorColorTag+ErrorPrefix+"%s{!}\n",
		formatMessage(message, ErrorPrefix, args),
	)
}

// Warn prints a formatted warning message in yellow to stderr
func Warn(message any, args ...any) {
	fmtc.Fprintf(
		os.Stderr, WarnColorTag+WarnPrefix+"%s{!}\n",
		formatMessage(message, WarnPrefix, args),
	)
}

// Info prints a formatted informational message in cyan to stdout
func Info(message any, args ...any) {
	fmtc.Fprintf(
		os.Stdout, InfoColorTag+InfoPrefix+"%s{!}\n",
		formatMessage(message, InfoPrefix, args),
	)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// formatMessage formats message based on it type
func formatMessage(message any, prefix string, args []any) string {
	var msg string

	switch m := message.(type) {
	case string:
		msg = m
	case error:
		msg = m.Error()
	case fmt.Stringer:
		msg = m.String()
	default:
		msg = fmt.Sprint(message)
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	if prefix != "" && strings.Contains(msg, "\n") {
		prefixOffset := strings.Repeat(" ", strutil.Len(prefix))
		msgText := strings.TrimRight(msg, "\n")
		tail := strutil.Substr(msg, len(msgText), 999999)
		msg = strings.ReplaceAll(msgText, "\n", "\n"+prefixOffset) + tail
	}

	return msg
}
