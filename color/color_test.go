package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"github.com/essentialkaos/ek/v13/mathutil"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ColorSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ColorSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ColorSuite) TestParse(c *C) {
	color, err := Parse("violet")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xEE82EE))

	color, err = Parse("#ff6347")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xFF6347))

	color, err = Parse("#FF6347")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xFF6347))

	color, err = Parse("FF6347")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xFF6347))

	color, err = Parse("#b3f")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xBB33FF))

	color, err = Parse("#FF6347AA")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xFF6347AA))

	color, err = Parse("#b3fa")
	c.Assert(err, IsNil)
	c.Assert(color.v, Equals, uint32(0xBB33FFAA))

	_, err = Parse("")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, "Color is empty")

	_, err = Parse("TEST")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, "strconv.ParseUint: parsing \"TTEESSTT\": invalid syntax")
}

func (s *ColorSuite) TestRGB(c *C) {
	c.Assert(RGB{0, 0, 0}.ToHex().v, DeepEquals, uint32(0x000000))
	c.Assert(RGB{255, 255, 255}.ToHex().v, DeepEquals, uint32(0xFFFFFF))
	c.Assert(RGB{127, 127, 127}.ToHex().v, DeepEquals, uint32(0x7F7F7F))

	c.Assert(RGB{23, 182, 89}.String(), Equals, "RGB{R:23 G:182 B:89}")
}

func (s *ColorSuite) TestHex(c *C) {
	c.Assert(NewHex(0x000000).ToRGB(), DeepEquals, RGB{0x00, 0x00, 0x00})
	c.Assert(NewHex(0xFFFFFF).ToRGB(), DeepEquals, RGB{0xFF, 0xFF, 0xFF})
	c.Assert(NewHex(0x49d62d).ToRGB(), DeepEquals, RGB{0x49, 0xd6, 0x2d})

	c.Assert(NewHex(0x49d62d).String(), Equals, "Hex{#49D62D}")
	c.Assert(NewHex(0x49d62df7).String(), Equals, "Hex{#49D62DF7}")
	c.Assert(NewHex(0x49d62d).ToWeb(false, false), Equals, "#49d62d")
	c.Assert(NewHex(0x49d62df7).ToWeb(true, false), Equals, "#49D62DF7")
	c.Assert(NewHex(0x49d62df7).ToWeb(false, false), Equals, "#49d62df7")
	c.Assert(NewHex(0xFFAA44).ToWeb(true, false), Equals, "#FFAA44")
	c.Assert(NewHex(0xFFAA44).ToWeb(true, true), Equals, "#FA4")
	c.Assert(NewHex(0xFFAA44CC).ToWeb(true, false), Equals, "#FFAA44CC")
	c.Assert(NewHex(0xFFAA44CC).ToWeb(true, true), Equals, "#FA4C")
	c.Assert(NewHex(0x0).ToWeb(true, false), Equals, "#000000")
	c.Assert(NewHex(0x0).ToWeb(true, true), Equals, "#000")
}

func (s *ColorSuite) TestRGBA(c *C) {
	c.Assert(RGBA{0, 0, 0, 0}.ToHex().v, DeepEquals, uint32(0x00000000))
	c.Assert(RGBA{255, 255, 255, 0}.ToHex().v, DeepEquals, uint32(0xFFFFFF00))
	c.Assert(RGBA{255, 255, 255, 255}.ToHex().v, DeepEquals, uint32(0xFFFFFFFF))
	c.Assert(RGBA{127, 127, 127, 127}.ToHex().v, DeepEquals, uint32(0x7F7F7F7F))

	c.Assert(NewHex(0x00000000).ToRGBA(), DeepEquals, RGBA{0, 0, 0, 0})
	c.Assert(NewHex(0xFFFFFFFF).ToRGBA(), DeepEquals, RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	c.Assert(NewHex(0x7F7F7F7F).ToRGBA(), DeepEquals, RGBA{0x7F, 0x7F, 0x7F, 0x7F})

	c.Assert(NewHex(0xFFAABB).IsRGBA(), Equals, false)
	c.Assert(NewHex(0xFFAABB01).IsRGBA(), Equals, true)

	c.Assert(RGBA{23, 182, 89, 130}.String(), Equals, "RGBA{R:23 G:182 B:89 A:0.51}")

	clr := RGBA{23, 182, 89, 0}
	c.Assert(clr.String(), Equals, "RGBA{R:23 G:182 B:89 A:0.00}")
	c.Assert(clr.WithAlpha(0.23).String(), Equals, "RGBA{R:23 G:182 B:89 A:0.23}")
}

func (s *ColorSuite) TestCMYK(c *C) {
	c.Assert(RGB{0, 0, 0}.ToCMYK(), DeepEquals, CMYK{0.0, 0.0, 0.0, 1.0})
	c.Assert(RGB{255, 255, 255}.ToCMYK(), DeepEquals, CMYK{0.0, 0.0, 0.0, 0.0})
	c.Assert(RGB{102, 102, 102}.ToCMYK(), DeepEquals, CMYK{0.0, 0.0, 0.0, 0.6})
	c.Assert(RGB{102, 0, 0}.ToCMYK(), DeepEquals, CMYK{0.0, 1.0, 1.0, 0.6})

	c.Assert(CMYK{0.0, 0.0, 0.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(CMYK{0.0, 0.0, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{0, 0, 0})
	c.Assert(CMYK{0.64, 0.77, 0, 0.17}.ToRGB(), DeepEquals, RGB{76, 48, 211})

	c.Assert(CMYK{0.64, 0.77, 0, 0.17}.String(), Equals, "CMYK{C:64% M:77% Y:0% K:17%}")
}

func (s *ColorSuite) TestHSV(c *C) {
	c.Assert(RGB{0, 0, 0}.ToHSV(), DeepEquals, HSV{0.0, 0.0, 0.0, 0.0})
	c.Assert(RGB{255, 255, 255}.ToHSV(), DeepEquals, HSV{0.0, 0.0, 1.0, 0.0})
	c.Assert(RGB{255, 0, 0}.ToHSV(), DeepEquals, HSV{0.0, 1.0, 1.0, 0.0})
	c.Assert(RGB{0, 255, 0}.ToHSV(), DeepEquals, HSV{1.0 / 3.0, 1.0, 1.0, 0.0})
	c.Assert(RGB{0, 0, 255}.ToHSV(), DeepEquals, HSV{2.0 / 3.0, 1.0, 1.0, 0.0})

	c.Assert(HSV{0.0, 0.0, 0.0, 0.0}.ToRGB(), DeepEquals, RGB{0, 0, 0})
	c.Assert(HSV{0.0, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})

	c.Assert(HSV{0.17, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.33, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.49, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.65, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.81, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.97, 0.0, 1.0, 0.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})

	c.Assert(HSV{0.32, 0.12, 0.76, 0.5}.ToRGBA(), DeepEquals, RGBA{172, 193, 170, 127})

	c.Assert(RGB{73, 158, 105}.ToHSV().String(), Equals, "HSV{H:143° S:54% V:62% A:0%}")
	c.Assert(RGBA{73, 158, 105, 127}.ToHSV().String(), Equals, "HSV{H:143° S:54% V:62% A:50%}")
}

func (s *ColorSuite) TestHSL(c *C) {
	c.Assert(RGB{255, 0, 0}.ToHSL(), DeepEquals, HSL{0.0, 1.0, 0.5, 0.0})
	c.Assert(RGB{0, 0, 0}.ToHSL(), DeepEquals, HSL{0.0, 0.0, 0.0, 0.0})
	c.Assert(RGB{63, 191, 191}.ToHSL(), DeepEquals, HSL{0.5, 0.5039370078740157, 0.4980392156862745, 0.0})
	c.Assert(RGB{255, 201, 201}.ToHSL(), DeepEquals, HSL{0.0, 1.0, 0.8941176470588235, 0.0})
	c.Assert(RGB{255, 130, 245}.ToHSL(), DeepEquals, HSL{0.8466666666666667, 1.0, 0.7549019607843137, 0.0})
	c.Assert(RGB{46, 196, 255}.ToHSL(), DeepEquals, HSL{0.5470494417862839, 1.0, 0.5901960784313726, 0.0})

	c.Assert(HSL{0.0, 0.0, 0.0, 0.0}.ToRGB(), DeepEquals, RGB{0, 0, 0})
	c.Assert(HSL{0.0, 1.0, 0.5, 0.0}.ToRGB(), DeepEquals, RGB{255, 0, 0})
	c.Assert(HSL{0.6833333333333332, 0.7079646017699115, 0.5568627450980392, 0.0}.ToRGB(), DeepEquals, RGB{77, 62, 222})
	c.Assert(HSL{0.6833333333333332, 0.7079646017699115, 0.5568627450980392, 0.50}.ToRGBA(), DeepEquals, RGBA{77, 62, 222, 127})

	c.Assert(RGB{146, 93, 176}.ToHSL().String(), Equals, "HSL{H:278° S:34% L:53% A:0%}")
	c.Assert(RGBA{146, 93, 176, 80}.ToHSL().String(), Equals, "HSL{H:278° S:34% L:53% A:31%}")
}

func (s *ColorSuite) TestHUE(c *C) {
	c.Assert(HUE2RGB(0.6, 0.3, 0.12), Equals, 0.384)
	c.Assert(HUE2RGB(0.6, 0.3, 0.35), Equals, 0.3)
	c.Assert(HUE2RGB(0.6, 0.3, 0.584), Equals, 0.4512)
	c.Assert(HUE2RGB(0.6, 0.3, 0.85), Equals, 0.6)
}

func (s *ColorSuite) TestTerm(c *C) {
	c.Assert(RGB{0, 175, 175}.ToTerm(), Equals, 37)
	c.Assert(RGB{255, 255, 0}.ToTerm(), Equals, 226)
	c.Assert(RGB{135, 175, 215}.ToTerm(), Equals, 110)
	c.Assert(RGB{175, 175, 175}.ToTerm(), Equals, 145)
	c.Assert(RGB{18, 18, 18}.ToTerm(), Equals, 233)
	c.Assert(RGB{48, 48, 48}.ToTerm(), Equals, 236)
	c.Assert(RGB{238, 238, 238}.ToTerm(), Equals, 255)

	c.Assert(Term2RGB(0), DeepEquals, RGB{0, 0, 0})
	c.Assert(Term2RGB(1), DeepEquals, RGB{255, 0, 0})
	c.Assert(Term2RGB(2), DeepEquals, RGB{0, 255, 0})
	c.Assert(Term2RGB(3), DeepEquals, RGB{255, 255, 0})
	c.Assert(Term2RGB(4), DeepEquals, RGB{0, 0, 255})
	c.Assert(Term2RGB(5), DeepEquals, RGB{255, 0, 255})
	c.Assert(Term2RGB(6), DeepEquals, RGB{0, 255, 255})
	c.Assert(Term2RGB(7), DeepEquals, RGB{191, 191, 191})
	c.Assert(Term2RGB(8), DeepEquals, RGB{64, 64, 64})
	c.Assert(Term2RGB(9), DeepEquals, RGB{255, 127, 127})
	c.Assert(Term2RGB(10), DeepEquals, RGB{127, 255, 127})
	c.Assert(Term2RGB(11), DeepEquals, RGB{255, 255, 127})
	c.Assert(Term2RGB(12), DeepEquals, RGB{127, 127, 255})
	c.Assert(Term2RGB(13), DeepEquals, RGB{255, 127, 255})
	c.Assert(Term2RGB(14), DeepEquals, RGB{127, 255, 255})
	c.Assert(Term2RGB(15), DeepEquals, RGB{127, 127, 127})

	c.Assert(Term2RGB(238), DeepEquals, RGB{68, 68, 68})
	c.Assert(Term2RGB(153), DeepEquals, RGB{175, 215, 255})

}

func (s *ColorSuite) TestLuminance(c *C) {
	c.Assert(Luminance(NewHex(0xffffff).ToRGB()), Equals, 1.0)
	c.Assert(Luminance(NewHex(0x000000).ToRGB()), Equals, 0.0)
	c.Assert(mathutil.Round(Luminance(NewHex(0x808080).ToRGB()), 8), Equals, 0.2158605)
	c.Assert(mathutil.Round(Luminance(NewHex(0x2861bd).ToRGB()), 8), Equals, 0.12674628)
}

func (s *ColorSuite) TestContrast(c *C) {
	c.Assert(Contrast(NewHex(0xffffff), NewHex(0x000000)), Equals, 21.0)
	c.Assert(Contrast(NewHex(0x000000), NewHex(0x000000)), Equals, 1.0)
	c.Assert(mathutil.Round(Contrast(NewHex(0x755757), NewHex(0x63547c)), 8), Equals, 1.0542372)
	c.Assert(mathutil.Round(Contrast(NewHex(0x333333), NewHex(0xd0b5ff)), 8), Equals, 7.0606833)
}
