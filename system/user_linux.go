// +build linux, !darwin, !windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getTimes is copy of fsutil.GetTimes
func getTimes(path string) (time.Time, time.Time, time.Time, error) {
	if path == "" {
		return time.Time{}, time.Time{}, time.Time{}, errors.New("Path is empty")
	}

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)),
		time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec)),
		time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)),
		nil
}

// getUserInfo return user info by name or id (name, id, gid, comment, home, shell)
//
func getUserInfo(nameOrID string) (string, int, int, string, string, string, error) {
	cmd := exec.Command("getent", "passwd", nameOrID)

	out, err := cmd.Output()

	if err != nil {
		return "", -1, -1, "", "", "", fmt.Errorf("User with this name/id %s is not exist", nameOrID)
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")
	aOut := strings.Split(sOut, ":")

	uid, _ := strconv.Atoi(aOut[2])
	gid, _ := strconv.Atoi(aOut[3])

	return aOut[0], uid, gid, aOut[4], aOut[5], aOut[6], nil
}
