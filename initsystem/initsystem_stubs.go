//go:build !linux && !darwin
// +build !linux,!darwin

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ SysV returns true if SysV is used on system
func SysV() bool {
	panic("UNSUPPORTED")
}

// ❗ Upstart returns true if Upstart is used on system
func Upstart() bool {
	panic("UNSUPPORTED")
}

// ❗ Systemd returns true if Systemd is used on system
func Systemd() bool {
	panic("UNSUPPORTED")
}

// ❗ Launchd returns true if Launchd is used on the system
func Launchd() bool {
	panic("UNSUPPORTED")
}

// ❗ IsPresent returns true if service is present in any init system
func IsPresent(name string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsWorks returns service state
func IsWorks(name string) (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ IsEnabled returns true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	panic("UNSUPPORTED")
}
