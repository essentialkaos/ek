package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"testing"
	"time"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TimeUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TimeUtilSuite{})

func (s *TimeUtilSuite) TestPretyDuration(c *C) {
	c.Assert(PrettyDuration(time.Duration(59000000000)), Equals, "59 seconds")
	c.Assert(PrettyDuration(120), Equals, "2 minutes")
	c.Assert(PrettyDuration(int8(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(int16(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(int32(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(int64(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(3720), Equals, "1 hour and 2 minutes")
	c.Assert(PrettyDuration(60), Equals, "1 minute")
	c.Assert(PrettyDuration(1370137), Equals, "2 weeks 1 day 20 hours 35 minutes and 37 seconds")
	c.Assert(PrettyDuration(0), Equals, "< 1 second")
	c.Assert(PrettyDuration("string"), Equals, "Wrong duration value")
	c.Assert(PrettyDuration(time.Second/3), Equals, "333.3 ms")
	c.Assert(PrettyDuration(time.Second/3669), Equals, "272.6 μs")
	c.Assert(PrettyDuration(time.Second/366995), Equals, "2.72 μs")
	c.Assert(PrettyDuration(time.Second/36698888), Equals, "27 ns")
}

func (s *TimeUtilSuite) TestDurationToSeconds(c *C) {
	c.Assert(DurationToSeconds(time.Minute), Equals, int64(60))
	c.Assert(DurationToSeconds(time.Hour), Equals, int64(3600))
}

func (s *TimeUtilSuite) TestFormat(c *C) {
	d := time.Unix(1388535309, 123456789).UTC()
	d1 := time.Unix(1389744909, 123456789).UTC()

	c.Assert(Format(d, "%%"), Equals, "%")
	c.Assert(Format(d, "%a"), Equals, "Wed")
	c.Assert(Format(d, "%A"), Equals, "Wednesday")
	c.Assert(Format(d, "%b"), Equals, "Jan")
	c.Assert(Format(d, "%B"), Equals, "January")
	c.Assert(Format(d, "%c"), Equals, "Wed 01 Jan 2014 12:15:09 AM UTC")
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
	c.Assert(Format(d, "%K"), Equals, "123")
	c.Assert(Format(d, "%l"), Equals, "12")
	c.Assert(Format(d, "%m"), Equals, "01")
	c.Assert(Format(d, "%M"), Equals, "15")
	c.Assert(Format(d, "%N"), Equals, "123456789")
	c.Assert(Format(d, "%n"), Equals, "\n")
	c.Assert(Format(d, "%p"), Equals, "am")
	c.Assert(Format(d, "%P"), Equals, "AM")
	c.Assert(Format(d, "%r"), Equals, "12:15:09 AM")
	c.Assert(Format(d, "%R"), Equals, "00:15")
	c.Assert(Format(d, "%s"), Equals, "1388535309")
	c.Assert(Format(d, "%S"), Equals, "09")
	c.Assert(Format(d, "%T"), Equals, "00:15:09")
	c.Assert(Format(d, "%u"), Equals, "3")
	c.Assert(Format(d, "%V"), Equals, "01")
	c.Assert(Format(d, "%w"), Equals, "3")
	c.Assert(Format(d, "%y"), Equals, "14")
	c.Assert(Format(d, "%Y"), Equals, "2014")
	c.Assert(Format(d, "%z"), Equals, "+0000")
	c.Assert(Format(d, "%Z"), Equals, "UTC")
	c.Assert(Format(d, "%:"), Equals, "+00:00")
	c.Assert(Format(d, "%Q"), Equals, "%Q")
	c.Assert(Format(d, "%1234"), Equals, "%1234")
	c.Assert(Format(d, "%SSec"), Equals, "09Sec")
	c.Assert(Format(d1, "%e"), Equals, "15")

	replaceDateTag(time.Now(), bytes.NewBufferString(""), bytes.NewBufferString(""))
}

func (s *TimeUtilSuite) TestTinyDate(c *C) {
	dt := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	td := TinyDate(dt.Unix())

	c.Assert(td.Unix(), Equals, dt.Unix())
	c.Assert(td.Time().Unix(), Equals, dt.Unix())
}

func (s *TimeUtilSuite) TestDateNames(c *C) {
	c.Assert(getShortWeekday(time.Sunday), Equals, "Sun")
	c.Assert(getShortWeekday(time.Monday), Equals, "Mon")
	c.Assert(getShortWeekday(time.Tuesday), Equals, "Tue")
	c.Assert(getShortWeekday(time.Wednesday), Equals, "Wed")
	c.Assert(getShortWeekday(time.Thursday), Equals, "Thu")
	c.Assert(getShortWeekday(time.Friday), Equals, "Fri")
	c.Assert(getShortWeekday(time.Saturday), Equals, "Sat")
	c.Assert(getShortWeekday(time.Weekday(7)), Equals, "")

	c.Assert(getLongWeekday(time.Sunday), Equals, "Sunday")
	c.Assert(getLongWeekday(time.Monday), Equals, "Monday")
	c.Assert(getLongWeekday(time.Tuesday), Equals, "Tuesday")
	c.Assert(getLongWeekday(time.Wednesday), Equals, "Wednesday")
	c.Assert(getLongWeekday(time.Thursday), Equals, "Thursday")
	c.Assert(getLongWeekday(time.Friday), Equals, "Friday")
	c.Assert(getLongWeekday(time.Saturday), Equals, "Saturday")
	c.Assert(getLongWeekday(time.Weekday(7)), Equals, "")

	c.Assert(getShortMonth(time.Month(0)), Equals, "")
	c.Assert(getShortMonth(time.January), Equals, "Jan")
	c.Assert(getShortMonth(time.February), Equals, "Feb")
	c.Assert(getShortMonth(time.March), Equals, "Mar")
	c.Assert(getShortMonth(time.April), Equals, "Apr")
	c.Assert(getShortMonth(time.May), Equals, "May")
	c.Assert(getShortMonth(time.June), Equals, "Jun")
	c.Assert(getShortMonth(time.July), Equals, "Jul")
	c.Assert(getShortMonth(time.August), Equals, "Aug")
	c.Assert(getShortMonth(time.September), Equals, "Sep")
	c.Assert(getShortMonth(time.October), Equals, "Oct")
	c.Assert(getShortMonth(time.November), Equals, "Nov")
	c.Assert(getShortMonth(time.December), Equals, "Dec")

	c.Assert(getLongMonth(time.Month(0)), Equals, "")
	c.Assert(getLongMonth(time.January), Equals, "January")
	c.Assert(getLongMonth(time.February), Equals, "February")
	c.Assert(getLongMonth(time.March), Equals, "March")
	c.Assert(getLongMonth(time.April), Equals, "April")
	c.Assert(getLongMonth(time.May), Equals, "May")
	c.Assert(getLongMonth(time.June), Equals, "June")
	c.Assert(getLongMonth(time.July), Equals, "July")
	c.Assert(getLongMonth(time.August), Equals, "August")
	c.Assert(getLongMonth(time.September), Equals, "September")
	c.Assert(getLongMonth(time.October), Equals, "October")
	c.Assert(getLongMonth(time.November), Equals, "November")
	c.Assert(getLongMonth(time.December), Equals, "December")

	c.Assert(getWeekdayNum(time.Unix(1448193600, 0)), Equals, 7)
}

func (s *TimeUtilSuite) TestAMPM(c *C) {
	c.Assert(getAMPMHour(time.Unix(1447838100, 0).UTC()), Equals, 9)
	c.Assert(getAMPM(time.Unix(1447838100, 0).UTC(), true), Equals, "AM")
	c.Assert(getAMPM(time.Unix(1447838100, 0).UTC(), false), Equals, "am")
	c.Assert(getAMPMHour(time.Unix(1447881300, 0).UTC()), Equals, 9)
	c.Assert(getAMPM(time.Unix(1447881300, 0).UTC(), true), Equals, "PM")
	c.Assert(getAMPM(time.Unix(1447881300, 0).UTC(), false), Equals, "pm")
}

func (s *TimeUtilSuite) TestTimezone(c *C) {
	ny, _ := time.LoadLocation("America/New_York")
	msk, _ := time.LoadLocation("Europe/Moscow")

	t := time.Unix(1447848900, 0)

	c.Assert(getTimezone(t.UTC().In(ny), false), Equals, "-0500")
	c.Assert(getTimezone(t.UTC().In(ny), true), Equals, "-05:00")
	c.Assert(getTimezone(t.UTC().In(msk), false), Equals, "+0300")
	c.Assert(getTimezone(t.UTC().In(msk), true), Equals, "+03:00")
}

func (s *TimeUtilSuite) TestDurationParsing(c *C) {
	c.Assert(ParseDuration(""), Equals, int64(0))
	c.Assert(ParseDuration("25s"), Equals, int64(25))
	c.Assert(ParseDuration("1m30s"), Equals, int64(90))
	c.Assert(ParseDuration("1h30m30s"), Equals, int64(5430))
	c.Assert(ParseDuration("1d3h30m30s"), Equals, int64(99030))
	c.Assert(ParseDuration("1w3d12h30m30s"), Equals, int64(909030))
	c.Assert(ParseDuration("10w"), Equals, int64(6048000))
	c.Assert(ParseDuration("180"), Equals, int64(180))
}
