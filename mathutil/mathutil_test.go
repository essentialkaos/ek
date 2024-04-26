package mathutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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
