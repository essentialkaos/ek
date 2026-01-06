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

func ExampleHasCodes() {
	input := "Hello"
	fmt.Println(HasCodes(input))

	input = "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(HasCodes(input))

	// Output:
	// false
	// true
}

func ExampleHasCodesBytes() {
	input := []byte("Hello")
	fmt.Println(HasCodesBytes(input))

	input = []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(HasCodesBytes(input))

	// Output:
	// false
	// true
}

func ExampleRemoveCodes() {
	input := "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(RemoveCodes(input))
	// Output:
	// Hello
}

func ExampleRemoveCodesBytes() {
	input := []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(string(RemoveCodesBytes(input)))
	// Output:
	// Hello
}
