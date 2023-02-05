package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsUserExist checks if user exist on system or not
func IsUserExist(name string) bool {
	cmd := exec.Command("dscl", ".", "-read", "/Users/"+name, "RecordName")

	return cmd.Run() == nil
}

// IsGroupExist checks if group exist on system or not
func IsGroupExist(name string) bool {
	cmd := exec.Command("dscl", ".", "-read", "/Groups/"+name, "RecordName")

	return cmd.Run() == nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getUserInfo tries to find user info by name
func getUserInfo(name string) (*User, error) {
	cmd := exec.Command(
		"dscl", ".", "-read", "/Users/"+name,
		"UniqueID", "PrimaryGroupID", "NFSHomeDirectory", "UserShell",
	)

	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return nil, fmt.Errorf("User with name %s does not exist", name)
	}

	user := &User{Name: name, RealName: name}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	for scanner.Scan() {
		line := scanner.Text()
		field := strutil.ReadField(line, 0, false, ":")

		switch field {
		case "UniqueID":
			user.UID, err = strconv.Atoi(strings.TrimSpace(strutil.ReadField(line, 1, false, ":")))
		case "PrimaryGroupID":
			user.GID, err = strconv.Atoi(strings.TrimSpace(strutil.ReadField(line, 1, false, ":")))
		case "NFSHomeDirectory":
			user.HomeDir = strutil.ReadField(
				strings.TrimSpace(strutil.ReadField(line, 1, false, ":")),
				0, false, " ",
			)
		case "UserShell":
			user.Shell = strings.TrimSpace(strutil.ReadField(line, 1, false, ":"))
		}

		if err != nil {
			return nil, fmt.Errorf("Can't parse field %s value: %v", field, err)
		}
	}

	user.RealUID = user.UID
	user.RealGID = user.GID

	return user, nil
}

// getGroupInfo returns group info by name
func getGroupInfo(name string) (*Group, error) {
	cmd := exec.Command("dscl", ".", "-read", "/Groups/"+name, "PrimaryGroupID")

	out, err := cmd.Output()

	if err != nil || len(out) == 0 {
		return nil, fmt.Errorf("Group with name %s does not exist", name)
	}

	group := &Group{Name: name}
	data := string(out)

	group.GID, err = strconv.Atoi(strings.TrimSpace(strutil.ReadField(data, 1, false, ":")))

	if err != nil {
		return nil, fmt.Errorf("Can't parse PrimaryGroupID field value: %v", err)
	}

	return group, nil
}
