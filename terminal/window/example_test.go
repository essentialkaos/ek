package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
