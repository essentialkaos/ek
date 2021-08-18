// +build 386 amd64

package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ColorSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ColorSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ColorSuite) TestRGB2Hex(c *C) {
	c.Assert(RGB2Hex(0, 0, 0), Equals, 0x000000)
	c.Assert(RGB2Hex(255, 255, 255), Equals, 0xFFFFFF)
	c.Assert(RGB2Hex(127, 127, 127), Equals, 0x7F7F7F)
}

func (s *ColorSuite) TestHex2RGB(c *C) {
	var r, g, b int

	r, g, b = Hex2RGB(0x000000)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0, 0, 0})

	r, g, b = Hex2RGB(0xFFFFFF)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0xFF, 0xFF, 0xFF})

	r, g, b = Hex2RGB(0x7F7F7F)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0x7F, 0x7F, 0x7F})
}

func (s *ColorSuite) TestRGBA2Hex(c *C) {
	c.Assert(RGBA2Hex(0, 0, 0, 0), Equals, int64(0x00000000))
	c.Assert(RGBA2Hex(255, 255, 255, 0), Equals, int64(0xFFFFFF00))
	c.Assert(RGBA2Hex(255, 255, 255, 255), Equals, int64(0xFFFFFFFF))
	c.Assert(RGBA2Hex(127, 127, 127, 127), Equals, int64(0x7F7F7F7F))
}

func (s *ColorSuite) TestHex2RGBA(c *C) {
	var r, g, b, a int

	r, g, b, a = Hex2RGBA(0x00000000)
	c.Assert([]int{r, g, b, a}, DeepEquals, []int{0, 0, 0, 0})

	r, g, b, a = Hex2RGBA(0xFFFFFFFF)
	c.Assert([]int{r, g, b, a}, DeepEquals, []int{0xFF, 0xFF, 0xFF, 0xFF})

	r, g, b, a = Hex2RGBA(0x7F7F7F7F)
	c.Assert([]int{r, g, b, a}, DeepEquals, []int{0x7F, 0x7F, 0x7F, 0x7F})
}

func (s *ColorSuite) TestHSB(c *C) {
	h, sv, b := RGB2HSB(0, 0, 0)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{0, 0, 0})

	h, sv, b = RGB2HSB(255, 255, 255)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{0, 0, 100})

	h, sv, b = RGB2HSB(255, 0, 0)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{0, 100, 100})

	h, sv, b = RGB2HSB(0, 255, 0)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{120, 100, 100})

	h, sv, b = RGB2HSB(0, 0, 255)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{240, 100, 100})

	h, sv, b = RGB2HSB(127, 127, 127)
	c.Assert([]int{h, sv, b}, DeepEquals, []int{0, 0, 50})

	r, g, b := HSB2RGB(0, 0, 0)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0, 0, 0})

	r, g, b = HSB2RGB(360, 100, 100)
	c.Assert([]int{r, g, b}, DeepEquals, []int{255, 0, 0})

	r, g, b = HSB2RGB(0, 100, 100)
	c.Assert([]int{r, g, b}, DeepEquals, []int{255, 0, 0})

	r, g, b = HSB2RGB(65, 50, 50)
	c.Assert([]int{r, g, b}, DeepEquals, []int{122, 128, 64})

	r, g, b = HSB2RGB(125, 50, 50)
	c.Assert([]int{r, g, b}, DeepEquals, []int{64, 128, 69})

	r, g, b = HSB2RGB(185, 50, 50)
	c.Assert([]int{r, g, b}, DeepEquals, []int{64, 122, 128})

	r, g, b = HSB2RGB(245, 50, 50)
	c.Assert([]int{r, g, b}, DeepEquals, []int{69, 64, 128})

	r, g, b = HSB2RGB(305, 50, 50)
	c.Assert([]int{r, g, b}, DeepEquals, []int{128, 64, 122})
}

func (s *ColorSuite) TestRGBACheck(c *C) {
	c.Assert(IsRGBA(int64(0xFFAABB)), Equals, false)
	c.Assert(IsRGBA(int64(0xFFAABB01)), Equals, true)
}

func (s *ColorSuite) TestRGB2Term(c *C) {
	c.Assert(RGB2Term(0, 175, 175), Equals, 37)
	c.Assert(RGB2Term(255, 255, 0), Equals, 226)
	c.Assert(RGB2Term(135, 175, 215), Equals, 110)
	// grayscale
	c.Assert(RGB2Term(175, 175, 175), Equals, 145)
	c.Assert(RGB2Term(18, 18, 18), Equals, 233)
	c.Assert(RGB2Term(48, 48, 48), Equals, 236)
	c.Assert(RGB2Term(238, 238, 238), Equals, 255)
}

func (s *ColorSuite) TestLuminance(c *C) {
	c.Assert(HEXLuminance(0xffffff), Equals, 1.0)
	c.Assert(HEXLuminance(0x000000), Equals, 0.0)
	c.Assert(HEXLuminance(0x808080), Equals, 0.2158605001138992)
	c.Assert(HEXLuminance(0x2861bd), Equals, 0.12674627666892935)
}

func (s *ColorSuite) TestCMYK(c *C) {
	C, M, Y, K := RGB2CMYK(0, 0, 0)
	c.Assert([]float64{C, M, Y, K}, DeepEquals, []float64{0.0, 0.0, 0.0, 1.0})

	C, M, Y, K = RGB2CMYK(255, 255, 255)
	c.Assert([]float64{C, M, Y, K}, DeepEquals, []float64{0.0, 0.0, 0.0, 0.0})

	C, M, Y, K = RGB2CMYK(102, 102, 102)
	c.Assert([]float64{C, M, Y, K}, DeepEquals, []float64{0.0, 0.0, 0.0, 0.6})

	C, M, Y, K = RGB2CMYK(102, 0, 0)
	c.Assert([]float64{C, M, Y, K}, DeepEquals, []float64{0.0, 1.0, 1.0, 0.6})

	r, g, b := CMYK2RGB(0.0, 0.0, 0.0, 0.0)
	c.Assert([]int{r, g, b}, DeepEquals, []int{255, 255, 255})

	r, g, b = CMYK2RGB(0.0, 0.0, 0.0, 1.0)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0, 0, 0})

	r, g, b = CMYK2RGB(0.64, 0.77, 0, 0.17)
	c.Assert([]int{r, g, b}, DeepEquals, []int{76, 48, 211})
}

func (s *ColorSuite) TestHSL(c *C) {
	H, S, L := RGB2HSL(255, 0, 0)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.0, 1.0, 0.5})

	H, S, L = RGB2HSL(0, 0, 0)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.0, 0.0, 0.0})

	H, S, L = RGB2HSL(63, 191, 191)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.5, 0.5039370078740157, 0.4980392156862745})

	H, S, L = RGB2HSL(255, 201, 201)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.0, 1.0, 0.8941176470588235})

	H, S, L = RGB2HSL(255, 130, 245)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.8466666666666667, 1.0, 0.7549019607843137})

	H, S, L = RGB2HSL(46, 196, 255)
	c.Assert([]float64{H, S, L}, DeepEquals, []float64{0.5470494417862839, 1.0, 0.5901960784313726})

	r, g, b := HSL2RGB(0.0, 0.0, 0.0)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0, 0, 0})

	r, g, b = HSL2RGB(0.0, 1.0, 0.5)
	c.Assert([]int{r, g, b}, DeepEquals, []int{255, 0, 0})

	r, g, b = HSL2RGB(0.6833333333333332, 0.7079646017699115, 0.5568627450980392)
	c.Assert([]int{r, g, b}, DeepEquals, []int{77, 62, 222})
}

func (s *ColorSuite) TestHUE(c *C) {
	c.Assert(HUE2RGB(0.6, 0.3, 0.12), Equals, 0.384)
	c.Assert(HUE2RGB(0.6, 0.3, 0.35), Equals, 0.3)
	c.Assert(HUE2RGB(0.6, 0.3, 0.584), Equals, 0.4512)
	c.Assert(HUE2RGB(0.6, 0.3, 0.85), Equals, 0.6)
}

func (s *ColorSuite) TestContrast(c *C) {
	c.Assert(Contrast(0xffffff, 0x000000), Equals, 21.0)
	c.Assert(Contrast(0x000000, 0x000000), Equals, 1.0)
	c.Assert(Contrast(0x755757, 0x63547c), Equals, 1.0542371982635754)
	c.Assert(Contrast(0x333333, 0xd0b5ff), Equals, 7.0606832983463805)
}
