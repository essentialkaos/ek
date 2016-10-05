// +build linux, darwin, !windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SudoExec execute some command with sudo
func SudoExec(user string, args ...string) error {
	var cmdArgs []string

	cmdArgs = append(cmdArgs, "-u", user)
	cmdArgs = append(cmdArgs, args...)

	return Exec("sudo", cmdArgs...)
}

// Exec execute some command
func Exec(command string, args ...string) error {
	var cmd = exec.Command(command)

	cmd.Args = append(cmd.Args, args...)

	return cmd.Run()
}

// RunAsUser run command as some user
func RunAsUser(user, logFile string, command string, args ...string) error {
	var log *os.File
	var err error

	if logFile != "" {
		log, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			return err
		}
	}

	defer log.Close()

	cmd := exec.Command(
		"/sbin/runuser",
		"-s", "/bin/bash",
		user, "-c",
		command, strings.Join(args, " "),
	)

	if logFile != "" && log != nil {
		cmd.Stderr = log
		cmd.Stdout = log
	}

	return cmd.Run()
}

// ////////////////////////////////////////////////////////////////////////////////// //
