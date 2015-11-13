package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type FmtUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FmtUtilSuite{})

func (s *FmtUtilSuite) TestPretyNum(c *C) {
	c.Assert(PrettyNum(999), Equals, "999")
	c.Assert(PrettyNum(1000), Equals, "1,000")
	c.Assert(PrettyNum(1234567890), Equals, "1,234,567,890")
	c.Assert(PrettyNum(0), Equals, "0")
	c.Assert(PrettyNum(2500.50), Equals, "2,500.50")
	c.Assert(PrettyNum(1.23), Equals, "1.23")
	c.Assert(PrettyNum(-1000), Equals, "-1,000")
}

func (s *FmtUtilSuite) TestPretySize(c *C) {
	c.Assert(PrettySize(0), Equals, "0B")
	c.Assert(PrettySize(345), Equals, "345B")
	c.Assert(PrettySize(1025), Equals, "1KB")
	c.Assert(PrettySize(3000125), Equals, "2.86MB")
	c.Assert(PrettySize(1024*1024), Equals, "1MB")
	c.Assert(PrettySize(1024*1024*1024), Equals, "1GB")
	c.Assert(PrettySize(1024*1024*1024*1024), Equals, "1TB")
	c.Assert(PrettySize(1052500000), Equals, "1003.7MB")
}

func (s *FmtUtilSuite) TestParseSize(c *C) {
	c.Assert(ParseSize("1 MB"), Equals, uint64(1024*1024))
	c.Assert(ParseSize("5gB"), Equals, uint64(5*1024*1024*1024))
	c.Assert(ParseSize("13kb"), Equals, uint64(13*1024))
	c.Assert(ParseSize("512"), Equals, uint64(512))
	c.Assert(ParseSize("0"), Equals, uint64(0))

	c.Assert(ParseSize(PrettySize(345)), Equals, uint64(345))
	c.Assert(ParseSize(PrettySize(1025)), Equals, uint64(1024))
	c.Assert(ParseSize(PrettySize(1024*1024)), Equals, uint64(1024*1024))
}

func (s *FmtUtilSuite) TestFloat(c *C) {
	c.Assert(Float(1.0), Equals, 1.0)
	c.Assert(Float(0.1), Equals, 0.1)
	c.Assert(Float(0.01), Equals, 0.01)
	c.Assert(Float(0.001), Equals, 0.0)
	c.Assert(Float(0.0001), Equals, 0.0)
}

func (s *FmtUtilSuite) TestWrap(c *C) {
	input := `The attack exploited two zero-day vulnerabilities, one in Microsoft's Internet Explorer and the other in Adobe's Flash Player, 
Invincea and iSight Partners said in their joint report released Tuesday. Adobe fixed the flaw back in December and Microsoft 
updated Internet Explorer as part of its Patch Tuesday release.

The cyber-espionage campaign appeared to last only a few days, but iSight and Invincea did not rule out the possibility of 
the campaign lasting a longer period of time.

The malware infection was inside the “Thought of the Day” Flash widget which appears whenever users try to access a Forbes.com 
page. Visitors didn't need to do anything other than to try to load Forbes.com in their browser to get infected. The 
demographics of the typical visitor to Forbes.com—senior executives, managers, and other professionals working for major 
corporations—indicate this campaign focused on cyber-espionage, not cybercrime, said Stephen Ward, an analyst with iSight 
Partners. Watering hole attacks are insidious because it wouldn't occur to anyone that these sites could be infected.
`
	result := `  The attack exploited two zero-day 
  vulnerabilities, one in Microsoft's 
  Internet Explorer and the other in 
  Adobe's Flash Player, Invincea and 
  iSight Partners said in their joint 
  report released Tuesday. Adobe fixed 
  the flaw back in December and 
  Microsoft updated Internet Explorer as 
  part of its Patch Tuesday release.

  Thecyber-espionage campaign appeared 
  to last only a few days, but iSight 
  and Invincea did not rule out the 
  possibility of the campaign lasting a 
  longer period of time.

  Themalware infection was inside the 
  “Thought of the Day” Flash widget 
  which appears whenever users try to 
  access a Forbes.com page. Visitors 
  didn't need to do anything other than 
  to try to load Forbes.com in their 
  browser to get infected. The 
  demographics of the typical visitor to 
  Forbes.com—senior executives, 
  managers, and other professionals 
  working for major 
  corporations—indicate this campaign 
  focused on cyber-espionage, not 
  cybercrime, said Stephen Ward, an 
  analyst with iSight Partners. Watering 
  hole attacks are insidious because it 
  wouldn't occur to anyone that these 
  sites could be infected.`

	c.Assert(Wrap(input, "  ", 40), Equals, result)
}

func (s *FmtUtilSuite) TestPluralize(c *C) {
	c.Assert(Pluralize(1, "test", "tests"), Equals, "1 test")
	c.Assert(Pluralize(2, "test", "tests"), Equals, "2 tests")
	c.Assert(Pluralize(100, "test", "tests"), Equals, "100 tests")
}
