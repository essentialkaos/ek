package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetCPUUsage returns info about CPU usage
func GetCPUUsage(duration time.Duration) (*CPUUsage, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ CalculateCPUUsage calcualtes CPU usage based on CPUStats
func CalculateCPUUsage(c1, c2 *CPUStats) *CPUUsage {
	panic("UNSUPPORTED")
	return nil
}

// ❗ GetCPUStats returns basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ GetCPUInfo returns slice with info about CPUs
func GetCPUInfo() ([]*CPUInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}
