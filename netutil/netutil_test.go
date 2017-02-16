package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
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
	c.Assert(GetIP6(), Not(Equals), "")
}
