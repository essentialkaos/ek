package rand

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "github.com/essentialkaos/check"
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

func (s *RandSuite) TestSlice(c *C) {
	t1 := strings.Join(Slice(256), "")
	t2 := strings.Join(Slice(256), "")

	c.Assert(t1, Not(Equals), t2)
	c.Assert(Slice(0), HasLen, 0)
}
