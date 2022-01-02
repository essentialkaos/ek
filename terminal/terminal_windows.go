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

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = errors.New("")

// Prompt is prompt string
var Prompt = "> "

// MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
var MaskSymbolColorTag = ""

// ////////////////////////////////////////////////////////////////////////////////// //

func ReadUI(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

func ReadAnswer(title, defaultAnswer string) (bool, error) {
	panic("UNSUPPORTED")
	return true, nil
}

func ReadPassword(title string, nonEmpty bool) (string, error) {
	panic("UNSUPPORTED")
	return "", nil
}

func PrintErrorMessage(message string, args ...interface{}) {
	panic("UNSUPPORTED")
}

func PrintWarnMessage(message string, args ...interface{}) {
	panic("UNSUPPORTED")
}

func PrintActionMessage(message string) {
	panic("UNSUPPORTED")
}

func PrintActionStatus(status int) {
	panic("UNSUPPORTED")
}

func AddHstory(ui string) {
	panic("UNSUPPORTED")
}

func SetCompletionHandler(h func(in string) []string) {
	panic("UNSUPPORTED")
}

func SetHintHandler(h func(input string) string) {
	panic("UNSUPPORTED")
}
