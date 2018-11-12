// +build windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime return uptime in seconds from 1/1/1970
func GetUptime() (uint64, error) {
	return 0, nil
}

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	return nil, nil
}

// GetMemUsage return memory usage info
func GetMemUsage() (*MemUsage, error) {
	return nil, nil
}

// GetCPUUsage return info about CPU usage
func GetCPUUsage(duration time.Duration) (*CPUUsage, error) {
	return nil, nil
}

// CalculateCPUUsage calcualtes CPU usage based on CPUStats
func CalculateCPUUsage(c1, c2 *CPUStats) *CPUUsage {
	return nil
}

// GetCPUStats return basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	return nil, nil
}

// GetCPUInfo returns slice with info about CPUs
func GetCPUInfo() ([]*CPUInfo, error) {
	return nil, nil
}

// GetFSUsage return info about mounted filesystems
func GetFSUsage() (map[string]*FSUsage, error) {
	return map[string]*FSUsage{"/": {}}, nil
}

// GetIOStats return I/O stats
func GetIOStats() (map[string]*IOStats, error) {
	return map[string]*IOStats{"/dev/sda1": {}}, nil
}

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	return nil, nil
}

// GetInterfacesStats return info about network interfaces
func GetInterfacesStats() (map[string]*InterfaceStats, error) {
	return map[string]*InterfaceStats{"eth0": {}}, nil
}

// GetNetworkSpeed return input/output speed in bytes per second
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	return 0, 0, nil
}

// CalculateNetworkSpeed calculate network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceStats, duration time.Duration) (uint64, uint64) {
	return 0, 0
}

// GetIOUtil return IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	return map[string]float64{"/": 0}, nil
}

// CalculateIOUtil calculate IO utilization for all devices
func CalculateIOUtil(io1, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
	return map[string]float64{"/": 0}
}

// ////////////////////////////////////////////////////////////////////////////////// //
