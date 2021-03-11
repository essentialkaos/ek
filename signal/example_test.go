package signal

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
