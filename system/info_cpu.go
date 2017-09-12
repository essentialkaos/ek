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
	"time"

	"pkg.re/essentialkaos/ek.v9/errutil"
	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with CPU info in procfs
var procStatFile = "/proc/stat"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUInfo return info about CPU usage
func GetCPUInfo(duration time.Duration) (*CPUInfo, error) {
	c1, err := GetCPUStats()

	if err != nil {
		return nil, err
	}

	time.Sleep(duration)

	c2, err := GetCPUStats()

	if err != nil {
		return nil, err
	}

	return CalculateCPUInfo(c1, c2), nil
}

func CalculateCPUInfo(c1, c2 *CPUStats) *CPUInfo {
	prevIdle := c1.Idle + c1.Wait
	idle := c2.Idle + c2.Wait

	prevNonIdle := c1.User + c1.Nice + c1.System + c1.IRQ + c1.SRQ + c1.Steal
	nonIdle := c2.User + c2.Nice + c2.System + c2.IRQ + c2.SRQ + c2.Steal

	prevTotal := prevIdle + prevNonIdle
	total := idle + nonIdle

	totalDiff := float64(total - prevTotal)
	idleDiff := float64(idle - prevIdle)
	allTotalDiff := float64(c2.Total - c1.Total)

	return &CPUInfo{
		System:  (float64(c2.System-c1.System) / allTotalDiff) * 100,
		User:    (float64(c2.User-c1.User) / allTotalDiff) * 100,
		Nice:    (float64(c2.Nice-c1.Nice) / allTotalDiff) * 100,
		Wait:    (float64(c2.Wait-c1.Wait) / allTotalDiff) * 100,
		Idle:    (float64(c2.Idle-c1.Idle) / allTotalDiff) * 100,
		Average: ((totalDiff - idleDiff) / totalDiff) * 100.0,
		Count:   c2.Count,
	}
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
				stats.User = parseUint(strutil.ReadField(text, 1, true), errs)
				stats.Nice = parseUint(strutil.ReadField(text, 2, true), errs)
				stats.System = parseUint(strutil.ReadField(text, 3, true), errs)
				stats.Idle = parseUint(strutil.ReadField(text, 4, true), errs)
				stats.Wait = parseUint(strutil.ReadField(text, 5, true), errs)
				stats.IRQ = parseUint(strutil.ReadField(text, 6, true), errs)
				stats.SRQ = parseUint(strutil.ReadField(text, 7, true), errs)
				stats.Steal = parseUint(strutil.ReadField(text, 8, true), errs)
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
