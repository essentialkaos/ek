package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"

	"github.com/essentialkaos/ek/v14/fsutil"
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

	c.Assert(Create(""), Equals, ErrEmptyName)
	c.Assert(Remove(""), Equals, ErrEmptyName)

	_, err := Get("")

	c.Assert(err, Equals, ErrEmptyName)

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "directory /_NOT_EXIST doesn't exist or not accessible")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = os.Args[0]

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, ".* is not a directory")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = "/"

	err = Create("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "directory / is not writable")

	// //////////////////////////////////////////////////////////////////////////////// //

	err = Remove("test")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "directory / is not writable")

	// //////////////////////////////////////////////////////////////////////////////// //

	pidNum, err := Get("test")

	c.Assert(pidNum, Equals, 0)
	c.Assert(err.Error(), Equals, "can't read PID file: open /test.pid: no such file or directory")

	// //////////////////////////////////////////////////////////////////////////////// //

	Dir = s.Dir

	err = Create("")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "PID file name can't be blank")

	// //////////////////////////////////////////////////////////////////////////////// //

	pidNum, err = Get("_test_")

	c.Assert(pidNum, Equals, 0)
	c.Assert(err, ErrorMatches, "can't read PID file: open .*/_test_.pid: no such file or directory")

	// //////////////////////////////////////////////////////////////////////////////// //

	os.WriteFile(s.Dir+"/bad.pid", []byte("ABCDE\n"), 0644)

	pidNum, err = Get("bad.pid")

	c.Assert(pidNum, Equals, 0)
	c.Assert(err.Error(), Equals, "can't parse PID: strconv.Atoi: parsing \"ABCDE\": invalid syntax")

	os.WriteFile(s.Dir+"/bad.pid", []byte("0\n"), 0644)

	pidNum, err = Get("bad.pid")

	c.Assert(pidNum, Equals, 0)
	c.Assert(err, Equals, ErrInvalidPID)

	// //////////////////////////////////////////////////////////////////////////////// //

	nonReadableDir := s.Dir + "/non-readable"

	os.Mkdir(nonReadableDir, 0200)

	Dir = nonReadableDir

	err = Create("test.pid")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `directory .*/non-readable is not readable`)

	pidNum, err = Get("test.pid")

	c.Assert(pidNum, Equals, 0)
	c.Assert(err, ErrorMatches, "directory .*/non-readable is not readable")
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

	pid, _ := Get("test")

	c.Assert(pid, Not(Equals), 0)
	c.Assert(os.Getpid(), Equals, pid)

	Remove("test")

	c.Assert(fsutil.IsExist(Dir+"/test.pid"), Equals, false)
	c.Assert(IsWorks("test"), Equals, false)
}

// ////////////////////////////////////////////////////////////////////////////////// //
