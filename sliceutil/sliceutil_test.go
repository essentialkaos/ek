package sliceutil

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

type SliceSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SliceSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SliceSuite) TestStr2Interface(c *C) {
	source := []string{"1", "2", "3"}
	result := StringToInterface(source)

	c.Assert(result, DeepEquals, []interface{}{"1", "2", "3"})
}

func (s *SliceSuite) TestInt2Interface(c *C) {
	source := []int{1, 2, 3}
	result := IntToInterface(source)

	c.Assert(result, DeepEquals, []interface{}{1, 2, 3})
}

func (s *SliceSuite) TestContains(c *C) {
	source := []string{"1", "2", "3"}

	c.Assert(Contains(source, "1"), Equals, true)
	c.Assert(Contains(source, "4"), Equals, false)
	c.Assert(Contains([]string{}, "1"), Equals, false)
}
