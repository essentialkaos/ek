//go:build linux || freebsd || darwin

package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TimeUtilSuite) TestLocalTimezone(c *C) {
	os.Unsetenv("TZ")
	c.Assert(LocalTimezone(), Not(Equals), "Local")

	os.Setenv("TZ", "Europe/Zurich")
	c.Assert(LocalTimezone(), Equals, "Europe/Zurich")
	os.Unsetenv("TZ")

	testDir := c.MkDir()
	os.Symlink("/usr/share/unknown", testDir+"/localtime")
	localZoneFile = testDir + "/localtime"

	c.Assert(LocalTimezone(), Equals, "Local")

	localZoneFile = "/_unknown_"
	c.Assert(LocalTimezone(), Equals, "Local")
}
