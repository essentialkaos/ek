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
	"fmt"
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
	PID        int    `json:"pid"`
	Comm       string `json:"comm"`
	State      string `json:"state"`
	PPID       int    `json:"ppid"`
	Session    int    `json:"session"`
	TTYNR      int    `json:"tty_nr"`
	TPGid      int    `json:"tpgid"`
	UTime      uint64 `json:"utime"`
	STime      uint64 `json:"stime"`
	CUTime     uint64 `json:"cutime"`
	CSTime     uint64 `json:"cstime"`
	Priority   int    `json:"priority"`
	Nice       int    `json:"nice"`
	NumThreads int    `json:"num_threads"`
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
