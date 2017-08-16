// +build linux

package process

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
	"os/exec"
	"strconv"
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
	fd, err := os.OpenFile("/proc/"+strconv.Itoa(pid)+"/stat", os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return nil, errors.New("Can't parse stat file for given process")
	}

	info := &ProcInfo{}
	errs := errutil.NewErrors()

	info.PID = parseInt(readField(text, 0), errs)
	info.Comm = readField(text, 1)
	info.State = readField(text, 2)
	info.PPID = parseInt(readField(text, 3), errs)
	info.Session = parseInt(readField(text, 5), errs)
	info.TTYNR = parseInt(readField(text, 6), errs)
	info.TPGid = parseInt(readField(text, 7), errs)
	info.UTime = parseUint(readField(text, 13), errs)
	info.STime = parseUint(readField(text, 14), errs)
	info.CUTime = parseUint(readField(text, 15), errs)
	info.CSTime = parseUint(readField(text, 16), errs)
	info.Priority = parseInt(readField(text, 17), errs)
	info.Nice = parseInt(readField(text, 18), errs)
	info.NumThreads = parseInt(readField(text, 19), errs)

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
