package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
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
