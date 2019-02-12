package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRGB2Hex() {
	fmt.Printf("%x\n", RGB2Hex(127, 25, 75))
	// Output: 7f194b
}

func ExampleHex2RGB() {
	r, g, b := Hex2RGB(0X7F194B)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:127 g:25 b:75
}

func ExampleHex2RGBA() {
	r, g, b, a := Hex2RGBA(0X7F194BCC)

	fmt.Printf("r:%d g:%d b:%d a:%d\n", r, g, b, a)

	// Output: r:127 g:25 b:75 a:204
}

func ExampleRGB2HSB() {
	h, s, v := RGB2HSB(127, 25, 75)

	fmt.Printf("h:%d s:%d v:%d\n", h, s, v)

	// Output: h:331 s:81 v:50
}

func ExampleHSB2RGB() {
	r, g, b := HSB2RGB(331, 81, 50)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:128 g:24 b:74
}

func ExampleIsRGBA() {
	fmt.Println(IsRGBA(0xAABBCC))
	fmt.Println(IsRGBA(0xAABBCCDD))

	// Output:
	// false
	// true
}
