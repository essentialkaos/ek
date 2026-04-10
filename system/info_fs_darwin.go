package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetFSUsage returns usage statistics for all currently mounted filesystems
func GetFSUsage() (map[string]*FSUsage, error) {
	panic("UNSUPPORTED")
	return map[string]*FSUsage{"/": {}}, nil
}

// ❗ GetIOStats returns current I/O counters keyed by block device name
func GetIOStats() (map[string]*IOStats, error) {
	panic("UNSUPPORTED")
	return map[string]*IOStats{"/dev/sda1": {}}, nil
}

// ❗ GetIOUtil measures I/O utilization per device over the given duration and returns
// values as percentages
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	panic("UNSUPPORTED")
	return map[string]float64{"/": 0}, nil
}

// ❗ CalculateIOUtil calculates I/O utilization percentages from two IOStats
// snapshots
func CalculateIOUtil(io1, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
	panic("UNSUPPORTED")
	return map[string]float64{"/": 0}
}
