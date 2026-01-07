package mathutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type MathUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&MathUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *MathUtilSuite) TestChecks(c *C) {
	c.Assert(IsNumber("1234567890.1"), Equals, true)
	c.Assert(IsNumber("1234567890"), Equals, true)
	c.Assert(IsInt(""), Equals, false)
	c.Assert(IsFloat(""), Equals, false)
	c.Assert(IsInt("1234567890a"), Equals, false)
	c.Assert(IsFloat("1234567890.1a"), Equals, false)
}

func (s *MathUtilSuite) TestBetween(c *C) {
	c.Assert(Between(5, 0, 100), Equals, 5)
	c.Assert(Between(10, 0, 5), Equals, 5)
	c.Assert(Between(-5, -10, 0), Equals, -5)
	c.Assert(Between(5, 10, 10), Equals, 10)
}

func (s *MathUtilSuite) TestMinMax(c *C) {
	c.Assert(Min(1, 10), Equals, 1)
	c.Assert(Min(-10, 10), Equals, -10)
	c.Assert(Min(10, -10), Equals, -10)
	c.Assert(Max(1, 10), Equals, 10)
	c.Assert(Max(-10, 10), Equals, 10)
	c.Assert(Max(10, -10), Equals, 10)
}

func (s *MathUtilSuite) TestAbs(c *C) {
	c.Assert(Abs(-10), Equals, 10)
	c.Assert(Abs(10), Equals, 10)
}

func (s *MathUtilSuite) TestRound(c *C) {
	c.Assert(Round(5.49, 0), Equals, 5.0)
	c.Assert(Round(5.50, 0), Equals, 6.0)
	c.Assert(Round(5.51, 0), Equals, 6.0)
}

func (s *MathUtilSuite) TestPerc(c *C) {
	c.Assert(Perc(0, 0), Equals, 0.0)
	c.Assert(Perc(0, 100), Equals, 0.0)
	c.Assert(Perc(25, 100), Equals, 25.0)
	c.Assert(Perc(100, 100), Equals, 100.0)
	c.Assert(Perc(200, 100), Equals, 200.0)
}

func (s *MathUtilSuite) TestFromPerc(c *C) {
	c.Assert(FromPerc(0, 0), Equals, 0.0)
	c.Assert(FromPerc(100, 0), Equals, 0.0)
	c.Assert(FromPerc(-1, 1000), Equals, 0.0)
	c.Assert(FromPerc(250, 100), Equals, 250.0)
	c.Assert(FromPerc(25.55, -100), Equals, -25.55)
}

func (s *MathUtilSuite) TestB(c *C) {
	c.Assert(B(true, 2.3, 3.7), Equals, 2.3)
	c.Assert(B(false, 2.3, 3.7), Equals, 3.7)
}
