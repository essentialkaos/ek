package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleHas() {
	input := "Hello"
	fmt.Println(Has(input))

	input = "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(Has(input))

	// Output:
	// false
	// true
}

func ExampleHasBytes() {
	input := []byte("Hello")
	fmt.Println(HasBytes(input))

	input = []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(HasBytes(input))

	// Output:
	// false
	// true
}

func ExampleRemove() {
	input := "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(Remove(input))
	// Output:
	// Hello
}

func ExampleRemoveBytes() {
	input := []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(string(RemoveBytes(input)))
	// Output:
	// Hello
}
