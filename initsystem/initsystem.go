//go:build linux || freebsd

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"bytes"
	"errors"
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
func IsPresent(serviceName string) bool {
	if hasSystemdService(serviceName) {
		return true
	}

	if hasSysVService(serviceName) {
		return true
	}

	if hasUpstartService(serviceName) {
		return true
	}

	return false
}

// IsWorks returns service state
func IsWorks(serviceName string) (bool, error) {
	if hasSystemdService(serviceName) {
		return getSystemdServiceState(serviceName)
	}

	if hasUpstartService(serviceName) {
		return getUpstartServiceState(serviceName)
	}

	if hasSysVService(serviceName) {
		return getSysVServiceState(serviceName)
	}

	return false, fmt.Errorf("can't find service %q state", serviceName)
}

// IsEnabled returns true if auto start enabled for given service
func IsEnabled(serviceName string) (bool, error) {
	if !IsPresent(serviceName) {
		return false, fmt.Errorf("service %q doesn't exist on this system", serviceName)
	}

	if hasSystemdService(serviceName) {
		return isSystemdEnabled(serviceName)
	}

	if hasUpstartService(serviceName) {
		return isUpstartEnabled(serviceName)
	}

	if hasSysVService(serviceName) {
		return isSysVEnabled(serviceName)
	}

	return false, fmt.Errorf("can't find service %q state", serviceName)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// hasSysVService checks if service exists in SysV init system
func hasSysVService(serviceName string) bool {
	// Default path for linux
	initDir := "/etc/rc.d/init.d"

	if fsutil.CheckPerms("FXS", initDir+"/"+serviceName) {
		return true
	}

	// Default path for BSD
	initDir = "/usr/local/etc/rc.d"

	return fsutil.CheckPerms("FXS", initDir+"/"+serviceName)
}

// hasUpstartService checks if service exists in Upstart init system
func hasUpstartService(serviceName string) bool {
	if !strings.HasSuffix(serviceName, ".conf") {
		serviceName += ".conf"
	}

	return fsutil.IsExist("/etc/init/" + serviceName)
}

// hasSystemdService checks if service exists in Systemd init system
func hasSystemdService(serviceName string) bool {
	if !strings.Contains(serviceName, ".") {
		serviceName += ".service"
	}

	if fsutil.IsExist("/etc/systemd/system/" + serviceName) {
		return true
	}

	if fsutil.IsExist("/etc/systemd/user/" + serviceName) {
		return true
	}

	return fsutil.IsExist("/usr/lib/systemd/system/" + serviceName)
}

// getSysVServiceState returns service state from SysV init system
func getSysVServiceState(serviceName string) (bool, error) {
	cmd := exec.Command("/sbin/service", serviceName, "status")
	output, err := cmd.Output()

	if err != nil {
		return false, fmt.Errorf("can't execute 'service' command to get current service status")
	}

	if bytes.Contains(output, []byte("ExecStart")) {
		return getSystemdServiceState(serviceName)
	}

	if cmd.ProcessState == nil {
		return false, fmt.Errorf("can't get service %q command process state", serviceName)
	}

	waitStatus := cmd.ProcessState.Sys()

	if waitStatus == nil {
		return false, fmt.Errorf("can't get service %q command process wait status", serviceName)
	}

	status, ok := waitStatus.(syscall.WaitStatus)

	if !ok {
		return false, fmt.Errorf("can't get service %q command exit code", serviceName)
	}

	exitStatus := status.ExitStatus()

	switch exitStatus {
	case 0:
		return true, nil
	case 3:
		return false, nil
	}

	return false, fmt.Errorf("'service' command returned unsupported exit code (%d)", exitStatus)
}

// getUpstartServiceState returns service state from Upstart init system
func getUpstartServiceState(serviceName string) (bool, error) {
	serviceName = strings.TrimSuffix(serviceName, ".conf")

	output, err := exec.Command("/sbin/status", serviceName).CombinedOutput()

	if err != nil {
		return false, wrapCommandOutputToError("upstart returned error", output)
	}

	return parseUpstartStatusOutput(string(output))
}

// getSystemdServiceState returns service state from Systemd init system
func getSystemdServiceState(serviceName string) (bool, error) {
	output, err := exec.Command(
		"/usr/bin/systemctl", "show", serviceName,
		"-p", "ActiveState",
		"-p", "LoadState",
	).CombinedOutput()

	if err != nil {
		return false, wrapCommandOutputToError("systemd returned error", output)
	}

	return parseSystemdStatusOutput(serviceName, string(output))
}

// isSysVEnabled checks if service is enabled in SysV init system
func isSysVEnabled(serviceName string) (bool, error) {
	output, err := exec.Command("/sbin/chkconfig", "--list", serviceName).CombinedOutput()

	if err != nil {
		return false, wrapCommandOutputToError("chkconfig returned error", output)
	}

	return parseSysvEnabledOutput(string(output))
}

// isUpstartEnabled checks if service is enabled in Upstart init system
func isUpstartEnabled(serviceName string) (bool, error) {
	if !strings.HasSuffix(serviceName, ".conf") {
		serviceName += ".conf"
	}

	return parseUpstartEnabledData(serviceName, "/etc/init/"+serviceName)
}

// isSystemdEnabled checks if service is enabled in Systemd init system
func isSystemdEnabled(serviceName string) (bool, error) {
	output, err := exec.Command("/usr/bin/systemctl", "is-enabled", serviceName).CombinedOutput()

	if err != nil {
		return false, wrapCommandOutputToError("systemd returned error", output)
	}

	return parseSystemdEnabledOutput(string(output)), nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseSystemdEnabledOutput parses output of 'systemctl is-enabled' command output
func parseSystemdEnabledOutput(data string) bool {
	return strings.TrimRight(data, "\n\r") == "enabled"
}

// parseUpstartEnabledData parses upstart configuration and extracts if service
// is enabled to start on system startup
func parseUpstartEnabledData(serviceName, file string) (bool, error) {
	fd, err := os.Open(file)

	if err != nil {
		return false, fmt.Errorf("can't read service %q unit file", serviceName)
	}

	defer fd.Close()

	s := bufio.NewScanner(fd)

	for s.Scan() {
		text := strings.TrimLeft(s.Text(), " \t")

		if strings.HasPrefix(text, "#") {
			continue
		}

		if strings.Contains(text, "start on") {
			return true, nil
		}
	}

	return false, nil
}

// parseSysvEnabledOutput parses output of 'chkconfig --list' command output
func parseSysvEnabledOutput(data string) (bool, error) {
	switch {
	case strings.Contains(data, ":on"):
		return true, nil

	case strings.Contains(data, ":off"):
		return false, nil

	default:
		return false, fmt.Errorf("can't parse chkconfig output")
	}
}

// parseSystemdStatusOutput parses output of 'systemctl show' command output
func parseSystemdStatusOutput(serviceName, data string) (bool, error) {
	loadState := strutil.ReadField(data, 0, false, '\n')
	loadStateValue := strutil.ReadField(loadState, 1, false, '=')

	if strings.Trim(loadStateValue, "\r\n") == "not-found" {
		return false, fmt.Errorf("unit %q could not be found", serviceName)
	}

	activeState := strutil.ReadField(data, 1, false, '\n')
	activeStateValue := strutil.ReadField(activeState, 1, false, '=')

	switch strings.Trim(activeStateValue, "\r\n") {
	case "active":
		return true, nil

	case "inactive", "failed":
		return false, nil
	}

	return false, fmt.Errorf("can't parse systemd output")
}

// parseUpstartStatusOutput parses output of 'initctl status' command output
func parseUpstartStatusOutput(data string) (bool, error) {
	data = strings.TrimRight(data, "\r\n")
	status := strutil.ReadField(data, 1, false, ' ')

	switch status {
	case "start/running":
		return true, nil

	case "stop/waiting":
		return false, nil

	default:
		return false, fmt.Errorf("can't parse upstart output")
	}
}

// wrapCommandOutputToError wraps command stderr output into error
func wrapCommandOutputToError(msg string, output []byte) error {
	if len(output) == 0 {
		return errors.New(msg + ": (no output)")
	}

	return fmt.Errorf("%s: %s", msg, strings.ReplaceAll(string(output), "\n", " "))
}
