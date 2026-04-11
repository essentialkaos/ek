//go:build !linux

// Package exec provides methods for executing commands
package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Sudo runs the given command as the specified user via sudo
func Sudo(user string, args ...string) error {
	panic("UNSUPPORTED")
}

// ❗ Run executes the given command with optional arguments
func Run(command string, args ...string) error {
	panic("UNSUPPORTED")
}

// ❗ RunAsUser runs the given command as the specified user via runuser,
// optionally redirecting stdout and stderr to logFile
func RunAsUser(user, logFile string, command string, args ...string) error {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //
