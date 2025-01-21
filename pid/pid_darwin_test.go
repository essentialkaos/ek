//go:build darwin
// +build darwin

package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"

	. "github.com/essentialkaos/check"
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
	os.WriteFile(s.Dir+"/test.pid", []byte("999999\n"), 0644)

	c.Assert(IsWorks("test"), Equals, false)
}

func (s *PidSuite) TestIsProcessWorks(c *C) {
	c.Assert(IsProcessWorks(os.Getpid()), Equals, true)
	c.Assert(IsProcessWorks(999999), Equals, false)
}
