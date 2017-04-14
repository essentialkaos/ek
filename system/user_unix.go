// +build linux freebsd

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserInfo find user info by name or id
func getUserInfo(nameOrID string) (*User, error) {
	cmd := exec.Command("getent", "passwd", nameOrID)

	out, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("User with this name/id %s does not exist", nameOrID)
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")
	aOut := strings.Split(sOut, ":")

	uid, _ := strconv.Atoi(aOut[2])
	gid, _ := strconv.Atoi(aOut[3])

	return &User{
		Name:     aOut[0],
		UID:      uid,
		GID:      gid,
		Comment:  aOut[4],
		HomeDir:  aOut[5],
		Shell:    aOut[6],
		RealName: aOut[0],
		RealUID:  uid,
		RealGID:  gid,
	}, nil
}
