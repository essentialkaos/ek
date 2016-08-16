package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRGB2Hex() {
	fmt.Printf("%x\n", color.RGB2Hex(127, 25, 75))
	// Output: 7f194b
}

func ExampleHex2RGB() {
	r, g, b := color.Hex2RGB(0x7f194b)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:127 g:25 b:75
}

func ExampleHex2RGBA() {
	r, g, b, a := color.Hex2RGBA(0x7f194bcc)

	fmt.Printf("r:%d g:%d b:%d a:%d\n", r, g, b, a)

	// Output: r:127 g:25 b:75 a:204
}

func ExampleRGB2HSB() {
	h, s, v := color.RGB2HSB(127, 25, 75)

	fmt.Printf("h:%d s:%d v:%d\n", h, s, v)

	// Output: h:331 s:81 v:50
}

func ExampleHSB2RGB() {
	r, g, b := color.HSB2RGB(331, 81, 50)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:128 g:24 b:74
}

func ExampleIsRGBA() {
	fmt.Println(color.IsRGBA(0xAABBCC))
	fmt.Println(color.IsRGBA(0xAABBCCDD))

	// Output:
	// false
	// true
}
