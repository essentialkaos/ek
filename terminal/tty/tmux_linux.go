//go:build linux
// +build linux

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var procFS = "/proc"

// ////////////////////////////////////////////////////////////////////////////////// //

// isTmuxAncestor returns true if the current process is an ancestor of tmux
func isTmuxAncestor() (bool, error) {
	pid := strconv.Itoa(os.Getppid())

	for {
		statFile := procFS + "/" + pid + "/stat"
		statData, err := os.ReadFile(statFile)

		if err != nil {
			return false, errors.New("Can't check process tree for tmux server")
		}

		statString := string(statData)
		processName := strutil.ReadField(statString, 1, false, " ")

		if strings.HasPrefix(processName, "(tmux:") {
			return true, nil
		}

		parentPID := strutil.ReadField(statString, 3, false, " ")

		if parentPID == "1" || parentPID == "0" {
			break
		}

		pid = parentPID
	}

	return false, nil
}
