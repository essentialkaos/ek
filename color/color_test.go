//go:build 386 || amd64
// +build 386 amd64

package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

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
	c.Assert(color, Equals, Hex(0xEE82EE))

	color, err = Parse("#ff6347")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xFF6347))

	color, err = Parse("#FF6347")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xFF6347))

	color, err = Parse("FF6347")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xFF6347))

	color, err = Parse("#b3f")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xBB33FF))

	color, err = Parse("#FF6347AA")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xFF6347AA))

	color, err = Parse("#b3fa")
	c.Assert(err, IsNil)
	c.Assert(color, Equals, Hex(0xBB33FFAA))

	_, err = Parse("")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, "Color is empty")

	_, err = Parse("TEST")
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, "strconv.ParseUint: parsing \"TTEESSTT\": invalid syntax")
}

func (s *ColorSuite) TestRGB(c *C) {
	c.Assert(RGB{0, 0, 0}.ToHex(), DeepEquals, Hex(0x000000))
	c.Assert(RGB{255, 255, 255}.ToHex(), DeepEquals, Hex(0xFFFFFF))
	c.Assert(RGB{127, 127, 127}.ToHex(), DeepEquals, Hex(0x7F7F7F))

	c.Assert(RGB{23, 182, 89}.String(), Equals, "RGB{R:23 G:182 B:89}")
}

func (s *ColorSuite) TestHex(c *C) {
	c.Assert(Hex(0x000000).ToRGB(), DeepEquals, RGB{0x00, 0x00, 0x00})
	c.Assert(Hex(0xFFFFFF).ToRGB(), DeepEquals, RGB{0xFF, 0xFF, 0xFF})
	c.Assert(Hex(0x49d62d).ToRGB(), DeepEquals, RGB{0x49, 0xd6, 0x2d})

	c.Assert(Hex(0x49d62d).String(), Equals, "Hex{#49D62D}")
	c.Assert(Hex(0x49d62df7).String(), Equals, "Hex{#49D62DF7}")
	c.Assert(Hex(0x49d62d).ToWeb(false), Equals, "#49d62d")
	c.Assert(Hex(0x49d62df7).ToWeb(true), Equals, "#49D62DF7")
	c.Assert(Hex(0xFFAA44).ToWeb(true), Equals, "#FA4")
	c.Assert(Hex(0x0).ToWeb(true), Equals, "#000")
}

func (s *ColorSuite) TestRGBA(c *C) {
	c.Assert(RGBA{0, 0, 0, 0}.ToHex(), DeepEquals, Hex(0x00000000))
	c.Assert(RGBA{255, 255, 255, 0}.ToHex(), DeepEquals, Hex(0xFFFFFF00))
	c.Assert(RGBA{255, 255, 255, 255}.ToHex(), DeepEquals, Hex(0xFFFFFFFF))
	c.Assert(RGBA{127, 127, 127, 127}.ToHex(), DeepEquals, Hex(0x7F7F7F7F))

	c.Assert(Hex(0x00000000).ToRGBA(), DeepEquals, RGBA{0, 0, 0, 0})
	c.Assert(Hex(0xFFFFFFFF).ToRGBA(), DeepEquals, RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	c.Assert(Hex(0x7F7F7F7F).ToRGBA(), DeepEquals, RGBA{0x7F, 0x7F, 0x7F, 0x7F})

	c.Assert(Hex(0xFFAABB).IsRGBA(), Equals, false)
	c.Assert(Hex(0xFFAABB01).IsRGBA(), Equals, true)

	c.Assert(RGBA{23, 182, 89, 130}.String(), Equals, "RGBA{R:23 G:182 B:89 A:0.51}")
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
	c.Assert(RGB{0, 0, 0}.ToHSV(), DeepEquals, HSV{0.0, 0.0, 0.0})
	c.Assert(RGB{255, 255, 255}.ToHSV(), DeepEquals, HSV{0.0, 0.0, 1.0})
	c.Assert(RGB{255, 0, 0}.ToHSV(), DeepEquals, HSV{0.0, 1.0, 1.0})
	c.Assert(RGB{0, 255, 0}.ToHSV(), DeepEquals, HSV{1.0 / 3.0, 1.0, 1.0})
	c.Assert(RGB{0, 0, 255}.ToHSV(), DeepEquals, HSV{2.0 / 3.0, 1.0, 1.0})

	c.Assert(HSV{0.0, 0.0, 0.0}.ToRGB(), DeepEquals, RGB{0, 0, 0})
	c.Assert(HSV{0.0, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})

	c.Assert(HSV{0.17, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.33, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.49, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.65, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.81, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})
	c.Assert(HSV{0.97, 0.0, 1.0}.ToRGB(), DeepEquals, RGB{255, 255, 255})

	c.Assert(RGB{73, 158, 105}.ToHSV().String(), Equals, "HSV{H:143° S:54% V:62%}")
}

func (s *ColorSuite) TestHSL(c *C) {
	c.Assert(RGB{255, 0, 0}.ToHSL(), DeepEquals, HSL{0.0, 1.0, 0.5})
	c.Assert(RGB{0, 0, 0}.ToHSL(), DeepEquals, HSL{0.0, 0.0, 0.0})
	c.Assert(RGB{63, 191, 191}.ToHSL(), DeepEquals, HSL{0.5, 0.5039370078740157, 0.4980392156862745})
	c.Assert(RGB{255, 201, 201}.ToHSL(), DeepEquals, HSL{0.0, 1.0, 0.8941176470588235})
	c.Assert(RGB{255, 130, 245}.ToHSL(), DeepEquals, HSL{0.8466666666666667, 1.0, 0.7549019607843137})
	c.Assert(RGB{46, 196, 255}.ToHSL(), DeepEquals, HSL{0.5470494417862839, 1.0, 0.5901960784313726})

	c.Assert(HSL{0.0, 0.0, 0.0}.ToRGB(), DeepEquals, RGB{0, 0, 0})
	c.Assert(HSL{0.0, 1.0, 0.5}.ToRGB(), DeepEquals, RGB{255, 0, 0})
	c.Assert(HSL{0.6833333333333332, 0.7079646017699115, 0.5568627450980392}.ToRGB(), DeepEquals, RGB{77, 62, 222})

	c.Assert(RGB{146, 93, 176}.ToHSL().String(), Equals, "HSL{H:278° S:34% L:53%}")
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
	// grayscale
	c.Assert(RGB{175, 175, 175}.ToTerm(), Equals, 145)
	c.Assert(RGB{18, 18, 18}.ToTerm(), Equals, 233)
	c.Assert(RGB{48, 48, 48}.ToTerm(), Equals, 236)
	c.Assert(RGB{238, 238, 238}.ToTerm(), Equals, 255)
}

func (s *ColorSuite) TestLuminance(c *C) {
	c.Assert(Luminance(Hex(0xffffff).ToRGB()), Equals, 1.0)
	c.Assert(Luminance(Hex(0x000000).ToRGB()), Equals, 0.0)
	c.Assert(Luminance(Hex(0x808080).ToRGB()), Equals, 0.2158605001138992)
	c.Assert(Luminance(Hex(0x2861bd).ToRGB()), Equals, 0.12674627666892935)
}

func (s *ColorSuite) TestContrast(c *C) {
	c.Assert(Contrast(Hex(0xffffff), Hex(0x000000)), Equals, 21.0)
	c.Assert(Contrast(Hex(0x000000), Hex(0x000000)), Equals, 1.0)
	c.Assert(Contrast(Hex(0x755757), Hex(0x63547c)), Equals, 1.0542371982635754)
	c.Assert(Contrast(Hex(0x333333), Hex(0xd0b5ff)), Equals, 7.0606832983463805)
}
