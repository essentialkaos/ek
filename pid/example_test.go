package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCreate() {
	// You can set default derectory for pid files
	Dir = "/home/user/my-pids"

	err := Create("servicename")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Pid file created!")
}

func ExampleRemove() {
	err := Remove("servicename")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Pid file removed!")
}

func ExampleGet() {
	pid := Get("servicename")

	if pid == -1 {
		fmt.Println("Can't read pid from pid file")
	}

	fmt.Printf("PID is %d\n", pid)
}
