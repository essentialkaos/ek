package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fsutil"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type PidSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PidSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ps *PidSuite) SetUpSuite(c *C) {
	Dir = c.MkDir()
}

func (ps *PidSuite) TestCreate(c *C) {
	err := Create("test")

	c.Assert(err, IsNil)

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsNonEmpty(Dir+"/test.pid"), Equals, true)
}

func (ps *PidSuite) TestGet(c *C) {
	pid := Get("test")

	c.Assert(pid, Not(Equals), -1)
	c.Assert(os.Getpid(), Equals, pid)
}

func (ps *PidSuite) TestRemove(c *C) {
	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, true)

	Remove("test")

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, false)
}

// ////////////////////////////////////////////////////////////////////////////////// //
