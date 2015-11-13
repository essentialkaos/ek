package timeutil

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

type TimeUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TimeUtilSuite{})

func (s *TimeUtilSuite) TestPretyDuration(c *C) {
	c.Assert(PrettyDuration(time.Duration(59000000000)), Equals, "59 seconds")
	c.Assert(PrettyDuration(120), Equals, "2 minutes")
	c.Assert(PrettyDuration(60), Equals, "1 minute")
	c.Assert(PrettyDuration(1234567), Equals, "2 weeks 6 hours 56 minutes and 7 seconds")
	c.Assert(PrettyDuration(0), Equals, "< 1 second")
}

func (s *TimeUtilSuite) TestFormat(c *C) {
	d := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)

	c.Assert(Format(d, "%%"), Equals, "%")
	c.Assert(Format(d, "%a"), Equals, "Wed")
	c.Assert(Format(d, "%A"), Equals, "Wednesday")
	c.Assert(Format(d, "%b"), Equals, "Jan")
	c.Assert(Format(d, "%B"), Equals, "January")
	c.Assert(Format(d, "%c"), Equals, "Wed 01 Jan 2014 12:00:00 AM UTC")
	c.Assert(Format(d, "%C"), Equals, "20")
	c.Assert(Format(d, "%d"), Equals, "01")
	c.Assert(Format(d, "%D"), Equals, "01/01/14")
	c.Assert(Format(d, "%e"), Equals, " 1")
	c.Assert(Format(d, "%F"), Equals, "2014-01-01")
	c.Assert(Format(d, "%G"), Equals, "2014")
	c.Assert(Format(d, "%H"), Equals, "00")
	c.Assert(Format(d, "%I"), Equals, "12")
	c.Assert(Format(d, "%j"), Equals, "001")
	c.Assert(Format(d, "%k"), Equals, " 0")
	c.Assert(Format(d, "%l"), Equals, "12")
	c.Assert(Format(d, "%m"), Equals, "01")
	c.Assert(Format(d, "%M"), Equals, "00")
	c.Assert(Format(d, "%N"), Equals, "000000000")
	c.Assert(Format(d, "%p"), Equals, "am")
	c.Assert(Format(d, "%P"), Equals, "AM")
	c.Assert(Format(d, "%r"), Equals, "12:00:00 AM")
	c.Assert(Format(d, "%R"), Equals, "00:00")
	c.Assert(Format(d, "%s"), Equals, "1388534400")
	c.Assert(Format(d, "%S"), Equals, "00")
	c.Assert(Format(d, "%T"), Equals, "00:00:00")
	c.Assert(Format(d, "%u"), Equals, "3")
	c.Assert(Format(d, "%V"), Equals, "01")
	c.Assert(Format(d, "%w"), Equals, "3")
	c.Assert(Format(d, "%y"), Equals, "14")
	c.Assert(Format(d, "%Y"), Equals, "2014")
	c.Assert(Format(d, "%z"), Equals, "+0000")
	c.Assert(Format(d, "%Z"), Equals, "UTC")
	c.Assert(Format(d, "%Q"), Equals, "%Q")
	c.Assert(Format(d, "%1234"), Equals, "%1234")
	c.Assert(Format(d, "%SSec"), Equals, "00Sec")
}

func (s *TimeUtilSuite) TestTinyDate(c *C) {
	dt := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	td := TinyDate(dt.Unix())

	c.Assert(td.Unix(), Equals, dt.Unix())
	c.Assert(td.Time().Unix(), Equals, dt.Unix())
}
