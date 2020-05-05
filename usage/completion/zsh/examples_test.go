package zsh

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"pkg.re/essentialkaos/ek.v12/options"
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

	opts := options.Map{
		"d:dir":       {},
		"nc:no-color": {Type: options.BOOL},
		"h:help":      {Type: options.BOOL},
		"v:version":   {Type: options.BOOL},
	}

	fmt.Println(Generate(info, opts, "app"))
}
