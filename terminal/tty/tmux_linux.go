//go:build linux

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// procFS is the path to the proc filesystem
var procFS = "/proc"

// ////////////////////////////////////////////////////////////////////////////////// //

// isTmuxAncestor walks the process tree to check whether any ancestor is a tmux server
func isTmuxAncestor() (bool, error) {
	pid := strconv.Itoa(os.Getppid())

	for {
		statFile := procFS + "/" + pid + "/stat"
		statData, err := os.ReadFile(statFile)

		if err != nil {
			return false, errors.New("can't check process tree for tmux server")
		}

		statString := string(statData)
		processName := strutil.ReadField(statString, 1, false, ' ')

		if strings.HasPrefix(processName, "(tmux:") {
			return true, nil
		}

		parentPID := strutil.ReadField(statString, 3, false, ' ')

		if parentPID == "1" || parentPID == "0" || parentPID == pid {
			break
		}

		pid = parentPID
	}

	return false, nil
}
