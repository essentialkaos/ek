//go:build !linux && !darwin

// Package netutil provides methods for working with network
package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetIP returns main IPv4 address
func GetIP() string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ GetIP6 returns main IPv6 address
func GetIP6() string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ GetAllIP returns all IPv4 addresses
func GetAllIP() []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ GetAllIP6 returns all IPv6 addresses
func GetAllIP6() []string {
	panic("UNSUPPORTED")
	return nil
}
