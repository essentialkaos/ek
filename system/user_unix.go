// +build linux freebsd

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v10/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserInfo find user info by name or ID
func getUserInfo(nameOrID string) (*User, error) {
	cmd := exec.Command("getent", "passwd", nameOrID)

	data, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("User with name/ID %s does not exist", nameOrID)
	}

	return parseGetentPasswdOutput(string(data))
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
