//+build windows

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// SysV if SysV is used on system
func SysV() bool {
	return false
}

// Upstart if Upstart is used on system
func Upstart() bool {
	return false
}

// Systemd if Systemd is used on system
func Systemd() bool {
	return false
}

// HasService return true if service is present in any init system
func HasService(name string) bool {
	return false
}

// IsServiceWorks return service state
func IsServiceWorks(name string) (bool, error) {
	return false, nil
}

// IsEnabled return true if auto start enabled for given service
func IsEnabled(name string) (bool, error) {
	return false, nil
}
