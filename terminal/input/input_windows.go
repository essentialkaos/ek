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

var (
	// ❗ MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
	MaskSymbolColorTag = ""

	// ❗ TitleColorTag is fmtc color tag used for input titles
	TitleColorTag = "{s}"
)

// ❗ HidePassword is flag for hiding password while typing
// Because of using the low-level linenoise method for this feature, we can not use a
// custom masking symbol, so it always will be an asterisk (*).
var HidePassword = false

// ❗ AlwaysYes is a flag, if set ReadAnswer will always return true (useful for working
// with option for forced actions)
var AlwaysYes = false

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Read reads user's input
func Read(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadAnswer reads user's answer for yes/no question
func ReadAnswer(title, defaultAnswer string) (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPassword reads password or some private input which will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPasswordSecure reads password or some private input which will be hidden
// after pressing Enter
func ReadPasswordSecure(title string, nonEmpty bool) (*secstr.String, error) {
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
