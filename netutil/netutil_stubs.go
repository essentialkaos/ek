//go:build !linux && !darwin

// Package netutil provides methods for working with network
package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetIP returns the primary IPv4 address of the host
func GetIP() string {
	panic("UNSUPPORTED")
}

// ❗ GetIP6 returns the primary IPv6 address of the host
func GetIP6() string {
	panic("UNSUPPORTED")
}

// ❗ GetAllIP returns all IPv4 addresses across all active network interfaces
func GetAllIP() []string {
	panic("UNSUPPORTED")
}

// ❗ GetAllIP6 returns all IPv6 addresses across all active network interfaces
func GetAllIP6() []string {
	panic("UNSUPPORTED")
}
