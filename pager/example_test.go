package pager

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleSetup() {
	// Use pager from PAGER env var or default (more)
	Setup("")

	// Or provide specific command.
	Setup("less -MQR")

	// Complete must be called at the end of the program work. You can call it with defer
	// in your main function.
	defer Complete()
}
