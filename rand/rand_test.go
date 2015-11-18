package rand

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type RandSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&RandSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *RandSuite) TestString(c *C) {
	c.Assert(len(String(100)), Equals, 100)
	c.Assert(len(String(1)), Equals, 1)
	c.Assert(len(String(0)), Equals, 0)
	c.Assert(len(String(-100)), Equals, 0)
}

func (s *RandSuite) TestInt(c *C) {
	n := 1000
	k := 0

	for i := 0; i < 1000; i++ {
		k += Int(n)
	}

	c.Assert(k/n, Not(Equals), n)
}

func (s *RandSuite) TestSlice(c *C) {
	t1 := strings.Join(Slice(256), "")
	t2 := strings.Join(Slice(256), "")

	c.Assert(t1, Not(Equals), t2)
	c.Assert(Slice(0), HasLen, 0)
}
