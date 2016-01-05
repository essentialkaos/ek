// +build !linux, !darwin, windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// User contains information about user
type User struct {
	UID      int
	GID      int
	Name     string
	Groups   []*Group
	Comment  string
	Shell    string
	HomeDir  string
	RealUID  int
	RealGID  int
	RealName string
}

// Group contains information about group
type Group struct {
	Name string
	GID  int
}

// SessionInfo contains information about all sessions
type SessionInfo struct {
	Name             string
	LoginTime        time.Time
	LastActivityTime time.Time
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUsername return current user name
func GetUsername() string {
	return ""
}

// GetGroupname return current user group name
func GetGroupname() string {
	return ""
}

// Who return info about all active sessions sorted by login time
func Who() ([]*SessionInfo, error) {
	return []*SessionInfo{}, nil
}

// CurrentUser return struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	return &User{}, nil
}

// LookupUser search user info by given name
func LookupUser(name string) (*User, error) {
	return &User{}, nil
}

// LookupGroup search group info by given name
func LookupGroup(name string) (*Group, error) {
	return &Group{}, nil
}

// IsUserExist check if user exist on system or not
func IsUserExist(name string) bool {
	return false
}

// IsGroupExist check if group exist on system or not
func IsGroupExist(name string) bool {
	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsRoot check if current user is root
func (u *User) IsRoot() bool {
	return false
}

// IsSudo check if it user over sudo command
func (u *User) IsSudo() bool {
	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //
