package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
