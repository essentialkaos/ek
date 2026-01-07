//go:build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetFSUsage() {
	fsInfo, err := GetFSUsage()

	if err != nil {
		return
	}

	// info is slice path -> info
	for path, info := range fsInfo {
		fmt.Printf(
			"Path: %s Type: %s Device: %s Used: %d Free: %d Total: %d\n",
			path, info.Type, info.Device, info.Used, info.Free, info.Total,
		)
	}
}

func ExampleGetIOStats() {
	ioStats, err := GetIOStats()

	if err != nil {
		return
	}

	// print info for each device
	for device, info := range ioStats {
		fmt.Printf("Device: %s", device)
		fmt.Printf(
			"  ReadComplete: %d ReadMerged: %d ReadSectors: %d ReadMs: %d\n",
			info.ReadComplete, info.ReadMerged, info.ReadSectors, info.ReadMs,
		)

		fmt.Printf(
			"  WriteComplete: %d WriteMerged: %d WriteSectors: %d WriteMs: %d\n",
			info.WriteComplete, info.WriteMerged, info.WriteSectors, info.WriteMs,
		)

		fmt.Printf(
			"  IOPending: %d IOMs: %d IOQueueMs: %d\n\n",
			info.IOPending, info.IOMs, info.IOQueueMs,
		)
	}
}

func ExampleGetIOUtil() {
	// get 5 sec IO utilization
	ioUtil, err := GetIOUtil(5 * time.Second)

	if err != nil {
		return
	}

	// print utilization for each device
	for device, utilization := range ioUtil {
		fmt.Printf("Device: %s Utilization: %g\n", device, utilization)
	}
}

func ExampleGetNetworkSpeed() {
	input, output, err := GetNetworkSpeed(5 * time.Second)

	if err != nil {
		return
	}

	// print input and output speed for all interfaces
	fmt.Printf("Input: %d kb/s\n Output: %d kb/s\n", input/1024, output/1024)
}

func ExampleGetUptime() {
	uptime, err := GetUptime()

	if err != nil {
		return
	}

	// print uptime
	fmt.Printf("Uptime: %d seconds\n", uptime)
}

func ExampleGetLA() {
	la, err := GetLA()

	if err != nil {
		return
	}

	// print 1, 5 and 15 min load average
	fmt.Printf("Min1: %g Min5: %g Min15: %g\n", la.Min1, la.Min5, la.Min15)
}

func ExampleGetMemUsage() {
	usage, err := GetMemUsage()

	if err != nil {
		return
	}

	// print all available memory info
	fmt.Printf("MemTotal: %d\n", usage.MemTotal)
	fmt.Printf("MemFree: %d\n", usage.MemFree)
	fmt.Printf("MemUsed: %d\n", usage.MemUsed)
	fmt.Printf("MemUsedPerc: %g\n", usage.MemUsedPerc())
	fmt.Printf("Buffers: %d\n", usage.Buffers)
	fmt.Printf("Cached: %d\n", usage.Cached)
	fmt.Printf("Active: %d\n", usage.Active)
	fmt.Printf("Inactive: %d\n", usage.Inactive)
	fmt.Printf("SwapTotal: %d\n", usage.SwapTotal)
	fmt.Printf("SwapFree: %d\n", usage.SwapFree)
	fmt.Printf("SwapUsed: %d\n", usage.SwapUsed)
	fmt.Printf("SwapUsedPerc: %g\n", usage.SwapUsedPerc())
	fmt.Printf("SwapCached: %d\n", usage.SwapCached)
	fmt.Printf("Dirty: %d\n", usage.Dirty)
	fmt.Printf("Slab: %d\n", usage.Slab)
}

func ExampleGetCPUUsage() {
	usage, err := GetCPUUsage(time.Minute)

	if err != nil {
		return
	}

	// print all available CPU info
	fmt.Printf("User: %f\n", usage.User)
	fmt.Printf("System: %f\n", usage.System)
	fmt.Printf("Nice: %f\n", usage.Nice)
	fmt.Printf("Idle: %f\n", usage.Idle)
	fmt.Printf("Wait: %f\n", usage.Wait)
	fmt.Printf("Average: %f\n", usage.Average)
	fmt.Printf("CPU Count: %d\n", usage.Count)
}

func ExampleGetSystemInfo() {
	sysInfo, err := GetSystemInfo()

	if err != nil {
		return
	}

	// print all available system info
	fmt.Printf("Hostname: %s\n", sysInfo.Hostname)
	fmt.Printf("OS: %s\n", sysInfo.OS)
	fmt.Printf("Kernel: %s\n", sysInfo.Kernel)
	fmt.Printf("Arch: %s\n", sysInfo.Arch)
}
