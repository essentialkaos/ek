// +build linux

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
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"pkg.re/essentialkaos/ek.v9/errutil"
	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with disc info in procfs
var procDiscStatsFile = "/proc/diskstats"

// Path to mtab file
var mtabFile = "/etc/mtab"

// Ticks per second
var hz = 0.0

// ////////////////////////////////////////////////////////////////////////////////// //

// GetFSInfo return info about mounted filesystems
func GetFSInfo() (map[string]*FSInfo, error) {
	fd, err := os.OpenFile(mtabFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	info := make(map[string]*FSInfo)

	for s.Scan() {
		text := s.Text()

		if text == "" || text[:1] == "#" || text[:1] != "/" {
			continue
		}

		device := strutil.ReadField(text, 0, true)
		path := strutil.ReadField(text, 1, true)
		fsInfo := &FSInfo{Type: strutil.ReadField(text, 2, true)}

		stats := &syscall.Statfs_t{}

		err = syscall.Statfs(path, stats)

		if err != nil {
			return nil, err
		}

		fsDevice, err := filepath.EvalSymlinks(device)

		if err == nil {
			fsInfo.Device = fsDevice
		} else {
			fsInfo.Device = device
		}

		fsInfo.Used = (stats.Blocks * uint64(stats.Bsize)) - (stats.Bfree * uint64(stats.Bsize))
		fsInfo.Total = fsInfo.Used + (stats.Bavail * uint64(stats.Bsize))
		fsInfo.Free = fsInfo.Total - fsInfo.Used

		info[path] = fsInfo
	}

	if len(info) == 0 {
		return nil, errors.New("Can't parse file " + mtabFile)
	}

	return info, nil
}

// GetIOStats return IO statistics as map device -> statistics
func GetIOStats() (map[string]*IOStats, error) {
	fd, err := os.OpenFile(procDiscStatsFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	iostats := make(map[string]*IOStats)
	errs := errutil.NewErrors()

	for s.Scan() {
		text := s.Text()
		device := strutil.ReadField(text, 2, true)

		if len(device) > 3 {
			if device[:3] == "ram" || device[:3] == "loo" {
				continue
			}
		}

		ios := &IOStats{}

		ios.ReadComplete = parseUint(strutil.ReadField(text, 3, true), errs)
		ios.ReadMerged = parseUint(strutil.ReadField(text, 4, true), errs)
		ios.ReadSectors = parseUint(strutil.ReadField(text, 5, true), errs)
		ios.ReadMs = parseUint(strutil.ReadField(text, 6, true), errs)
		ios.WriteComplete = parseUint(strutil.ReadField(text, 7, true), errs)
		ios.WriteMerged = parseUint(strutil.ReadField(text, 8, true), errs)
		ios.WriteSectors = parseUint(strutil.ReadField(text, 9, true), errs)
		ios.WriteMs = parseUint(strutil.ReadField(text, 10, true), errs)
		ios.IOPending = parseUint(strutil.ReadField(text, 11, true), errs)
		ios.IOMs = parseUint(strutil.ReadField(text, 12, true), errs)
		ios.IOQueueMs = parseUint(strutil.ReadField(text, 13, true), errs)

		if errs.HasErrors() {
			return nil, errs.Last()
		}

		iostats["/dev/"+device] = ios
	}

	return iostats, nil
}

// GetIOUtil return slice (device -> utilization) with IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	io1, err := GetIOStats()

	if err != nil {
		return nil, err
	}

	time.Sleep(duration)

	io2, err := GetIOStats()

	if err != nil {
		return nil, err
	}

	return CalculateIOUtil(io1, io2, duration), nil
}

// CalculateIOUtil calculate IO utilization for all devices
func CalculateIOUtil(io1, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
	if io1 == nil || io2 == nil {
		return nil
	}

	result := make(map[string]float64)

	// convert duration to jiffies
	itv := uint64(duration / (time.Millisecond * 10))

	for device := range io1 {
		if io1[device] == nil || io2[device] == nil {
			continue
		}

		util := float64(io2[device].IOMs-io1[device].IOMs) / float64(itv) * getHZ()

		util /= 10.0

		if util > 100.0 {
			util = 100.0
		}

		result[device] = util
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getHZ return number of processor clock ticks per second
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
