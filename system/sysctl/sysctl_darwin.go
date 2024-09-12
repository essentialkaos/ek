package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os/exec"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// getKernelParam reads kernel parameter from procfs
func getKernelParam(param string) (string, error) {
	output, err := exec.Command(binary, "-n", param).Output()
	return strings.Trim(string(output), "\n\r"), err
}
