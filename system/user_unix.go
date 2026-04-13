//go:build linux || freebsd

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type utmpRecord struct {
	Type   int16
	_      [2]byte
	_      int32
	Device [32]byte
	_      [4]byte
	User   [32]byte
	Host   [256]byte
	_      utmpExitStatus
	_      int32
	Time   utmpTimeval
	_      [16]byte
	_      [20]byte
}

type utmpExitStatus struct {
	Termination int16
	Exit        int16
}

type utmpTimeval struct {
	Sec  int32
	Usec int32
}

// ////////////////////////////////////////////////////////////////////////////////// //

// devDir is path to dir with pts data
var devDir = "/dev"

// utmpFile is path to utmp file with sessions info
var utmpFile = "/var/run/utmp"

// ////////////////////////////////////////////////////////////////////////////////// //

// Who returns information about all active login sessions
func Who() ([]*SessionInfo, error) {
	fd, err := os.Open(utmpFile)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	var sessions []*SessionInfo

	for {
		var rec utmpRecord

		err = binary.Read(fd, binary.LittleEndian, &rec)

		if err == io.EOF {
			break
		}

		if err != nil {
			return sessions, fmt.Errorf("can't read utmp record: %w", err)
		}

		if rec.Type != 0x7 {
			continue
		}

		pts := string(bytes.TrimRight(rec.Device[:], "\x00"))
		_, mtime, _, _ := getTimes(devDir + "/" + pts)

		sessions = append(sessions, &SessionInfo{
			Username:         string(bytes.TrimRight(rec.User[:], "\x00")),
			Host:             string(bytes.TrimRight(rec.Host[:], "\x00")),
			LoginTime:        time.Unix(int64(rec.Time.Sec), 0),
			LastActivityTime: mtime,
		})
	}

	return sessions, nil
}

// IsUserExist returns true if a user with the given name or UID exists on the system
func IsUserExist(name string) bool {
	return exec.Command("getent", "passwd", name).Run() == nil
}

// IsGroupExist returns true if a group with the given name or GID exists on the system
func IsGroupExist(name string) bool {
	return exec.Command("getent", "group", name).Run() == nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserInfo tries to find user info by name or UID
func getUserInfo(nameOrID string) (*User, error) {
	data, err := exec.Command("getent", "passwd", nameOrID).Output()

	if err != nil {
		return nil, fmt.Errorf("user with name/ID %q does not exist", nameOrID)
	}

	return parseGetentPasswdOutput(string(data))
}

// getGroupInfo returns group info by name or id
func getGroupInfo(nameOrID string) (*Group, error) {
	data, err := exec.Command("getent", "group", nameOrID).Output()

	if err != nil {
		return nil, fmt.Errorf("group with name/ID %q does not exist", nameOrID)
	}

	return parseGetentGroupOutput(string(data))
}

// parseGetentGroupOutput parse 'getent passwd' command output
func parseGetentPasswdOutput(data string) (*User, error) {
	data = strings.TrimRight(data, "\r\n")

	uid, err := strconv.Atoi(strutil.ReadField(data, 2, false, ':'))

	if err != nil {
		return nil, ErrCantParseGetentOutput
	}

	gid, err := strconv.Atoi(strutil.ReadField(data, 3, false, ':'))

	if err != nil {
		return nil, ErrCantParseGetentOutput
	}

	return &User{
		Name:     strutil.ReadField(data, 0, false, ':'),
		UID:      uid,
		GID:      gid,
		Comment:  strutil.ReadField(data, 4, false, ':'),
		HomeDir:  strutil.ReadField(data, 5, false, ':'),
		Shell:    strutil.ReadField(data, 6, false, ':'),
		RealName: strutil.ReadField(data, 0, false, ':'),
		RealUID:  uid,
		RealGID:  gid,
	}, nil
}

// parseGetentGroupOutput parse 'getent group' command output
func parseGetentGroupOutput(data string) (*Group, error) {
	name := strutil.ReadField(data, 0, false, ':')
	id := strutil.ReadField(data, 2, false, ':')

	if name == "" || id == "" {
		return nil, ErrCantParseGetentOutput
	}

	gid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return &Group{name, gid}, nil
}
