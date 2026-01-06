package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCreate() {
	// You can set default directory for pid files
	Dir = "/home/user/my-pids"

	err := Create("servicename")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("PID file created!")
}

func ExampleRemove() {
	err := Remove("servicename")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("PID file removed!")
}

func ExampleGet() {
	pid := Get("servicename")

	if pid == -1 {
		fmt.Println("Can't read PID from PID file")
	}

	fmt.Printf("PID is %d\n", pid)
}

func ExampleRead() {
	pid := Read("/var/run/httpd.pid")

	if pid == -1 {
		fmt.Println("Can't read PID from PID file")
	}

	fmt.Printf("PID is %d\n", pid)
}

func ExampleIsProcessWorks() {
	pid := 1234

	if IsProcessWorks(pid) {
		fmt.Printf("Process with PID %d is works\n", pid)
	} else {
		fmt.Printf("Process with PID %d isn't working\n", pid)
	}
}

func ExampleIsWorks() {
	if IsWorks("servicename") {
		fmt.Println("Process is works")
	} else {
		fmt.Println("Process isn't working")
	}
}
