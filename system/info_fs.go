// +build linux

package system

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
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"pkg.re/essentialkaos/ek.v9/errutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with disc info in procfs
var procDiscStatsFile = "/proc/diskstats"

// Path to mtab file
var mtabFile = "/etc/mtab"

// Ticks per second
var hz = 0.0

// ////////////////////////////////////////////////////////////////////////////////// //

// FSInfo contains info about fs usage
type FSInfo struct {
	Type    string   `json:"type"`    // FS type (ext4/ntfs/etc...)
	Device  string   `json:"device"`  // Device spec
	Used    uint64   `json:"used"`    // Used space
	Free    uint64   `json:"free"`    // Free space
	Total   uint64   `json:"total"`   // Total space
	IOStats *IOStats `json:"iostats"` // IO statistics
}

// IOStats contains information about I/O
type IOStats struct {
	ReadComplete  uint64 `json:"read_complete"`  // Reads completed successfully
	ReadMerged    uint64 `json:"read_merged"`    // Reads merged
	ReadSectors   uint64 `json:"read_sectors"`   // Sectors read
	ReadMs        uint64 `json:"read_ms"`        // Time spent reading (ms)
	WriteComplete uint64 `json:"write_complete"` // Writes completed
	WriteMerged   uint64 `json:"write_merged"`   // Writes merged
	WriteSectors  uint64 `json:"write_sectors"`  // Sectors written
	WriteMs       uint64 `json:"write_ms"`       // Time spent writing (ms)
	IOPending     uint64 `json:"io_pending"`     // I/Os currently in progress
	IOMs          uint64 `json:"io_ms"`          // Time spent doing I/Os (ms)
	IOQueueMs     uint64 `json:"io_queue_ms"`    // Weighted time spent doing I/Os (ms)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetFSInfo return info about mounted filesystems
func GetFSInfo() (map[string]*FSInfo, error) {
	iostats, err := GetIOStats()

	if err != nil {
		return nil, err
	}

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

		device := readField(text, 0)
		path := readField(text, 1)
		fsInfo := &FSInfo{Type: readField(text, 2)}

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
		fsInfo.IOStats = iostats[strings.Replace(fsInfo.Device, "/dev/", "", 1)]

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
		device := readField(text, 2)

		if len(device) > 3 {
			if device[:3] == "ram" || device[:3] == "loo" {
				continue
			}
		}

		ios := &IOStats{}

		ios.ReadComplete = parseUint(readField(text, 3), errs)
		ios.ReadMerged = parseUint(readField(text, 4), errs)
		ios.ReadSectors = parseUint(readField(text, 5), errs)
		ios.ReadMs = parseUint(readField(text, 6), errs)
		ios.WriteComplete = parseUint(readField(text, 7), errs)
		ios.WriteMerged = parseUint(readField(text, 8), errs)
		ios.WriteSectors = parseUint(readField(text, 9), errs)
		ios.WriteMs = parseUint(readField(text, 10), errs)
		ios.IOPending = parseUint(readField(text, 11), errs)
		ios.IOMs = parseUint(readField(text, 12), errs)
		ios.IOQueueMs = parseUint(readField(text, 13), errs)

		if errs.HasErrors() {
			return nil, errs.Last()
		}

		iostats[device] = ios
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
func CalculateIOUtil(io1 map[string]*IOStats, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
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
