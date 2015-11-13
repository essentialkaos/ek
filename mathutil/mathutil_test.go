package mathutil

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

type MathUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&MathUtilSuite{})

func (s *MathUtilSuite) TestBetween(c *C) {
	// Ints
	c.Assert(Between(5, 0, 100), Equals, 5)
	c.Assert(Between(10, 0, 5), Equals, 5)
	c.Assert(Between(-5, -10, 0), Equals, -5)
	c.Assert(Between(5, 10, 10), Equals, 10)

	// Floats
	c.Assert(BetweenF(0.5, 0.1, 0.7), Equals, 0.5)
	c.Assert(BetweenF(0.01, 0.1, 0.7), Equals, 0.1)
	c.Assert(BetweenF(5.0, 0.1, 0.7), Equals, 0.7)
}

func (s *MathUtilSuite) TestRound(c *C) {
	c.Assert(Round(5.49, 0), Equals, 5.0)
	c.Assert(Round(5.50, 0), Equals, 6.0)
	c.Assert(Round(5.51, 0), Equals, 6.0)
}
