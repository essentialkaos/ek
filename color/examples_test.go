package color

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

func ExampleParse() {
	fmt.Println(Parse("#ff6347"))
	fmt.Println(Parse("#B8F"))
	fmt.Println(Parse("#FF3B21A6"))
	fmt.Println(Parse("mintcream"))

	// Output:
	// Hex{#FF6347} <nil>
	// Hex{#BB88FF} <nil>
	// Hex{#FF3B21A6} <nil>
	// Hex{#F5FFFA} <nil>
}

func ExampleRGB2Hex() {
	fmt.Printf("%s\n", RGB2Hex(RGB{127, 25, 75}).ToWeb(true))

	// Output: #7F194B
}

func ExampleHex2RGB() {
	fmt.Printf("%s\n", Hex2RGB(0x7F194B))

	// Output: RGB{R:127 G:25 B:75}
}

func ExampleHex2RGBA() {
	fmt.Printf("%s\n", Hex2RGBA(0x7F194BCC))

	// Output: RGBA{R:127 G:25 B:75 A:0.80}
}

func ExampleRGBA2Hex() {
	c := RGBA2Hex(RGBA{127, 25, 75, 204})

	fmt.Printf("%s\n", c.ToWeb(true))

	// Output: #7F194BCC
}

func ExampleRGB2Term() {
	c := RGB{255, 0, 0}

	fmt.Printf("%s → \\e[38;5;%dm\n", c, RGB2Term(c))

	// Output: RGB{R:255 G:0 B:0} → \e[38;5;196m
}

func ExampleRGB2CMYK() {
	fmt.Printf("%s\n", RGB2CMYK(RGB{127, 25, 75}))

	// Output: CMYK{C:0% M:80% Y:41% K:50%}
}

func ExampleCMYK2RGB() {
	fmt.Printf("%s\n", CMYK2RGB(CMYK{0, 0.8, 0.41, 0.5}))

	// Output: RGB{R:127 G:25 B:75}
}

func ExampleRGB2HSV() {
	fmt.Printf("%s\n", RGB2HSV(RGB{127, 25, 75}))

	// Output: HSV{H:331° S:80% V:50%}
}

func ExampleHSV2RGB() {
	c := HSV2RGB(HSV{331.0 / 360.0, 80.0 / 100.0, 50.0 / 100.0})

	fmt.Printf("%s\n", c)

	// Output: RGB{R:127 G:25 B:74}
}

func ExampleRGB2HSL() {
	fmt.Printf("%s\n", RGB2HSL(RGB{127, 25, 75}))

	// Output: HSL{H:331° S:67% L:30%}
}

func ExampleHSL2RGB() {
	c := HSL2RGB(HSL{331.0 / 360.0, 67.0 / 100.0, 30.0 / 100.0})

	fmt.Printf("%s\n", c)

	// Output: RGB{R:127 G:25 B:74}
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
	c := Contrast(0x222222, 0x40abf7)

	fmt.Printf("ratio: %.2f:1\n", c)

	// Output: ratio: 6.35:1
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRGB_ToHex() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHex().ToWeb(true))

	// Output: #7F194B
}

func ExampleRGB_ToCMYK() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToCMYK())

	// Output: CMYK{C:0% M:80% Y:41% K:50%}
}

func ExampleRGB_ToHSV() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHSV())

	// Output: HSV{H:331° S:80% V:50%}
}

func ExampleRGB_ToHSL() {
	fmt.Printf("%s\n", RGB{127, 25, 75}.ToHSL())

	// Output: HSL{H:331° S:67% L:30%}
}

func ExampleRGB_ToTerm() {
	c := RGB{255, 0, 0}

	fmt.Printf("%s → \\e[38;5;%dm\n", c, c.ToTerm())

	// Output: RGB{R:255 G:0 B:0} → \e[38;5;196m
}

func ExampleRGBA_ToHex() {
	c := RGBA{127, 25, 75, 204}.ToHex()

	fmt.Printf("%s\n", c.ToWeb(true))

	// Output: #7F194BCC
}

func ExampleCMYK_ToRGB() {
	fmt.Printf("%s\n", CMYK{0, 0.8, 0.41, 0.5}.ToRGB())

	// Output: RGB{R:127 G:25 B:75}
}

func ExampleHSV_ToRGB() {
	c := HSV{331.0 / 360.0, 80.0 / 100.0, 50.0 / 100.0}.ToRGB()

	fmt.Printf("%s\n", c)

	// Output: RGB{R:127 G:25 B:74}
}

func ExampleHSL_ToRGB() {
	c := HSL{331.0 / 360.0, 67.0 / 100.0, 30.0 / 100.0}.ToRGB()

	fmt.Printf("%s\n", c)

	// Output: RGB{R:127 G:25 B:74}
}

func ExampleHex_IsRGBA() {
	c1 := Hex(0x7F194B)
	c2 := Hex(0x7F194B5F)

	fmt.Printf("%s → %t\n", c1.ToWeb(true), c1.IsRGBA())
	fmt.Printf("%s → %t\n", c2.ToWeb(true), c2.IsRGBA())

	// Output:
	// #7F194B → false
	// #7F194B5F → true
}

func ExampleHex_ToRGB() {
	fmt.Printf("%s\n", Hex(0x7F194B).ToRGB())

	// Output: RGB{R:127 G:25 B:75}
}

func ExampleHex_ToRGBA() {
	fmt.Printf("%s\n", Hex(0x7F194B5F).ToRGBA())

	// Output: RGBA{R:127 G:25 B:75 A:0.37}
}

func ExampleHex_ToWeb() {
	fmt.Printf("%s\n", Hex(0x7F194B).ToWeb(true))
	fmt.Printf("%s\n", Hex(0x7F194B).ToWeb(false))

	// Output:
	// #7F194B
	// #7f194b
}
