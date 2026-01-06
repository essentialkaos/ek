//go:build linux || darwin

// Package sysctl provides methods for reading kernel parameters
package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// binary is sysctl binary name
var binary = "sysctl"

// procFS is path to procfs
var procFS = "/proc/sys"

// ////////////////////////////////////////////////////////////////////////////////// //

// Params contains all kernel parameters
type Params map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// All returns all kernel parameters
func All() (Params, error) {
	return getParams()
}

// Get returns kernel parameter value as a string
func Get(param string) (string, error) {
	switch {
	case param == "":
		return "", fmt.Errorf("Kernel parameter name cannot be empty")
	case !strings.Contains(param, ".") || strings.ContainsAny(param, " /\n\t"):
		return "", fmt.Errorf("Invalid parameter name %q", param)
	}

	return getParam(param)
}

// GetI returns kernel parameter value as an int
func GetI(param string) (int, error) {
	p, err := Get(param)

	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(p)

	if err != nil {
		return 0, fmt.Errorf("Can't parse %q parameter as int: %w", param, err)
	}

	return i, nil
}

// GetI64 returns kernel parameter value as an int64
func GetI64(param string) (int64, error) {
	p, err := Get(param)

	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(p, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Can't parse %q parameter as int64: %w", param, err)
	}

	return i, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns kernel parameter value as a string
func (p Params) Get(name string) string {
	if len(p) == 0 {
		return ""
	}

	return p[name]
}

// GetI returns kernel parameter value as an int
func (p Params) GetI(param string) (int, error) {
	i, err := strconv.Atoi(p.Get(param))

	if err != nil {
		return 0, fmt.Errorf("Can't parse %q parameter as int: %w", param, err)
	}

	return i, nil
}

// GetI64 returns kernel parameter value as an int64
func (p Params) GetI64(param string) (int64, error) {
	i, err := strconv.ParseInt(p.Get(param), 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Can't parse %q parameter as int64: %w", param, err)
	}

	return i, nil
}
