package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNew() {
	pb := New(1000, "Downloading...")

	pbs := DefaultSettings
	pbs.RefreshRate = 50 * time.Millisecond
	pbs.ShowSpeed = false
	pbs.ShowRemaining = false

	defer pb.Finish()

	pb.Start()

	for range time.NewTicker(time.Second).C {
		pb.Add(1)
	}
}
