//+build linux freebsd

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strings"
	"syscall"

	"pkg.re/essentialkaos/ek.v9/env"
	"pkg.re/essentialkaos/ek.v9/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Service states codes
const (
	STATE_STOPPED       = 0
	STATE_WORKS         = 1
	STATE_UNKNOWN uint8 = 255
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SysV if SysV is used on system
func SysV() bool {
	return !Systemd()
}

// Upstart if Upstart is used on system
func Upstart() bool {
	return env.Which("initctl") != ""
}

// Systemd if Systemd is used on system
func Systemd() bool {
	return env.Which("systemctl") != ""
}

// HasService return true if service is present in any init system
func HasService(name string) bool {
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

// GetServiceState return service state
func GetServiceState(name string) uint8 {
	if hasSystemdService(name) {
		return getSystemdServiceState(name)
	}

	if hasUpstartService(name) {
		return getUpstartServiceState(name)
	}

	if hasSysVService(name) {
		return getSysVServiceState(name)
	}

	return STATE_UNKNOWN
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
	if !strings.HasSuffix(name, ".service") {
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

func getSysVServiceState(name string) uint8 {
	cmd := exec.Command("/sbin/service", name, "status")

	cmd.Run()

	if cmd.ProcessState == nil {
		return STATE_UNKNOWN
	}

	waitStatus := cmd.ProcessState.Sys()

	if waitStatus == nil {
		return STATE_UNKNOWN
	}

	status, ok := waitStatus.(syscall.WaitStatus)

	if !ok {
		return STATE_UNKNOWN
	}

	switch status.ExitStatus() {
	case 0:
		return STATE_WORKS
	case 3:
		return STATE_STOPPED
	}

	return STATE_UNKNOWN
}

func getUpstartServiceState(name string) uint8 {
	if strings.HasSuffix(name, ".conf") {
		name = strings.Replace(name, ".conf", "", -1)
	}

	output, err := exec.Command("/sbin/status", name).Output()

	if err != nil {
		return STATE_UNKNOWN
	}

	if strings.Contains(string(output), "start/running") {
		return STATE_WORKS
	}

	if strings.Contains(string(output), "stop/waiting") {
		return STATE_STOPPED
	}

	return STATE_UNKNOWN
}

func getSystemdServiceState(name string) uint8 {
	output, err := exec.Command("/usr/bin/systemctl", "is-active", name).Output()

	if err != nil {
		return STATE_UNKNOWN
	}

	switch strings.TrimRight(string(output), "\n\r") {
	case "active":
		return STATE_WORKS
	case "inactive":
		return STATE_STOPPED
	}

	return STATE_UNKNOWN
}
