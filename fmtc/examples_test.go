package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePrint() {
	// print colored text
	// {!} is tag for style reset
	Print("{d}black{!}\n")
	Print("{r}red{!}\n")
	Print("{y}yellow{!}\n")
	Print("{b}blue{!}\n")
	Print("{c}cyan{!}\n")
	Print("{m}magenta{!}\n")
	Print("{g}green{!}\n")
	Print("{s}light grey{!}\n")

	// use text modificators

	// light colors
	Print("{r-}light red{!}\n")
	Print("{r-}dark grey{!}\n")

	// bold + color
	Print("{r*}red{!}\n")
	Print("{g*}green{!}\n")

	// dim
	Print("{r^}red{!}\n")
	Print("{g^}green{!}\n")

	// underline
	Print("{r_}red{!}\n")
	Print("{g_}green{!}\n")

	// blink
	Print("{r~}red{!}\n")
	Print("{g~}green{!}\n")

	// reverse
	Print("{r@}red{!}\n")
	Print("{g@}green{!}\n")

	// background color
	Print("{D}black{!}\n")
	Print("{R}red{!}\n")
	Print("{Y}yellow{!}\n")
	Print("{B}blue{!}\n")
	Print("{C}cyan{!}\n")
	Print("{M}magenta{!}\n")
	Print("{G}green{!}\n")
	Print("{S}light grey{!}\n")

	// many tags at once
	// underline, cyan text with the red background
	Print("{cR_}text{!}\n")

	// many tags in once
	Print("{r}{*}red and bold{!}\n")

	// modificator reset
	Print("{r}{*}red and bold {!*}just red{!}\n")

	// 256 colors (# for foreground, % for background)
	Print("{#201}pink text{!}\n")
	Print("{%201}pink background{!}\n")

	// 24-bit colors (# for foreground, % for background)
	Print("{#7cfc00}lawngreen text{!}\n")
	Print("{%6a5acd}slateblue background{!}\n")

	// Named colors
	// All color names must match the next regex
	// pattern: [a-zA-Z0-9_]+

	// Add new color with name "error"
	NameColor("error", "{r}")

	// Print using named color
	Print("{?error}lawngreen text{!}\n")

	// Redefine "error" color to 24-bit color
	NameColor("error", "{#ff0000}")

	// Remove named color
	RemoveColor("error")
}

func ExamplePrintf() {
	// print colored text
	// {!} is tag for style reset
	Printf("{d}%s{!}\n", "black")
	Printf("{r}%s{!}\n", "red")
	Printf("{y}%s{!}\n", "yellow")
	Printf("{b}%s{!}\n", "blue")
	Printf("{c}%s{!}\n", "cyan")
	Printf("{m}%s{!}\n", "magenta")
	Printf("{g}%s{!}\n", "green")
	Printf("{s}%s{!}\n", "light grey")

	// use text modificators

	// light colors
	Printf("{r-}%s{!}\n", "light red")
	Printf("{r-}%s{!}\n", "dark grey")

	// bold + color
	Printf("{r*}%s{!}\n", "red")
	Printf("{g*}%s{!}\n", "green")

	// dim
	Printf("{r^}%s{!}\n", "red")
	Printf("{g^}%s{!}\n", "green")

	// underline
	Printf("{r_}%s{!}\n", "red")
	Printf("{g_}%s{!}\n", "green")

	// blink
	Printf("{r~}%s{!}\n", "red")
	Printf("{g~}%s{!}\n", "green")

	// reverse
	Printf("{r@}%s{!}\n", "red")
	Printf("{g@}%s{!}\n", "green")

	// background color
	Printf("{D}%s{!}\n", "black")
	Printf("{R}%s{!}\n", "red")
	Printf("{Y}%s{!}\n", "yellow")
	Printf("{B}%s{!}\n", "blue")
	Printf("{C}%s{!}\n", "cyan")
	Printf("{M}%s{!}\n", "magenta")
	Printf("{G}%s{!}\n", "green")
	Printf("{S}%s{!}\n", "light grey")

	// many tags at once
	// underline, cyan text with the red background
	Printf("{cR_}%s{!}\n", "text")

	// many tags in once
	Printf("{r}{*}%s{!}\n", "red and bold")

	// 256 colors (# for foreground, % for background)
	Printf("{#201}%s{!}\n", "pink text")
	Printf("{%201}%s{!}\n", "pink background")

	// 24-bit colors (# for foreground, % for background)
	Printf("{#7cfc00}%s{!}\n", "lawngreen text")
	Printf("{%6a5acd}%s{!}\n", "slateblue background")
}

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

	// many tags in once
	Println("{r}{*}red and bold{!}")

	// modificator reset
	Println("{r}{*}red and bold {!*}just red{!}")

	// 256 colors (# for foreground, % for background)
	Println("{#201}pink text{!}")
	Println("{%201}pink background{!}")

	// 24-bit colors (# for foreground, % for background)
	Println("{#7cfc00}lawngreen text{!}")
	Println("{%6a5acd}slateblue background{!}")
}

func ExampleFprint() {
	Fprint(os.Stderr, "{r}This is error message{!}\n")
	Fprint(os.Stdout, "{g}This is normal message{!}\n")
}

func ExampleFprintf() {
	Fprintf(os.Stderr, "{r}%s{!}\n", "This is error message")
	Fprintf(os.Stdout, "{g}%s{!}\n", "This is normal message")
}

func ExampleFprintln() {
	Fprintln(os.Stderr, "{r}This is error message{!}")
	Fprintln(os.Stdout, "{g}This is normal message{!}")
}

func ExampleSprint() {
	msg := Sprint("{r}This is error message{!}\n")
	fmt.Print(msg)
}

func ExampleSprintf() {
	msg := Sprintf("{r}%s{!}\n", "This is error message")
	fmt.Print(msg)
}

func ExampleSprintln() {
	msg := Sprintln("{r}This is error message{!}")
	fmt.Print(msg)
}

func ExampleErrorf() {
	err := Errorf("This is error")
	fmt.Print(err.Error())
}

func ExampleTPrint() {
	TPrint("{s}This is temporary text{!}\n")
	time.Sleep(time.Second)
	TPrint("{*}This message replace previous message after 1 sec{!}\n")
}

func ExampleTPrintf() {
	TPrintf("{s}%s{!}\n", "This is temporary text")
	time.Sleep(time.Second)
	TPrintf("{*}%s{!}\n", "This message replace previous message after 1 sec")
}

func ExampleTPrintln() {
	TPrintln("{s}This is temporary text{!}")
	time.Sleep(time.Second)
	TPrintln("{*}This message replace previous message after 1 sec{!}")
}

func ExampleLPrint() {
	// Only "This is text" will be shown
	LPrint(12, "{r}This is text {g} with colors{!}")
}

func ExampleLPrintf() {
	// Only "This is text" will be shown
	LPrintf(12, "{r}This is %s {g} with colors{!}", "text")
}

func ExampleLPrintln() {
	// Only "This is text" will be shown
	LPrintln(12, "{r}This is %s {g} with colors{!}")
}

func ExampleTLPrint() {
	// Power of TPrint and LPrint in one method
	TLPrint(15, "{s}This is temporary text{!}")
	time.Sleep(time.Second)
	TLPrint(15, "{*}This message replace previous message after 1 sec{!}")
}

func ExampleTLPrintf() {
	// Power of TPrintf and LPrintf in one method
	TLPrintf(15, "{s}%s{!}", "This is temporary text")
	time.Sleep(time.Second)
	TLPrintf(15, "{*}%s{!}", "This message replace previous message after 1 sec")
}

func ExampleTLPrintln() {
	// Power of TPrintln and LPrintln in one method
	TLPrintln(15, "{s}This is temporary text{!}")
	time.Sleep(time.Second)
	TLPrintln(15, "{*}This message replace previous message after 1 sec{!}")
}

func ExampleBell() {
	// terminal bell
	Bell()
}

func ExampleNewLine() {
	// just print a new line
	NewLine()
	// Print 3 new lines
	NewLine(3)
}

func ExampleClean() {
	// Remove color tags from text
	fmt.Println(Clean("{r}Text{!}"))

	// Output: Text
}

func ExampleIs256ColorsSupported() {
	fmt.Printf("256 Colors Supported: %t\n", Is256ColorsSupported())
}

func ExampleIsTrueColorSupported() {
	fmt.Printf("TrueColor Supported: %t\n", IsTrueColorSupported())
}

func ExampleIsTag() {
	fmt.Printf("%s is tag: %t\n", "{r}", IsTag("{r}"))
	fmt.Printf("%s is tag: %t\n", "[r]", IsTag("[r]"))

	// Output:
	// {r} is tag: true
	// [r] is tag: false
}

func ExampleIf() {
	userIsAdmin := true

	If(userIsAdmin).Println("You are admin!")
}

func ExampleNameColor() {
	// Add new color with name "error"
	NameColor("error", "{r}")

	// Print a message with named color
	Print("{?error}lawngreen text{!}\n")

	// Redefine "error" color to 24-bit color
	NameColor("error", "{#ff0000}")

	// Print a message with new color
	Print("{?error}lawngreen text{!}\n")
}

func ExampleRemoveColor() {
	// Add new color with name "error"
	NameColor("error", "{r}")

	// Print a message with named color
	Print("{?error}lawngreen text{!}\n")

	// Remove named color
	RemoveColor("error")
}
