package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"testing"

	. "github.com/essentialkaos/check"

	"github.com/essentialkaos/ek/v13/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type PidSuite struct {
	Dir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PidSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PidSuite) SetUpSuite(c *C) {
	s.Dir = c.MkDir()
}

func (s *PidSuite) TestErrors(c *C) {
	Dir = "/_NOT_EXIST"

	err := Create("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Directory /_NOT_EXIST doesn't exist or not accessible")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = os.Args[0]

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, ".* is not a directory")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = "/"

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Directory / is not writable")

	// //////////////////////////////////////////////////////////////////////////////// //

	err = Remove("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Directory / is not writable")

	// //////////////////////////////////////////////////////////////////////////////// //

	pidNum := Get("test")

	c.Assert(pidNum, Equals, -1)

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = s.Dir

	err = Create("")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "PID file name can't be blank")

	// //////////////////////////////////////////////////////////////////////////////// //

	pidNum = Get("_test_")

	c.Assert(pidNum, Equals, -1)

	// //////////////////////////////////////////////////////////////////////////////// //

	os.WriteFile(s.Dir+"/bad.pid", []byte("ABCDE\n"), 0644)

	pidNum = Get("bad.pid")

	c.Assert(pidNum, Equals, -1)

	// //////////////////////////////////////////////////////////////////////////////// //

	nonReadableDir := s.Dir + "/non-readable"

	os.Mkdir(nonReadableDir, 0200)

	Dir = nonReadableDir

	err = Create("test.pid")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, fmt.Sprintf("Directory %s is not readable", nonReadableDir))
	c.Assert(Get("test.pid"), Equals, -1)
}

func (s *PidSuite) TestPidFuncs(c *C) {
	Dir = s.Dir

	err := Create("test.pid")

	c.Assert(err, IsNil)

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsEmpty(Dir+"/test.pid"), Equals, false)

	err = Create("test")

	c.Assert(err, IsNil)

	pid := Get("test")

	c.Assert(pid, Not(Equals), -1)
	c.Assert(os.Getpid(), Equals, pid)

	Remove("test")

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, false)
	c.Assert(IsWorks("test"), Equals, false)
}

// ////////////////////////////////////////////////////////////////////////////////// //
