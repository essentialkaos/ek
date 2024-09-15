package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"

	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/system/sysctl"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetCPUUsage returns info about CPU usage
func GetCPUUsage(duration time.Duration) (*CPUUsage, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ CalculateCPUUsage calculates CPU usage based on CPUStats
func CalculateCPUUsage(c1, c2 *CPUStats) *CPUUsage {
	panic("UNSUPPORTED")
	return nil
}

// ❗ GetCPUStats returns basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUInfo returns slice with info about CPUs
func GetCPUInfo() ([]*CPUInfo, error) {
	params, err := sysctl.All()

	if err != nil {
		return nil, fmt.Errorf("Can't get kernel parameters: %w", err)
	}

	p0cache, _ := params.GetI("hw.perflevel0.l2cachesize")
	p1cache, _ := params.GetI("hw.perflevel1.l2cachesize")

	cores, _ := params.GetI("machdep.cpu.core_count")
	threads, _ := params.GetI("machdep.cpu.thread_count")

	return []*CPUInfo{
		{
			Vendor:    strutil.Q(params.Get("machdep.cpu.vendor"), "Apple"),
			Model:     params.Get("machdep.cpu.brand_string"),
			Cores:     cores,
			Siblings:  threads,
			CacheSize: uint64(p0cache + p1cache),
		},
	}, nil
}
