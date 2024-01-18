package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v12/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGenerate() {
	info := usage.NewInfo()

	info.AddCommand("test", "Test data")
	info.AddCommand("clean", "Clean data")

	info.AddOption("d:dir", "Path to working directory", "dir")

	info.BoundOptions("test", "d:dir")
	info.BoundOptions("clean", "d:dir")

	info.AddOption("nc:no-color", "Disable colors in output")
	info.AddOption("h:help", "Show this help message")
	info.AddOption("v:version", "Show version")

	fmt.Println(Generate(info, "app", "txt"))
}
