package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
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
	User             *User     `json:"user"`
	LoginTime        time.Time `json:"login_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Who return info about all active sessions sorted by login time
func Who() ([]*SessionInfo, error) {
	return nil, nil
}

// CurrentUser return struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	return nil, nil
}

// LookupUser search user info by given name
func LookupUser(name string) (*User, error) {
	return nil, nil
}

// LookupGroup search group info by given name
func LookupGroup(name string) (*Group, error) {
	return nil, nil
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

// GroupList return slice with user groups names
func (u *User) GroupList() []string {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
