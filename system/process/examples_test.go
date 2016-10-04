// +build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

	// process is top process in tree
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
