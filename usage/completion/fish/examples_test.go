package fish

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"pkg.re/essentialkaos/ek.v12/usage"
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

	fmt.Println(Generate(info, "app"))
}
