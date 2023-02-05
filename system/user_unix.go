//go:build linux || freebsd
// +build linux freebsd

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsUserExist checks if user exist on system or not
func IsUserExist(name string) bool {
	cmd := exec.Command("getent", "passwd", name)

	return cmd.Run() == nil
}

// IsGroupExist checks if group exist on system or not
func IsGroupExist(name string) bool {
	cmd := exec.Command("getent", "group", name)

	return cmd.Run() == nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserInfo tries to find user info by name or UID
func getUserInfo(nameOrID string) (*User, error) {
	cmd := exec.Command("getent", "passwd", nameOrID)

	data, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("User with name/ID %s does not exist", nameOrID)
	}

	return parseGetentPasswdOutput(string(data))
}

// getGroupInfo returns group info by name or id
func getGroupInfo(nameOrID string) (*Group, error) {
	cmd := exec.Command("getent", "group", nameOrID)

	data, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("Group with name/ID %s does not exist", nameOrID)
	}

	return parseGetentGroupOutput(string(data))
}

// parseGetentGroupOutput parse 'getent passwd' command output
func parseGetentPasswdOutput(data string) (*User, error) {
	data = strings.TrimRight(data, "\r\n")

	uid, err := strconv.Atoi(strutil.ReadField(data, 2, false, ":"))

	if err != nil {
		return nil, ErrCantParseGetentOutput
	}

	gid, err := strconv.Atoi(strutil.ReadField(data, 3, false, ":"))

	if err != nil {
		return nil, ErrCantParseGetentOutput
	}

	return &User{
		Name:     strutil.ReadField(data, 0, false, ":"),
		UID:      uid,
		GID:      gid,
		Comment:  strutil.ReadField(data, 4, false, ":"),
		HomeDir:  strutil.ReadField(data, 5, false, ":"),
		Shell:    strutil.ReadField(data, 6, false, ":"),
		RealName: strutil.ReadField(data, 0, false, ":"),
		RealUID:  uid,
		RealGID:  gid,
	}, nil
}

// parseGetentGroupOutput parse 'getent group' command output
func parseGetentGroupOutput(data string) (*Group, error) {
	name := strutil.ReadField(data, 0, false, ":")
	id := strutil.ReadField(data, 2, false, ":")

	if name == "" || id == "" {
		return nil, ErrCantParseGetentOutput
	}

	gid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return &Group{name, gid}, nil
}
