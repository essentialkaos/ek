package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"testing"

	. "pkg.re/check.v1"
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

func (s *SliceSuite) TestString2Error(c *C) {
	source := []string{"A", "B", "C"}
	result := StringToError(source)

	c.Assert(result, DeepEquals,
		[]error{
			errors.New("A"),
			errors.New("B"),
			errors.New("C"),
		})
}

func (s *SliceSuite) TestError2String(c *C) {
	source := []error{
		errors.New("A"),
		errors.New("B"),
		errors.New("C"),
	}

	result := ErrorToString(source)

	c.Assert(result, DeepEquals, []string{"A", "B", "C"})
}

func (s *SliceSuite) TestContains(c *C) {
	source := []string{"1", "2", "3"}

	c.Assert(Contains(source, "1"), Equals, true)
	c.Assert(Contains(source, "4"), Equals, false)
	c.Assert(Contains([]string{}, "1"), Equals, false)
}

func (s *SliceSuite) TestExclude(c *C) {
	source := []string{"1", "2", "3", "4", "5", "6"}

	c.Assert(Exclude(source, []string{"1"}), DeepEquals, []string{"2", "3", "4", "5", "6"})
	c.Assert(Exclude(source, []string{"1", "3", "6"}), DeepEquals, []string{"2", "4", "5"})
	c.Assert(Exclude(source, []string{"1", "2", "3", "4", "5", "6"}), IsNil)
}

func (s *SliceSuite) TestDeduplicate(c *C) {
	source := []string{"1", "2", "2", "2", "3", "4", "5", "5", "6", "6"}

	c.Assert(Deduplicate(source), DeepEquals, []string{"1", "2", "3", "4", "5", "6"})
	c.Assert(Deduplicate([]string{"1"}), DeepEquals, []string{"1"})
}
