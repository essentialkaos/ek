//go:build linux || darwin

// Package kernel provides methods for collecting information from OS kernel
package kernel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
	"strings"

	"github.com/essentialkaos/ek/v14/sortutil"
	"github.com/essentialkaos/ek/v14/strutil"
	"github.com/essentialkaos/ek/v14/support"
	"github.com/essentialkaos/ek/v14/system/sysctl"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect returns kernel parameters matching the given names or prefix patterns.
// Patterns ending with "*" are treated as prefix globs (e.g. "vm.*" matches all vm
// parameters). Returns nil if no params match or if the kernel cannot be queried.
func Collect(params ...string) []support.KernelParam {
	kernelParams, err := sysctl.All()

	if err != nil {
		return nil
	}

	sort.Slice(kernelParams, func(i, j int) bool {
		return sortutil.NaturalLess(kernelParams[i].Name, kernelParams[j].Name)
	})

	var result []support.KernelParam

	for _, pattern := range params {
		isGlob := strings.HasSuffix(pattern, "*")
		param := strings.TrimRight(pattern, "*")

		for _, p := range kernelParams {
			if isGlob {
				if !strings.HasPrefix(p.Name, param) {
					continue
				}
			} else {
				if p.Name != param {
					continue
				}
			}

			value := strings.ReplaceAll(p.Value, "\t", " ")
			value = strutil.SqueezeRepeats(value, " ")

			result = append(result, support.KernelParam{
				Key:   p.Name,
				Value: value,
			})
		}
	}

	return result
}
