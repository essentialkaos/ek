package events

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewDispatcher() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)

	d.Dispatch("myEvent", "Hello!")
}

func ExampleDispatcher_AddHandler() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)
	d.AddHandler("unknownEvent", testHandler)

	err := d.Dispatch("myEvent", "Hello!")

	fmt.Println(err) // nil
}

func ExampleDispatcher_RemoveHandler() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)
	d.RemoveHandler("myEvent", testHandler)

	err := d.Dispatch("myEvent", "Hello!")

	fmt.Println(err) // not nil
}

func ExampleDispatcher_HasHandler() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)

	fmt.Println("myEvent:", d.HasHandler("myEvent", testHandler))
	fmt.Println("unknownEvent:", d.HasHandler("unknownEvent", testHandler))

	// Output:
	// myEvent: true
	// unknownEvent: false
}

func ExampleDispatcher_Dispatch() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)

	err := d.Dispatch("myEvent", "Hello!")

	fmt.Println(err) // nil
}

func ExampleDispatcher_DispatchAndWait() {
	d := NewDispatcher()

	d.AddHandler("myEvent", testHandler)

	err := d.DispatchAndWait("myEvent", "Hello!")

	fmt.Println(err) // nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func testHandler(payload interface{}) {
	fmt.Printf("Got payload: %v\n", payload)
}
