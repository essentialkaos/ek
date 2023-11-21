// Package pager provides methods for pager setup (more/less)
package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
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

// DEFAULT is default pager command
const DEFAULT = "more"

// ////////////////////////////////////////////////////////////////////////////////// //

var pagerCmd *exec.Cmd
var pagerOut *os.File

var stdout *os.File
var stderr *os.File

// ////////////////////////////////////////////////////////////////////////////////// //

var ErrAlreadySet = errors.New("Pager already set")

// ////////////////////////////////////////////////////////////////////////////////// //

// Setup set up pager for work. After calling this method, any data sent to Stdout and
// Stderr (using fmt, fmtc, or terminal packages) will go to the pager.
func Setup(pager string) error {
	if pagerCmd != nil {
		return ErrAlreadySet
	}

	pagerCmd = getPagerCommand(pager)

	pagerCmd.Stdout, stdout = os.Stdout, os.Stdout
	pagerCmd.Stderr, stderr = os.Stderr, os.Stderr

	w, err := pagerCmd.StdinPipe()

	if err != nil {
		return err
	}

	pagerOut = w.(*os.File)
	os.Stdout = pagerOut

	return pagerCmd.Start()
}

// Complete finishes pager work
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
	if pager == "" {
		pager = os.Getenv("PAGER")
	}

	if pager == "" {
		pager = DEFAULT
	}

	if strings.Contains(pager, " ") {
		cmdSlice := strings.Fields(pager)
		return exec.Command(cmdSlice[0], cmdSlice[1:]...)
	}

	return exec.Command(pager)
}
