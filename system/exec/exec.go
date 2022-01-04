//go:build !windows
// +build !windows

// Package exec provides methods for executing commands
package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Sudo executes command with sudo
func Sudo(user string, args ...string) error {
	var cmdArgs []string

	cmdArgs = append(cmdArgs, "-u", user)
	cmdArgs = append(cmdArgs, args...)

	return Run("sudo", cmdArgs...)
}

// Run executes command
func Run(command string, args ...string) error {
	var cmd = exec.Command(command)

	cmd.Args = append(cmd.Args, args...)

	return cmd.Run()
}

// RunAsUser runs command as a given user
func RunAsUser(user, logFile, command string, args ...string) error {
	var logFd *os.File
	var err error

	if logFile != "" {
		logFd, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			return err
		}
	}

	defer logFd.Close()

	cmd := exec.Command(
		"/sbin/runuser",
		"-s", "/bin/bash",
		user, "-c",
		command, strings.Join(args, " "),
	)

	if logFd != nil {
		cmd.Stderr = logFd
		cmd.Stdout = logFd
	}

	return cmd.Run()
}

// ////////////////////////////////////////////////////////////////////////////////// //
