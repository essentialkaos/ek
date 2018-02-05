// +build !linux, !darwin, windows

// Package exec provides methods for executing commands
package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Sudo execute some command with sudo
func Sudo(user string, args ...string) error {
	return nil
}

// Run execute some command
func Run(command string, args ...string) error {
	return nil
}

// RunAsUser run command as some user
func RunAsUser(user, logFile string, command string, args ...string) error {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
