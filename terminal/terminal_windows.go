package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ErrKillSignal is error type when user cancel input
var ErrKillSignal = errors.New("")

// ❗ Prompt is prompt string
var Prompt = "> "

// ❗ MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// ❗ MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
var MaskSymbolColorTag = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ReadUI reads user input
func ReadUI(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

// ❗ ReadAnswer reads user answer for yes/no question
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

// ❗ PrintErrorMessage prints error message
func PrintErrorMessage(message string, args ...interface{}) {
	panic("UNSUPPORTED")
}

// ❗ PrintWarnMessage prints warning message
func PrintWarnMessage(message string, args ...interface{}) {
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
func AddHstory(ui string) {
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
