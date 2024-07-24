//go:build !linux || !darwin || !freebsd
// +build !linux !darwin !freebsd

// Package env provides methods for working with environment variables
package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Which find full path to some app
func Which(name string) string {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //
