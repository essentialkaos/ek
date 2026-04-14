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
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// binary is sysctl binary name
var binary = "sysctl"

// procFS is path to procfs
var procFS = "/proc/sys"

// ////////////////////////////////////////////////////////////////////////////////// //

// All returns all kernel parameters available on the current system
func All() (Params, error) {
	return getParams()
}

// Get returns the kernel parameter with the given name. The name must be dot-separated
// (e.g. "kernel.pid_max") and must not contain spaces or slashes.
func Get(name string) (Param, error) {
	switch {
	case name == "":
		return Param{}, fmt.Errorf("parameter name is empty")
	case !strings.Contains(name, ".") || strings.ContainsAny(name, " /\n\t"):
		return Param{}, fmt.Errorf("invalid parameter name %q", name)
	}

	return getParam(name)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getParams reads all kernel parameters from sysctl binary output
func getParams() (Params, error) {
	output, err := exec.Command(binary, "-a").Output()

	if err != nil {
		return nil, fmt.Errorf("can't get kernel parameters from sysctl")
	}

	var result Params

	buf := bytes.NewBuffer(output)

	for {
		line, err := buf.ReadString('\n')
		name, value, ok := readParam(line)

		if ok && value != "" {
			result = append(result, Param{name, value})
		}

		if err != nil {
			break
		}
	}

	return result, nil
}
