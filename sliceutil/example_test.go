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

func ExampleCopy() {
	s1 := []string{"A", "B", "C"}
	s2 := Copy(s1)

	fmt.Printf("%v\n", s2)
	// Output: [A B C]
}

func ExampleIsEqual() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{1, 3, 2, 4}

	fmt.Printf("%v == %v → %t\n", s1, s2, IsEqual(s1, s2))
	fmt.Printf("%v == %v → %t\n", s2, s3, IsEqual(s2, s3))
	// Output:
	// [1 2 3 4] == [1 2 3 4] → true
	// [1 2 3 4] == [1 3 2 4] → false
}

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
