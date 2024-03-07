package value

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ValuesSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ValuesSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ValuesSuite) TestParseInt64(c *C) {
	c.Assert(ParseInt64(""), Equals, int64(0))
	c.Assert(ParseInt64("", 234), Equals, int64(234))

	c.Assert(ParseInt64("123"), Equals, int64(123))
	c.Assert(ParseInt64("0xFF"), Equals, int64(255))

	c.Assert(ParseInt64("ABCD"), Equals, int64(0))
	c.Assert(ParseInt64("0xZZ"), Equals, int64(0))
}

func (s *ValuesSuite) TestParseInt(c *C) {
	c.Assert(ParseInt(""), Equals, 0)
	c.Assert(ParseInt("", 234), Equals, 234)

	c.Assert(ParseInt("123"), Equals, 123)
	c.Assert(ParseInt("0xFF"), Equals, 255)

	c.Assert(ParseInt("ABCD"), Equals, 0)
	c.Assert(ParseInt("0xZZ"), Equals, 0)

	maxCheckFail = true
	c.Assert(ParseInt("9999"), Equals, 0)
	maxCheckFail = false
}

func (s *ValuesSuite) TestParseUint(c *C) {
	c.Assert(ParseUint(""), Equals, uint(0))
	c.Assert(ParseUint("", 234), Equals, uint(234))

	c.Assert(ParseUint("123"), Equals, uint(123))
	c.Assert(ParseUint("0xFF"), Equals, uint(255))

	c.Assert(ParseUint("ABCD"), Equals, uint(0))
	c.Assert(ParseUint("0xZZ"), Equals, uint(0))

	maxCheckFail = true
	c.Assert(ParseUint("9999"), Equals, uint(0))
	maxCheckFail = false
}

func (s *ValuesSuite) TestParseUint64(c *C) {
	c.Assert(ParseUint64(""), Equals, uint64(0))
	c.Assert(ParseUint64("", 234), Equals, uint64(234))

	c.Assert(ParseUint64("123"), Equals, uint64(123))
	c.Assert(ParseUint64("0xFF"), Equals, uint64(255))

	c.Assert(ParseUint64("ABCD"), Equals, uint64(0))
	c.Assert(ParseUint64("0xZZ"), Equals, uint64(0))
}

func (s *ValuesSuite) TestParseFloat(c *C) {
	c.Assert(ParseFloat(""), Equals, 0.0)
	c.Assert(ParseFloat("", 234.0), Equals, 234.0)

	c.Assert(ParseFloat("123"), Equals, 123.0)

	c.Assert(ParseFloat("ABCD"), Equals, 0.0)
}

func (s *ValuesSuite) TestParseBool(c *C) {
	c.Assert(ParseBool(""), Equals, false)
	c.Assert(ParseBool("", true), Equals, true)

	c.Assert(ParseBool("0"), Equals, false)
	c.Assert(ParseBool("No"), Equals, false)
	c.Assert(ParseBool("False"), Equals, false)

	c.Assert(ParseBool("true"), Equals, true)
	c.Assert(ParseBool("abcd"), Equals, true)
}

func (s *ValuesSuite) TestParseMode(c *C) {
	c.Assert(ParseMode(""), Equals, os.FileMode(0))
	c.Assert(ParseMode("", 0600), Equals, os.FileMode(0600))

	c.Assert(ParseMode("0600"), Equals, os.FileMode(0600))
	c.Assert(ParseMode("600"), Equals, os.FileMode(0600))

	c.Assert(ParseMode("ABCD"), Equals, os.FileMode(0))
}

func (s *ValuesSuite) TestParseDuration(c *C) {
	c.Assert(ParseDuration("", time.Minute), Equals, time.Duration(0))
	c.Assert(ParseDuration("", time.Minute, time.Hour), Equals, time.Hour)

	c.Assert(ParseDuration("3", time.Minute, time.Hour), Equals, 3*time.Minute)

	c.Assert(ParseDuration("ABCD", time.Minute), Equals, time.Duration(0))
}

func (s *ValuesSuite) TestParseTimeDuration(c *C) {
	c.Assert(ParseTimeDuration(""), Equals, time.Duration(0))
	c.Assert(ParseTimeDuration("", time.Hour), Equals, time.Hour)

	c.Assert(ParseTimeDuration("7s"), Equals, 7*time.Second)
	c.Assert(ParseTimeDuration("6m"), Equals, 6*time.Minute)
	c.Assert(ParseTimeDuration("3h"), Equals, 3*time.Hour)
	c.Assert(ParseTimeDuration("2d"), Equals, 48*time.Hour)
	c.Assert(ParseTimeDuration("3w"), Equals, 3*7*24*time.Hour)

	c.Assert(ParseTimeDuration("ABCD"), Equals, time.Duration(0))
}

func (s *ValuesSuite) TestParseTimestamp(c *C) {
	c.Assert(ParseTimestamp("").IsZero(), Equals, true)
	c.Assert(ParseTimestamp("", time.Unix(1709627257, 0)).Unix(), Equals, int64(1709627257))

	c.Assert(ParseTimestamp("1709627257").Unix(), Equals, int64(1709627257))

	c.Assert(ParseTimestamp("ABCD").IsZero(), Equals, true)
}

func (s *ValuesSuite) TestParseTimezone(c *C) {
	l, _ := time.LoadLocation("America/Los_Angeles")

	c.Assert(ParseTimezone(""), IsNil)
	c.Assert(ParseTimezone("", l), NotNil)

	c.Assert(ParseTimezone("Europe/Vienna"), NotNil)

	c.Assert(ParseTimezone("ABCD"), IsNil)
}

func (s *ValuesSuite) TestParseList(c *C) {
	c.Assert(ParseList(""), IsNil)
	c.Assert(ParseList("", []string{"A", "B"}), DeepEquals, []string{"A", "B"})

	c.Assert(ParseList("A, B"), DeepEquals, []string{"A", "B"})
}
