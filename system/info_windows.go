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

// GetMemInfo return memory info
func GetMemInfo() (*MemInfo, error) {
	return nil, nil
}

// GetCPUInfo return info about CPU usage
func GetCPUInfo() (*CPUInfo, error) {
	return nil, nil
}

// GetCPUStats return basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	return nil, nil
}

// GetFSInfo return info about mounted filesystems
func GetFSInfo() (map[string]*FSInfo, error) {
	return map[string]*FSInfo{"/": {}}, nil
}

// GetIOStats return I/O stats
func GetIOStats() (map[string]*IOStats, error) {
	return map[string]*IOStats{"/dev/sda1": {}}, nil
}

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	return nil, nil
}

// GetInterfacesInfo return info about network interfaces
func GetInterfacesInfo() (map[string]*InterfaceInfo, error) {
	return map[string]*InterfaceInfo{"eth0": {}}, nil
}

// GetNetworkSpeed return input/output speed in bytes per second
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	return 0, 0, nil
}

// CalculateNetworkSpeed calculate network input/output speed in bytes per second for
// all network interfaces
func CalculateNetworkSpeed(ii1, ii2 map[string]*InterfaceInfo, duration time.Duration) (uint64, uint64) {
	return 0, 0
}

// GetIOUtil return IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	return map[string]float64{"/": 0}, nil
}

// CalculateIOUtil calculate IO utilization for all devices
func CalculateIOUtil(fi1 map[string]*FSInfo, fi2 map[string]*FSInfo, duration time.Duration) map[string]float64 {
	return map[string]float64{"/": 0}
}

// ////////////////////////////////////////////////////////////////////////////////// //
