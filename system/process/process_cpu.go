//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInfo returns the parsed /proc/[pid]/stat fields for the given process
func GetInfo(pid int) (*ProcInfo, error) {
	text, err := readStatData(pid)

	if err != nil {
		return nil, err
	}

	return parseStatData(text)
}

// GetSample returns a CPU-time snapshot of the given process for use with
// [CalculateCPUUsage]
func GetSample(pid int) (ProcSample, error) {
	text, err := readStatData(pid)

	if err != nil {
		return 0, err
	}

	return parseSampleData(text)
}

// CalculateCPUUsage returns the percentage of CPU used between two samples over
// the given duration
func CalculateCPUUsage(s1, s2 ProcSample, duration time.Duration) float64 {
	total := float64(s2 - s1)
	seconds := float64(duration) / float64(time.Second)

	return 100.0 * ((total / getHZ()) / seconds)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// readStatData reads stat file data and returns it as a string
func readStatData(pid int) (string, error) {
	statFile := path.Join(procFS, strconv.Itoa(pid), "stat")
	data, err := os.ReadFile(statFile)

	if err != nil {
		return "", err
	}

	if len(data) < 20 {
		return "", fmt.Errorf("stat file %s has not valid data", statFile)
	}

	return strings.Trim(string(data), "\n\r"), nil
}

// parseStatData parses CPU stats data
func parseStatData(text string) (*ProcInfo, error) {
	var err error

	info := &ProcInfo{}

	for i := range 20 {
		switch i {
		case 0:
			info.PID, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 1:
			info.Comm = strutil.ReadField(text, i, true)
		case 2:
			info.State = strutil.ReadField(text, i, true)
		case 3:
			info.PPID, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 5:
			info.Session, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 6:
			info.TTYNR, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 7:
			info.TPGid, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 13:
			info.UTime, err = parseUint64Field(strutil.ReadField(text, i, true), i)
		case 14:
			info.STime, err = parseUint64Field(strutil.ReadField(text, i, true), i)
		case 15:
			info.CUTime, err = parseUint64Field(strutil.ReadField(text, i, true), i)
		case 16:
			info.CSTime, err = parseUint64Field(strutil.ReadField(text, i, true), i)
		case 17:
			info.Priority, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 18:
			info.Nice, err = parseIntField(strutil.ReadField(text, i, true), i)
		case 19:
			info.NumThreads, err = parseIntField(strutil.ReadField(text, i, true), i)
		}

		if err != nil {
			return nil, err
		}
	}

	return info, nil
}

// parseSampleData extracts CPU sample info
func parseSampleData(text string) (ProcSample, error) {
	var err error
	var utime, stime, cutime, cstime uint64

	for i := 13; i < 17; i++ {
		value := strutil.ReadField(text, i, true)

		switch i {
		case 13:
			utime, err = parseUint64Field(value, i)
		case 14:
			stime, err = parseUint64Field(value, i)
		case 15:
			cutime, err = parseUint64Field(value, i)
		case 16:
			cstime, err = parseUint64Field(value, i)
		}

		if err != nil {
			return 0, err
		}
	}

	return ProcSample(utime + stime + cutime + cstime), nil
}

// parseIntField parses int value of field
func parseIntField(s string, field int) (int, error) {
	v, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("can't parse stat field %d: %w", field, err)
	}

	return v, nil
}

// parseIntField parses uint value of field
func parseUint64Field(s string, field int) (uint64, error) {
	v, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("can't parse stat field %d: %w", field, err)
	}

	return v, nil
}

// getHZ returns number of processor clock ticks per second
func getHZ() float64 {
	// CLK_TCK is a constant on Linux
	// https://git.musl-libc.org/cgit/musl/tree/src/conf/sysconf.c#n30
	return 100.0
}
