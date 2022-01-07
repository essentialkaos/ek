package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleShow() {
	file := "file.txt"
	Show("Removing file %s", file)
	time.Sleep(time.Second)
	Done(true)
}

func ExampleUpdate() {
	Show("My long running task")
	time.Sleep(time.Second)
	Show("My long running task still working")
	time.Sleep(time.Second)
	Done(true)
}

func ExampleDone() {
	Show("My long running task with good result")
	time.Sleep(time.Second)
	Done(true)

	Show("My long running task with bad result")
	time.Sleep(time.Second)
	Done(false)
}
