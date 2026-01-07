//go:build !linux && !darwin

// Package sysctl provides methods for reading kernel parameters
package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Params contains all kernel parameters
type Params map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// All returns all kernel parameters
func All() (Params, error) {
	return nil, nil
}

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

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns kernel parameter value as a string
func (p Params) Get(name string) string {
	return ""
}

// GetI returns kernel parameter value as an int
func (p Params) GetI(param string) (int, error) {
	return 0, nil
}

// GetI64 returns kernel parameter value as an int64
func (p Params) GetI64(param string) (int64, error) {
	return 0, nil
}
