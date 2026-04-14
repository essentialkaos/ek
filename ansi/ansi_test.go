package ansi

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

type ANSISuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ANSISuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ANSISuite) TestHas(c *C) {
	c.Assert(Has(""), Equals, false)
	c.Assert(Has("ABCD^]"), Equals, false)
	c.Assert(Has("\033ABCD"), Equals, true)
	c.Assert(Has("\x1bABCD"), Equals, true)
	c.Assert(Has("\x1BABCD"), Equals, true)
}

func (s *ANSISuite) TestHasBytes(c *C) {
	c.Assert(HasBytes([]byte{}), Equals, false)
	c.Assert(HasBytes([]byte("ABCD^]")), Equals, false)
	c.Assert(HasBytes([]byte("\033ABCD")), Equals, true)
	c.Assert(HasBytes([]byte("\x1bABCD")), Equals, true)
	c.Assert(HasBytes([]byte("\x1BABCD")), Equals, true)
}

func (s *ANSISuite) TestRemove(c *C) {
	c.Assert(Remove(""), Equals, "")
	c.Assert(Remove("ABCD"), Equals, "ABCD")
	c.Assert(Remove("\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m"), Equals, "Hello World!")
}

func (s *ANSISuite) TestRemoveBytes(c *C) {
	c.Assert(RemoveBytes([]byte{}), DeepEquals, []byte{})
	c.Assert(RemoveBytes([]byte("ABCD")), DeepEquals, []byte("ABCD"))
	c.Assert(RemoveBytes(
		[]byte("\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m")),
		DeepEquals, []byte("Hello World!"),
	)
}

func (s *ANSISuite) BenchmarkRemove(c *C) {
	data := "\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m"

	for range c.N {
		Remove(data)
	}
}

func (s *ANSISuite) BenchmarkRemoveBytes(c *C) {
	data := []byte("\033[40;38;5;82mHello \x1b[30;48;5;82mWorld!\x1B[0m")

	for range c.N {
		RemoveBytes(data)
	}
}
