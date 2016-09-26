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
			"Path: %s Type: %s Device: %s Used: %d Free: %d Total: %d\n",
			path, info.Type, info.Device, info.Used, info.Free, info.Total,
		)
	}
}

func Example_getIOStats() {
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
			"  IOPending: %d IOMs: %d IOQueueMs: %d\n",
			info.IOPending, info.IOMs, info.IOQueueMs,
		)
	}
}
