// +build linux

package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////

func (s *FmtUtilSuite) TestTermSize(c *C) {
	w, h := GetTermSize()

	c.Assert(w, Not(Equals), 0)
	c.Assert(h, Not(Equals), 0)
}
