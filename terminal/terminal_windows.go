package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"

	"github.com/essentialkaos/ek/v12/secstr"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ErrKillSignal is error type when user cancel input
var ErrKillSignal = errors.New("")

// ❗ Prompt is prompt string
var Prompt = "> "

// ❗ MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// ❗ HideLength is flag for hiding password length
var HideLength = false

// ❗ HidePassword is flag for hiding password while typing
// Because of using the low-level linenoise method for this feature, we can not use a
// custom masking symbol, so it always will be an asterisk (*).
var HidePassword = false

var (
	// ❗ MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
	MaskSymbolColorTag = ""

	// ❗ TitleColorTag is fmtc color tag used for input titles
	TitleColorTag = "{s}"

	// ❗ ErrorColorTag is fmtc color tag used for error messages
	ErrorColorTag = "{r}"

	// ❗ WarnColorTag is fmtc color tag used for warning messages
	WarnColorTag = "{y}"

	// ❗ InfoColorTag is fmtc color tag used for info messages
	InfoColorTag = "{c-}"
)

var (
	// ❗ ErrorPrefix is prefix for error messages
	ErrorPrefix = ""

	// ❗ WarnPrefix is prefix for warning messages
	WarnPrefix = ""

	// ❗ InfoPrefix is prefix for info messages
	InfoPrefix = ""
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Read reads user's input
func Read(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

// ❗ ReadAnswer reads user's answer for yes/no question
func ReadAnswer(title, defaultAnswer string) (bool, error) {
	panic("UNSUPPORTED")
	return true, nil
}

// ❗ ReadPassword reads password or some private input which will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

// ❗ ReadPasswordSecure reads password or some private input which will be hidden
// after pressing Enter
func ReadPasswordSecure(title string, nonEmpty bool) (*secstr.String, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ PrintActionMessage prints message about action currently in progress
func PrintActionMessage(message string) {
	panic("UNSUPPORTED")
}

// ❗ PrintActionStatus prints message with action execution status
func PrintActionStatus(status int) {
	panic("UNSUPPORTED")
}

// ❗ Error prints error message
func Error(message string, args ...any) {
	panic("UNSUPPORTED")
}

// ❗ Warn prints warning message
func Warn(message string, args ...any) {
	panic("UNSUPPORTED")
}

// ❗ Info prints info message
func Info(message string, args ...any) {
	panic("UNSUPPORTED")
}

// ❗ AddHistory adds line to input history
func AddHistory(ui string) {
	panic("UNSUPPORTED")
}

// ❗ SetCompletionHandler adds function for autocompletion
func SetCompletionHandler(h func(in string) []string) {
	panic("UNSUPPORTED")
}

// ❗ SetHintHandler adds function for input hints
func SetHintHandler(h func(input string) string) {
	panic("UNSUPPORTED")
}
