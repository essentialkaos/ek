package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type SortSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SortSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SortSuite) TestVersionSorting(c *C) {
	v1 := []string{"1", "2.1", "2", "2.3.4", "1.3", "1.6.5", "2.3.3", "14.0", "6"}

	Versions(v1)

	c.Assert(v1, DeepEquals, []string{"1", "1.3", "1.6.5", "2", "2.1", "2.3.3", "2.3.4", "6", "14.0"})

	v2 := []string{"1-2", "2", "2.2.3", "2.2.3", "1"}

	Versions(v2)

	c.Assert(v2, DeepEquals, []string{"1", "1-2", "2", "2.2.3", "2.2.3"})
}

func (s *SortSuite) TestStringSorting(c *C) {
	s1 := []string{"Apple", "auto", "image", "Monica", "7", "flower", "moon"}
	s2 := []string{"Apple", "auto", "image", "Monica", "7", "flower", "moon"}

	Strings(s1, false)
	Strings(s2, true)

	c.Assert(s1, DeepEquals, []string{"7", "Apple", "Monica", "auto", "flower", "image", "moon"})
	c.Assert(s2, DeepEquals, []string{"7", "Apple", "auto", "flower", "image", "Monica", "moon"})
}

func (s *SortSuite) TestNaturalSorting(c *C) {
	s0 := []string{"1"}

	StringsNatural(s0)

	s1 := []string{"abc5", "abc1", "abc01", "ab", "abc10", "abc2"}

	StringsNatural(s1)

	c.Assert(s1, DeepEquals, []string{"ab", "abc1", "abc01", "abc2", "abc5", "abc10"})

	c.Assert(NaturalLess("0", "00"), Equals, true)
	c.Assert(NaturalLess("00", "0"), Equals, false)
	c.Assert(NaturalLess("aa", "ab"), Equals, true)
	c.Assert(NaturalLess("ab", "abc"), Equals, true)
	c.Assert(NaturalLess("abc", "ad"), Equals, true)
	c.Assert(NaturalLess("ab1", "ab2"), Equals, true)
	c.Assert(NaturalLess("ab1c", "ab1c"), Equals, false)
	c.Assert(NaturalLess("ab12", "abc"), Equals, true)
	c.Assert(NaturalLess("ab2a", "ab10"), Equals, true)
	c.Assert(NaturalLess("a0001", "a0000001"), Equals, true)
	c.Assert(NaturalLess("a10", "abcdefgh2"), Equals, true)
	c.Assert(NaturalLess("аб2аб", "аб10аб"), Equals, true)
	c.Assert(NaturalLess("2аб", "3аб"), Equals, true)
	c.Assert(NaturalLess("a1b", "a01b"), Equals, true)
	c.Assert(NaturalLess("a01b", "a1b"), Equals, false)
	c.Assert(NaturalLess("ab01b", "ab010b"), Equals, true)
	c.Assert(NaturalLess("ab010b", "ab01b"), Equals, false)
	c.Assert(NaturalLess("a001b01", "a01b001"), Equals, false)
	c.Assert(NaturalLess("a1", "a1x"), Equals, true)
	c.Assert(NaturalLess("1ax", "1b"), Equals, true)
	c.Assert(NaturalLess("1b", "1ax"), Equals, false)
	c.Assert(NaturalLess("082", "83"), Equals, true)
	c.Assert(NaturalLess("083a", "9a"), Equals, false)
	c.Assert(NaturalLess("9a", "083a"), Equals, true)
}
