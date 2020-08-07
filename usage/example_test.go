package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleAbout_Render() {
	about := About{
		App:     "MySupperApp",
		Desc:    "My super golang utility",
		Version: "1.0.1",
		Release: "-44",
		Build:   "17746", // Number of build or commit hash
		Year:    2009,    // Year when company was founded
		License: "MIT",
		Owner:   "John Dow <john@domain.com>",
	}

	about.Render()
}

func ExampleInfo_AddGroup() {
	info := NewInfo("myapp", "items...")

	// You can add custom commands groups
	info.AddGroup("External Commands")

	info.AddCommand("publish", "Publish items")

	// You can define option argument name
	info.AddOption("o:output", "Output", "file")

	// render all data
	info.Render()
}

func ExampleInfo_AddCommand() {
	info := NewInfo("myapp", "items...")

	// You can define command arguments names
	info.AddCommand("add", "Add item", "file")

	// Also, you can mark optional arguments using ? prefix
	info.AddCommand("remove", "Remove item", "file", "?mode")
	info.AddCommand("list", "List items")

	// You can add custom commands groups
	info.AddGroup("External Commands")

	info.AddCommand("publish", "Publish items")

	// render all data
	info.Render()
}

func ExampleInfo_AddOption() {
	info := NewInfo("myapp", "items...")

	// AddOption support options in format used in ek.arg package
	info.AddOption("v:version", "Print version")

	// You can define option argument name
	info.AddOption("o:output", "Output", "file")

	// render all data
	info.Render()
}

func ExampleInfo_AddExample() {
	info := NewInfo("myapp", "items...")

	info.AddCommand("add", "Add item", "file")
	info.AddCommand("remove", "Remove item", "file", "?mode")

	// First part with application name will be automatically added
	info.AddExample("add file.dat")

	// This is example with description
	info.AddExample("remove file.dat", "Remove file.dat")

	// render all data
	info.Render()
}

func ExampleInfo_AddRawExample() {
	info := NewInfo("myapp", "items...")

	info.AddCommand("add", "Add item", "file")
	info.AddCommand("remove", "Remove item", "file", "?mode")

	// Raw example (without application name) without description
	info.AddRawExample("add file.dat")

	// Raw example (without application name) with description
	info.AddRawExample("remove file.dat", "Remove file.dat")

	// render all data
	info.Render()
}

func ExampleInfo_AddSpoiler() {
	info := NewInfo("myapp", "items...")

	// Spoiler will be shown before all commands and options
	info.AddSpoiler("This is my supadupa utility")

	// render all data
	info.Render()
}

func ExampleInfo_BoundOptions() {
	info := NewInfo("myapp", "items...")

	info.AddCommand("publish", "Publish items")

	info.AddOption("o:output", "Output", "file")

	// Link command and options (will be used for completion generation)
	info.BoundOptions("publish", "o:output")

	// render all data
	info.Render()
}

func ExampleInfo_GetCommand() {
	info := NewInfo("myapp", "items...")

	// You can define command arguments names
	info.AddCommand("add", "Add item", "file")

	// Also, you can mark optional arguments using ? prefix
	info.AddCommand("remove", "Remove item", "file", "?mode")
	info.AddCommand("list", "List items")

	cmd := info.GetCommand("list")

	fmt.Println(cmd.Desc)
	// Output: List items
}

func ExampleInfo_GetOption() {
	info := NewInfo("myapp", "items...")

	// AddOption support options in format used in ek.arg package
	info.AddOption("v:version", "Print version")

	// You can define option argument name
	info.AddOption("o:output", "Output file", "file")

	opt := info.GetOption("o:output")

	fmt.Println(opt.Desc)
	// Output: Output file
}

func ExampleInfo_Render() {
	info := NewInfo("myapp", "items...")

	// Spoiler will be shown before all commands and options
	info.AddSpoiler("This is my supadupa utility")

	// You can define command arguments names
	info.AddCommand("add", "Add item", "file")

	// Also, you can mark optional arguments using ? prefix
	info.AddCommand("remove", "Remove item", "file", "?mode")
	info.AddCommand("list", "List items")

	// You can add custom commands groups
	info.AddGroup("External Commands")

	info.AddCommand("publish", "Publish items")

	info.AddOption("--help", "Print help content")

	// AddOption support options in format used in ek.arg package
	info.AddOption("v:version", "Print version")

	// You can define option argument name
	info.AddOption("o:output", "Output", "file")

	// Link command and options (will be used for completion generation)
	info.BoundOptions("publish", "o:output")

	// First part with application name will be automatically added
	info.AddExample("add file.dat")

	// This is example with description
	info.AddExample("remove file.dat", "Remove file.dat")

	// Raw example without description
	info.AddRawExample("add file.dat")

	// Raw example with description
	info.AddRawExample("remove file.dat", "Remove file.dat")

	// render all data
	info.Render()
}
