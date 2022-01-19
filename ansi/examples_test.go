package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Example_HasCodes() {
	input := "Hello"
	fmt.Println(HasCodes(input))

	input = "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(HasCodes(input))

	// Output:
	// false
	// true
}

func Example_HasCodesBytes() {
	input := []byte("Hello")
	fmt.Println(HasCodesBytes(input))

	input = []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(HasCodesBytes(input))

	// Output:
	// false
	// true
}

func Example_RemoveCodes() {
	input := "\033[40;38;5;82mHello\x1B[0m"
	fmt.Println(RemoveCodes(input))
	// Output:
	// Hello
}

func Example_RemoveCodesBytes() {
	input := []byte("\033[40;38;5;82mHello\x1B[0m")
	fmt.Println(string(RemoveCodesBytes(input)))
	// Output:
	// Hello
}
