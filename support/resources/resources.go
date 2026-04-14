//go:build linux || darwin

// Package resources provides methods for collecting information about system
// resources (cpu/memory)
package resources

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v14/support"
	"github.com/essentialkaos/ek/v14/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects CPU and memory usage info
func Collect() *support.ResourcesInfo {
	result := &support.ResourcesInfo{}

	cpuInfo, err1 := system.GetCPUInfo()

	if err1 == nil {
		for _, p := range cpuInfo {
			threads := 0

			if p.Cores > 0 {
				threads = p.Siblings / p.Cores
			}

			result.CPU = append(result.CPU, support.CPUInfo{
				Model:   p.Model,
				Cores:   p.Cores,
				Threads: threads,
			})
		}
	}

	memInfo, err2 := system.GetMemUsage()

	if err2 == nil {
		result.MemTotal = memInfo.MemTotal
		result.MemFree = memInfo.MemFree
		result.MemUsed = memInfo.MemUsed
		result.SwapTotal = memInfo.SwapTotal
		result.SwapFree = memInfo.SwapFree
		result.SwapUsed = memInfo.SwapUsed
	}

	if err1 != nil && err2 != nil {
		return nil
	}

	return result
}
