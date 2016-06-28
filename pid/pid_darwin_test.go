package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PidSuite) TestIsWorks(c *C) {
	Dir = s.Dir

	err := Create("test")

	c.Assert(err, IsNil)

	c.Assert(IsWorks("test"), Equals, true)

	Remove("test")

	c.Assert(IsWorks("test"), Equals, false)

	// Write fake pid to pid file
	ioutil.WriteFile(s.Dir+"/test.pid", []byte("9736163"), 0644)

	c.Assert(IsWorks("test"), Equals, true)
}
