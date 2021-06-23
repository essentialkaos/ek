// +build linux freebsd

package pid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "pkg.re/essentialkaos/check.v1"
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
	ioutil.WriteFile(s.Dir+"/test.pid", []byte("69999\n"), 0644)

	c.Assert(IsWorks("test"), Equals, false)

	procfsDir = c.MkDir()
	err = os.Mkdir(procfsDir+"/69999", 0755)

	if err != nil {
		c.Fatal(fmt.Sprintf("Can't create directory %s", procfsDir+"/69999"))
	}

	time.Sleep(2 * time.Second)

	err = os.Mkdir(procfsDir+"/1", 0755)

	if err != nil {
		c.Fatal(fmt.Sprintf("Can't create directory %s", procfsDir+"/1"))
	}

	c.Assert(IsWorks("test"), Equals, false)
}

func (s *PidSuite) TestIsProcessWorks(c *C) {
	c.Assert(IsProcessWorks(os.Getpid()), Equals, true)
	c.Assert(IsProcessWorks(999999), Equals, false)
}
