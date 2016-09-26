package system

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

func Example_exec() {
	err := Exec("/bin/echo", "abc", "123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func Example_getFSInfo() {
	fsInfo, err := GetFSInfo()

	if err != nil {
		return
	}

	// info is slice path -> info
	for path, info := range fsInfo {
		fmt.Printf(
			"Path: %s Type: %s Device: %s Used: %d Free: %d Total: %d",
			path, info.Type, info.Device, info.Used, info.Free, info.Total,
		)
	}
}
