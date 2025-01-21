package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"sort"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SliceSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SliceSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SliceSuite) TestCopy(c *C) {
	var ss []string

	c.Assert(Copy(ss), IsNil)

	c.Assert(Copy([]string{"A"}), DeepEquals, []string{"A"})
}

func (s *SliceSuite) TestIsEqual(c *C) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{1, 3, 2, 4}
	s4 := []int{1, 2, 3}

	var s5, s6 []int

	c.Assert(IsEqual(s1, nil), Equals, false)
	c.Assert(IsEqual(nil, s2), Equals, false)
	c.Assert(IsEqual(s1, s2), Equals, true)
	c.Assert(IsEqual(s1, s3), Equals, false)
	c.Assert(IsEqual(s1, s4), Equals, false)
	c.Assert(IsEqual(s5, s6), Equals, true)
}

func (s *SliceSuite) TestStringToInterface(c *C) {
	source := []string{"1", "2", "3"}

	c.Assert(StringToInterface(nil), IsNil)
	c.Assert(StringToInterface(source), DeepEquals, []any{"1", "2", "3"})
}

func (s *SliceSuite) TestIntToInterface(c *C) {
	source := []int{1, 2, 3}

	c.Assert(IntToInterface(nil), IsNil)
	c.Assert(IntToInterface(source), DeepEquals, []any{1, 2, 3})
}

func (s *SliceSuite) TestStringToError(c *C) {
	source := []string{"A", "B", "C"}

	c.Assert(StringToError(nil), IsNil)
	c.Assert(StringToError(source), DeepEquals,
		[]error{
			errors.New("A"),
			errors.New("B"),
			errors.New("C"),
		})
}

func (s *SliceSuite) TestErrorToString(c *C) {
	source := []error{
		errors.New("A"),
		errors.New("B"),
		errors.New("C"),
	}

	c.Assert(ErrorToString(nil), IsNil)
	c.Assert(ErrorToString(source), DeepEquals, []string{"A", "B", "C"})
}

func (s *SliceSuite) TestIndex(c *C) {
	source := []string{"1", "2", "3"}

	c.Assert(Index(source, "2"), Equals, 1)
	c.Assert(Index(source, "4"), Equals, -1)
	c.Assert(Index([]string{}, "1"), Equals, -1)
}

func (s *SliceSuite) TestContains(c *C) {
	source := []string{"1", "2", "3"}

	c.Assert(Contains(source, "1"), Equals, true)
	c.Assert(Contains(source, "4"), Equals, false)
	c.Assert(Contains([]string{}, "1"), Equals, false)
}

func (s *SliceSuite) TestExclude(c *C) {
	source := []string{"1", "2", "3", "4", "5", "6"}

	c.Assert(Exclude(source), DeepEquals, []string{"1", "2", "3", "4", "5", "6"})
	c.Assert(Exclude(source, "1"), DeepEquals, []string{"2", "3", "4", "5", "6"})
	c.Assert(Exclude(source, "1", "3", "6"), DeepEquals, []string{"2", "4", "5"})
	c.Assert(Exclude(source, "1", "2", "3", "4", "5", "6"), DeepEquals, []string{})
}

func (s *SliceSuite) TestDeduplicate(c *C) {
	source1 := []string{"1", "2", "2", "2", "3", "4", "5", "5", "6", "6"}
	source2 := []string{"abc", "ABC", "A", "B", "C", "abc", "D", "E", "ABC"}

	sort.Strings(source1)
	sort.Strings(source2)

	c.Assert(Deduplicate(source1), DeepEquals, []string{"1", "2", "3", "4", "5", "6"})
	c.Assert(Deduplicate(source2), DeepEquals, []string{"A", "ABC", "B", "C", "D", "E", "abc"})
	c.Assert(Deduplicate([]string{"1"}), DeepEquals, []string{"1"})
}

func (s *SliceSuite) TestJoin(c *C) {
	c.Assert(Join([]int{1, 2, 3, 4, 5}, ";"), Equals, "1;2;3;4;5")
	c.Assert(Join([]string{"test1", "test2", "test3"}, "--"), Equals, "test1--test2--test3")
	c.Assert(Join([]any{"test", 34, 12.50}, ","), Equals, "test,34,12.5")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SliceSuite) BenchmarkStringToInterface(c *C) {
	source := []string{"1", "2", "3"}

	for i := 0; i < c.N; i++ {
		StringToInterface(source)
	}
}

func (s *SliceSuite) BenchmarkIntToInterface(c *C) {
	source := []int{1, 2, 3}

	for i := 0; i < c.N; i++ {
		IntToInterface(source)
	}
}

func (s *SliceSuite) BenchmarkStringToError(c *C) {
	source := []string{"A", "B", "C"}

	for i := 0; i < c.N; i++ {
		StringToError(source)
	}
}

func (s *SliceSuite) BenchmarkErrorToString(c *C) {
	source := []error{
		errors.New("A"),
		errors.New("B"),
		errors.New("C"),
	}

	for i := 0; i < c.N; i++ {
		ErrorToString(source)
	}
}

func (s *SliceSuite) BenchmarkDeduplicate(c *C) {
	source := []string{"1", "2", "2", "2", "3", "4", "5", "5", "6", "6"}

	for i := 0; i < c.N; i++ {
		Deduplicate(source)
	}
}

func (s *SliceSuite) BenchmarkExclude(c *C) {
	source := []string{"1", "2", "3", "4", "5", "6"}

	for i := 0; i < c.N; i++ {
		Exclude(source, "1", "3", "6")
	}
}
