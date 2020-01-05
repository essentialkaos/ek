package emoji

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
