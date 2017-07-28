//+build windows

// Package initsystem provides methods for working with different init systems
package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Service states codes
const (
	STATE_STOPPED       = 0
	STATE_WORKS         = 1
	STATE_UNKNOWN uint8 = 255
)

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

// GetServiceState return service state
func GetServiceState(name string) uint8 {
	return STATE_UNKNOWN
}
