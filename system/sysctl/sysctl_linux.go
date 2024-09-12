package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"

	"github.com/essentialkaos/ek/v13/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getKernelParam reads kernel parameter from procfs
func getKernelParam(param string) (string, error) {
	p, err := os.ReadFile(path.Clean(path.Join(
		procFS, strings.ReplaceAll(param, ".", "/"),
	)))

	if err != nil {
		return "", fmt.Errorf("Can't read parameter %q: %w", param, err)
	}

	return strings.Trim(string(p), "\n\r"), nil
}
