package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type CronSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CronSuite{})

func (s *CronSuite) TestParsing(c *C) {
	e1, _ := Parse("* * * * *")

	c.Assert(e1.IsDue(time.Now()), Equals, true)
	c.Assert(e1.IsDue(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)), Equals, true)

	e2, _ := Parse("45 17 7 6 *")

	c.Assert(e2.IsDue(time.Date(2015, 6, 7, 17, 45, 0, 0, time.Local)), Equals, true)
	c.Assert(e2.IsDue(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)), Equals, false)

	e3, _ := Parse("0,15,30,45 0,6,12,18 1,15,31 * Mon-Fri")

	c.Assert(e3.IsDue(time.Date(2015, 1, 1, 12, 45, 0, 0, time.Local)), Equals, true)
	c.Assert(e3.IsDue(time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)), Equals, false)

	e4, _ := Parse("*/15 */6 1,15,31 * 1-5")

	c.Assert(e4.IsDue(time.Date(2015, 1, 15, 18, 15, 0, 0, time.Local)), Equals, true)
	c.Assert(e4.IsDue(time.Date(2015, 2, 15, 18, 15, 0, 0, time.Local)), Equals, false)

	e5, _ := Parse("* * * 1,3,5,7,9,11 *")

	c.Assert(e5.IsDue(time.Date(2015, 1, 15, 18, 15, 0, 0, time.Local)), Equals, true)
	c.Assert(e5.IsDue(time.Date(2015, 2, 15, 18, 15, 0, 0, time.Local)), Equals, false)

	e6, _ := Parse("* * * Jan,Feb,Mar *")

	c.Assert(e6.IsDue(time.Date(2015, 2, 15, 20, 22, 0, 0, time.Local)), Equals, true)
	c.Assert(e6.IsDue(time.Date(2015, 4, 15, 20, 22, 0, 0, time.Local)), Equals, false)
}
