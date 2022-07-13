package lock

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type LockSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&LockSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *LockSuite) TestErrors(c *C) {
	Dir = "/_NOT_EXIST"

	err := Create("test")
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Directory /_NOT_EXIST doesn't exist or not accessible")

	err = Remove("test")
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Directory /_NOT_EXIST doesn't exist or not accessible")

	c.Assert(Expired("test", time.Second), Equals, false)
}

func (s *LockSuite) TestBasics(c *C) {
	Dir = c.MkDir()

	c.Assert(Has("test"), Equals, false)

	err := Create("test")
	c.Assert(err, IsNil)

	c.Assert(Has("test"), Equals, true)
	c.Assert(Expired("test", time.Minute), Equals, false)

	time.Sleep(time.Second)

	c.Assert(Expired("test", time.Millisecond), Equals, true)

	err = Remove("test")
	c.Assert(err, IsNil)

	c.Assert(Has("test"), Equals, false)
}
