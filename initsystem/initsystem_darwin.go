// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SysV returns true if SysV is used on system
func SysV() bool {
	return false
}

// Upstart returns true if Upstart is used on system
func Upstart() bool {
	return false
}

// Systemd returns true if Systemd is used on system
func Systemd() bool {
	return false
}

// Launchd returns true if Launchd is used on the system
func Launchd() bool {
	return true
}

// IsPresent returns true if service is present in any init system
func IsPresent(name string) bool {
	isExist, _, _ := getLaunchdStatus(name)
	return isExist
}

// IsWorks returns service state
func IsWorks(name string) (bool, error) {
	_, isWorks, err := getLaunchdStatus(name)
	return isWorks, err
}

// IsEnabled returns true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	return false, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getLaunchdStatus(name string) (bool, bool, error) {
	output, err := exec.Command("launchctl", "list").Output()

	if err != nil {
		return false, false, fmt.Errorf("launchd returned error")
	}

	isExist, isWorks := parseLaunchdOutput(output, name)

	return isExist, isWorks, nil
}

func parseLaunchdOutput(data []byte, name string) (bool, bool) {
	buf := bytes.NewBuffer(data)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		procName := strutil.ReadField(line, 2, false, '\t')

		if !strings.HasPrefix(procName, name) {
			continue
		}

		procPID := strutil.ReadField(line, 0, false, '\t')

		if procPID == "-" {
			return true, false
		} else {
			return true, true
		}
	}

	return false, false
}
