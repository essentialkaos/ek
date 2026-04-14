//go:build linux

// Package exec provides methods for executing commands

package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/exec"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultLogMode default mode for log files
var DefaultLogMode os.FileMode = 0600

// ////////////////////////////////////////////////////////////////////////////////// //

// Sudo runs the given command as the specified user via sudo
func Sudo(user string, args ...string) error {
	var cmdArgs []string

	cmdArgs = append(cmdArgs, "-u", user)
	cmdArgs = append(cmdArgs, args...)

	return Run("sudo", cmdArgs...)
}

// Run executes the given command with optional arguments
func Run(command string, args ...string) error {
	return exec.Command(command, args...).Run()
}

// RunAsUser runs the given command as the specified user via runuser,
// optionally redirecting stdout and stderr to logFile
func RunAsUser(user, logFile, command string, args ...string) error {
	var logFd *os.File
	var err error

	if logFile != "" {
		logFd, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, DefaultLogMode)

		if err != nil {
			return err
		}

		defer logFd.Close()
	}

	cmd := exec.Command("/sbin/runuser", "-s", "/bin/bash", user, "--", command)

	cmd.Args = append(cmd.Args, args...)

	if logFd != nil {
		cmd.Stderr = logFd
		cmd.Stdout = logFd
	}

	return cmd.Run()
}

// ////////////////////////////////////////////////////////////////////////////////// //
