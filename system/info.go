// +build !windows

// Package system provides methods for working with system data (metrics/users)
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _HZ = 100.0

const (
	_PROC_UPTIME    = "/proc/uptime"
	_PROC_LOADAVG   = "/proc/loadavg"
	_PROC_MEMINFO   = "/proc/meminfo"
	_PROC_CPUINFO   = "/proc/stat"
	_PROC_NET       = "/proc/net/dev"
	_PROC_DISCSTATS = "/proc/diskstats"
	_MTAB_FILE      = "/etc/mtab"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	LINUX_ARCH   = "arch"
	LINUX_CENTOS = "centos"
	LINUX_DEBIAN = "debian"
	LINUX_FEDORA = "fedora"
	LINUX_GENTOO = "gentoo"
	LINUX_RHEL   = "rhel"
	LINUX_SUSE   = "suse"
	LINUX_UBUNTU = "ubuntu"
	DARWIN_OSX   = "osx"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// LoadAvg contains information about average system load
type LoadAvg struct {
	Min1  float64 `json:"min1"`  // LA in last 1 minute
	Min5  float64 `json:"min5"`  // LA in last 5 minutes
	Min15 float64 `json:"min15"` // LA in last 15 minutes
	RProc int     `json:"rproc"` // Number of currently runnable kernel scheduling entities
	TProc int     `json:"tproc"` // Number of kernel scheduling entities that currently exist on the system
}

// MemInfo contains info about system memory
type MemInfo struct {
	MemTotal   uint64 `json:"total"`       // Total usable ram (i.e. physical ram minus a few reserved bits and the kernel binary code)
	MemFree    uint64 `json:"free"`        // The sum of MemFree - (Buffers + Cached)
	MemUsed    uint64 `json:"used"`        // MemTotal - MemFree
	Buffers    uint64 `json:"buffers"`     // Relatively temporary storage for raw disk blocks shouldn't get tremendously large (20MB or so)
	Cached     uint64 `json:"cached"`      // In-memory cache for files read from the disk (the pagecache).  Doesn't include SwapCached
	Active     uint64 `json:"active"`      // Memory that has been used more recently and usually not reclaimed unless absolutely necessary
	Inactive   uint64 `json:"inactive"`    // Memory which has been less recently used.  It is more eligible to be reclaimed for other purposes
	SwapTotal  uint64 `json:"swap_total"`  // Total amount of swap space available
	SwapFree   uint64 `json:"swap_free"`   // Memory which has been evicted from RAM, and is temporarily on the disk still also is in the swapfile
	SwapUsed   uint64 `json:"swap_used"`   // SwapTotal - SwapFree
	SwapCached uint64 `json:"spaw_cached"` // Memory that once was swapped out, is swapped back in but
	Dirty      uint64 `json:"dirty"`       // Memory which is waiting to get written back to the disk
	Slab       uint64 `json:"slab"`        // In-kernel data structures cache
}

// CPUInfo contains info about CPU usage
type CPUInfo struct {
	User   float64 `json:"user"`   // Normal processes executing in user mode
	System float64 `json:"system"` // Processes executing in kernel mode
	Nice   float64 `json:"nice"`   // Niced processes executing in user mode
	Idle   float64 `json:"idle"`   // Twiddling thumbs
	Wait   float64 `json:"wait"`   // Waiting for I/O to complete
	Count  int     `json:"count"`  // Number of CPU cores
}

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

// SystemInfo contains info about system (hostname, OS, arch...)
type SystemInfo struct {
	Hostname     string `json:"hostname"`     // Hostname
	OS           string `json:"os"`           // OS name
	Distribution string `json:"distribution"` // OS distribution
	Version      string `json:"version"`      // OS version
	Kernel       string `json:"kernel"`       // Kernel version
	Arch         string `json:"arch"`         // System architecture (i386/i686/x86_64/etc...)
}

// InterfaceInfo contains info about network interfaces
type InterfaceInfo struct {
	ReceivedBytes      uint64 `json:"received_bytes"`
	ReceivedPackets    uint64 `json:"received_packets"`
	TransmittedBytes   uint64 `json:"transmitted_bytes"`
	TransmittedPackets uint64 `json:"transmitted_packets"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

type basicCPUInfo struct {
	User   uint64
	Nice   uint64
	System uint64
	Idle   uint64
	Wait   uint64
	IRQ    uint64
	SRQ    uint64
	Steal  uint64
	Total  uint64
	Count  int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime return system uptime in seconds
func GetUptime() (uint64, error) {
	content, err := readFileContent(_PROC_UPTIME)

	if err != nil {
		return 0, err
	}

	ca := strings.Split(content[0], " ")

	if len(ca) != 2 {
		return 0, errors.New("Can't parse file " + _PROC_UPTIME)
	}

	up, _ := strconv.ParseFloat(ca[0], 64)

	return uint64(up), nil
}

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	result := &LoadAvg{}
	content, err := readFileContent(_PROC_LOADAVG)

	if err != nil {
		return nil, err
	}

	contentSlice := strings.Split(content[0], " ")

	if len(contentSlice) != 5 {
		return nil, errors.New("Can't parse file " + _PROC_LOADAVG)
	}

	procSlice := strings.Split(contentSlice[3], "/")

	result.Min1, _ = strconv.ParseFloat(contentSlice[0], 64)
	result.Min5, _ = strconv.ParseFloat(contentSlice[1], 64)
	result.Min15, _ = strconv.ParseFloat(contentSlice[2], 64)
	result.RProc, _ = strconv.Atoi(procSlice[0])
	result.TProc, _ = strconv.Atoi(procSlice[1])

	return result, nil
}

// GetMemInfo return memory info
func GetMemInfo() (*MemInfo, error) {
	var props = map[string]bool{
		"MemTotal":   true,
		"MemFree":    true,
		"Buffers":    true,
		"Cached":     true,
		"SwapCached": true,
		"Active":     true,
		"Inactive":   true,
		"SwapTotal":  true,
		"SwapFree":   true,
		"Dirty":      true,
		"Slab":       true,
	}

	result := &MemInfo{}
	content, err := readFileContent(_PROC_MEMINFO)

	if err != nil {
		return nil, err
	}

	for _, line := range content {
		if line == "" {
			continue
		}

		lineSlice := strings.Split(line, ":")

		if len(lineSlice) != 2 {
			return nil, errors.New("Can't parse file " + _PROC_MEMINFO)
		}

		if props[lineSlice[0]] != true {
			continue
		}

		strValue := strings.TrimRight(lineSlice[1], " kB")
		strValue = strings.Replace(strValue, " ", "", -1)
		uintValue, err := strconv.ParseUint(strValue, 10, 64)

		if err != nil {
			return nil, err
		}

		switch lineSlice[0] {
		case "MemTotal":
			result.MemTotal = uintValue * 1024
		case "MemFree":
			result.MemFree = uintValue * 1024
		case "Buffers":
			result.Buffers = uintValue * 1024
		case "Cached":
			result.Cached = uintValue * 1024
		case "SwapCached":
			result.SwapCached = uintValue * 1024
		case "Active":
			result.Active = uintValue * 1024
		case "Inactive":
			result.Inactive = uintValue * 1024
		case "SwapTotal":
			result.SwapTotal = uintValue * 1024
		case "SwapFree":
			result.SwapFree = uintValue * 1024
		case "Dirty":
			result.Dirty = uintValue * 1024
		case "Slab":
			result.Slab = uintValue * 1024
		}
	}

	result.MemFree += result.Cached + result.Buffers
	result.MemUsed = result.MemTotal - result.MemFree
	result.SwapUsed = result.SwapTotal - result.SwapFree

	return result, nil
}

// GetCPUInfo return info about CPU usage
func GetCPUInfo() (*CPUInfo, error) {
	info, err := getCPUStats()

	if err != nil {
		return nil, err
	}

	return &CPUInfo{
		System: (float64(info.System) / float64(info.Total)) * 100,
		User:   (float64(info.User) / float64(info.Total)) * 100,
		Nice:   (float64(info.Nice) / float64(info.Total)) * 100,
		Wait:   (float64(info.Wait) / float64(info.Total)) * 100,
		Idle:   (float64(info.Idle) / float64(info.Total)) * 100,
		Count:  info.Count,
	}, nil
}

// GetFSInfo return info about mounted filesystems
func GetFSInfo() (map[string]*FSInfo, error) {
	result := make(map[string]*FSInfo)

	content, err := readFileContent(_MTAB_FILE)

	if err != nil {
		return nil, err
	}

	ios, err := GetIOStats()

	if err != nil {
		return nil, err
	}

	for _, line := range content {
		if line == "" || line[0:1] == "#" || line[0:1] != "/" {
			continue
		}

		values := strings.Split(line, " ")

		if len(values) < 4 {
			return nil, errors.New("Can't parse file " + _MTAB_FILE)
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

		result[path] = fsInfo
	}

	return result, nil
}

// GetIOStats return IO statistics as map device -> statistics
func GetIOStats() (map[string]*IOStats, error) {
	result := make(map[string]*IOStats)

	content, err := readFileContent(_PROC_DISCSTATS)

	if err != nil {
		return nil, err
	}

	for _, line := range content {
		if line == "" {
			continue
		}

		values := cleanSlice(strings.Split(line, " "))

		if len(values) != 14 {
			return nil, errors.New("Can't parse file " + _PROC_DISCSTATS)
		}

		device := values[2]

		if len(device) > 3 {
			if device[0:3] == "ram" || device[0:3] == "loo" {
				continue
			}
		}

		metrics := stringSliceToUintSlice(values[3:])

		result[device] = &IOStats{
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

	return result, nil
}

// GetInterfacesInfo return info about network interfaces
func GetInterfacesInfo() (map[string]*InterfaceInfo, error) {
	result := make(map[string]*InterfaceInfo)

	content, err := readFileContent(_PROC_NET)

	if err != nil {
		return nil, err
	}

	if len(content) <= 2 {
		return result, nil
	}

	for _, line := range content[2:] {
		lineSlice := strings.Split(line, ":")

		if len(lineSlice) != 2 {
			continue
		}

		metrics := cleanSlice(strings.Split(lineSlice[1], " "))
		name := strings.TrimLeft(lineSlice[0], " ")
		receivedBytes, _ := strconv.ParseUint(metrics[0], 10, 64)
		receivedPackets, _ := strconv.ParseUint(metrics[1], 10, 64)
		transmittedBytes, _ := strconv.ParseUint(metrics[8], 10, 64)
		transmittedPackets, _ := strconv.ParseUint(metrics[9], 10, 64)

		result[name] = &InterfaceInfo{
			receivedBytes,
			receivedPackets,
			transmittedBytes,
			transmittedPackets,
		}
	}

	return result, nil
}

// GetNetworkSpeed return network input/output speed in bytes per second for
// all network interfaces
func GetNetworkSpeed(duration time.Duration) (uint64, uint64, error) {
	intInfo1, err := GetInterfacesInfo()

	if err != nil {
		return 0, 0, err
	}

	time.Sleep(duration)

	intInfo2, err := GetInterfacesInfo()

	if err != nil {
		return 0, 0, err
	}

	rb1, tb1 := getActiveInterfacesBytes(intInfo1)
	rb2, tb2 := getActiveInterfacesBytes(intInfo2)

	if rb1+tb1 == 0 || rb2+tb2 == 0 {
		return 0, 0, nil
	}

	durationSec := uint64(duration / time.Second)

	return (rb2 - rb1) / durationSec, (tb2 - tb1) / durationSec, nil
}

// GetIOUtil return slice (device -> utilization) with IO utilization
func GetIOUtil(duration time.Duration) (map[string]float64, error) {
	result := make(map[string]float64)

	fsInfoPrev, err := GetFSInfo()

	if err != nil {
		return nil, err
	}

	prevCPUInfo, err := getCPUStats()

	if err != nil {
		return nil, err
	}

	time.Sleep(duration)

	fsInfoCur, err := GetFSInfo()

	if err != nil {
		return nil, err
	}

	curCPUInfo, err := getCPUStats()

	if err != nil {
		return nil, err
	}

	curCPUUsage := float64(curCPUInfo.User + curCPUInfo.System + curCPUInfo.Idle + curCPUInfo.Wait)
	prevCPUUsage := float64(prevCPUInfo.User + prevCPUInfo.System + prevCPUInfo.Idle + prevCPUInfo.Wait)

	deltams := 1000.0 * (curCPUUsage - prevCPUUsage) / float64(curCPUInfo.Count) / _HZ

	for n, f := range fsInfoPrev {
		if fsInfoPrev[n].IOStats != nil && fsInfoCur[n].IOStats != nil {
			ticks := float64(fsInfoCur[n].IOStats.IOQueueMs - fsInfoPrev[n].IOStats.IOQueueMs)
			util := 100.0 * ticks / deltams

			if util > 100.0 {
				util = 100.0
			}

			result[f.Device] = util
		}
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readFileContent(file string) ([]string, error) {
	result := []string{}

	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	result = strings.Split(string(content), "\n")

	if len(result) == 0 {
		return nil, errors.New("File " + file + " is empty")
	}

	return result, nil
}

func cleanSlice(s []string) []string {
	var result []string

	for _, item := range s {
		if item != "" {
			result = append(result, item)
		}
	}

	return result
}

func byteSliceToString(s [65]int8) string {
	result := ""

	for _, r := range s {
		if r == 0 {
			break
		}

		result += string(r)
	}

	return result
}

func stringSliceToUintSlice(s []string) []uint64 {
	var result []uint64

	for _, i := range s {
		iu, _ := strconv.ParseUint(i, 10, 64)
		result = append(result, iu)
	}

	return result
}

func getActiveInterfacesBytes(is map[string]*InterfaceInfo) (uint64, uint64) {
	var (
		received    uint64
		transmitted uint64
	)

	for name, info := range is {
		if len(name) < 3 || name[0:3] != "eth" {
			continue
		}

		if info.ReceivedBytes == 0 && info.TransmittedBytes == 0 {
			continue
		}

		received += info.ReceivedBytes
		transmitted += info.TransmittedBytes
	}

	return received, transmitted
}

// getCPUStats return basicCPUInfo
func getCPUStats() (basicCPUInfo, error) {
	content, err := readFileContent(_PROC_CPUINFO)

	if err != nil || len(content) <= 1 {
		return basicCPUInfo{}, errors.New("Can't parse file " + _PROC_CPUINFO)
	}

	result := basicCPUInfo{}

	for _, line := range content {
		if strings.HasPrefix(line, "cpu") {
			result.Count++
		}
	}

	result.Count--

	cpu := strings.Replace(content[0], "cpu  ", "", -1)
	cpua := strings.Split(cpu, " ")

	result.User, _ = strconv.ParseUint(cpua[0], 10, 64)
	result.Nice, _ = strconv.ParseUint(cpua[1], 10, 64)
	result.System, _ = strconv.ParseUint(cpua[2], 10, 64)
	result.Idle, _ = strconv.ParseUint(cpua[3], 10, 64)
	result.Wait, _ = strconv.ParseUint(cpua[4], 10, 64)
	result.IRQ, _ = strconv.ParseUint(cpua[5], 10, 64)
	result.SRQ, _ = strconv.ParseUint(cpua[6], 10, 64)
	result.Steal, _ = strconv.ParseUint(cpua[7], 10, 64)

	result.Total = result.User + result.System + result.Nice + result.Idle + result.Wait + result.IRQ + result.SRQ + result.Steal

	return result, nil
}

func isFileExist(path string) bool {
	if path == "" {
		return false
	}

	return syscall.Access(path, syscall.F_OK) == nil
}
