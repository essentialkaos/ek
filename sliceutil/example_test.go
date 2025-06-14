package sliceutil

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

func ExampleStringToInterface() {
	s := []string{"A", "B"}

	fmt.Printf("%v\n", StringToInterface(s))
	// Output: [A B]
}

func ExampleIntToInterface() {
	s := []int{1, 2}

	fmt.Printf("%v\n", IntToInterface(s))
	// Output: [1 2]
}

func ExampleErrorToString() {
	s := []error{fmt.Errorf("error1")}

	fmt.Printf("%v\n", ErrorToString(s))
	// Output: [error1]
}

func ExampleExclude() {
	s := []string{"A", "B", "C", "D"}

	fmt.Println(Exclude(s, "B", "D"))
	// Output: [A C]
}

func ExampleJoin() {
	s1 := []string{"A", "B", "C", "D"}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []any{"John", 183, 98.123, false}

	fmt.Println(Join(s1, ":"))
	fmt.Println(Join(s2, ","))
	fmt.Println(Join(s3, ";"))

	// Output:
	// A:B:C:D
	// 1,2,3,4,5
	// John;183;98.123;false
}

func ExampleDiff() {
	s1 := []int{1, 2, 3, 5, 7, 8}
	s2 := []int{2, 3, 5, 6, 7}

	added, deleted := Diff(s1, s2)

	fmt.Printf("Added: %v\n", Join(added, ", "))
	fmt.Printf("Deleted: %v\n", Join(deleted, ", "))

	// Output:
	// Added: 6
	// Deleted: 1, 8
}

func ExampleShuffle() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Shuffle(s)

	fmt.Printf("Result: %v\n", s)
}

func ExampleFilter() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	r := Filter(s, func(v int, index int) bool {
		return v%2 == 0
	})

	fmt.Printf("Result: %v\n", r)

	// Output:
	// Result: [2 4 6 8 10]
}
