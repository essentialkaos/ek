// +build linux, darwin, !windows

// Package exec provides methods for executing commands
package exec

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Sudo execute some command with sudo
func Sudo(user string, args ...string) error {
	var cmdArgs []string

	cmdArgs = append(cmdArgs, "-u", user)
	cmdArgs = append(cmdArgs, args...)

	return Run("sudo", cmdArgs...)
}

// Run execute some command
func Run(command string, args ...string) error {
	var cmd = exec.Command(command)

	cmd.Args = append(cmd.Args, args...)

	return cmd.Run()
}

// RunAsUser run command as some user
func RunAsUser(user, logFile string, command string, args ...string) error {
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
