package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

// getParam reads kernel parameter from /proc/sys
func getParam(name string) (Param, error) {
	p, err := os.ReadFile(path.Clean(path.Join(
		procFS, strings.ReplaceAll(name, ".", "/"),
	)))

	if err != nil {
		return Param{}, fmt.Errorf("Can't read parameter %q: %w", name, err)
	}

	return Param{
		Name:  name,
		Value: strings.Trim(string(p), "\n\r"),
	}, nil
}

// readParam returns parameter name and value from sysctl output
func readParam(text string) (string, string, bool) {
	return strings.Cut(strings.Trim(text, "\n\r"), " = ")
}
