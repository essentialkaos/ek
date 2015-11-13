package rand

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

type RandSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&RandSuite{})

func (s *RandSuite) TestString(c *C) {
	c.Assert(len(String(100)), Equals, 100)
	c.Assert(len(String(1)), Equals, 1)
	c.Assert(len(String(0)), Equals, 0)
	c.Assert(len(String(-100)), Equals, 0)
}
