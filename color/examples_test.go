package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleParse() {
	fmt.Println(Parse("#ff6347"))
	fmt.Println(Parse("#B8F"))
	fmt.Println(Parse("#FF3B21A6"))
	fmt.Println(Parse("mintcream"))

	// Output:
	// #FF6347 <nil>
	// #BB88FF <nil>
	// #FF3B21A6 <nil>
	// #F5FFFA <nil>
}

func ExampleRGB2Hex() {
	fmt.Printf("%s\n", RGB2Hex(RGB{127, 25, 75}).ToWeb(true, false))

	// Output: #7F194B
}

func ExampleHex2RGB() {
	fmt.Printf("%#v\n", Hex2RGB(NewHex(0x7F194B)))

	// Output: RGB{R:127, G:25, B:75}
}

func ExampleHex2RGBA() {
	fmt.Printf("%#v\n", Hex2RGBA(NewHex(0x7F194BCC)))

	// Output: RGBA{R:127, G:25, B:75, A:204}
}

func ExampleRGBA2Hex() {
	c := RGBA2Hex(RGBA{127, 25, 75, 204})

	fmt.Printf("%s\n", c.ToWeb(true, false))

	// Output: #7F194BCC
}

func ExampleRGB2Term() {
	c := RGB{255, 0, 0}

	fmt.Printf("%#v → \\e[38;5;%dm\n", c, RGB2Term(c))

	// Output: RGB{R:255, G:0, B:0} → \e[38;5;196m
}

func ExampleTerm2RGB() {
	c := uint8(162)

	fmt.Printf("%d → %#v\n", c, Term2RGB(c))

	// Output: 162 → RGB{R:215, G:0, B:135}
}

func ExampleRGB2CMYK() {
	fmt.Printf("%s\n", RGB2CMYK(RGB{127, 25, 75}))

	// Output: 0%,80%,41%,50%
}

func ExampleCMYK2RGB() {
	fmt.Printf("%s\n", CMYK2RGB(CMYK{0, 0.8, 0.41, 0.5}))

	// Output: 127,25,75
}

func ExampleRGB2HSV() {
	fmt.Printf("%s\n", RGB2HSV(RGB{127, 25, 75}))

	// Output: 331°,80%,50%,0%
}

func ExampleHSV2RGB() {
	c := HSV2RGB(HSV{H: 331.0 / 360.0, S: 80.0 / 100.0, V: 50.0 / 100.0})

	fmt.Printf("%s\n", c)

	// Output: 127,25,74
}

func ExampleRGB2HSL() {
	fmt.Printf("%s\n", RGB2HSL(RGB{127, 25, 75}))

	// Output: 331°,67%,30%,0%
}

func ExampleHSL2RGB() {
	c := HSL2RGB(HSL{H: 331.0 / 360.0, S: 67.0 / 100.0, L: 30.0 / 100.0})

	fmt.Printf("%s\n", c)

	// Output: 127,25,74
}

func ExampleHUE2RGB() {
	hue := HUE2RGB(0.3, 0.12, 0.56)

	fmt.Printf("Hue: %.4f\n", hue)

	// Output: Hue: 0.1848
}

func ExampleLuminance() {
	fmt.Printf("Lum: %.7f\n", Luminance(RGB{135, 85, 189}))

	// Output: Lum: 0.1532202
}

func ExampleContrast() {
	c := Contrast(NewHex(0x222222), NewHex(0x40abf7))

	fmt.Printf("ratio: %.2f:1\n", c)

	// Output: ratio: 6.35:1
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRGB_ToHex() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHex().ToWeb(true, false))

	// Output: #7F194B
}

func ExampleRGB_ToCMYK() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToCMYK())

	// Output: 0%,80%,41%,50%
}

func ExampleRGB_ToHSV() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHSV())

	// Output: 331°,80%,50%,0%
}

func ExampleRGB_ToHSL() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHSL())

	// Output: 331°,67%,30%,0%
}

func ExampleRGB_ToTerm() {
	c := RGB{255, 0, 0}

	fmt.Printf("%s → \\e[38;5;%dm\n", c, c.ToTerm())

	// Output: 255,0,0 → \e[38;5;196m
}

func ExampleRGBA_ToHex() {
	c := RGBA{127, 25, 75, 204}.ToHex()

	fmt.Printf("%s\n", c.ToWeb(true, false))

	// Output: #7F194BCC
}

func ExampleCMYK_ToRGB() {
	fmt.Printf("%s\n", CMYK{0, 0.8, 0.41, 0.5}.ToRGB())

	// Output: 127,25,75
}

func ExampleHSV_ToRGB() {
	c := HSV{H: 331.0 / 360.0, S: 80.0 / 100.0, V: 50.0 / 100.0}.ToRGB()

	fmt.Printf("%s\n", c)

	// Output: 127,25,74
}

func ExampleHSL_ToRGB() {
	c := HSL{H: 331.0 / 360.0, S: 67.0 / 100.0, L: 30.0 / 100.0}.ToRGB()

	fmt.Printf("%s\n", c)

	// Output: 127,25,74
}

func ExampleHex_IsRGBA() {
	c1 := NewHex(0x7F194B)
	c2 := NewHex(0x7F194B5F)

	fmt.Printf("%s → %t\n", c1.ToWeb(true, false), c1.IsRGBA())
	fmt.Printf("%s → %t\n", c2.ToWeb(true, false), c2.IsRGBA())

	// Output:
	// #7F194B → false
	// #7F194B5F → true
}

func ExampleHex_ToRGB() {
	fmt.Printf("%s\n", NewHex(0x7F194B).ToRGB())

	// Output: 127,25,75
}

func ExampleHex_ToRGBA() {
	fmt.Printf("%s\n", NewHex(0x7F194B5F).ToRGBA())

	// Output: 127,25,75,0.37
}

func ExampleHex_ToWeb() {
	fmt.Printf("%s\n", NewHex(0xFFAA44CC).ToWeb(true, false))
	fmt.Printf("%s\n", NewHex(0xFFAA44CC).ToWeb(false, true))

	// Output:
	// #FFAA44CC
	// #fa4c
}
