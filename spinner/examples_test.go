package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleShow() {
	file := "file.txt"

	// Customize design
	OkSymbol = "🤟"     // Set success status symbol
	OkColorTag = "{b}" // Set success status symbol color

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

func ExampleSkip() {
	Show("My skipped running task")
	time.Sleep(time.Second)
	Skip()
}
