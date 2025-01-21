//go:build !linux && !darwin
// +build !linux,!darwin

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Push changes current working directory and add previous working directory to stack
func Push(dir string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ Pop changes current working directory to previous in stack
func Pop() string {
	panic("UNSUPPORTED")
	return ""
}
