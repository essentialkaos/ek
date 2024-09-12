//go:build !linux && !darwin
// +build !linux,!darwin

// Package sysctl provides methods for reading kernel parameters
package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns kernel parameter value as a string
func Get(param string) (string, error) {
	return "", nil
}

// GetI returns kernel parameter value as an int
func GetI(param string) (int, error) {
	return 0, nil
}

// GetI64 returns kernel parameter value as an int64
func GetI64(param string) (int64, error) {
	return 0, nil
}
