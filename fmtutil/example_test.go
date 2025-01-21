package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleSeparator() {
	// You can change color of separator symbols and title using fmtc color tags

	SeparatorColorTag = "{r}"       // Set color to red
	SeparatorTitleColorTag = "{r*}" // Set color to red with bold weight

	// Or you can remove colors
	SeparatorColorTag = ""
	SeparatorTitleColorTag = ""

	// This is tiny separator (just 1 line)
	Separator(true)

	// This is wide separator with newlines before and after separator
	Separator(false)

	// You can set title of separator
	Separator(true, "MY SEPARATOR")
}

func ExamplePrettyNum() {
	var (
		n1 int     = 10
		n2 uint    = 5000
		n3 float64 = 6128750.26
	)

	fmt.Printf("%d → %s\n", n1, PrettyNum(n1))
	fmt.Printf("%d → %s\n", n2, PrettyNum(n2))
	fmt.Printf("%.2f → %s\n", n3, PrettyNum(n3))

	// Set default order separator
	OrderSeparator = " "

	fmt.Printf("%.2f → %s\n", n3, PrettyNum(n3))

	// Use custom order separator
	fmt.Printf("%.2f → %s\n", n3, PrettyNum(n3, "|"))

	// Output:
	// 10 → 10
	// 5000 → 5,000
	// 6128750.26 → 6,128,750.26
	// 6128750.26 → 6 128 750.26
	// 6128750.26 → 6|128|750.26
}

func ExamplePrettyBool() {
	fmt.Printf("%t → %s\n", true, PrettyBool(true))
	fmt.Printf("%t → %s\n", false, PrettyBool(false))
	fmt.Printf("%t → %s\n", true, PrettyBool(true, "Yep", "Nope"))
	fmt.Printf("%t → %s\n", false, PrettyBool(false, "Yep", "Nope"))

	// Output:
	// true → Y
	// false → N
	// true → Yep
	// false → Nope
}

func ExamplePrettyPerc() {
	var (
		n1 float64 = 0.123
		n2 float64 = 10.24
		n3 float64 = 1294.193
	)

	OrderSeparator = ","

	fmt.Printf("%f → %s\n", n1, PrettyPerc(n1))
	fmt.Printf("%f → %s\n", n2, PrettyPerc(n2))
	fmt.Printf("%f → %s\n", n3, PrettyPerc(n3))

	// Output:
	// 0.123000 → 0.12%
	// 10.240000 → 10.2%
	// 1294.193000 → 1,294.2%

}

func ExamplePrettySize() {
	s1 := 193
	s2 := 184713
	s3 := int64(46361936461)

	fmt.Printf("%d → %s\n", s1, PrettySize(s1))
	fmt.Printf("%d → %s\n", s2, PrettySize(s2))
	fmt.Printf("%d → %s\n", s3, PrettySize(s3))

	// Set size separator
	SizeSeparator = " "

	fmt.Printf("%d → %s\n", s3, PrettySize(s3))

	// Use custom order separator
	fmt.Printf("%d → %s\n", s3, PrettySize(s3, "|"))

	// Output:
	// 193 → 193B
	// 184713 → 180.4KB
	// 46361936461 → 43.2GB
	// 46361936461 → 43.2 GB
	// 46361936461 → 43.2|GB
}

func ExampleParseSize() {
	s1 := "160"
	s2 := "34Mb"
	s3 := "2.2 GB"

	fmt.Printf("%s → %d\n", s1, ParseSize(s1))
	fmt.Printf("%s → %d\n", s2, ParseSize(s2))
	fmt.Printf("%s → %d\n", s3, ParseSize(s3))

	// Output:
	// 160 → 160
	// 34Mb → 35651584
	// 2.2 GB → 2362232012
}

func ExampleFloat() {
	f1 := 0.3145
	f2 := 3.452
	f3 := 135.5215

	fmt.Printf("%f → %g\n", f1, Float(f1))
	fmt.Printf("%f → %g\n", f2, Float(f2))
	fmt.Printf("%f → %g\n", f3, Float(f3))

	// Output:
	// 0.314500 → 0.31
	// 3.452000 → 3.45
	// 135.521500 → 135.5
}

func ExampleWrap() {
	text := "Aenean tincidunt metus a tortor aramus, ut bibendum magna fringilla."

	fmt.Println(
		Wrap(text, "", 36),
	)
}

func ExampleColorizePassword() {
	password := ">+XY!b3Rog"

	fmt.Println(ColorizePassword(password, "{r}", "{y}", "{c}"))

	// Output:
	// {c}>+{r}XY{c}!{r}b{y}3{r}Rog{!}
}
