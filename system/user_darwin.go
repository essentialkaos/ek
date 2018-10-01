package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
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

// getUserInfo find user info by name
func getUserInfo(nameOrID string) (*User, error) {
	cmd := exec.Command("dscl", ".", "-read", "/Users/"+nameOrID)

	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return nil, fmt.Errorf("User with this name %s does not exist", nameOrID)
	}

	var (
		lineStart = 0
		uid       int
		gid       int
		home      string
		shell     string
	)

	for i, b := range out {
		if b != '\n' {
			continue
		}

		// Skip long lines
		if i-lineStart > 128 {
			lineStart = i + 1
			continue
		}

		line := string(out[lineStart:i])

		lineStart = i + 1

		sepIndex := strings.Index(line, ":")

		if sepIndex == -1 {
			continue
		}

		rec := line[0:sepIndex]

		switch rec {
		case "UniqueID":
			uid, _ = strconv.Atoi(line[sepIndex+2:])
		case "PrimaryGroupID":
			gid, _ = strconv.Atoi(line[sepIndex+2:])
		case "NFSHomeDirectory":
			home = line[sepIndex+2:]
		case "UserShell":
			shell = line[sepIndex+2:]
		}
	}

	return &User{
		Name:     nameOrID,
		UID:      uid,
		GID:      gid,
		HomeDir:  home,
		Shell:    shell,
		RealName: nameOrID,
		RealUID:  uid,
		RealGID:  gid,
	}, nil
}
