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
	"sync"
	"syscall"
	"time"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyPath is returned when a required path argument is empty
	ErrEmptyPath = errors.New("path is empty")

	// ErrEmptyUserName is returned when a required user name or UID is empty
	ErrEmptyUserName = errors.New("user name/ID can't be blank")

	// ErrEmptyGroupName is returned when a required group name or GID is empty
	ErrEmptyGroupName = errors.New("group name/ID can't be blank")

	// ErrCantParseIdOutput is returned when the output of the id command has an unexpected format
	ErrCantParseIdOutput = errors.New("can't parse id command output")

	// ErrCantParseGetentOutput is returned when the output of getent has an unexpected format
	ErrCantParseGetentOutput = errors.New("can't parse getent command output")
)

// CurrentUserCachePeriod is the duration for which the current user info is cached
var CurrentUserCachePeriod = 5 * time.Minute

// ////////////////////////////////////////////////////////////////////////////////// //

// Current is user info cache
var curUser *User

// curUserUpdateDate is date when user data was updated
var curUserUpdateDate time.Time

// curUserMu mutex to control access to current user data
var curUserMu sync.RWMutex

// ////////////////////////////////////////////////////////////////////////////////// //

// CurrentUser returns information about the user running the current process.
// Results are cached for CurrentUserCachePeriod; pass true to bypass the cache.
func CurrentUser(avoidCache ...bool) (*User, error) {
	if len(avoidCache) == 0 || !avoidCache[0] {
		if curUser != nil && time.Since(curUserUpdateDate) < CurrentUserCachePeriod {
			curUserMu.RLock()
			defer curUserMu.RUnlock()
			return curUser, nil
		}
	}

	curUserMu.Lock()
	defer curUserMu.Unlock()

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

// LookupUser returns user information for the given username or UID string
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

// LookupGroup returns group information for the given group name or GID string
func LookupGroup(nameOrID string) (*Group, error) {
	if nameOrID == "" {
		return nil, ErrEmptyGroupName
	}

	return getGroupInfo(nameOrID)
}

// CurrentTTY returns the path of the TTY attached to the current process, or an empty
// string if none
func CurrentTTY() string {
	pid := strconv.Itoa(os.Getpid())
	fdLink, err := os.Readlink("/proc/" + pid + "/fd/0")

	if err != nil {
		return ""
	}

	return fdLink
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRoot returns true if the user is root (UID 0 and GID 0)
func (u *User) IsRoot() bool {
	return u.UID == 0 && u.GID == 0
}

// IsSudo returns true if the process is running under sudo elevation
func (u *User) IsSudo() bool {
	return u.IsRoot() && u.RealUID != 0 && u.RealGID != 0
}

// GroupList returns the names of all groups the user belongs to
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
	userName := os.Getenv("SUDO_USER")
	userUID := os.Getenv("SUDO_UID")
	userGID := os.Getenv("SUDO_GID")

	if userName == "" || userUID == "" || userGID == "" {
		return "", -1, -1
	}

	uid, err := strconv.Atoi(userUID)

	if err != nil {
		return "", -1, -1
	}

	gid, err := strconv.Atoi(userGID)

	if err != nil {
		return "", -1, -1
	}

	return userName, uid, gid
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
