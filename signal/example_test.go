package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleHandlers_Track() {
	hupHandler := func() {
		fmt.Println("Got HUP signal")
	}

	usr1Handler := func() {
		fmt.Println("Got USR1 signal")
	}

	Handlers{
		HUP:  hupHandler,
		USR1: usr1Handler,
	}.Track()
}

func ExampleHandlers_TrackAsync() {
	hupHandler := func() {
		fmt.Println("Got HUP signal")
	}

	usr1Handler := func() {
		fmt.Println("Got USR1 signal")
	}

	Handlers{
		HUP:  hupHandler,
		USR1: usr1Handler,
	}.TrackAsync()

	time.Sleep(time.Hour)
}
