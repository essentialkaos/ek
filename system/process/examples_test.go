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

func Example_getTree() {
	process, err := GetTree()

	if err != nil {
		return
	}

	// process is a top process in the tree
	fmt.Println(process.PID)
}

func Example_getList() {
	processes, err := GetList()

	if err != nil {
		return
	}

	// processes is slice with info about all active processes
	fmt.Println(processes)
}

func Example_calculateCPUUsage() {
	pid := 1345
	duration := time.Second * 15

	sample1, _ := GetSample(pid)
	time.Sleep(duration)
	sample2, _ := GetSample(pid)

	fmt.Printf("CPU Usage: %g%%\n", CalculateCPUUsage(sample1, sample2, duration))
}
