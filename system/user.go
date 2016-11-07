// +build linux, darwin, !windows

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
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"pkg.re/essentialkaos/ek.v5/env"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _PTS_DIR = "/dev/pts"

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
	User             *User     `json:"user"`
	LoginTime        time.Time `json:"login_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}

// sessionsInfo is slice with SessionInfo
type sessionsInfo []*SessionInfo

// ////////////////////////////////////////////////////////////////////////////////// //

func (s sessionsInfo) Len() int {
	return len(s)
}

func (s sessionsInfo) Less(i, j int) bool {
	return s[i].LoginTime.Unix() < s[j].LoginTime.Unix()
}

func (s sessionsInfo) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Current user info cache
var curUser *User

// ////////////////////////////////////////////////////////////////////////////////// //

// Who return info about all active sessions sorted by login time
func Who() ([]*SessionInfo, error) {
	var result []*SessionInfo

	ptsList := readDir(_PTS_DIR)

	if len(ptsList) == 0 {
		return result, nil
	}

	for _, file := range ptsList {
		if file == "ptmx" {
			continue
		}

		info, err := getSessionInfo(file)

		if err != nil {
			continue
		}

		result = append(result, info)
	}

	if len(result) != 0 {
		sort.Sort(sessionsInfo(result))
	}

	return result, nil
}

// CurrentUser return struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	if len(avoidCache) == 0 && curUser != nil {
		return curUser, nil
	}

	user, err := LookupUser(getCurrentUserName())

	if err != nil {
		return user, err
	}

	if user.Name == "root" {
		appendRealUserInfo(user)
	}

	curUser = user

	return user, nil
}

// LookupUser search user info by given name
func LookupUser(nameOrID string) (*User, error) {
	if nameOrID == "" {
		return nil, errors.New("User name/id can't be blank")
	}

	user, err := getUserInfo(nameOrID)

	if err != nil {
		return nil, err
	}

	appendGroupInfo(user)

	return user, nil
}

// LookupGroup search group info by given name
func LookupGroup(nameOrID string) (*Group, error) {
	if nameOrID == "" {
		return nil, errors.New("Group name/id can't be blank")
	}

	name, gid, err := getGroupInfo(nameOrID)

	return &Group{Name: name, GID: gid}, err
}

// IsUserExist check if user exist on system or not
func IsUserExist(name string) bool {
	cmd := exec.Command("getent", "passwd", name)

	err := cmd.Run()

	if err == nil {
		return true
	}

	return false
}

// IsGroupExist check if group exist on system or not
func IsGroupExist(name string) bool {
	cmd := exec.Command("getent", "group", name)

	err := cmd.Run()

	if err == nil {
		return true
	}

	return false
}

// CurrentTTY return current tty or empty string if error occurred
func CurrentTTY() string {
	pid := strconv.Itoa(os.Getpid())
	fdLink, err := os.Readlink("/proc/" + pid + "/fd/0")

	if err != nil {
		return ""
	}

	return fdLink
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRoot check if current user is root
func (u *User) IsRoot() bool {
	return u.UID == 0 && u.GID == 0
}

// IsSudo check if it user over sudo command
func (u *User) IsSudo() bool {
	return u.IsRoot() && u.RealUID != 0 && u.RealGID != 0
}

// GroupList return slice with user groups names
func (u *User) GroupList() []string {
	var result []string

	for _, group := range u.Groups {
		result = append(result, group.Name)
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getCurrentUserName() string {
	cmd := exec.Command("id", "-un")

	out, err := cmd.Output()

	if err != nil {
		return ""
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")

	return sOut
}

// appendGroupInfo append info about groups
func appendGroupInfo(user *User) {
	cmd := exec.Command("id", user.Name)

	out, err := cmd.Output()

	if err != nil {
		return
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")
	aOut := strings.Split(sOut, "=")

	if len(aOut) < 4 {
		return
	}

	for _, info := range strings.Split(aOut[3], ",") {
		user.Groups = append(user.Groups, parseGroupInfo(info))
	}
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

// getUserInfo return uid associated with current tty
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

// getGroupInfo return group info by name or id
func getGroupInfo(nameOrID string) (string, int, error) {
	cmd := exec.Command("getent", "group", nameOrID)

	out, err := cmd.Output()

	if err != nil {
		return "", -1, fmt.Errorf("Group with this name/id %s is not exist", nameOrID)
	}

	sOut := string(out[:])
	sOut = strings.Trim(sOut, "\n")
	aOut := strings.Split(sOut, ":")

	gid, _ := strconv.Atoi(aOut[1])

	return aOut[0], gid, nil
}

// parseGroupInfo remove bracket symbols, parse value as number and return result
func parseGroupInfo(info string) *Group {
	ai := strings.Split(info, "(")

	gid, _ := strconv.Atoi(ai[0])
	name := strings.TrimRight(ai[1], ")")

	return &Group{name, gid}
}

// getOwner return file or dir owner uid
func getOwner(path string) (int, error) {
	if path == "" {
		return -1, errors.New("Path is empty")
	}

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, err
	}

	return int(stat.Uid), nil
}

func readDir(dir string) []string {
	fd, err := syscall.Open(dir, syscall.O_CLOEXEC, 0644)

	if err != nil {
		return []string{}
	}

	var size = 100
	var n = -1

	var nbuf int
	var bufp int

	var buf = make([]byte, 4096)
	var names = make([]string, 0, size)

	for n != 0 {
		if bufp >= nbuf {
			bufp = 0

			var errno error

			nbuf, errno = fixCount(syscall.ReadDirent(fd, buf))

			if errno != nil {
				return names
			}

			if nbuf <= 0 {
				break
			}
		}

		var nb, nc int
		nb, nc, names = syscall.ParseDirent(buf[bufp:nbuf], n, names)
		bufp += nb
		n -= nc
	}

	return names
}

func fixCount(n int, err error) (int, error) {
	if n < 0 {
		n = 0
	}
	return n, err
}

func getSessionInfo(pts string) (*SessionInfo, error) {
	ptsFile := _PTS_DIR + "/" + pts
	uid, err := getOwner(ptsFile)

	if err != nil {
		return nil, err
	}

	user, err := getUserInfo(strconv.Itoa(uid))

	if err != nil {
		return nil, err
	}

	_, mtime, ctime, err := getTimes(ptsFile)

	if err != nil {
		return nil, err
	}

	return &SessionInfo{
		User:             user,
		LoginTime:        ctime,
		LastActivityTime: mtime,
	}, nil
}
