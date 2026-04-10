package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with CPU stats in procfs
var procStatFile = "/proc/stat"

// Path to file with info about CPU
var cpuInfoFile = "/proc/cpuinfo"

// Files with CPU info
var (
	cpuPossibleFile = "/sys/devices/system/cpu/possible"
	cpuPresentFile  = "/sys/devices/system/cpu/present"
	cpuOnlineFile   = "/sys/devices/system/cpu/online"
	cpuOfflineFile  = "/sys/devices/system/cpu/offline"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUUsage measures CPU usage over the given duration and returns a usage breakdown
func GetCPUUsage(duration time.Duration) (*CPUUsage, error) {
	c1, err := GetCPUStats()

	if err != nil {
		return nil, err
	}

	time.Sleep(duration)

	c2, err := GetCPUStats()

	if err != nil {
		return nil, err
	}

	return CalculateCPUUsage(c1, c2), nil
}

// CalculateCPUUsage calculates CPU usage percentages from two consecutive
// CPUStats snapshots
func CalculateCPUUsage(c1, c2 *CPUStats) *CPUUsage {
	prevIdle := c1.Idle + c1.Wait
	idle := c2.Idle + c2.Wait

	prevNonIdle := c1.User + c1.Nice + c1.System + c1.IRQ + c1.SRQ + c1.Steal
	nonIdle := c2.User + c2.Nice + c2.System + c2.IRQ + c2.SRQ + c2.Steal

	prevTotal := prevIdle + prevNonIdle
	total := idle + nonIdle

	totalDiff := float64(total - prevTotal)
	idleDiff := float64(idle - prevIdle)
	allTotalDiff := float64(c2.Total - c1.Total)

	if totalDiff == 0 || allTotalDiff == 0 {
		return &CPUUsage{
			Count: c2.Count,
		}
	}

	return &CPUUsage{
		System:  (float64(c2.System-c1.System) / allTotalDiff) * 100,
		User:    (float64(c2.User-c1.User) / allTotalDiff) * 100,
		Nice:    (float64(c2.Nice-c1.Nice) / allTotalDiff) * 100,
		Wait:    (float64(c2.Wait-c1.Wait) / allTotalDiff) * 100,
		Idle:    (float64(c2.Idle-c1.Idle) / allTotalDiff) * 100,
		Average: ((totalDiff - idleDiff) / totalDiff) * 100.0,
		Count:   c2.Count,
	}
}

// GetCPUStats returns a snapshot of raw cumulative CPU time counters
func GetCPUStats() (*CPUStats, error) {
	s, closer, err := getFileScanner(procStatFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseCPUStats(s)
}

// GetCPUInfo returns static information about each physical CPU package
func GetCPUInfo() ([]*CPUInfo, error) {
	s, closeFunc, err := getFileScanner(cpuInfoFile)

	if err != nil {
		return nil, err
	}

	defer closeFunc()

	return parseCPUInfo(s)
}

// GetCPUCount returns the number of CPUs in each availability state
func GetCPUCount() (CPUCount, error) {
	var err error

	result := CPUCount{}

	result.Possible, err = parseCPUCountInfo(cpuPossibleFile)

	if err != nil {
		return CPUCount{}, err
	}

	result.Present, err = parseCPUCountInfo(cpuPresentFile)

	if err != nil {
		return CPUCount{}, err
	}

	result.Online, err = parseCPUCountInfo(cpuOnlineFile)

	if err != nil {
		return CPUCount{}, err
	}

	result.Offline, err = parseCPUCountInfo(cpuOfflineFile)

	if err != nil {
		return CPUCount{}, err
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseCPUStats parses cpu stats data
func parseCPUStats(s *bufio.Scanner) (*CPUStats, error) {
	stats := &CPUStats{}

	for s.Scan() {
		text := s.Text()

		if !strings.HasPrefix(text, "cpu") {
			continue
		}

		if !strings.HasPrefix(text, "cpu ") {
			stats.Count++
			continue
		}

		for i := range 8 {
			v, err := strconv.ParseUint(strutil.ReadField(text, i+1, true), 10, 64)

			if err != nil {
				return nil, fmt.Errorf(
					"can't parse field %d as unsigned integer in %s: %w",
					i+1, procStatFile, err,
				)
			}

			switch i {
			case 0:
				stats.User = v
			case 1:
				stats.Nice = v
			case 2:
				stats.System = v
			case 3:
				stats.Idle = v
			case 4:
				stats.Wait = v
			case 5:
				stats.IRQ = v
			case 6:
				stats.SRQ = v
			case 7:
				stats.Steal = v
			}

			stats.Total += v
		}
	}

	if stats.Count == 0 {
		return nil, fmt.Errorf("procfs file %s has no valid data", procStatFile)
	}

	return stats, nil
}

// parseCPUInfo parses cpu info data
func parseCPUInfo(s *bufio.Scanner) ([]*CPUInfo, error) {
	var (
		err      error
		info     []*CPUInfo
		vendor   string
		model    string
		siblings int
		cores    int
		cache    uint64
		speed    float64
		id       int
	)

	for s.Scan() {
		text := s.Text()

		switch {
		case strings.HasPrefix(text, "vendor_id"):
			vendor = strings.Trim(strutil.ReadField(text, 1, false, ':'), " ")

		case strings.HasPrefix(text, "model name"):
			model = strings.Trim(strutil.ReadField(text, 1, false, ':'), " ")

		case strings.HasPrefix(text, "cache size"):
			cache, err = parseSize(strings.Trim(strutil.ReadField(text, 1, false, ':'), " KB"))

		case strings.HasPrefix(text, "cpu MHz"):
			speed, err = strconv.ParseFloat(strings.Trim(strutil.ReadField(text, 1, false, ':'), " "), 64)

		case strings.HasPrefix(text, "physical id"):
			id, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ':'), " "))

		case strings.HasPrefix(text, "siblings"):
			siblings, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ':'), " "))

		case strings.HasPrefix(text, "cpu cores"):
			cores, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ':'), " "))

		case strings.HasPrefix(text, "flags"):
			if len(info) < id+1 {
				info = append(info, &CPUInfo{vendor, model, cores, siblings, cache, nil})
			}

			if id < len(info) && info[id] != nil {
				info[id].Speed = append(info[id].Speed, speed)
			}
		}

		if err != nil {
			return nil, err
		}
	}

	if vendor == "" && model == "" {
		return nil, errors.New("can't parse cpuinfo file")
	}

	return info, nil
}

// parseCPUCountInfo parses CPU count data
func parseCPUCountInfo(sourceFile string) (uint32, error) {
	var count uint32

	data, err := os.ReadFile(sourceFile)

	if err != nil {
		return 0, err
	}

	line := strings.Trim(string(data), " \n\r")

	for i := range 100 {
		token := strutil.ReadField(line, i, false, ',')

		if token == "" {
			break
		}

		if !strings.ContainsRune(token, '-') {
			_, err := strconv.Atoi(token)

			if err != nil {
				return 0, fmt.Errorf("invalid cpu index %q: %w", token, err)
			}

			count++

			continue
		}

		startNum, endNum, _ := strings.Cut(token, "-")

		start, err := strconv.Atoi(startNum)

		if err != nil {
			return 0, fmt.Errorf("invalid cpu range start %q: %w", token, err)
		}

		end, err := strconv.Atoi(endNum)

		if err != nil {
			return 0, fmt.Errorf("invalid cpu range end %q: %w", token, err)
		}

		if end >= start {
			count += uint32(end-start) + 1
		}
	}

	return count, nil
}

// getCPUArchBits returns CPU architecture bits (32/64)
func getCPUArchBits() int {
	s, closer, err := getFileScanner(cpuInfoFile)

	if err != nil {
		return 0
	}

	defer closer()

	for s.Scan() {
		text := s.Text()

		if !strings.HasPrefix(text, "flags") {
			continue
		}

		// lm - 64 / tm - 32 / rm - 16
		if strings.Contains(text, " lm ") {
			return 64
		}
	}

	return 32
}
