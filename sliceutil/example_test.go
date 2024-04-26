package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCopy() {
	s1 := []string{"A", "B", "C"}
	s2 := Copy(s1)

	fmt.Printf("%v", s2)
	// Output: [A B C]
}

func ExampleStringToInterface() {
	s := []string{"A", "B"}

	fmt.Printf("%v", StringToInterface(s))
	// Output: [A B]
}

func ExampleIntToInterface() {
	s := []int{1, 2}

	fmt.Printf("%v", IntToInterface(s))
	// Output: [1 2]
}

func ExampleErrorToString() {
	s := []error{fmt.Errorf("error1")}

	fmt.Printf("%v", ErrorToString(s))
	// Output: [error1]
}

func ExampleIndex() {
	s := []string{"A", "B", "C"}

	fmt.Println(Index(s, "C"))
	fmt.Println(Index(s, "D"))
	// Output: 2
	// -1
}

func ExampleExclude() {
	s := []string{"A", "B", "C", "D"}

	fmt.Println(Exclude(s, "B", "D"))
	// Output: [A C]
}

func ExampleDeduplicate() {
	s := []string{"A", "A", "A", "B", "C", "C", "D"}

	fmt.Println(Deduplicate(s))
	// Output: [A B C D]
}
