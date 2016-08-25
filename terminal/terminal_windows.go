// +build windows, !linux, !darwin

package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// KillSignalError is error type when user cancel input
var KillSignalError = errors.New("")

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

func SetCompletionHandler(compfunc func(in string) []string) {
	return
}
