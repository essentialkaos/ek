package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

type NetUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&NetUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *NetUtilSuite) TestGetIP(c *C) {
	c.Assert(GetIP(), Not(Equals), "")
}

func (s *NetUtilSuite) TestGetIP6(c *C) {
	if os.Getenv("CI") == "" {
		c.Assert(GetIP6(), Not(Equals), "")
	}
}

func (s *NetUtilSuite) TestGetAllIP(c *C) {
	c.Assert(GetAllIP(), Not(HasLen), 0)
}

func (s *NetUtilSuite) TestGetAllIP6(c *C) {
	if os.Getenv("CI") == "" {
		c.Assert(GetAllIP6(), Not(HasLen), 0)
	}
}
