// +build windows, !linux, !darwin

package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
	return "", nil
}

func ReadAnswer(title, defaultAnswer string) (bool, error) {
	return true, nil
}

func ReadPassword(title string, nonEmpty bool) (string, error) {
	return "", nil
}

func PrintErrorMessage(message string, args ...interface{}) {
	return
}

func PrintWarnMessage(message string, args ...interface{}) {
	return
}

func PrintActionMessage(message string) {
	return
}

func PrintActionStatus(status int) {
	return
}

func AddHstory(ui string) {
	return
}

func SetCompletionHandler(h func(in string) []string) {
	return
}

func SetHintHandler(h func(input string) string) {
	return
}
