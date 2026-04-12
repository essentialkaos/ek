package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getParam reads kernel parameter by name
func getParam(name string) (Param, error) {
	output, err := exec.Command(binary, "-n", name).Output()

	if err != nil {
		return Param{}, fmt.Errorf("can't get kernel parameters from sysctl")
	}

	return Param{
		Name:  name,
		Value: strings.Trim(string(p), "\n\r"),
	}, nil
}

// readParam returns parameter name and value from sysctl output
func readParam(text string) (string, string, bool) {
	return strings.Cut(strings.Trim(text, "\n\r"), ": ")
}
