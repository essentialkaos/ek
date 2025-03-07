package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TimeUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TimeUtilSuite{})

func (s *TimeUtilSuite) TestPrettyDuration(c *C) {
	c.Assert(PrettyDuration(time.Duration(59000000000)), Equals, "59 seconds")
	c.Assert(PrettyDuration(120), Equals, "2 minutes")
	c.Assert(PrettyDuration(int16(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(int32(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(int64(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(uint16(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(uint32(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(uint64(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(uint(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(float32(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(float64(120)), Equals, "2 minutes")
	c.Assert(PrettyDuration(3720), Equals, "1 hour and 2 minutes")
	c.Assert(PrettyDuration(3720), Equals, "1 hour and 2 minutes")
	c.Assert(PrettyDuration(3720), Equals, "1 hour and 2 minutes")
	c.Assert(PrettyDuration(3720), Equals, "1 hour and 2 minutes")
	c.Assert(PrettyDuration(60), Equals, "1 minute")
	c.Assert(PrettyDuration(1370137), Equals, "2 weeks 1 day 20 hours 35 minutes and 37 seconds")
	c.Assert(PrettyDuration(0), Equals, "< 1 second")
	c.Assert(PrettyDuration("string"), Equals, "")
	c.Assert(PrettyDuration(time.Second/3), Equals, "333.3 ms")
	c.Assert(PrettyDuration(time.Second/3669), Equals, "272.6 μs")
	c.Assert(PrettyDuration(time.Second/366995), Equals, "2.72 μs")
	c.Assert(PrettyDuration(time.Second/36698888), Equals, "27 ns")
}

func (s *TimeUtilSuite) TestPrettyDurationSimple(c *C) {
	c.Assert(PrettyDurationSimple(true), Equals, "")
	c.Assert(PrettyDurationSimple(0), Equals, "< 1 second")
	c.Assert(PrettyDurationSimple(1), Equals, "1 second")
	c.Assert(PrettyDurationSimple(123), Equals, "2 minutes")
	c.Assert(PrettyDurationSimple(12134), Equals, "3 hours")
	c.Assert(PrettyDurationSimple(112345), Equals, "1 day")
	c.Assert(PrettyDurationSimple(4523412), Equals, "52 days")
}

func (s *TimeUtilSuite) TestPrettyDurationInDays(c *C) {
	c.Assert(PrettyDurationInDays("ABC"), Equals, "")
	c.Assert(PrettyDurationInDays(120), Equals, "1 day")
	c.Assert(PrettyDurationInDays(7200), Equals, "1 day")
	c.Assert(PrettyDurationInDays(90000), Equals, "1 day")
	c.Assert(PrettyDurationInDays(1296000), Equals, "15 days")
}

func (s *TimeUtilSuite) TestShortDuration(c *C) {
	c.Assert(ShortDuration(time.Duration(0)), Equals, "0:00")
	c.Assert(ShortDuration(time.Duration(3546*time.Millisecond), false), Equals, "0:03")
	c.Assert(ShortDuration(time.Duration(3546*time.Millisecond), true), Equals, "0:03.546")
	c.Assert(ShortDuration(time.Duration(59000000000)), Equals, "0:59")
	c.Assert(ShortDuration(60), Equals, "1:00")
	c.Assert(ShortDuration(120), Equals, "2:00")
	c.Assert(ShortDuration(2235), Equals, "37:15")
	c.Assert(ShortDuration(1234567), Equals, "342:56:07")
	c.Assert(ShortDuration(int16(1234)), Equals, "20:34")
	c.Assert(ShortDuration(int32(1234)), Equals, "20:34")
	c.Assert(ShortDuration(int64(1234)), Equals, "20:34")
	c.Assert(ShortDuration(uint16(1234)), Equals, "20:34")
	c.Assert(ShortDuration(uint32(1234)), Equals, "20:34")
	c.Assert(ShortDuration(uint64(1234)), Equals, "20:34")
	c.Assert(ShortDuration(uint(1234)), Equals, "20:34")
	c.Assert(ShortDuration(float32(1234)), Equals, "20:34")
	c.Assert(ShortDuration(float64(1234)), Equals, "20:34")
	c.Assert(ShortDuration("ABCD"), Equals, "")
}

func (s *TimeUtilSuite) TestMiniDuration(c *C) {
	c.Assert(MiniDuration(89*time.Hour, ""), Equals, "4d")
	c.Assert(MiniDuration(time.Duration(0)), Equals, "0 ns")
	c.Assert(MiniDuration(89*time.Hour), Equals, "4 d")
	c.Assert(MiniDuration(15*time.Hour), Equals, "15 h")
	c.Assert(MiniDuration(80*time.Second), Equals, "1.3 m")
	c.Assert(MiniDuration(3*time.Second), Equals, "3 s")
	c.Assert(MiniDuration(3*time.Millisecond), Equals, "3 ms")
	c.Assert(MiniDuration(3*time.Microsecond), Equals, "3 μs")
	c.Assert(MiniDuration(3*time.Nanosecond), Equals, "3 ns")
}

func (s *TimeUtilSuite) TestDurationToSeconds(c *C) {
	c.Assert(SecondsToDuration(1), Equals, time.Second)
	c.Assert(SecondsToDuration(1.5), Equals, 1500*time.Millisecond)
	c.Assert(SecondsToDuration(3600), Equals, time.Hour)
}

func (s *TimeUtilSuite) TestSecondsToDuration(c *C) {
	c.Assert(DurationToSeconds(time.Second/4), Equals, 0.25)
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
	d, _ := ParseDuration("")
	c.Assert(d, Equals, time.Duration(0))

	d, _ = ParseDuration("25s")
	c.Assert(d, Equals, time.Duration(25)*time.Second)

	d, _ = ParseDuration("1m30s")
	c.Assert(d, Equals, time.Duration(90)*time.Second)

	d, _ = ParseDuration("1h30m30s")
	c.Assert(d, Equals, time.Duration(5430)*time.Second)

	d, _ = ParseDuration("1d3h30m30s")
	c.Assert(d, Equals, time.Duration(99030)*time.Second)

	d, _ = ParseDuration("1w3d12h30m30s")
	c.Assert(d, Equals, time.Duration(909030)*time.Second)

	d, _ = ParseDuration("10w")
	c.Assert(d, Equals, time.Duration(6048000)*time.Second)

	d, _ = ParseDuration("180")
	c.Assert(d, Equals, time.Duration(180)*time.Second)

	_, err := ParseDuration("180k")
	c.Assert(err, NotNil)

	_, err = ParseDuration("wm")
	c.Assert(err, NotNil)

	_, err = ParseDuration("9999999999999999999999999999999999999s")
	c.Assert(err, NotNil)

	_, err = ParseDuration("1h35m56")
	c.Assert(err, NotNil)

	d, _ = ParseDuration("30", 's')
	c.Assert(d, Equals, time.Duration(30)*time.Second)

	d, _ = ParseDuration("25", 'm')
	c.Assert(d, Equals, time.Duration(1500)*time.Second)

	d, _ = ParseDuration("14", 'h')
	c.Assert(d, Equals, time.Duration(50400)*time.Second)

	d, _ = ParseDuration("5", 'd')
	c.Assert(d, Equals, time.Duration(432000)*time.Second)

	d, _ = ParseDuration("2", 'w')
	c.Assert(d, Equals, time.Duration(1209600)*time.Second)

	_, err = ParseDuration("9999999999999999999999999999999999999", 's')
	c.Assert(err, NotNil)
}

func (s *TimeUtilSuite) TestPeriod(c *C) {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2023, 6, 15, 18, 45, 30, 0, time.Local),
	}

	c.Assert(p.Contains(time.Date(2021, 1, 1, 9, 30, 15, 0, time.Local)), Equals, false)
	c.Assert(p.Contains(time.Date(2023, 6, 15, 18, 50, 30, 0, time.Local)), Equals, false)
	c.Assert(p.Contains(time.Date(2021, 1, 1, 12, 30, 16, 0, time.Local)), Equals, true)
	c.Assert(p.Contains(time.Date(2023, 6, 15, 18, 45, 29, 0, time.Local)), Equals, true)

	c.Assert(p.Duration().String(), Equals, "21486h15m15s")
	c.Assert(p.Seconds(), Equals, 77350515)
	c.Assert(p.Minutes(), Equals, 1289175)
	c.Assert(p.Hours(), Equals, 21486)
	c.Assert(p.Days(), Equals, 895)
	c.Assert(p.Weeks(), Equals, 127)
	c.Assert(p.Years(), Equals, 2)

	c.Assert(Period{}.String(), Equals, "0ns")
	c.Assert(p.String(), Equals, "895d")
}

func (s *TimeUtilSuite) TestHelpers(c *C) {
	d := time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local)

	c.Assert(PrevDay(d), DeepEquals, time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local))
	c.Assert(PrevWeek(d, time.Monday), DeepEquals, time.Date(2020, 12, 21, 0, 0, 0, 0, time.Local))
	c.Assert(PrevMonth(d), DeepEquals, time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local))
	c.Assert(PrevYear(d), DeepEquals, time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))
	c.Assert(NextDay(d), DeepEquals, time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local))
	c.Assert(NextWeek(d, time.Monday), DeepEquals, time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local))
	c.Assert(NextMonth(d), DeepEquals, time.Date(2021, 2, 1, 0, 0, 0, 0, time.Local))
	c.Assert(NextYear(d), DeepEquals, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local))

	d = time.Date(2021, 8, 1, 12, 30, 15, 0, time.Local)
	c.Assert(PrevWorkday(d), DeepEquals, time.Date(2021, 7, 30, 0, 0, 0, 0, time.Local))
	c.Assert(PrevWeekend(d), DeepEquals, time.Date(2021, 7, 31, 0, 0, 0, 0, time.Local))
	c.Assert(NextWorkday(d), DeepEquals, time.Date(2021, 8, 2, 0, 0, 0, 0, time.Local))
	c.Assert(NextWeekend(d), DeepEquals, time.Date(2021, 8, 7, 0, 0, 0, 0, time.Local))

	d = time.Time{}

	c.Assert(StartOfHour(d).IsZero(), Equals, true)
	c.Assert(StartOfDay(d).IsZero(), Equals, true)
	c.Assert(StartOfWeek(d, time.Monday).IsZero(), Equals, true)
	c.Assert(StartOfMonth(d).IsZero(), Equals, true)
	c.Assert(StartOfYear(d).IsZero(), Equals, true)
	c.Assert(EndOfHour(d).IsZero(), Equals, true)
	c.Assert(EndOfDay(d).IsZero(), Equals, true)
	c.Assert(EndOfWeek(d, time.Monday).IsZero(), Equals, true)
	c.Assert(EndOfMonth(d).IsZero(), Equals, true)
	c.Assert(EndOfYear(d).IsZero(), Equals, true)

	d = time.Date(2021, 8, 13, 12, 30, 15, 0, time.Local)
	c.Assert(StartOfHour(d), DeepEquals, time.Date(2021, 8, 13, 12, 0, 0, 0, time.Local))
	c.Assert(StartOfDay(d), DeepEquals, time.Date(2021, 8, 13, 0, 0, 0, 0, time.Local))
	c.Assert(StartOfWeek(d, time.Monday), DeepEquals, time.Date(2021, 8, 9, 0, 0, 0, 0, time.Local))
	c.Assert(StartOfMonth(d), DeepEquals, time.Date(2021, 8, 1, 0, 0, 0, 0, time.Local))
	c.Assert(StartOfYear(d), DeepEquals, time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local))
	c.Assert(EndOfHour(d), DeepEquals, time.Date(2021, 8, 13, 12, 59, 59, 999999999, time.Local))
	c.Assert(EndOfDay(d), DeepEquals, time.Date(2021, 8, 13, 23, 59, 59, 999999999, time.Local))
	c.Assert(EndOfWeek(d, time.Monday), DeepEquals, time.Date(2021, 8, 15, 23, 59, 59, 999999999, time.Local))
	c.Assert(EndOfMonth(d), DeepEquals, time.Date(2021, 8, 31, 23, 59, 59, 999999999, time.Local))
	c.Assert(EndOfYear(d), DeepEquals, time.Date(2021, 12, 31, 23, 59, 59, 999999999, time.Local))

	y := time.Now().In(time.Local).Year()
	c.Assert(FromISOWeek(0, 0, time.Local), DeepEquals, time.Date(y, 1, 1, 0, 0, 0, 0, time.Local))
	c.Assert(FromISOWeek(100, 2021, time.Local), DeepEquals, time.Date(2021, 12, 31, 0, 0, 0, 0, time.Local))
	c.Assert(FromISOWeek(23, 2021, time.Local), DeepEquals, time.Date(2021, 6, 4, 0, 0, 0, 0, time.Local))

	d = time.Date(2021, 8, 1, 12, 30, 15, 0, time.Local)
	c.Assert(IsWeekend(d), Equals, true)
	d = d.AddDate(0, 0, 3)
	c.Assert(IsWeekend(d), Equals, false)

	d = time.Date(2030, 1, 1, 12, 0, 0, 0, time.Local)
	c.Assert(Until(d, DAY), Not(Equals), 0)

	d = time.Date(2012, 1, 1, 12, 0, 0, 0, time.Local)
	c.Assert(Since(d, DAY), Not(Equals), 0)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TimeUtilSuite) BenchmarkParseDuration(c *C) {
	for range c.N {
		ParseDuration("1w2d3h10m12s")
	}
}

func (s *TimeUtilSuite) BenchmarkFormat(c *C) {
	ts := time.Now()

	for range c.N {
		Format(ts, "%Y/%m/%d %H:%M:%S")
	}
}
