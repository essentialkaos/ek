// +build linux

package process

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
