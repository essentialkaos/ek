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

	c.Assert(IsExpired("test", time.Second), Equals, false)
}

func (s *LockSuite) TestBasics(c *C) {
	Dir = c.MkDir()

	c.Assert(Has("test"), Equals, false)

	err := Create("test")
	c.Assert(err, IsNil)

	c.Assert(Has("test"), Equals, true)
	c.Assert(IsExpired("test", time.Minute), Equals, false)

	time.Sleep(time.Second)

	c.Assert(IsExpired("test", time.Millisecond), Equals, true)

	err = Remove("test")
	c.Assert(err, IsNil)

	c.Assert(Has("test"), Equals, false)
}

func (s *LockSuite) TestWait(c *C) {
	Dir = c.MkDir()

	c.Assert(Wait("test", time.Now().Add(time.Second)), Equals, true)

	err := Create("test")
	c.Assert(err, IsNil)

	go func() {
		time.Sleep(time.Second)
		Remove("test")
	}()

	c.Assert(Wait("test", time.Now().Add(time.Microsecond)), Equals, false)
	c.Assert(Wait("test", time.Now().Add(3*time.Second)), Equals, true)
}
