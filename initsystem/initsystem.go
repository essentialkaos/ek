//go:build linux || freebsd
// +build linux freebsd

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/essentialkaos/ek/v13/env"
	"github.com/essentialkaos/ek/v13/fsutil"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_STATUS_UNKNOWN     = 0
	_STATUS_PRESENT     = 1
	_STATUS_NOT_PRESENT = 2
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	sysvStatus    = _STATUS_UNKNOWN
	upstartStatus = _STATUS_UNKNOWN
	systemdStatus = _STATUS_UNKNOWN
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SysV returns true if SysV is used on the system
func SysV() bool {
	if sysvStatus != _STATUS_UNKNOWN {
		return sysvStatus == _STATUS_PRESENT
	}

	switch Systemd() {
	case true:
		sysvStatus = _STATUS_NOT_PRESENT
	default:
		sysvStatus = _STATUS_PRESENT
	}

	return sysvStatus == _STATUS_PRESENT
}

// Upstart returns true if Upstart is used on the system
func Upstart() bool {
	if upstartStatus != _STATUS_UNKNOWN {
		return upstartStatus == _STATUS_PRESENT
	}

	switch env.Which("initctl") {
	case "":
		upstartStatus = _STATUS_NOT_PRESENT
	default:
		upstartStatus = _STATUS_PRESENT
	}

	return upstartStatus == _STATUS_PRESENT
}

// Systemd returns true if Systemd is used on the system
func Systemd() bool {
	if systemdStatus != _STATUS_UNKNOWN {
		return systemdStatus == _STATUS_PRESENT
	}

	switch env.Which("systemctl") {
	case "":
		systemdStatus = _STATUS_NOT_PRESENT
	default:
		systemdStatus = _STATUS_PRESENT
	}

	return systemdStatus == _STATUS_PRESENT
}

// Launchd returns true if Launchd is used on the system
func Launchd() bool {
	return false
}

// IsPresent returns true if service is present in any init system
func IsPresent(name string) bool {
	if hasSystemdService(name) {
		return true
	}

	if hasSysVService(name) {
		return true
	}

	if hasUpstartService(name) {
		return true
	}

	return false
}

// IsWorks returns service state
func IsWorks(name string) (bool, error) {
	if hasSystemdService(name) {
		return getSystemdServiceState(name)
	}

	if hasUpstartService(name) {
		return getUpstartServiceState(name)
	}

	if hasSysVService(name) {
		return getSysVServiceState(name)
	}

	return false, fmt.Errorf("Can't find service state")
}

// IsEnabled returns true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	if !IsPresent(name) {
		return false, fmt.Errorf("Service doesn't exist on this system")
	}

	if hasSystemdService(name) {
		return isSystemdEnabled(name)
	}

	if hasUpstartService(name) {
		return isUpstartEnabled(name)
	}

	if hasSysVService(name) {
		return isSysVEnabled(name)
	}

	return false, fmt.Errorf("Can't find service state")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func hasSysVService(name string) bool {
	// Default path for linux
	initDir := "/etc/rc.d/init.d"

	if fsutil.CheckPerms("FXS", initDir+"/"+name) {
		return true
	}

	// Default path for BSD
	initDir = "/usr/local/etc/rc.d"

	return fsutil.CheckPerms("FXS", initDir+"/"+name)
}

func hasUpstartService(name string) bool {
	if !strings.HasSuffix(name, ".conf") {
		name = name + ".conf"
	}

	return fsutil.IsExist("/etc/init/" + name)
}

func hasSystemdService(name string) bool {
	if !strings.Contains(name, ".") {
		name = name + ".service"
	}

	if fsutil.IsExist("/etc/systemd/system/" + name) {
		return true
	}

	if fsutil.IsExist("/etc/systemd/user/" + name) {
		return true
	}

	return fsutil.IsExist("/usr/lib/systemd/system/" + name)
}

func getSysVServiceState(name string) (bool, error) {
	cmd := exec.Command("/sbin/service", name, "status")

	output, _ := cmd.Output()

	if bytes.Contains(output, []byte("ExecStart")) {
		return getSystemdServiceState(name)
	}

	if cmd.ProcessState == nil {
		return false, fmt.Errorf("Can't get service command process state")
	}

	waitStatus := cmd.ProcessState.Sys()

	if waitStatus == nil {
		return false, fmt.Errorf("Can't get service command process state")
	}

	status, ok := waitStatus.(syscall.WaitStatus)

	if !ok {
		return false, fmt.Errorf("Can't get service command exit code")
	}

	exitStatus := status.ExitStatus()

	switch exitStatus {
	case 0:
		return true, nil
	case 3:
		return false, nil
	}

	return false, fmt.Errorf("service command return unsupported exit code (%d)", exitStatus)
}

func getUpstartServiceState(name string) (bool, error) {
	if strings.HasSuffix(name, ".conf") {
		name = strings.Replace(name, ".conf", "", -1)
	}

	output, err := exec.Command("/sbin/status", name).Output()

	if err != nil {
		return false, fmt.Errorf("upstart returned an error")
	}

	return parseUpstartStatusOutput(string(output))
}

func getSystemdServiceState(name string) (bool, error) {
	output, err := exec.Command("/usr/bin/systemctl", "show", name, "-p", "ActiveState", "-p", "LoadState").Output()

	if err != nil {
		return false, fmt.Errorf("systemd return an error")
	}

	return parseSystemdStatusOutput(name, string(output))
}

func isSysVEnabled(name string) (bool, error) {
	output, err := exec.Command("/sbin/chkconfig", "--list", name).Output()

	if err != nil {
		return false, fmt.Errorf("chkconfig returned an error")
	}

	return parseSysvEnabledOutput(string(output))
}

func isUpstartEnabled(name string) (bool, error) {
	if !strings.HasSuffix(name, ".conf") {
		name = name + ".conf"
	}

	return parseUpstartEnabledData("/etc/init/" + name)
}

func isSystemdEnabled(name string) (bool, error) {
	output, err := exec.Command("/usr/bin/systemctl", "is-enabled", name).Output()

	if err != nil {
		return false, fmt.Errorf("systemd return error: %v", err)
	}

	return parseSystemdEnabledOutput(string(output)), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseSystemdEnabledOutput(data string) bool {
	return strings.TrimRight(data, "\n\r") == "enabled"
}

func parseUpstartEnabledData(file string) (bool, error) {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return false, fmt.Errorf("Can't read service unit file")
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	for s.Scan() {
		text := strings.TrimLeft(s.Text(), " ")

		if strings.HasPrefix(text, "#") {
			continue
		}

		if strings.Contains(text, "start on") {
			return true, nil
		}
	}

	return false, nil
}

func parseSysvEnabledOutput(data string) (bool, error) {
	switch {
	case strings.Contains(data, ":on"):
		return true, nil

	case strings.Contains(data, ":off"):
		return false, nil

	default:
		return false, fmt.Errorf("Can't parse chkconfig output")
	}
}

func parseSystemdStatusOutput(name, data string) (bool, error) {
	loadState := strutil.ReadField(data, 0, false, '\n')
	loadStateValue := strutil.ReadField(loadState, 1, false, '=')

	if strings.Trim(loadStateValue, "\r\n") == "not-found" {
		return false, fmt.Errorf("Unit %s could not be found", name)
	}

	activeState := strutil.ReadField(data, 1, false, '\n')
	activeStateValue := strutil.ReadField(activeState, 1, false, '=')

	switch strings.Trim(activeStateValue, "\r\n") {
	case "active":
		return true, nil

	case "inactive", "failed":
		return false, nil
	}

	return false, fmt.Errorf("Can't parse systemd output")
}

func parseUpstartStatusOutput(data string) (bool, error) {
	data = strings.TrimRight(data, "\r\n")
	status := strutil.ReadField(data, 1, false, ' ')

	switch status {
	case "start/running":
		return true, nil

	case "stop/waiting":
		return false, nil

	default:
		return false, fmt.Errorf("Can't parse upstart output")
	}
}
