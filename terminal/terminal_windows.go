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

// ❗ MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
var MaskSymbolColorTag = ""

// ❗ TitleColorTag is fmtc color tag used for input titles
var TitleColorTag = "{c}"

// ❗ ErrorColorTag is fmtc color tag used for error messages
var ErrorColorTag = "{r}"

// ❗ WarnColorTag is fmtc color tag used for warning messages
var WarnColorTag = "{y}"

// ❗ ErrorPrefix is prefix for error messages
var ErrorPrefix = ""

// ❗ WarnPrefix is prefix for warning messages
var WarnPrefix = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Read reads user's input
func Read(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

// ❗ ReadUI reads user's input
//
// Deprecated: Use method Read instead
func ReadUI(title string, nonEmpty bool) (string, error) {
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
	return "", nil
}

// ❗ PrintErrorMessage prints error message
func PrintErrorMessage(message string, args ...any) {
	panic("UNSUPPORTED")
}

// ❗ PrintWarnMessage prints warning message
func PrintWarnMessage(message string, args ...any) {
	panic("UNSUPPORTED")
}

// ❗ PrintActionMessage prints message about action currently in progress
func PrintActionMessage(message string) {
	panic("UNSUPPORTED")
}

// ❗ PrintActionStatus prints message with action execution status
func PrintActionStatus(status int) {
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
