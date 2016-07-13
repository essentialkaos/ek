package color

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
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

	r, g, b = Hex2RGB(0xFFFFFFFF)
	c.Assert([]int{r, g, b}, DeepEquals, []int{0xFF, 0xFF, 0xFF})
}

func (s *ColorSuite) TestRGBA2Hex(c *C) {
	c.Assert(RGBA2Hex(0, 0, 0, 0), Equals, 0x00000000)
	c.Assert(RGBA2Hex(255, 255, 255, 0), Equals, 0xFFFFFF00)
	c.Assert(RGBA2Hex(255, 255, 255, 255), Equals, 0xFFFFFFFF)
	c.Assert(RGBA2Hex(127, 127, 127, 127), Equals, 0x7F7F7F7F)
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

func (s *ColorSuite) TestRGB2HSB(c *C) {
	var h, sv, b int

	h, sv, b = RGB2HSB(0, 0, 0)
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
}

func (s *ColorSuite) TestHSB2RGB(c *C) {
	var r, g, b int

	r, g, b = HSB2RGB(0, 0, 0)
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
	c.Assert(IsRGBA(0xFFAABB), Equals, false)
	c.Assert(IsRGBA(0xFFAABB01), Equals, true)
}
