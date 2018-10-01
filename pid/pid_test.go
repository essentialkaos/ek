package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "pkg.re/check.v1"

	"pkg.re/essentialkaos/ek.v10/fsutil"
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
	c.Assert(err.Error(), Equals, "Directory /_NOT_EXIST does not exist")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = os.Args[0]

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, fmt.Sprintf("%s is not directory", os.Args[0]))

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
	c.Assert(err.Error(), Equals, "Pid file name can't be blank")

	// //////////////////////////////////////////////////////////////////////////////// //

	pidNum = Get("_test_")

	c.Assert(pidNum, Equals, -1)

	// //////////////////////////////////////////////////////////////////////////////// //

	ioutil.WriteFile(s.Dir+"/bad.pid", []byte("ABCDE\n"), 0644)

	pidNum = Get("bad.pid")

	c.Assert(pidNum, Equals, -1)

	// //////////////////////////////////////////////////////////////////////////////// //

	nonReadableDir := s.Dir + "/non-readable"

	os.Mkdir(nonReadableDir, 0200)

	Dir = nonReadableDir

	err = Create("test.pid")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, fmt.Sprintf("Directory %s is not readable", nonReadableDir))
}

func (s *PidSuite) TestPidFuncs(c *C) {
	Dir = s.Dir

	err := Create("test.pid")

	c.Assert(err, IsNil)

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsReadable(Dir+"/test.pid"), Equals, true)
	c.Assert(fsutil.IsNonEmpty(Dir+"/test.pid"), Equals, true)

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
