package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SystemSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SystemSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SystemSuite) TestUptime(c *C) {
	uptime, err := GetUptime()

	c.Assert(err, IsNil)
	c.Assert(uptime, Not(Equals), 0)
}

func (s *SystemSuite) TestLoadAvg(c *C) {
	la, err := GetLA()

	c.Assert(err, IsNil)
	c.Assert(la, NotNil)
}

func (s *SystemSuite) TestUser(c *C) {
	c.Assert(IsUserExist("root"), Equals, true)
	c.Assert(IsUserExist("_unknown_"), Equals, false)
	c.Assert(IsGroupExist("wheel"), Equals, true)
	c.Assert(IsGroupExist("_unknown_"), Equals, false)
}
