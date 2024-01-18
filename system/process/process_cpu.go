//go:build linux
// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Process state flags
const (
	STATE_RUNNING   = "R"
	STATE_SLEEPING  = "S"
	STATE_DISK_WAIT = "D"
	STATE_ZOMBIE    = "Z"
	STATE_STOPPED   = "T"
	STATE_DEAD      = "X"
	STATE_WAKEKILL  = "K"
	STATE_WAKING    = "W"
	STATE_PARKED    = "P"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcInfo contains partial info from /proc/[PID]/stat
type ProcInfo struct {
	PID        int    `json:"pid"`         // The process ID
	Comm       string `json:"comm"`        // The filename of the executable, in parentheses
	State      string `json:"state"`       // Process state
	PPID       int    `json:"ppid"`        // The PID of the parent of this process
	Session    int    `json:"session"`     // The session ID of the process
	TTYNR      int    `json:"tty_nr"`      // The controlling terminal of the process
	TPGid      int    `json:"tpgid"`       // The ID of the foreground process group of the controlling terminal of the process
	UTime      uint64 `json:"utime"`       // Amount of time that this process has been scheduled in user mode, measured in clock ticks
	STime      uint64 `json:"stime"`       // Amount of time that this process has been scheduled in kernel mode, measured in clock ticks
	CUTime     uint64 `json:"cutime"`      // Amount of time that this process's waited-for children have been scheduled in user mode, measured in clock ticks
	CSTime     uint64 `json:"cstime"`      // Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks
	Priority   int    `json:"priority"`    // Priority
	Nice       int    `json:"nice"`        // The nice value
	NumThreads int    `json:"num_threads"` // Number of threads in this process
}

// ProcSample contains value for usage calculation
type ProcSample uint

// ////////////////////////////////////////////////////////////////////////////////// //

// ToSample converts ProcInfo to ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	return ProcSample(pi.UTime + pi.STime + pi.CUTime + pi.CSTime)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInfo returns process info from procfs
func GetInfo(pid int) (*ProcInfo, error) {
	fd, err := os.OpenFile(procFS+"/"+strconv.Itoa(pid)+"/stat", os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return nil, errors.New("Can't parse stat file for given process")
	}

	return parseStatData(text)
}

// codebeat:disable[LOC,ABC]

// GetSample returns ProcSample for CPU usage calculation
func GetSample(pid int) (ProcSample, error) {
	fd, err := os.OpenFile(procFS+"/"+strconv.Itoa(pid)+"/stat", os.O_RDONLY, 0)

	if err != nil {
		return 0, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return 0, errors.New("Can't parse stat file for given process")
	}

	return parseSampleData(text)
}

// codebeat:enable[LOC,ABC]

// CalculateCPUUsage calculates CPU usage
func CalculateCPUUsage(s1, s2 ProcSample, duration time.Duration) float64 {
	total := float64(s2 - s1)
	seconds := float64(duration) / float64(time.Second)

	return 100.0 * ((total / getHZ()) / seconds)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

// parseStatData parses CPU stats data
func parseStatData(text string) (*ProcInfo, error) {
	var err error

	info := &ProcInfo{}

	for i := 0; i < 20; i++ {
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
			return ProcSample(0), err
		}
	}

	return ProcSample(utime + stime + cutime + cstime), nil
}

// codebeat:enable[LOC,ABC]

// parseIntField parses int value of field
func parseIntField(s string, field int) (int, error) {
	v, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("Can't parse stat field %d: %w", field, err)
	}

	return v, nil
}

// parseIntField parses uint value of field
func parseUint64Field(s string, field int) (uint64, error) {
	v, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Can't parse stat field %d: %w", field, err)
	}

	return v, nil
}

// getHZ returns number of processor clock ticks per second
func getHZ() float64 {
	// CLK_TCK is a constant on Linux
	// https://git.musl-libc.org/cgit/musl/tree/src/conf/sysconf.c#n30
	return 100.0
}
