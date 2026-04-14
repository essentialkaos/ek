package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"

	"github.com/essentialkaos/ek/v14/strutil"
	"github.com/essentialkaos/ek/v14/system/sysctl"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetCPUUsage measures CPU usage over the given duration and returns a usage
// breakdown
func GetCPUUsage(duration time.Duration) (*CPUUsage, error) {
	panic("UNSUPPORTED")
}

// ❗ CalculateCPUUsage calculates CPU usage percentages from two consecutive
// CPUStats snapshots
func CalculateCPUUsage(c1, c2 *CPUStats) *CPUUsage {
	panic("UNSUPPORTED")
}

// ❗ GetCPUStats returns a snapshot of raw cumulative CPU time counters
func GetCPUStats() (*CPUStats, error) {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUInfo returns static information about each physical CPU package
func GetCPUInfo() ([]*CPUInfo, error) {
	params, err := sysctl.All()

	if err != nil {
		return nil, fmt.Errorf("can't get kernel parameters: %w", err)
	}

	p0cache, _ := params.Get("hw.perflevel0.l2cachesize").Int()
	p1cache, _ := params.Get("hw.perflevel1.l2cachesize").Int()
	cores, _ := params.Get("machdep.cpu.core_count").Int()
	threads, _ := params.Get("machdep.cpu.thread_count").Int()

	return []*CPUInfo{
		{
			Vendor:    strutil.Q(params.Get("machdep.cpu.vendor").String(), "Apple"),
			Model:     params.Get("machdep.cpu.brand_string").String(),
			Cores:     cores,
			Siblings:  threads,
			CacheSize: uint64(p0cache + p1cache),
		},
	}, nil
}
