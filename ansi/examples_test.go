package ansi

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Example_HasCodes() {
	fmt.Println(HasCodes("Hello"))
	fmt.Println(HasCodes("\033[40;38;5;82mHello\x1B[0m"))

	// Output:
	// false
	// true
}

func Example_RemoveCodes() {
	fmt.Println(RemoveCodes("\033[40;38;5;82mHello\x1B[0m"))
	// Output:
	// Hello
}
