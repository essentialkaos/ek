//go:build linux
// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetTree() {
	process, err := GetTree()

	if err != nil {
		return
	}

	// process is a top process in the tree
	fmt.Printf("%v\n", process.PID)
}

func ExampleGetList() {
	processes, err := GetList()

	if err != nil {
		return
	}

	// processes is slice with info about all active processes
	fmt.Printf("%v\n", processes)
}

func ExampleCalculateCPUUsage() {
	pid := 1345
	duration := time.Second * 15

	sample1, _ := GetSample(pid)
	time.Sleep(duration)
	sample2, _ := GetSample(pid)

	fmt.Printf("CPU Usage: %g%%\n", CalculateCPUUsage(sample1, sample2, duration))
}

func ExampleGetMemInfo() {
	info, err := GetMemInfo(1000)

	if err != nil {
		return
	}

	fmt.Printf("%v\n", info)
}

func ExampleGetMountInfo() {
	info, err := GetMountInfo(1000)

	if err != nil {
		return
	}

	fmt.Printf("%v\n", info)
}

func ExampleGetCPUPriority() {
	pid := os.Getpid()
	pr, ni, err := GetCPUPriority(pid)

	if err != nil {
		return
	}

	fmt.Printf("PR: %d | NI: %d\n", pr, ni)
}

func ExampleSetCPUPriority() {
	pid := os.Getpid()
	err := SetCPUPriority(pid, -20)

	if err != nil {
		return
	}

	pr, ni, err := GetCPUPriority(pid)

	if err != nil {
		return
	}

	fmt.Printf("PR: %d | NI: %d\n", pr, ni)
}

func ExampleGetIOPriority() {
	pid := os.Getpid()
	class, classdata, err := GetIOPriority(pid)

	if err != nil {
		return
	}

	fmt.Printf("Class: %d | Classdata: %d\n", class, classdata)
}

func ExampleSetIOPriority() {
	pid := os.Getpid()
	err := SetIOPriority(pid, PRIO_CLASS_REAL_TIME, 5)

	if err != nil {
		return
	}

	class, classdata, err := GetIOPriority(pid)

	if err != nil {
		return
	}

	fmt.Printf("Class: %d | Classdata: %d\n", class, classdata)
}
