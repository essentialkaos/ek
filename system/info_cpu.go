// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"

	"pkg.re/essentialkaos/ek.v9/errutil"
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
	fd, err := os.OpenFile(procStatFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	stats := &CPUStats{}
	errs := errutil.NewErrors()

	for s.Scan() {
		text := s.Text()

		if len(text) > 4 && text[:3] == "cpu" {
			if text[:4] == "cpu " {
				stats.User = parseUint(readField(text, 1), errs)
				stats.Nice = parseUint(readField(text, 2), errs)
				stats.System = parseUint(readField(text, 3), errs)
				stats.Idle = parseUint(readField(text, 4), errs)
				stats.Wait = parseUint(readField(text, 5), errs)
				stats.IRQ = parseUint(readField(text, 6), errs)
				stats.SRQ = parseUint(readField(text, 7), errs)
				stats.Steal = parseUint(readField(text, 8), errs)
			} else {
				stats.Count++
				continue
			}

			if errs.HasErrors() {
				return nil, errs.Last()
			}
		}
	}

	if stats.Count == 0 {
		return nil, errors.New("Can't parse file " + procStatFile)
	}

	stats.Total = stats.User + stats.System + stats.Nice + stats.Idle + stats.Wait + stats.IRQ + stats.SRQ + stats.Steal

	return stats, nil
}
