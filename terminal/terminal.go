// Package terminal provides methods for printing messages to terminal
package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/strutil"
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

var (
	// ErrorPrefix is prefix for error messages
	ErrorPrefix = ""

	// WarnPrefix is prefix for warning messages
	WarnPrefix = ""

	// InfoPrefix is prefix for info messages
	InfoPrefix = ""
)

// ////////////////////////////////////////////////////////////////////////////////// //

// PrintActionMessage prints message about action currently in progress
func PrintActionMessage(message string) {
	fmtc.Printf("{*}%s:{!} ", message)
}

// PrintActionStatus prints message with action execution status
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

// Error prints error message
func Error(message any, args ...any) {
	fmtc.Fprintf(
		os.Stdout, ErrorColorTag+ErrorPrefix+"%s{!}\n",
		formatMessage(message, ErrorPrefix, args),
	)
}

// Warn prints warning message
func Warn(message any, args ...any) {
	fmtc.Fprintf(
		os.Stdout, WarnColorTag+WarnPrefix+"%s{!}\n",
		formatMessage(message, WarnPrefix, args),
	)
}

// Info prints info message
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
