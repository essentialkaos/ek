// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v9/errutil"
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

// ////////////////////////////////////////////////////////////////////////////////// //

// Ticks per second
var hz = 0.0

// ////////////////////////////////////////////////////////////////////////////////// //

// GetInfo return process info from procfs
func GetInfo(pid int) (*ProcInfo, error) {
	data, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")

	if err != nil {
		return nil, err
	}

	dataSlice := strings.Split(string(data), " ")

	if len(dataSlice) < 20 {
		return nil, errors.New("Can't parse stat file for given process")
	}

	info := &ProcInfo{}
	errs := errutil.NewErrors()

	info.PID = parseIntField(dataSlice, 0, errs)
	info.Comm = dataSlice[1]
	info.State = dataSlice[2]
	info.PPID = parseIntField(dataSlice, 3, errs)
	info.Session = parseIntField(dataSlice, 5, errs)
	info.TTYNR = parseIntField(dataSlice, 6, errs)
	info.TPGid = parseIntField(dataSlice, 7, errs)
	info.UTime = parseUintField(dataSlice, 13, errs)
	info.STime = parseUintField(dataSlice, 14, errs)
	info.CUTime = parseUintField(dataSlice, 15, errs)
	info.CSTime = parseUintField(dataSlice, 16, errs)
	info.Priority = parseIntField(dataSlice, 17, errs)
	info.Nice = parseIntField(dataSlice, 18, errs)
	info.NumThreads = parseIntField(dataSlice, 19, errs)

	if errs.HasErrors() {
		return nil, errs.Last()
	}

	return info, nil
}

// CalculateCPUUsage calculate CPU usage
func CalculateCPUUsage(i1, i2 *ProcInfo, duration time.Duration) float64 {
	if i1 == nil || i2 == nil {
		return 0.0
	}

	i1Total := i1.UTime + i1.STime + i1.CUTime + i1.CSTime
	i2Total := i2.UTime + i2.STime + i2.CUTime + i2.CSTime

	total := float64(i2Total - i1Total)
	seconds := float64(duration) / float64(time.Second)

	return 100.0 * ((total / getHZ()) / seconds)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getHZ() float64 {
	if hz != 0.0 {
		return hz
	}

	output, err := exec.Command("/usr/bin/getconf", "CLK_TCK").Output()

	if err != nil {
		hz = 100.0
		return hz
	}

	hz, _ = strconv.ParseFloat(string(output), 64)

	if hz == 0.0 {
		hz = 100.0
		return hz
	}

	return hz
}

func parseIntField(data []string, index int, errs *errutil.Errors) int {
	value, err := strconv.Atoi(data[index])

	if err != nil {
		errs.Add(err)
	}

	return value
}

func parseUintField(data []string, index int, errs *errutil.Errors) uint64 {
	value, err := strconv.ParseUint(data[index], 10, 64)

	if err != nil {
		errs.Add(err)
	}

	return value
}
