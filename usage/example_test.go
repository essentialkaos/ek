package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
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

func ExampleInfo_Render() {
	info := NewInfo("myapp", "items...")

	info.AddSpoiler("This is my supadupa utility")

	// You can define command arguments names
	info.AddCommand("add", "Add item", "file")

	// Also, you can mark optional arguments using ? prefix
	info.AddCommand("remove", "Remove item", "file", "?mode")
	info.AddCommand("list", "List items")

	// You can add custom commands or options groups
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

	// render all data
	info.Render()
}
