package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePrintln() {
	// print colored text
	Println("{d}black{!}")
	Println("{r}red{!}")
	Println("{y}yellow{!}")
	Println("{b}blue{!}")
	Println("{c}cyan{!}")
	Println("{m}magenta{!}")
	Println("{g}green{!}")
	Println("{s}light grey{!}")

	// use text modificators

	// light colors
	Println("{r-}light red{!}")
	Println("{r-}dark grey{!}")

	// bold + color
	Println("{r*}red{!}")
	Println("{g*}green{!}")

	// dim
	Println("{r^}red{!}")
	Println("{g^}green{!}")

	// underline
	Println("{r_}red{!}")
	Println("{g_}green{!}")

	// blink
	Println("{r~}red{!}")
	Println("{g~}green{!}")

	// reverse
	Println("{r@}red{!}")
	Println("{g@}green{!}")

	// background color
	Println("{D}black{!}")
	Println("{R}red{!}")
	Println("{Y}yellow{!}")
	Println("{B}blue{!}")
	Println("{C}cyan{!}")
	Println("{M}magenta{!}")
	Println("{G}green{!}")
	Println("{S}light grey{!}")
}

func ExamplePrintf() {
	var (
		user         = "Bob"
		isUserOnline = true
	)

	if isUserOnline {
		Printf("User {c}%s{!} is {g}online{!}\n")
	} else {
		Printf("User {c}%s{!} is {r}offline{!}\n")
	}
}

func ExampleBell() {
	// terminal bell
	Bell()
}

func ExampleNewLine() {
	// just print new line
	NewLine()
}

func ExampleClean() {
	// Remove color tags from text
	fmt.Println(Clean("{r}Text{!}"))

	// Output: Text
}

func ExampleT_printf() {
	t := T{}

	t.Printf("This is temporary text")
	t.Printf("This message replace previous message")
}
