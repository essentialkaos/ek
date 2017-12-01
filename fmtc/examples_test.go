package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePrintln() {

	// print colored text
	// {!} is tag for style reset
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

	// many tags at once
	// underline, cyan text with the red background
	Println("{cR_}text{!}")
}

func ExamplePrintf() {
	var (
		user         = "Bob"
		isUserOnline = true
	)

	if isUserOnline {
		Printf("User {c}%s{!} is {g}online{!}\n", user)
	} else {
		Printf("User {c}%s{!} is {r}offline{!}\n", user)
	}
}

func ExampleBell() {
	// terminal bell
	Bell()
}

func ExampleNewLine() {
	// just print a new line
	NewLine()
}

func ExampleClean() {
	// Remove color tags from text
	fmt.Println(Clean("{r}Text{!}"))

	// Output: Text
}

func ExampleTprintf() {
	TPrintf("This is temporary text")
	TPrintf("This message replace previous message")
}
