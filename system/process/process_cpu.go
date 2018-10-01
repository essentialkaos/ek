// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
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

	"pkg.re/essentialkaos/ek.v10/strutil"
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

// Ticks per second
var hz = 0.0

// ////////////////////////////////////////////////////////////////////////////////// //

// ToSample convert ProcInfo to ProcSample for CPU usage calculation
func (pi *ProcInfo) ToSample() ProcSample {
	return ProcSample(pi.UTime + pi.STime + pi.CUTime + pi.CSTime)
}

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

	return parseStatData(text)
}

// codebeat:disable[LOC,ABC]

// GetSample return ProcSample for CPU usage calculation
func GetSample(pid int) (ProcSample, error) {
	fd, err := os.OpenFile("/proc/"+strconv.Itoa(pid)+"/stat", os.O_RDONLY, 0)

	if err != nil {
		return 0, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return 0, errors.New("Can't parse stat file for given process")
	}

	utime, err := strconv.ParseUint(strutil.ReadField(text, 13, true), 10, 64)

	if err != nil {
		return 0, errors.New("Can't parse stat field 13")
	}

	stime, err := strconv.ParseUint(strutil.ReadField(text, 14, true), 10, 64)

	if err != nil {
		return 0, errors.New("Can't parse stat field 14")
	}

	cutime, err := strconv.ParseUint(strutil.ReadField(text, 15, true), 10, 64)

	if err != nil {
		return 0, errors.New("Can't parse stat field 15")
	}

	cstime, err := strconv.ParseUint(strutil.ReadField(text, 16, true), 10, 64)

	if err != nil {
		return 0, errors.New("Can't parse stat field 16")
	}

	return ProcSample(utime + stime + cutime + cstime), nil
}

// codebeat:enable[LOC,ABC]

// CalculateCPUUsage calculate CPU usage
func CalculateCPUUsage(s1, s2 ProcSample, duration time.Duration) float64 {
	total := float64(s2 - s1)
	seconds := float64(duration) / float64(time.Second)

	return 100.0 * ((total / getHZ()) / seconds)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

func parseStatData(text string) (*ProcInfo, error) {
	var err error

	info := &ProcInfo{}

	info.Comm = strutil.ReadField(text, 1, true)
	info.State = strutil.ReadField(text, 2, true)

	info.PID, err = strconv.Atoi(strutil.ReadField(text, 0, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 0")
	}

	info.PPID, err = strconv.Atoi(strutil.ReadField(text, 3, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 3")
	}

	info.Session, err = strconv.Atoi(strutil.ReadField(text, 5, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 5")
	}

	info.TTYNR, err = strconv.Atoi(strutil.ReadField(text, 6, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 6")
	}

	info.TPGid, err = strconv.Atoi(strutil.ReadField(text, 7, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 7")
	}

	info.UTime, err = strconv.ParseUint(strutil.ReadField(text, 13, true), 10, 64)

	if err != nil {
		return nil, errors.New("Can't parse stat field 13")
	}

	info.STime, err = strconv.ParseUint(strutil.ReadField(text, 14, true), 10, 64)

	if err != nil {
		return nil, errors.New("Can't parse stat field 14")
	}

	info.CUTime, err = strconv.ParseUint(strutil.ReadField(text, 15, true), 10, 64)

	if err != nil {
		return nil, errors.New("Can't parse stat field 15")
	}

	info.CSTime, err = strconv.ParseUint(strutil.ReadField(text, 16, true), 10, 64)

	if err != nil {
		return nil, errors.New("Can't parse stat field 16")
	}

	info.Priority, err = strconv.Atoi(strutil.ReadField(text, 17, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 17")
	}

	info.Nice, err = strconv.Atoi(strutil.ReadField(text, 18, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 18")
	}

	info.NumThreads, err = strconv.Atoi(strutil.ReadField(text, 19, true))

	if err != nil {
		return nil, errors.New("Can't parse stat field 19")
	}

	return info, nil
}

// codebeat:enable[LOC,ABC]

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
