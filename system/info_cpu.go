// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CPUInfo contains info about CPU usage
type CPUInfo struct {
	User   float64 `json:"user"`   // Normal processes executing in user mode
	System float64 `json:"system"` // Processes executing in kernel mode
	Nice   float64 `json:"nice"`   // Niced processes executing in user mode
	Idle   float64 `json:"idle"`   // Twiddling thumbs
	Wait   float64 `json:"wait"`   // Waiting for I/O to complete
	Count  int     `json:"count"`  // Number of CPU cores
}

// CPUStats contains basic CPU stats
type CPUStats struct {
	User   uint64 `json:"user"`
	Nice   uint64 `json:"nice"`
	System uint64 `json:"system"`
	Idle   uint64 `json:"idle"`
	Wait   uint64 `json:"wait"`
	IRQ    uint64 `json:"irq"`
	SRQ    uint64 `json:"srq"`
	Steal  uint64 `json:"steal"`
	Total  uint64 `json:"total"`
	Count  int    `json:"count"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with CPU info in procfs
var procStatFile = "/proc/stat"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUInfo return info about CPU usage
func GetCPUInfo() (*CPUInfo, error) {
	stats, err := GetCPUStats()

	if err != nil {
		return nil, err
	}

	return &CPUInfo{
		System: (float64(stats.System) / float64(stats.Total)) * 100,
		User:   (float64(stats.User) / float64(stats.Total)) * 100,
		Nice:   (float64(stats.Nice) / float64(stats.Total)) * 100,
		Wait:   (float64(stats.Wait) / float64(stats.Total)) * 100,
		Idle:   (float64(stats.Idle) / float64(stats.Total)) * 100,
		Count:  stats.Count,
	}, nil
}

// GetCPUStats return basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	stats := &CPUStats{}

	content, err := readFileContent(procStatFile)

	if err != nil {
		return nil, errors.New("Can't parse file " + procStatFile)
	}

	for _, line := range content {
		if strings.HasPrefix(line, "cpu") {
			stats.Count++
		}
	}

	stats.Count--

	cpuInfo := strings.Replace(content[0], "cpu  ", "", -1)
	cpuInfoSlice := strings.Split(cpuInfo, " ")

	if len(cpuInfoSlice) < 8 {
		return nil, errors.New("Can't parse file " + procStatFile)
	}

	stats.User, _ = strconv.ParseUint(cpuInfoSlice[0], 10, 64)
	stats.Nice, _ = strconv.ParseUint(cpuInfoSlice[1], 10, 64)
	stats.System, _ = strconv.ParseUint(cpuInfoSlice[2], 10, 64)
	stats.Idle, _ = strconv.ParseUint(cpuInfoSlice[3], 10, 64)
	stats.Wait, _ = strconv.ParseUint(cpuInfoSlice[4], 10, 64)
	stats.IRQ, _ = strconv.ParseUint(cpuInfoSlice[5], 10, 64)
	stats.SRQ, _ = strconv.ParseUint(cpuInfoSlice[6], 10, 64)
	stats.Steal, _ = strconv.ParseUint(cpuInfoSlice[7], 10, 64)

	stats.Total = stats.User + stats.System + stats.Nice + stats.Idle + stats.Wait + stats.IRQ + stats.SRQ + stats.Steal

	return stats, nil
}
