package pager

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

func Test(t *testing.T) { TestingT(t) }

type PagerSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PagerSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PagerSuite) TearDownSuite(c *C) {
	Complete()
}

func (s *PagerSuite) TestPager(c *C) {
	Redirect()()
	c.Assert(pagerCmd, IsNil)

	c.Assert(Setup(), IsNil)
	Complete()

	c.Assert(Setup("cat"), IsNil)
	c.Assert(Setup("cat"), DeepEquals, ErrAlreadySet)
	Complete()

	c.Assert(pagerCmd, IsNil)
	c.Assert(pagerOut, IsNil)

	os.Setenv("PAGER", "")

	binMore = "_unknown_"
	binLess = "_unknown_"

	c.Assert(Setup(""), DeepEquals, ErrNoPager)
}

func (s *PagerSuite) TestPagerSearch(c *C) {
	os.Setenv("PAGER", "")
	os.Setenv("LESS", "")
	os.Setenv("MORE", "")

	cmd := getPagerCommand("cat")
	c.Assert(cmd.Args, DeepEquals, []string{"cat"})

	binMore = "echo"
	binLess = "echo"

	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"more", "-f"})

	os.Setenv("MORE", "-l -s")

	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"more", "-l", "-s"})

	binMore = "_unknown_"

	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"less", "-R"})

	os.Setenv("LESS", "-MRQ")

	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"less", "-MRQ"})

	binMore = "more"
	binLess = "less"
}
