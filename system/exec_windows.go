// +build !linux, !darwin, windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
