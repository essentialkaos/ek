package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v12/strutil"
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

// GetCPUUsage returns info about CPU usage
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

// It's ok to have so complex method for calculation
// codebeat:disable[CYCLO]

// CalculateCPUUsage calcualtes CPU usage based on CPUStats
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

// codebeat:enable[CYCLO]

// GetCPUStats returns basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	s, closer, err := getFileScanner(procStatFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseCPUStats(s)
}

// GetCPUInfo returns slice with info about CPUs
func GetCPUInfo() ([]*CPUInfo, error) {
	s, closer, err := getFileScanner(cpuInfoFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseCPUInfo(s)
}

// GetCPUCount returns info about CPU
func GetCPUCount() (CPUCount, error) {
	possible, err := os.ReadFile(cpuPossibleFile)

	if err != nil {
		return CPUCount{}, err
	}

	present, err := os.ReadFile(cpuPresentFile)

	if err != nil {
		return CPUCount{}, err
	}

	online, err := os.ReadFile(cpuOnlineFile)

	if err != nil {
		return CPUCount{}, err
	}

	offline, err := os.ReadFile(cpuOfflineFile)

	if err != nil {
		return CPUCount{}, err
	}

	return CPUCount{
		Possible: parseCPUCountInfo(string(possible)),
		Present:  parseCPUCountInfo(string(present)),
		Online:   parseCPUCountInfo(string(online)),
		Offline:  parseCPUCountInfo(string(offline)),
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC,CYCLO]

// parseCPUStats parses cpu stats data
func parseCPUStats(s *bufio.Scanner) (*CPUStats, error) {
	var err error

	stats := &CPUStats{}

	for s.Scan() {
		text := s.Text()

		if len(text) < 4 || text[:3] != "cpu" {
			continue
		}

		if text[:4] != "cpu " {
			stats.Count++
			continue
		}

		stats.User, err = strconv.ParseUint(strutil.ReadField(text, 1, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 1 as unsigned integer in " + procStatFile)
		}

		stats.Nice, err = strconv.ParseUint(strutil.ReadField(text, 2, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 2 as unsigned integer in " + procStatFile)
		}

		stats.System, err = strconv.ParseUint(strutil.ReadField(text, 3, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 3 as unsigned integer in " + procStatFile)
		}

		stats.Idle, err = strconv.ParseUint(strutil.ReadField(text, 4, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 4 as unsigned integer in " + procStatFile)
		}

		stats.Wait, err = strconv.ParseUint(strutil.ReadField(text, 5, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 5 as unsigned integer in " + procStatFile)
		}

		stats.IRQ, err = strconv.ParseUint(strutil.ReadField(text, 6, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 6 as unsigned integer in " + procStatFile)
		}

		stats.SRQ, err = strconv.ParseUint(strutil.ReadField(text, 7, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 7 as unsigned integer in " + procStatFile)
		}

		stats.Steal, err = strconv.ParseUint(strutil.ReadField(text, 8, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 8 as unsigned integer in " + procStatFile)
		}

	}

	if stats.Count == 0 {
		return nil, errors.New("Can't parse file " + procStatFile)
	}

	stats.Total = stats.User + stats.System + stats.Nice + stats.Idle + stats.Wait + stats.IRQ + stats.SRQ + stats.Steal

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
			vendor = strings.Trim(strutil.ReadField(text, 1, false, ":"), " ")

		case strings.HasPrefix(text, "model name"):
			model = strings.Trim(strutil.ReadField(text, 1, false, ":"), " ")

		case strings.HasPrefix(text, "cache size"):
			cache, err = parseSize(strings.Trim(strutil.ReadField(text, 1, false, ":"), " KB"))

		case strings.HasPrefix(text, "cpu MHz"):
			speed, err = strconv.ParseFloat(strings.Trim(strutil.ReadField(text, 1, false, ":"), " "), 64)

		case strings.HasPrefix(text, "physical id"):
			id, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ":"), " "))

		case strings.HasPrefix(text, "siblings"):
			siblings, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ":"), " "))

		case strings.HasPrefix(text, "cpu cores"):
			cores, err = strconv.Atoi(strings.Trim(strutil.ReadField(text, 1, false, ":"), " "))

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
		return nil, errors.New("Can't parse cpuinfo file")
	}

	return info, nil
}

// codebeat:enable[LOC,ABC,CYCLO]

// parseCPUCountInfo parses CPU count data
func parseCPUCountInfo(data string) uint32 {
	startNum := strings.Trim(strutil.ReadField(data, 0, false, "-"), "\n\r")
	endNum := strings.Trim(strutil.ReadField(data, 1, false, "-"), "\n\r")

	start, _ := strconv.ParseUint(startNum, 10, 32)
	end, _ := strconv.ParseUint(endNum, 10, 32)

	return uint32(end-start) + 1
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
