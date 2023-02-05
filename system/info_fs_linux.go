package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with disk info in procfs
var procDiskStatsFile = "/proc/diskstats"

// Path to mtab file
var mtabFile = "/etc/mtab"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetFSUsage returns info about mounted filesystems
func GetFSUsage() (map[string]*FSUsage, error) {
	s, closer, err := getFileScanner(mtabFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseFSInfo(s, true)
}

// GetIOStats returns IO statistics as map device -> statistics
func GetIOStats() (map[string]*IOStats, error) {
	s, closer, err := getFileScanner(procDiskStatsFile)

	if err != nil {
		return nil, err
	}

	defer closer()

	return parseIOStats(s)
}

// GetIOUtil returns slice (device -> utilization) with IO utilization
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

// CalculateIOUtil calculates IO utilization for all devices
func CalculateIOUtil(io1, io2 map[string]*IOStats, duration time.Duration) map[string]float64 {
	if io1 == nil || io2 == nil {
		return nil
	}

	result := make(map[string]float64)

	// convert duration to jiffies
	itv := uint64(duration / (time.Millisecond * 10))

	for device := range io1 {
		if io2[device] == nil {
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

// codebeat:disable[LOC,ABC,CYCLO]

// parseIOStats parses IO stats data
func parseIOStats(s *bufio.Scanner) (map[string]*IOStats, error) {
	var err error

	iostats := make(map[string]*IOStats)

	for s.Scan() {
		text := s.Text()
		device := strutil.ReadField(text, 2, true)

		if len(device) > 3 {
			// Skip devices with names like ram* and loop*
			if device[:3] == "ram" || device[:3] == "loo" {
				continue
			}
		}

		ios := &IOStats{}

		ios.ReadComplete, err = strconv.ParseUint(strutil.ReadField(text, 3, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 3 as unsigned integer in " + procDiskStatsFile)
		}

		ios.ReadMerged, err = strconv.ParseUint(strutil.ReadField(text, 4, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 4 as unsigned integer in " + procDiskStatsFile)
		}

		ios.ReadSectors, err = strconv.ParseUint(strutil.ReadField(text, 5, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 5 as unsigned integer in " + procDiskStatsFile)
		}

		ios.ReadMs, err = strconv.ParseUint(strutil.ReadField(text, 6, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 6 as unsigned integer in " + procDiskStatsFile)
		}

		ios.WriteComplete, err = strconv.ParseUint(strutil.ReadField(text, 7, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 7 as unsigned integer in " + procDiskStatsFile)
		}

		ios.WriteMerged, err = strconv.ParseUint(strutil.ReadField(text, 8, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 8 as unsigned integer in " + procDiskStatsFile)
		}

		ios.WriteSectors, err = strconv.ParseUint(strutil.ReadField(text, 9, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 9 as unsigned integer in " + procDiskStatsFile)
		}

		ios.WriteMs, err = strconv.ParseUint(strutil.ReadField(text, 10, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 10 as unsigned integer in " + procDiskStatsFile)
		}

		ios.IOPending, err = strconv.ParseUint(strutil.ReadField(text, 11, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 11 as unsigned integer in " + procDiskStatsFile)
		}

		ios.IOMs, err = strconv.ParseUint(strutil.ReadField(text, 12, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 12 as unsigned integer in " + procDiskStatsFile)
		}

		ios.IOQueueMs, err = strconv.ParseUint(strutil.ReadField(text, 13, true), 10, 64)

		if err != nil {
			return nil, errors.New("Can't parse field 13 as unsigned integer in " + procDiskStatsFile)
		}

		iostats["/dev/"+device] = ios
	}

	return iostats, nil
}

// parseFSInfo parses fs info data
func parseFSInfo(s *bufio.Scanner, calculateStats bool) (map[string]*FSUsage, error) {
	var err error

	info := make(map[string]*FSUsage)

	for s.Scan() {
		text := s.Text()

		if text == "" || text[:1] == "#" || text[:1] != "/" {
			continue
		}

		device := strutil.ReadField(text, 0, true)
		path := strutil.ReadField(text, 1, true)
		fsInfo := &FSUsage{Type: strutil.ReadField(text, 2, true)}

		if !calculateStats {
			continue
		}

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

// enable:disable[LOC,ABC,CYCLO]

// getHZ returns number of processor clock ticks per second
func getHZ() float64 {
	// CLK_TCK is a constant on Linux
	// https://git.musl-libc.org/cgit/musl/tree/src/conf/sysconf.c#n30
	return 100
}
