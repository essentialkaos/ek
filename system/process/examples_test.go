// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
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
