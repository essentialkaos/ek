//go:build !linux || !darwin
// +build !linux !darwin

// Package input provides methods for reading user input
package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"

	"github.com/essentialkaos/ek/v13/secstr"
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
var TitleColorTag = "{s}"

// ❗ HidePassword is flag for hiding password while typing
// Because of using the low-level linenoise method for this feature, we can not use a
// custom masking symbol, so it always will be an asterisk (*).
var HidePassword = false

// ❗ AlwaysYes is a flag, if set ReadAnswer will always return true (useful for working
// with option for forced actions)
var AlwaysYes = false

// ❗ NewLine is a flag for extra new line after inputs
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Read reads user input
func Read(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadAnswer reads user's answer to yes/no question
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPassword reads password or some private input that will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPasswordSecure reads password or some private input that will be hidden
// after pressing Enter
func ReadPasswordSecure(title string, nonEmpty bool) (*secstr.String, error) {
	panic("UNSUPPORTED")
}

// ❗ AddHistory adds line to input history
func AddHistory(data string) {
	panic("UNSUPPORTED")
}

// ❗ SetCompletionHandler adds autocompletion function (using Tab key)
func SetCompletionHandler(h func(input string) []string) {
	panic("UNSUPPORTED")
}

// ❗ SetHintHandler adds function for input hints
func SetHintHandler(h func(input string) string) {
	panic("UNSUPPORTED")
}
