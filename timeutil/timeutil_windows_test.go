//go:build windows

package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TimeUtilSuite) TestLocalTimezone(c *C) {
	c.Assert(LocalTimezone(), Not(Equals), "Local")

	registryKeyName = "Unknown"
	c.Assert(LocalTimezone(), Equals, "Local")

	registryKey = "Unknown"
	c.Assert(LocalTimezone(), Equals, "Local")
}
