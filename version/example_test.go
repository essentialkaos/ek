package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleParse() {
	v, err := Parse("6.12.1-beta2+exp.sha.5114f85")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Major: %d\n", v.Major())
	fmt.Printf("Minor: %d\n", v.Minor())
	fmt.Printf("Patch: %d\n", v.Patch())
	fmt.Printf("PreRelease: %s\n", v.PreRelease())
	fmt.Printf("Build: %s\n", v.Build())

	// Output:
	// Major: 6
	// Minor: 12
	// Patch: 1
	// PreRelease: beta2
	// Build: exp.sha.5114f85
}

func ExampleVersion_Int() {
	v, _ := Parse("14.8.22")

	fmt.Println(v.Int())

	// Output:
	// 14008022
}

func ExampleVersion_Equal() {
	v1, _ := Parse("1")
	v2, _ := Parse("1.0.0")
	v3, _ := Parse("1.0.1")

	fmt.Printf("%s = %s → %t\n", v1.String(), v2.String(), v1.Equal(v2))
	fmt.Printf("%s = %s → %t\n", v1.String(), v3.String(), v1.Equal(v3))

	// Output:
	// 1 = 1.0.0 → true
	// 1 = 1.0.1 → false
}

func ExampleVersion_Less() {
	v1, _ := Parse("1")
	v2, _ := Parse("1.0.0")
	v3, _ := Parse("1.0.1")

	fmt.Printf("%s < %s → %t\n", v1.String(), v2.String(), v1.Less(v2))
	fmt.Printf("%s < %s → %t\n", v1.String(), v3.String(), v1.Less(v3))

	// Output:
	// 1 < 1.0.0 → false
	// 1 < 1.0.1 → true
}

func ExampleVersion_Greater() {
	v1, _ := Parse("1")
	v2, _ := Parse("1.0.0")
	v3, _ := Parse("1.0.1")

	fmt.Printf("%s > %s → %t\n", v1.String(), v2.String(), v1.Greater(v2))
	fmt.Printf("%s > %s → %t\n", v3.String(), v1.String(), v3.Greater(v1))

	// Output:
	// 1 > 1.0.0 → false
	// 1.0.1 > 1 → true
}

func ExampleVersion_Contains() {
	v1, _ := Parse("1.0")
	v2, _ := Parse("1.0.1")
	v3, _ := Parse("1.1")

	fmt.Printf("%s contains %s → %t\n", v1.String(), v2.String(), v1.Contains(v2))
	fmt.Printf("%s contains %s → %t\n", v1.String(), v3.String(), v1.Contains(v3))

	// Output:
	// 1.0 contains 1.0.1 → true
	// 1.0 contains 1.1 → false
}
