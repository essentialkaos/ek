//go:build linux
// +build linux

// Package resources provides methods for collecting information about system
// resources (cpu/memory)
package resources

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/support"
	"github.com/essentialkaos/ek/v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Collect collects info about CPU and memory
func Collect(extMemInfo bool) *support.ResourcesInfo {
	result := &support.ResourcesInfo{}

	cpuInfo, err1 := system.GetCPUInfo()

	if err1 == nil {
		for _, p := range cpuInfo {
			result.CPU = append(result.CPU, support.CPUInfo{
				Model:   p.Model,
				Cores:   p.Cores,
				Threads: p.Siblings / p.Cores,
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
