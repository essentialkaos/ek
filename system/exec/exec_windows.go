// +build !linux, !darwin, windows

// Package exec provides methods for executing commands
package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
