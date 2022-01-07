package ansi

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

type ANSISuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ANSISuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ANSISuite) TestHasCodes(c *C) {
	c.Assert(HasCodes(""), Equals, false)
	c.Assert(HasCodes("ABCD^]"), Equals, false)
	c.Assert(HasCodes("\033ABCD"), Equals, true)
	c.Assert(HasCodes("\x1bABCD"), Equals, true)
	c.Assert(HasCodes("\x1BABCD"), Equals, true)
}

func (s *ANSISuite) TestRemoveCodes(c *C) {
	c.Assert(RemoveCodes(""), Equals, "")
	c.Assert(RemoveCodes("ABCD"), Equals, "ABCD")
	c.Assert(RemoveCodes("\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m"), Equals, "Hello World!")
}

func (s *ANSISuite) BenchmarkRemoveCodes(c *C) {
	for i := 0; i < c.N; i++ {
		RemoveCodes("\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m")
	}
}
