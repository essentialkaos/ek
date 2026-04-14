//go:build !linux && !darwin

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Who returns info about all active sessions sorted by login time
func Who() ([]*SessionInfo, error) {
	panic("UNSUPPORTED")
}

// ❗ CurrentUser returns struct with info about current user
func CurrentUser(avoidCache ...bool) (*User, error) {
	panic("UNSUPPORTED")
}

// ❗ LookupUser searches user info by given name
func LookupUser(name string) (*User, error) {
	panic("UNSUPPORTED")
}

// ❗ LookupGroup searches group info by given name
func LookupGroup(name string) (*Group, error) {
	panic("UNSUPPORTED")
}

// ❗ IsUserExist checks if user exist on system or not
func IsUserExist(name string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsGroupExist checks if group exist on system or not
func IsGroupExist(name string) bool {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsRoot checks if current user is root
func (u *User) IsRoot() bool {
	panic("UNSUPPORTED")
}

// ❗ IsSudo checks if it user over sudo command
func (u *User) IsSudo() bool {
	panic("UNSUPPORTED")
}

// ❗ GroupList returns slice with user groups names
func (u *User) GroupList() []string {
	panic("UNSUPPORTED")
}
