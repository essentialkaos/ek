package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/essentialkaos/ek/v13/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getParam reads kernel parameter from procfs
func getParam(param string) (string, error) {
	p, err := os.ReadFile(path.Clean(path.Join(
		procFS, strings.ReplaceAll(param, ".", "/"),
	)))

	if err != nil {
		return "", fmt.Errorf("Can't read parameter %q: %w", param, err)
	}

	return strings.Trim(string(p), "\n\r"), nil
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

		name, value, ok := strings.Cut(strings.Trim(line, "\n\r"), " = ")

		if ok {
			params[name] = value
		}
	}

	return params, nil
}
