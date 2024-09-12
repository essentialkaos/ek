//go:build linux || darwin
// +build linux darwin

// Package sysctl provides methods for reading kernel parameters
package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

// Get returns kernel parameter value as a string
func Get(param string) (string, error) {
	switch {
	case param == "":
		return "", fmt.Errorf("Kernel parameter name cannot be empty")
	case !strings.Contains(param, ".") || strings.ContainsAny(param, " /\n\t"):
		return "", fmt.Errorf("Invalid parameter name %q", param)
	}

	return getKernelParam(param)
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
