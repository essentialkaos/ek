package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

// getParam reads kernel parameter from procfs
func getParam(param string) (string, error) {
	output, err := exec.Command(binary, "-n", param).Output()

	if err != nil {
		return "", fmt.Errorf("Can't get kernel parameters from sysctl")
	}

	return strings.Trim(string(output), "\n\r"), err
}

// getParams reads all kernel parameters
func getParams() (Params, error) {
	output, err := exec.Command(binary, "-a").Output()

	if err != nil {
		return nil, fmt.Errorf("Can't get kernel parameters from sysctl")
	}

	params := make(Params)
	buf := bytes.NewBuffer(output)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		name, value, ok := strings.Cut(strings.Trim(line, "\n\r"), ": ")

		if ok {
			params[name] = value
		}
	}

	return params, nil
}
