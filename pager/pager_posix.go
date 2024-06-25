//go:build !windows
// +build !windows

// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var pagerCmd *exec.Cmd
var pagerOut *os.File

var stdout *os.File
var stderr *os.File

// ////////////////////////////////////////////////////////////////////////////////// //

var binLess = "less"
var binMore = "more"

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrAlreadySet = errors.New("Pager already set")
	ErrNoPager    = errors.New("There is no pager on the system")
	ErrStdinPipe  = errors.New("Can't get pager stdin")
)

// AllowEnv is a flag that allows to user to define pager binary using PAGER environment
// variable
var AllowEnv bool

// ////////////////////////////////////////////////////////////////////////////////// //

// Setup set up pager for work. After calling this method, any data sent to Stdout and
// Stderr (using fmt, fmtc, or terminal packages) will go to the pager.
func Setup(pager ...string) error {
	if pagerCmd != nil {
		return ErrAlreadySet
	}

	if len(pager) == 0 {
		pagerCmd = getPagerCommand("")
	} else {
		pagerCmd = getPagerCommand(pager[0])
	}

	if pagerCmd == nil {
		return ErrNoPager
	}

	pagerCmd.Stdout, stdout = os.Stdout, os.Stdout
	pagerCmd.Stderr, stderr = os.Stderr, os.Stderr

	w, err := pagerCmd.StdinPipe()

	if err != nil {
		return err
	}

	switch t := w.(type) {
	case *os.File:
		pagerOut = t
		os.Stdout = pagerOut
	default:
		return ErrStdinPipe
	}

	return pagerCmd.Start()
}

// Complete finishes output redirect to pager
func Complete() {
	if pagerOut != nil {
		pagerOut.Close()
		pagerOut = nil
	}

	if pagerCmd != nil {
		pagerCmd.Wait()
		pagerCmd = nil
	}

	os.Stdout = stdout
	os.Stderr = stderr
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getPagerCommand creates command for pager
func getPagerCommand(pager string) *exec.Cmd {
	if pager == "" && AllowEnv {
		pager = os.Getenv("PAGER")
	}

	if pager == "" {
		pager = findPager()
	}

	if pager == "" {
		return nil
	}

	if strings.Contains(pager, " ") {
		cmdSlice := strings.Fields(pager)
		return exec.Command(cmdSlice[0], cmdSlice[1:]...)
	}

	return exec.Command(pager)
}

// findPager tries to find pager an it options
func findPager() string {
	_, err := exec.LookPath(binMore)

	if err == nil {
		moreOpts := os.Getenv("MORE")

		if moreOpts != "" {
			return "more " + moreOpts
		}

		return "more -f"
	}

	_, err = exec.LookPath(binLess)

	if err == nil {
		lessOpts := os.Getenv("LESS")

		if lessOpts != "" {
			return "less " + lessOpts
		}

		return "less -R"
	}

	return ""
}
