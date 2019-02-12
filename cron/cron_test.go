package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type CronSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CronSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CronSuite) TestParsing(c *C) {

	e0, err := Parse("* * * *")

	c.Assert(err, NotNil)
	c.Assert(e0, IsNil)

	e1, err := Parse("* * * * *")

	c.Assert(err, IsNil)
	c.Assert(e1, NotNil)
	c.Assert(e1.IsDue(time.Now()), Equals, true)
	c.Assert(e1.IsDue(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)), Equals, true)

	c.Assert(
		e1.Prev(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)),
		Equals,
		time.Date(2014, 12, 31, 23, 59, 0, 0, time.Local),
	)

	c.Assert(
		e1.Next(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)),
		Equals,
		time.Date(2015, 1, 1, 0, 1, 0, 0, time.Local),
	)

	c.Assert(e1.Next(), Not(Equals), time.Unix(0, 0))
	c.Assert(e1.Prev(), Not(Equals), time.Unix(0, 0))

	e2, err := Parse("45 17 7 6 *")

	c.Assert(err, IsNil)
	c.Assert(e2, NotNil)
	c.Assert(e2.IsDue(time.Date(2015, 6, 7, 17, 45, 0, 0, time.Local)), Equals, true)
	c.Assert(e2.IsDue(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)), Equals, false)

	e3, err := Parse("0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	c.Assert(err, IsNil)
	c.Assert(e3, NotNil)
	c.Assert(e3.IsDue(time.Date(2015, 1, 1, 12, 45, 0, 0, time.Local)), Equals, true)
	c.Assert(e3.IsDue(time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)), Equals, false)

	e4, err := Parse("*/15 */6 1,15,31 * 1-5")

	c.Assert(err, IsNil)
	c.Assert(e4, NotNil)
	c.Assert(e4.IsDue(time.Date(2015, 1, 15, 18, 15, 0, 0, time.Local)), Equals, true)
	c.Assert(e4.IsDue(time.Date(2015, 2, 15, 18, 15, 0, 0, time.Local)), Equals, false)

	e5, err := Parse("* * * 1,3,5,7,9,11 *")

	c.Assert(err, IsNil)
	c.Assert(e5, NotNil)
	c.Assert(e5.IsDue(time.Date(2015, 1, 15, 18, 15, 0, 0, time.Local)), Equals, true)
	c.Assert(e5.IsDue(time.Date(2015, 2, 15, 18, 15, 0, 0, time.Local)), Equals, false)

	e6, err := Parse("* * * Jan,Feb,Mar *")

	c.Assert(err, IsNil)
	c.Assert(e6, NotNil)
	c.Assert(e6.IsDue(time.Date(2015, 2, 15, 20, 22, 0, 0, time.Local)), Equals, true)
	c.Assert(e6.IsDue(time.Date(2015, 4, 15, 20, 22, 0, 0, time.Local)), Equals, false)

	e7, err := Parse("* * * * *")

	c.Assert(err, IsNil)
	c.Assert(e7, NotNil)
	c.Assert(e7.IsDue(), Equals, true)

	e8, err := Parse("15 12 10 * *")

	c.Assert(err, IsNil)
	c.Assert(e8, NotNil)
	c.Assert(e8.IsDue(time.Date(2015, 1, 1, 0, 15, 0, 0, time.Local)), Equals, false)
	c.Assert(e8.IsDue(time.Date(2015, 1, 1, 12, 15, 0, 0, time.Local)), Equals, false)

	c.Assert(e3.String(), Equals, "0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	e9, err := Parse("0 12 1 1 Mon")

	c.Assert(err, IsNil)
	c.Assert(e9, NotNil)

	c.Assert(
		e9.Prev(time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)),
		Equals,
		time.Unix(0, 0),
	)

	e10, err := Parse("0 12 1 1 Wed")

	c.Assert(err, IsNil)
	c.Assert(e10, NotNil)

	c.Assert(
		e10.Next(time.Date(2015, 6, 1, 0, 0, 0, 0, time.Local)),
		Equals,
		time.Unix(0, 0),
	)

	e11, err := Parse("45 17 7 0-5 1")

	c.Assert(err, IsNil)
	c.Assert(e11, NotNil)

	c.Assert(between(1, 5, 10), Equals, uint8(5))
	c.Assert(between(15, 5, 10), Equals, uint8(10))
	c.Assert(between(7, 5, 10), Equals, uint8(7))
}

func (s *CronSuite) TestAliases(c *C) {
	c.Assert(getAliasExpression("@yearly"), Equals, YEARLY)
	c.Assert(getAliasExpression("@annually"), Equals, ANNUALLY)
	c.Assert(getAliasExpression("@monthly"), Equals, MONTHLY)
	c.Assert(getAliasExpression("@weekly"), Equals, WEEKLY)
	c.Assert(getAliasExpression("@daily"), Equals, DAILY)
	c.Assert(getAliasExpression("@hourly"), Equals, HOURLY)

	dn1, dn1Ok := getDayNumByName("sun")
	dn2, dn2Ok := getDayNumByName("mon")
	dn3, dn3Ok := getDayNumByName("tue")
	dn4, dn4Ok := getDayNumByName("wed")
	dn5, dn5Ok := getDayNumByName("thu")
	dn6, dn6Ok := getDayNumByName("fri")
	dn7, dn7Ok := getDayNumByName("sat")
	dn8, dn8Ok := getDayNumByName("???")

	c.Assert(dn1, Equals, uint8(0))
	c.Assert(dn1Ok, Equals, true)
	c.Assert(dn2, Equals, uint8(1))
	c.Assert(dn2Ok, Equals, true)
	c.Assert(dn3, Equals, uint8(2))
	c.Assert(dn3Ok, Equals, true)
	c.Assert(dn4, Equals, uint8(3))
	c.Assert(dn4Ok, Equals, true)
	c.Assert(dn5, Equals, uint8(4))
	c.Assert(dn5Ok, Equals, true)
	c.Assert(dn6, Equals, uint8(5))
	c.Assert(dn6Ok, Equals, true)
	c.Assert(dn7, Equals, uint8(6))
	c.Assert(dn7Ok, Equals, true)
	c.Assert(dn8, Equals, uint8(0))
	c.Assert(dn8Ok, Equals, false)

	mn1, mn1Ok := getMonthNumByName("jan")
	mn2, mn2Ok := getMonthNumByName("feb")
	mn3, mn3Ok := getMonthNumByName("mar")
	mn4, mn4Ok := getMonthNumByName("apr")
	mn5, mn5Ok := getMonthNumByName("may")
	mn6, mn6Ok := getMonthNumByName("jun")
	mn7, mn7Ok := getMonthNumByName("jul")
	mn8, mn8Ok := getMonthNumByName("aug")
	mn9, mn9Ok := getMonthNumByName("sep")
	mn10, mn10Ok := getMonthNumByName("oct")
	mn11, mn11Ok := getMonthNumByName("nov")
	mn12, mn12Ok := getMonthNumByName("dec")
	mn13, mn13Ok := getMonthNumByName("???")

	c.Assert(mn1, Equals, uint8(1))
	c.Assert(mn1Ok, Equals, true)
	c.Assert(mn2, Equals, uint8(2))
	c.Assert(mn2Ok, Equals, true)
	c.Assert(mn3, Equals, uint8(3))
	c.Assert(mn3Ok, Equals, true)
	c.Assert(mn4, Equals, uint8(4))
	c.Assert(mn4Ok, Equals, true)
	c.Assert(mn5, Equals, uint8(5))
	c.Assert(mn5Ok, Equals, true)
	c.Assert(mn6, Equals, uint8(6))
	c.Assert(mn6Ok, Equals, true)
	c.Assert(mn7, Equals, uint8(7))
	c.Assert(mn7Ok, Equals, true)
	c.Assert(mn8, Equals, uint8(8))
	c.Assert(mn8Ok, Equals, true)
	c.Assert(mn9, Equals, uint8(9))
	c.Assert(mn9Ok, Equals, true)
	c.Assert(mn10, Equals, uint8(10))
	c.Assert(mn10Ok, Equals, true)
	c.Assert(mn11, Equals, uint8(11))
	c.Assert(mn11Ok, Equals, true)
	c.Assert(mn12, Equals, uint8(12))
	c.Assert(mn12Ok, Equals, true)
	c.Assert(mn13, Equals, uint8(0))
	c.Assert(mn13Ok, Equals, false)
}

func (s *CronSuite) TestNearIndex(c *C) {
	items := []uint8{1, 2, 3, 4, 6, 7, 8, 9}

	c.Assert(getNearNextIndex(items, 5), Equals, 4)
	c.Assert(getNearNextIndex(items, 6), Equals, 4)
	c.Assert(getNearNextIndex(items, 10), Equals, 0)
	c.Assert(getNearPrevIndex(items, 5), Equals, 3)
	c.Assert(getNearPrevIndex(items, 6), Equals, 4)
	c.Assert(getNearPrevIndex(items, 0), Equals, 7)
}

func (s *CronSuite) TestErrors(c *C) {
	e, err := Parse("0-A * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("A-1 * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("*/A * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("*/0 * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("A * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("0,1,A * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)

	e, err = Parse("0,1,2-A * * * *")

	c.Assert(e, IsNil)
	c.Assert(err, NotNil)
}
