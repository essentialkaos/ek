//go:build !windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/essentialkaos/ek/v13/env"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// User contains information about user
type User struct {
	UID      int      `json:"uid"`
	GID      int      `json:"gid"`
	Name     string   `json:"name"`
	Groups   []*Group `json:"groups"`
	Comment  string   `json:"comment"`
	Shell    string   `json:"shell"`
	HomeDir  string   `json:"home_dir"`
	RealUID  int      `json:"real_uid"`
	RealGID  int      `json:"real_gid"`
	RealName string   `json:"real_name"`
}

// Group contains information about group
type Group struct {
	Name string `json:"name"`
	GID  int    `json:"gid"`
}

// SessionInfo contains information about all sessions
type SessionInfo struct {
	Username         string    `json:"username"`
	Host             string    `json:"host"`
	LoginTime        time.Time `json:"login_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors
var (
	// ErrEmptyPath is returned if given path is empty
	ErrEmptyPath = errors.New("Path is empty")

	// ErrEmptyUserName is returned if given user name or uid is empty
	ErrEmptyUserName = errors.New("User name/ID can't be blank")

	// ErrEmptyGroupName is returned if given group name of gid is empty
	ErrEmptyGroupName = errors.New("Group name/ID can't be blank")

	// ErrCantParseIdOutput is returned if id command output has unsupported format
	ErrCantParseIdOutput = errors.New("Can't parse id command output")

	// ErrCantParseGetentOutput is returned if getent command output has unsupported format
	ErrCantParseGetentOutput = errors.New("Can't parse getent command output")
)

// CurrentUserCachePeriod is cache period for current user info
var CurrentUserCachePeriod = 5 * time.Minute

// ////////////////////////////////////////////////////////////////////////////////// //

// Current is user info cache
var curUser *User

// curUserUpdateDate is date when user data was updated
var curUserUpdateDate time.Time

// ////////////////////////////////////////////////////////////////////////////////// //

// CurrentUser returns struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	if len(avoidCache) == 0 || !avoidCache[0] {
		if curUser != nil && time.Since(curUserUpdateDate) < CurrentUserCachePeriod {
			return curUser, nil
		}
	}

	username, err := getCurrentUserName()

	if err != nil {
		return nil, err
	}

	user, err := LookupUser(username)

	if err != nil {
		return user, err
	}

	if user.Name == "root" {
		appendRealUserInfo(user)
	}

	curUser = user
	curUserUpdateDate = time.Now()

	return user, nil
}

// LookupUser searches user info by given name
func LookupUser(nameOrID string) (*User, error) {
	if nameOrID == "" {
		return nil, ErrEmptyUserName
	}

	user, err := getUserInfo(nameOrID)

	if err != nil {
		return nil, err
	}

	err = appendGroupInfo(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// LookupGroup searches group info by given name
func LookupGroup(nameOrID string) (*Group, error) {
	if nameOrID == "" {
		return nil, ErrEmptyGroupName
	}

	return getGroupInfo(nameOrID)
}

// CurrentTTY returns current tty or empty string if error occurred
func CurrentTTY() string {
	pid := strconv.Itoa(os.Getpid())
	fdLink, err := os.Readlink("/proc/" + pid + "/fd/0")

	if err != nil {
		return ""
	}

	return fdLink
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRoot checks if current user is root
func (u *User) IsRoot() bool {
	return u.UID == 0 && u.GID == 0
}

// IsSudo checks if it user over sudo command
func (u *User) IsSudo() bool {
	return u.IsRoot() && u.RealUID != 0 && u.RealGID != 0
}

// GroupList returns slice with user groups names
func (u *User) GroupList() []string {
	var result []string

	for _, group := range u.Groups {
		result = append(result, group.Name)
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getCurrentUserName returns name of current user
func getCurrentUserName() (string, error) {
	cmd := exec.Command("id", "-un")

	data, err := cmd.Output()

	if err != nil {
		return "", ErrCantParseIdOutput
	}

	username := strings.TrimRight(string(data), "\n")

	return username, nil
}

// appendGroupInfo append info about groups
func appendGroupInfo(user *User) error {
	cmd := exec.Command("id", user.Name)

	data, err := cmd.Output()

	if err != nil {
		return ErrCantParseIdOutput
	}

	user.Groups = extractGroupsInfo(string(data))

	return nil
}

// appendRealUserInfo append real user info when user under sudo
func appendRealUserInfo(user *User) {
	username, uid, gid := getRealUserByPTY()

	if username == "" {
		username, uid, gid = getRealUserFromEnv()
	}

	user.RealName = username
	user.RealUID = uid
	user.RealGID = gid
}

// getUserInfo returns UID associated with current TTY
func getTDOwnerID() (int, bool) {
	tty := CurrentTTY()

	if tty == "" {
		return -1, false
	}

	ownerID, err := getOwner(tty)

	return ownerID, err == nil
}

// getRealUserByPTY try to find info about real user from real user PTY
func getRealUserByPTY() (string, int, int) {
	ownerID, ok := getTDOwnerID()

	if !ok {
		return "", -1, -1
	}

	realUser, err := getUserInfo(strconv.Itoa(ownerID))

	if err != nil {
		return "", -1, -1
	}

	return realUser.Name, realUser.UID, realUser.GID
}

// getRealUserFromEnv try to find info about real user in environment variables
func getRealUserFromEnv() (string, int, int) {
	envMap := env.Get()

	if envMap["SUDO_USER"] == "" || envMap["SUDO_UID"] == "" || envMap["SUDO_GID"] == "" {
		return "", -1, -1
	}

	user := envMap["SUDO_USER"]
	uid, _ := strconv.Atoi(envMap["SUDO_UID"])
	gid, _ := strconv.Atoi(envMap["SUDO_GID"])

	return user, uid, gid
}

// getOwner returns file or directory owner UID
func getOwner(path string) (int, error) {
	if path == "" {
		return -1, ErrEmptyPath
	}

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, err
	}

	return int(stat.Uid), nil
}

// extractGroupsFromIdInfo extracts info from id command output
func extractGroupsInfo(data string) []*Group {
	var field int
	var result []*Group

	data = strings.TrimRight(data, "\n")
	groupsInfo := strutil.ReadField(data, 3, false, '=')

	if groupsInfo == "" {
		return nil
	}

	for {
		groupInfo := strutil.ReadField(groupsInfo, field, false, ',')

		if groupInfo == "" {
			break
		}

		group, err := parseGroupInfo(groupInfo)

		if err == nil {
			result = append(result, group)
		}

		field++
	}

	return result
}

// parseGroupInfo parse group info from 'id' command
func parseGroupInfo(data string) (*Group, error) {
	id := strutil.ReadField(data, 0, false, '(')
	name := strutil.ReadField(data, 1, false, '(')
	gid, _ := strconv.Atoi(id)

	if len(name) == 0 {
		group, err := LookupGroup(id)

		if err != nil {
			return nil, err
		}

		return group, nil
	}

	return &Group{GID: gid, Name: strutil.Substring(name, 0, -1)}, nil
}
