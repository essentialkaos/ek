//go:build !linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetTree returns root process with all subprocesses on the system
func GetTree(pid ...int) (*ProcessInfo, error) {
	panic("UNSUPPORTED")
}

// ❗ GetList returns slice with all active processes on the system
func GetList() ([]*ProcessInfo, error) {
	panic("UNSUPPORTED")
}
