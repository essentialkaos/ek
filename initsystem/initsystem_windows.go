// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// SysV if SysV is used on system
func SysV() bool {
	panic("UNSUPPORTED")
	return false
}

// Upstart if Upstart is used on system
func Upstart() bool {
	panic("UNSUPPORTED")
	return false
}

// Systemd if Systemd is used on system
func Systemd() bool {
	panic("UNSUPPORTED")
	return false
}

// IsPresent returns true if service is present in any init system
func IsPresent(name string) bool {
	panic("UNSUPPORTED")
	return false
}

// IsWorks returns service state
func IsWorks(name string) (bool, error) {
	panic("UNSUPPORTED")
	return false, nil
}

// IsEnabled returns true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	panic("UNSUPPORTED")
	return false, nil
}
