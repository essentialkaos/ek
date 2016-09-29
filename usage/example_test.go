package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleAbout_Render() {
	about := About{
		App:     "MySupperApp",
		Desc:    "My super golang utility",
		Version: "1.0.1",
		Release: "-44",
		Build:   "17746",
		Year:    2009, // Year when company was founded
		License: "MIT",
		Owner:   "John Dow <john@domain.com>",
	}

	about.Render()
}

func ExampleInfo_Render() {
	info := NewInfo("myapp", "file...")

	info.AddSpoiler("This is my supadupa utility")

	// you can add commands or options groups
	info.AddGroup("Basic Commands")

	info.AddCommand("add", "Add item")
	info.AddCommand("remove", "Remove item")
	info.AddCommand("list", "List items")

	info.AddGroup("External Commands")

	info.AddCommand("publish", "Publish items")

	info.AddOption("--help", "Print help content")

	// AddOption support options in format used in ek.arg package
	info.AddOption("v:version", "Print version")

	// first part with application name will be automatically added
	info.AddExample("add file.dat")

	// this is example with description
	info.AddExample("remove file.dat", "Remove file.da")

	// render all data
	info.Render()
}
