package sortutil

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

func ExampleVersions() {
	versionSlice := []string{
		"2.0-5",
		"1.3b",
		"1.1",
		"1.3",
		"1.1.6",
		"1",
		"2.0",
		"2.0-1",
	}

	Versions(versionSlice)

	fmt.Println(versionSlice)

	// Output:
	// [1 1.1 1.1.6 1.3 1.3b 2.0 2.0-1 2.0-5]
}

func ExampleStrings() {
	stringSlice := []string{
		"Alisa",
		"Luna",
		"remedios",
		"Ona",
		"Eugene",
		"lorriane",
		"Zachariah",
		"cecily",
		"eleonora",
		"Dotty",
	}

	// Case insensitive sorting
	Strings(stringSlice, false)

	fmt.Println(stringSlice)

	// Case sensitive sorting
	Strings(stringSlice, true)

	fmt.Println(stringSlice)

	// Output:
	// [Alisa Dotty Eugene Luna Ona Zachariah cecily eleonora lorriane remedios]
	// [Alisa cecily Dotty eleonora Eugene lorriane Luna Ona remedios Zachariah]
}
