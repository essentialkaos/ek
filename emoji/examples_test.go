package emoji

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGet() {
	fmt.Println(Get("zap"))
	// Output: ⚡️
}

func ExampleGetName() {
	fmt.Println(GetName("⚡️"))
	// Output: zap
}

func ExampleFind() {
	fmt.Printf("%#v\n", Find("baby_"))
}

func ExampleEmojize() {
	fmt.Println(Emojize("Hi :smile:!"))
	// Output: Hi 😄!
}
