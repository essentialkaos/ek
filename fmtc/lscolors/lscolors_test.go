package lscolors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type LSCSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&LSCSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ls *LSCSuite) TestColorize(c *C) {
	colorMap, initialized = nil, false

	os.Setenv("LS_COLORS", "rs=0:di=01;38;5;75:ln=38;5;141:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=38;5;202:*.txt=38;5;178:*.bz=38;5;105:")

	c.Assert(GetColor("test.log"), Equals, "")
	c.Assert(GetColor("test.txt"), Equals, "\x1b[38;5;178m")
	c.Assert(GetColor("test.tar.bz"), Equals, "\x1b[38;5;105m")

	c.Assert(Colorize("test.log"), Equals, "test.log")
	c.Assert(Colorize("test.txt"), Equals, "\x1b[38;5;178mtest.txt\x1b[0m")
	c.Assert(Colorize("test.tar.bz"), Equals, "\x1b[38;5;105mtest.tar.bz\x1b[0m")

	c.Assert(ColorizePath("/home/john/test.log"), Equals, "/home/john/test.log")
	c.Assert(ColorizePath("/home/john/test.txt"), Equals, "\x1b[38;5;178m/home/john/test.txt\x1b[0m")
	c.Assert(ColorizePath("/home/john/test.tar.bz"), Equals, "\x1b[38;5;105m/home/john/test.tar.bz\x1b[0m")

	colorMap, initialized = nil, false

	os.Setenv("LS_COLORS", "")

	c.Assert(Colorize("test.log"), Equals, "test.log")
	c.Assert(Colorize("test.txt"), Equals, "test.txt")
	c.Assert(Colorize("test.tar.bz"), Equals, "test.tar.bz")
	c.Assert(ColorizePath("/home/john/test.log"), Equals, "/home/john/test.log")
	c.Assert(ColorizePath("/home/john/test.txt"), Equals, "/home/john/test.txt")
	c.Assert(ColorizePath("/home/john/test.tar.bz"), Equals, "/home/john/test.tar.bz")
}
