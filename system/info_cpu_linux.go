package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v10/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with CPU stats in procfs
var procStatFile = "/proc/stat"

// Path to file with info about CPU
var cpuInfoFile = "/proc/cpuinfo"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetCPUUsage return info about CPU usage
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

// GetCPUStats return basic CPU stats
func GetCPUStats() (*CPUStats, error) {
	fd, err := os.OpenFile(procStatFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return parseCPUStats(bufio.NewReader(fd))
}

// GetCPUInfo returns slice with info about CPUs
func GetCPUInfo() ([]*CPUInfo, error) {
	fd, err := os.OpenFile(cpuInfoFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return parseCPUInfo(bufio.NewReader(fd))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC,CYCLO]

// parseCPUStats parses cpu stats data
func parseCPUStats(r io.Reader) (*CPUStats, error) {
	var err error

	s := bufio.NewScanner(r)
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
func parseCPUInfo(r io.Reader) ([]*CPUInfo, error) {
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

	s := bufio.NewScanner(r)

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

			info[id].Speed = append(info[id].Speed, speed)
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
