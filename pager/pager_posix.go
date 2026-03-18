//go:build linux || darwin || freebsd

// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
	// ErrAlreadySet is returned by [Setup] if the pager has already been initialized
	ErrAlreadySet = errors.New("pager already set")

	// ErrNoPager is returned by [Setup] if no supported pager binary (less, more)
	// was found on the system and no explicit pager was provided
	ErrNoPager = errors.New("no pager found on the system")

	// ErrStdinPipe is returned by [Setup] if the pager process stdin pipe could not
	// be obtained as an *os.File, which is required for stdout redirection
	ErrStdinPipe = errors.New("can't get pager stdin pipe")

	// ErrPagerError is returned by [Complete] if the pager process exited with a
	// non-zero exit code
	ErrPagerError = errors.New("pager exited with an error")
)

// AllowEnv allows the user to define the pager binary via the PAGER environment
// variable
var AllowEnv bool

// ////////////////////////////////////////////////////////////////////////////////// //

// Setup redirects os.Stdout and os.Stderr through the given pager process.
// If no pager is provided, less or more is located automatically.
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
		pagerCmd = nil
		return ErrStdinPipe
	}

	err = pagerCmd.Start()

	if err != nil {
		os.Stdout, os.Stderr = stdout, stderr
		pagerOut.Close()
		pagerOut, pagerCmd = nil, nil
		return ErrPagerError
	}

	return nil
}

// Complete closes the pager stdin pipe, waits for the pager process to exit,
// and restores [os.Stdout] and [os.Stderr] to their original values
func Complete() {
	if pagerCmd == nil {
		return
	}

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
