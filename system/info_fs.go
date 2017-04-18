// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
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
	content, err := readFileContent(mtabFile)

	if err != nil {
		return nil, err
	}

	ios, err := GetIOStats()

	if err != nil {
		return nil, err
	}

	info := make(map[string]*FSInfo)

	for _, line := range content {
		if line == "" || line[0:1] == "#" || line[0:1] != "/" {
			continue
		}

		values := strings.Split(line, " ")

		if len(values) < 4 {
			return nil, errors.New("Can't parse file " + mtabFile)
		}

		path := values[1]
		fsInfo := &FSInfo{Type: values[2]}
		stats := &syscall.Statfs_t{}

		err = syscall.Statfs(path, stats)

		if err != nil {
			return nil, err
		}

		fsDevice, err := filepath.EvalSymlinks(values[0])

		if err == nil {
			fsInfo.Device = fsDevice
		} else {
			fsInfo.Device = values[0]
		}

		fsInfo.Total = stats.Blocks * uint64(stats.Bsize)
		fsInfo.Free = uint64(stats.Bavail) * uint64(stats.Bsize)
		fsInfo.Used = fsInfo.Total - (stats.Bfree * uint64(stats.Bsize))
		fsInfo.IOStats = ios[strings.Replace(fsInfo.Device, "/dev/", "", 1)]

		info[path] = fsInfo
	}

	return info, nil
}

// GetIOStats return IO statistics as map device -> statistics
func GetIOStats() (map[string]*IOStats, error) {
	content, err := readFileContent(procDiscStatsFile)

	if err != nil {
		return nil, err
	}

	iostats := make(map[string]*IOStats)

	for _, line := range content {
		if line == "" {
			continue
		}

		lineSlice := splitLine(line)

		if len(lineSlice) != 14 {
			return nil, errors.New("Can't parse file " + procDiscStatsFile)
		}

		device := lineSlice[2]

		if len(device) > 3 {
			if device[0:3] == "ram" || device[0:3] == "loo" {
				continue
			}
		}

		metrics := stringSliceToUintSlice(lineSlice[3:])

		iostats[device] = &IOStats{
			ReadComplete:  metrics[0],
			ReadMerged:    metrics[1],
			ReadSectors:   metrics[2],
			ReadMs:        metrics[3],
			WriteComplete: metrics[4],
			WriteMerged:   metrics[5],
			WriteSectors:  metrics[6],
			WriteMs:       metrics[7],
			IOPending:     metrics[8],
			IOMs:          metrics[9],
			IOQueueMs:     metrics[10],
		}
	}

	return iostats, nil
}

// GetIOUtil return slice (device -> utilization) with IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	fi1, err := GetFSInfo()

	if err != nil {
		return nil, err
	}

	ci1, err := GetCPUInfo()

	if err != nil {
		return nil, err
	}

	time.Sleep(duration)

	fi2, err := GetFSInfo()

	if err != nil {
		return nil, err
	}

	ci2, err := GetCPUInfo()

	if err != nil {
		return nil, err
	}

	return CalculateIOUtil(ci1, fi1, ci2, fi2), nil
}

// CalculateIOUtil calculate IO utilization for all devices
func CalculateIOUtil(ci1 *CPUInfo, fi1 map[string]*FSInfo, ci2 *CPUInfo, fi2 map[string]*FSInfo) map[string]float64 {
	result := make(map[string]float64)

	curCPUUsage := float64(ci2.User + ci2.System + ci2.Idle + ci2.Wait)
	prevCPUUsage := float64(ci1.User + ci1.System + ci1.Idle + ci1.Wait)

	deltams := 1000.0 * (curCPUUsage - prevCPUUsage) / float64(ci2.Count) / getHZ()

	for n, f := range fi1 {
		if fi1[n].IOStats != nil && fi2[n].IOStats != nil {
			ticks := float64(fi2[n].IOStats.IOQueueMs - fi1[n].IOStats.IOQueueMs)
			util := 100.0 * ticks / deltams

			if util > 100.0 {
				util = 100.0
			}

			result[f.Device] = util
		}
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// stringSliceToUintSlice convert string slice to uint64 slice
func stringSliceToUintSlice(s []string) []uint64 {
	var result []uint64

	for _, i := range s {
		iu, _ := strconv.ParseUint(i, 10, 64)
		result = append(result, iu)
	}

	return result
}

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
