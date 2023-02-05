package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ User contains information about user
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

// ❗ Group contains information about group
type Group struct {
	Name string `json:"name"`
	GID  int    `json:"gid"`
}

// ❗ SessionInfo contains information about all sessions
type SessionInfo struct {
	User             *User     `json:"user"`
	LoginTime        time.Time `json:"login_time"`
	LastActivityTime time.Time `json:"last_activity_time"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Who returns info about all active sessions sorted by login time
func Who() ([]*SessionInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ CurrentUser returns struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ LookupUser searches user info by given name
func LookupUser(name string) (*User, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ LookupGroup searches group info by given name
func LookupGroup(name string) (*Group, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ IsUserExist checks if user exist on system or not
func IsUserExist(name string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsGroupExist checks if group exist on system or not
func IsGroupExist(name string) bool {
	panic("UNSUPPORTED")
	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsRoot checks if current user is root
func (u *User) IsRoot() bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsSudo checks if it user over sudo command
func (u *User) IsSudo() bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ GroupList returns slice with user groups names
func (u *User) GroupList() []string {
	panic("UNSUPPORTED")
	return nil
}
