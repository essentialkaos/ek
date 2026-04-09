//go:build !linux && !darwin

// Package input provides methods for reading user input
package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"

	"github.com/essentialkaos/ek/v13/secstr"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ CompletionHandler is the function signature for Tab-completion callbacks
type CompletionHandler = func(input string) []string

// ❗ HintHandler is the function signature for inline input hint callbacks
type HintHandler = func(input string) string

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ErrKillSignal is error type when user cancel input
var ErrKillSignal = errors.New("")

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Prompt is the string displayed before each user input line
var Prompt = "> "

// ❗ MaskSymbol is the character used to replace each typed character when masking
// passwords
var MaskSymbol = "*"

// ❗ HideLength conceals the password length by replacing the mask with a fixed
// "[hidden]" label
var HideLength = false

// ❗ HidePassword hides typed characters in real time using linenoise's built-in
// mask mode. When enabled, the masking symbol is always an asterisk regardless
// of [MaskSymbol].
var HidePassword = false

// ❗ MaskSymbolColorTag is the fmtc color tag applied when rendering the mask symbol
var MaskSymbolColorTag = ""

// ❗ TitleColorTag is the fmtc color tag applied to input title lines
var TitleColorTag = "{s}"

// ❗ AlwaysYes causes [ReadAnswer] to return true without prompting the user,
// useful for non-interactive or forced-action modes
var AlwaysYes = false

// ❗ NewLine controls whether an extra blank line is printed before and after each
// input
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

// InvalidAnswerMesssage is returned when the user's answer is neither Y nor N
var InvalidAnswerMesssage = "Please enter Y or N"

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Read displays an optional title, prompts for a line of text, and applies
// any provided validators before returning the input
func Read(title string, validators ...Validator) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadAnswer prompts the user with a yes/no question and returns their answer
// as a bool. An optional defaultAnswer ("Y" or "N") is used when the user submits an
// empty response.
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPassword prompts for a line of text that is masked on screen after entry
func ReadPassword(title string, validators ...Validator) (string, error) {
	panic("UNSUPPORTED")
}

// ❗ ReadPasswordSecure works like ReadPassword but returns the value wrapped in a
// secstr.String to prevent the secret from lingering in plain memory
func ReadPasswordSecure(title string, validators ...Validator) (*secstr.String, error) {
	panic("UNSUPPORTED")
}

// ❗ AddHistory appends a line to the linenoise input history
func AddHistory(data string) {
	panic("UNSUPPORTED")
}

// ❗ SetHistoryCapacity sets the maximum number of entries retained in input history.
// Returns an error if capacity is less than 1.
func SetHistoryCapacity(capacity int) error {
	panic("UNSUPPORTED")
}

// ❗ SetCompletionHandler registers a callback invoked on Tab key press to provide
// completions
func SetCompletionHandler(h CompletionHandler) {
	panic("UNSUPPORTED")
}

// ❗ SetHintHandler registers a callback that provides an inline hint shown while the
// user types
func SetHintHandler(h HintHandler) {
	panic("UNSUPPORTED")
}
