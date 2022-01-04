package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type WindowSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&WindowSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *WindowSuite) TestGetSize(c *C) {
	w, h := GetSize()

	c.Assert(w, Not(Equals), -1)
	c.Assert(w, Not(Equals), 0)
	c.Assert(h, Not(Equals), -1)
	c.Assert(h, Not(Equals), 0)
}

func (s *WindowSuite) TestGetWidth(c *C) {
	c.Assert(GetWidth(), Not(Equals), -1)
	c.Assert(GetWidth(), Not(Equals), 0)
}

func (s *WindowSuite) TestGetHeight(c *C) {
	c.Assert(GetHeight(), Not(Equals), -1)
	c.Assert(GetHeight(), Not(Equals), 0)
}

func (s *WindowSuite) TestErrors(c *C) {
	tty = "/non-exist"

	w, h := GetSize()

	c.Assert(w, Equals, -1)
	c.Assert(h, Equals, -1)

	tty = "/dev/tty"
}
