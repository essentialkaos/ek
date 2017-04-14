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
)

// ////////////////////////////////////////////////////////////////////////////////// //

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

	info.PID, _ = strconv.Atoi(dataSlice[0])
	info.Comm = dataSlice[1]
	info.State = dataSlice[2]
	info.PPID, _ = strconv.Atoi(dataSlice[3])
	info.Session, _ = strconv.Atoi(dataSlice[5])
	info.TTYNR, _ = strconv.Atoi(dataSlice[6])
	info.TPGid, _ = strconv.Atoi(dataSlice[7])
	info.UTime, _ = strconv.ParseUint(dataSlice[13], 10, 64)
	info.STime, _ = strconv.ParseUint(dataSlice[14], 10, 64)
	info.CUTime, _ = strconv.ParseUint(dataSlice[15], 10, 64)
	info.CSTime, _ = strconv.ParseUint(dataSlice[16], 10, 64)
	info.Priority, _ = strconv.Atoi(dataSlice[17])
	info.Nice, _ = strconv.Atoi(dataSlice[18])
	info.NumThreads, _ = strconv.Atoi(dataSlice[19])

	return info, nil
}

// CalculateCPUUsage calculate cpu usage
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
