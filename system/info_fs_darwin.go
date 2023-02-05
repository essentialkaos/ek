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

// ❗ GetFSUsage returns info about mounted filesystems
func GetFSUsage() (map[string]*FSUsage, error) {
	panic("UNSUPPORTED")
	return map[string]*FSUsage{"/": {}}, nil
}

// ❗ GetIOStats returns I/O stats
func GetIOStats() (map[string]*IOStats, error) {
	panic("UNSUPPORTED")
	return map[string]*IOStats{"/dev/sda1": {}}, nil
}

// ❗ GetIOUtil returns IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	panic("UNSUPPORTED")
	return map[string]float64{"/": 0}, nil
}

// ❗ CalculateIOUtil calculates IO utilization for all devices
func CalculateIOUtil(io1, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
	panic("UNSUPPORTED")
	return map[string]float64{"/": 0}
}
