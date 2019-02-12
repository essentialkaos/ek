// +build darwin

package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
