// +build !windows

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

// basicCPUInfo contains basic CPU metrics
type basicCPUInfo struct {
	User   uint64
	Nice   uint64
	System uint64
	Idle   uint64
	Wait   uint64
	IRQ    uint64
	SRQ    uint64
	Steal  uint64
	Total  uint64
	Count  int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with CPU info in procfs
var procStatFile = "/proc/stat"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUInfo return info about CPU usage
func GetCPUInfo() (*CPUInfo, error) {
	info, err := getCPUStats()

	if err != nil {
		return nil, err
	}

	return &CPUInfo{
		System: (float64(info.System) / float64(info.Total)) * 100,
		User:   (float64(info.User) / float64(info.Total)) * 100,
		Nice:   (float64(info.Nice) / float64(info.Total)) * 100,
		Wait:   (float64(info.Wait) / float64(info.Total)) * 100,
		Idle:   (float64(info.Idle) / float64(info.Total)) * 100,
		Count:  info.Count,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getCPUStats return basicCPUInfo
func getCPUStats() (basicCPUInfo, error) {
	info := basicCPUInfo{}

	content, err := readFileContent(procStatFile)

	if err != nil {
		return info, errors.New("Can't parse file " + procStatFile)
	}

	for _, line := range content {
		if strings.HasPrefix(line, "cpu") {
			info.Count++
		}
	}

	info.Count--

	cpuInfo := strings.Replace(content[0], "cpu  ", "", -1)
	cpuInfoSlice := strings.Split(cpuInfo, " ")

	if len(cpuInfoSlice) < 8 {
		return info, errors.New("Can't parse file " + procStatFile)
	}

	info.User, _ = strconv.ParseUint(cpuInfoSlice[0], 10, 64)
	info.Nice, _ = strconv.ParseUint(cpuInfoSlice[1], 10, 64)
	info.System, _ = strconv.ParseUint(cpuInfoSlice[2], 10, 64)
	info.Idle, _ = strconv.ParseUint(cpuInfoSlice[3], 10, 64)
	info.Wait, _ = strconv.ParseUint(cpuInfoSlice[4], 10, 64)
	info.IRQ, _ = strconv.ParseUint(cpuInfoSlice[5], 10, 64)
	info.SRQ, _ = strconv.ParseUint(cpuInfoSlice[6], 10, 64)
	info.Steal, _ = strconv.ParseUint(cpuInfoSlice[7], 10, 64)

	info.Total = info.User + info.System + info.Nice + info.Idle + info.Wait + info.IRQ + info.SRQ + info.Steal

	return info, nil
}
