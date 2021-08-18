package color

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

func ExampleRGB2CMYK() {
	c, m, y, k := RGB2CMYK(68, 133, 148)

	fmt.Printf(
		"C:%.0f%% M:%.0f%% Y:%.0f%% K:%.0f%%\n",
		c*100.0, m*100.0, y*100.0, k*100.0,
	)

	// Output: C:54% M:10% Y:0% K:42%
}

func ExampleCMYK2RGB() {
	r, g, b := CMYK2RGB(0.0, 0.25, 0.77, 0.17)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:211 g:158 b:48
}

func ExampleRGB2HSL() {
	h, s, l := RGB2HSL(111, 128, 173)

	fmt.Printf(
		"h:%.0f° s:%.0f%% l:%.0f%%\n",
		h*360, s*100, l*100,
	)

	// Output: h:224° s:27% l:56%
}

func ExampleHSL2RGB() {
	r, g, b := HSL2RGB(0.6209677419354839, 0.27433628318584075, 0.5568627450980392)

	fmt.Printf("r:%d g:%d b:%d\n", r, g, b)

	// Output: r:111 g:128 b:173
}

func ExampleHUE2RGB() {
	hue := HUE2RGB(0.3, 0.12, 0.56)

	fmt.Printf("hue:%.4f\n", hue)

	// Output: hue:0.1848
}

func ExampleRGBLuminance() {
	l := RGBLuminance(135, 85, 189)

	fmt.Printf("lum: %.7f\n", l)

	// Output: lum: 0.1532202
}

func ExampleHEXLuminance() {
	l := HEXLuminance(0x8755bd)

	fmt.Printf("lum: %.7f\n", l)

	// Output: lum: 0.1532202
}

func ExampleContrast() {
	c := Contrast(0x222222, 0x40abf7)

	fmt.Printf("ratio: %.2f:1\n", c)

	// Output: ratio: 6.35:1
}

func ExampleIsRGBA() {
	fmt.Println(IsRGBA(0xAABBCC))
	fmt.Println(IsRGBA(0xAABBCCDD))

	// Output:
	// false
	// true
}
