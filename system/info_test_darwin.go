package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "pkg.re/essentialkaos/check.v1"
)

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
