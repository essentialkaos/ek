// +build !linux, !darwin, windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// SudoExec execute some command with sudo
func SudoExec(user string, args ...string) error {
	return nil
}

// Exec execute some command
func Exec(command string, args ...string) error {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
