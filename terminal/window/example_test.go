package window

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

func ExampleGetSize() {
	width, height := GetSize()

	if width == -1 && height == -1 {
		fmt.Println("Can't detect window size")
		return
	}

	fmt.Printf("Window size: %d x %d\n", width, height)
}

func ExampleGetWidth() {
	width := GetWidth()

	if width == -1 {
		fmt.Println("Can't detect window size")
		return
	}

	fmt.Printf("Window width: %d\n", width)
}

func ExampleGetHeight() {
	height := GetHeight()

	if height == -1 {
		fmt.Println("Can't detect window size")
		return
	}

	fmt.Printf("Window height: %d\n", height)
}
