//+build linux freebsd

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"pkg.re/essentialkaos/ek.v10/env"
	"pkg.re/essentialkaos/ek.v10/fsutil"
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

// IsServiceWorks return service state
func IsServiceWorks(name string) (bool, error) {
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

// IsEnabled return true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	if !HasService(name) {
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

func getSysVServiceState(name string) (bool, error) {
	cmd := exec.Command("/sbin/service", name, "status")

	cmd.Run()

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

	if strings.Contains(string(output), "start/running") {
		return true, nil
	}

	if strings.Contains(string(output), "stop/waiting") {
		return false, nil
	}

	return false, fmt.Errorf("Can't parse upstart output")
}

func getSystemdServiceState(name string) (bool, error) {
	output, err := exec.Command("/usr/bin/systemctl", "show", name, "-p", "ActiveState", "-p", "LoadState").Output()

	if err != nil {
		return false, fmt.Errorf("systemd return an error")
	}

	switch {
	case strings.Contains(string(output), "LoadState=not-found"):
		return false, fmt.Errorf("Unit %s could not be found ", name)

	case strings.Contains(string(output), "ActiveState=active"):
		return true, nil

	case strings.Contains(string(output), "ActiveState=inactive"):
		return false, nil

	}

	return false, fmt.Errorf("Can't parse systemd output")
}

func isSysVEnabled(name string) (bool, error) {
	output, err := exec.Command("/sbin/chkconfig", "--list", name).Output()

	if err != nil {
		return false, fmt.Errorf("chkconfig returned an error")
	}

	if strings.Contains(string(output), ":on") {
		return true, nil
	}

	if strings.Contains(string(output), ":off") {
		return false, nil
	}

	return false, fmt.Errorf("Can't parse chkconfig output")
}

func isUpstartEnabled(name string) (bool, error) {
	if !strings.HasSuffix(name, ".conf") {
		name = name + ".conf"
	}

	fd, err := os.OpenFile("/etc/init/"+name, os.O_RDONLY, 0)

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

func isSystemdEnabled(name string) (bool, error) {
	output, err := exec.Command("/usr/bin/systemctl", "is-enabled", name).Output()

	if err != nil {
		return false, fmt.Errorf("systemd return error: %v", err)
	}

	if strings.TrimRight(string(output), "\n\r") == "enabled" {
		return true, nil
	}

	return false, nil
}
