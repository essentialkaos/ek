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
	c.Assert(Setup("cat"), IsNil)
	c.Assert(Setup("cat"), NotNil)

	Complete()

	c.Assert(pagerCmd, IsNil)
	c.Assert(pagerOut, IsNil)
}

func (s *PagerSuite) TestPagerSearch(c *C) {
	os.Setenv("PAGER", "")

	cmd := getPagerCommand("cat")
	c.Assert(cmd.Args, DeepEquals, []string{"cat"})

	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"more"})

	os.Setenv("PAGER", "less -MQR")
	cmd = getPagerCommand("")
	c.Assert(cmd.Args, DeepEquals, []string{"less", "-MQR"})
}
