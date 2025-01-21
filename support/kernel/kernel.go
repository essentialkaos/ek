//go:build linux || darwin
// +build linux darwin

// Package kernel provides methods for collecting information from OS kernel
package kernel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/support"
	"github.com/essentialkaos/ek/v13/system/sysctl"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects info from OS kernel
func Collect(params ...string) []support.KernelParam {
	kernelParams, err := sysctl.All()

	if err != nil {
		return nil
	}

	var result []support.KernelParam

	for _, param := range params {
		isGlob := strings.HasSuffix(param, "*")
		param = strings.TrimRight(param, "*")

		for k, v := range kernelParams {
			if isGlob {
				if !strings.HasPrefix(k, param) {
					continue
				}
			} else {
				if k != param {
					continue
				}
			}

			value := strings.ReplaceAll(v, "\t", " ")
			value = strutil.SqueezeRepeats(value, " ")

			result = append(result, support.KernelParam{
				Key:   k,
				Value: value,
			})
		}
	}

	return result
}
