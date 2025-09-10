package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type InitSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&InitSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *InitSuite) TestLaunchdIsPresent(c *C) {
	c.Assert(IsPresent("com.apple.unknown"), Equals, false)
	c.Assert(IsPresent("com.apple.homed"), Equals, true)
	c.Assert(IsPresent("com.apple.avatarsd"), Equals, true)
}

func (s *InitSuite) TestLaunchdIsWorks(c *C) {
	isWorks, err := IsWorks("com.apple.homed")

	c.Assert(err, IsNil)
	c.Assert(isWorks, Equals, true)

	isWorks, err = IsWorks("com.apple.avatarsd")

	c.Assert(err, IsNil)
	c.Assert(isWorks, Equals, false)
}
