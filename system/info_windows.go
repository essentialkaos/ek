// Package system provides methods for working with system data (metrics/users)
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

// ❗ GetUptime returns uptime in seconds from 1/1/1970
func GetUptime() (uint64, error) {
	panic("UNSUPPORTED")
	return 0, nil
}

// ❗ GetLA returns loadavg
func GetLA() (*LoadAvg, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ GetMemUsage returns memory usage info
func GetMemUsage() (*MemUsage, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

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

// ❗ GetSystemInfo returns system info
func GetSystemInfo() (*SystemInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ GetOSInfo returns info about OS
func GetOSInfo() (*OSInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ GetInterfacesStats returns info about network interfaces
func GetInterfacesStats() (map[string]*InterfaceStats, error) {
	panic("UNSUPPORTED")
	return map[string]*InterfaceStats{"eth0": {}}, nil
}

// ❗ GetNetworkSpeed returns input/output speed in bytes per second
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	panic("UNSUPPORTED")
	return 0, 0, nil
}

// ❗ CalculateNetworkSpeed calculates network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceStats, duration time.Duration) (uint64, uint64) {
	panic("UNSUPPORTED")
	return 0, 0
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
