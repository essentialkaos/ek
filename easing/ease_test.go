package easing

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type EaseSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&EaseSuite{})

func (s *EaseSuite) TestLinear(c *C) {
	c.Assert(Linear(0, 0, 10, 10), Equals, 0.0)
	c.Assert(Linear(10, 0, 10, 10), Equals, 10.0)
	c.Assert(Linear(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestQuad(c *C) {
	c.Assert(QuadIn(5, 0, 10, 10), Equals, 2.5)
	c.Assert(QuadOut(5, 0, 10, 10), Equals, 7.5)
	c.Assert(QuadInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(QuadIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(QuadOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(QuadInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestCubic(c *C) {
	c.Assert(CubicIn(5, 0, 10, 10), Equals, 1.25)
	c.Assert(CubicOut(5, 0, 10, 10), Equals, 8.75)
	c.Assert(CubicInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(CubicIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(CubicOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(CubicInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestQuint(c *C) {
	c.Assert(QuintIn(5, 0, 10, 10), Equals, 0.3125)
	c.Assert(QuintOut(5, 0, 10, 10), Equals, 9.6875)
	c.Assert(QuintInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(QuintIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(QuintOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(QuintInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestSine(c *C) {
	c.Assert(SineIn(5, 0, 10, 10), Equals, 3.0978992192751917)
	c.Assert(SineOut(5, 0, 10, 10), Equals, 7.236090437019012)
	c.Assert(SineInOut(5, 0, 10, 10), Equals, 4.999999999999999)
	c.Assert(SineIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(SineOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(SineInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestExpo(c *C) {
	c.Assert(ExpoIn(5, 0, 10, 10), Equals, 0.3125)
	c.Assert(ExpoOut(5, 0, 10, 10), Equals, 9.6875)
	c.Assert(ExpoInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(ExpoIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(ExpoOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(ExpoInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestCirc(c *C) {
	c.Assert(CircIn(5, 0, 10, 10), Equals, 1.339745962155614)
	c.Assert(CircOut(5, 0, 10, 10), Equals, 8.660254037844386)
	c.Assert(CircInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(CircIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(CircOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(CircInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestElastic(c *C) {
	c.Assert(ElasticIn(1, 0, 10, 10), Equals, 0.01953125)
	c.Assert(ElasticOut(5, 0, 10, 10), Equals, 10.15625)
	c.Assert(ElasticInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(ElasticIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(ElasticOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(ElasticInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestBack(c *C) {
	c.Assert(BackIn(5, 0, 10, 10), Equals, -0.9655673136318949)
	c.Assert(BackOut(5, 0, 10, 10), Equals, 10.965567313631894)
	c.Assert(BackInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(BackIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(BackOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(BackInOut(20, 0, 10, 10), Equals, 10.0)
}

func (s *EaseSuite) TestBounce(c *C) {
	c.Assert(BounceIn(5, 0, 10, 10), Equals, 2.34375)
	c.Assert(BounceOut(5, 0, 10, 10), Equals, 7.65625)
	c.Assert(BounceInOut(5, 0, 10, 10), Equals, 5.0)
	c.Assert(BounceIn(20, 0, 10, 10), Equals, 10.0)
	c.Assert(BounceOut(20, 0, 10, 10), Equals, 10.0)
	c.Assert(BounceInOut(20, 0, 10, 10), Equals, 10.0)
}
