package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TableSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TableSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TableSuite) TestNewTable(c *C) {
	c.Assert(NewTable(), NotNil)
}

func (s *TableSuite) TestSetHeaders(c *C) {
	var t *Table

	c.Assert(t.SetHeaders("abc"), IsNil)

	t = NewTable("abc")

	c.Assert(t.SetHeaders("abc", "def", "123"), NotNil)
	c.Assert(t.Headers, HasLen, 3)
}

func (s *TableSuite) TestRenderHeaders(c *C) {
	t := NewTable("abc", "def")
	t.RenderHeaders()
}

func (s *TableSuite) TestSetSizes(c *C) {
	var t *Table

	c.Assert(t.SetSizes(10, 20, 30), IsNil)

	t = NewTable()

	c.Assert(t.SetSizes(10, 20, 30), NotNil)
	c.Assert(t.Sizes, HasLen, 3)
}

func (s *TableSuite) TestSetAlignments(c *C) {
	var t *Table

	c.Assert(t.SetAlignments(0, 1, 2), IsNil)

	t = NewTable()

	c.Assert(t.SetAlignments(0, 1, 2), NotNil)
	c.Assert(t.Alignment, HasLen, 3)
}

func (s *TableSuite) TestHasData(c *C) {
	t := NewTable("1", "2", "3")

	c.Assert(t.HasData(), Equals, false)

	c.Assert(t.Add(10, "abc", 3.14), NotNil)

	c.Assert(t.HasData(), Equals, true)
}

func (s *TableSuite) TestAdd(c *C) {
	var t *Table

	c.Assert(t.Add(10, "abc", 3.14), IsNil)

	t = NewTable()

	c.Assert(t.Add(), NotNil)
	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.data, HasLen, 1)
}

func (s *TableSuite) TestPrint(c *C) {
	var t *Table

	c.Assert(t.Print(10, "abc", 3.14), IsNil)

	t = NewTable("ABCD", "ABCDEF", "ABCD")

	c.Assert(t.Print(10, "abc", 3.14), NotNil)
	c.Assert(t.Print(10, "abµ", 3.14, 400), NotNil)
}

func (s *TableSuite) TestSeparator(c *C) {
	var t *Table

	c.Assert(t.Separator(), IsNil)

	t = NewTable()

	c.Assert(t.Separator(), NotNil)
}

func (s *TableSuite) TestRender(c *C) {
	var t *Table

	c.Assert(t.Render(), IsNil)

	t = NewTable()

	c.Assert(t.Render(), NotNil)
	c.Assert(t.Print(), NotNil)

	t = NewTable()

	c.Assert(t.Render(), NotNil)

	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.Add(11, "ABC", 2.28), NotNil)
	c.Assert(t.Add(11, "ABC", strings.Repeat("ABC123_", 20)), NotNil)

	c.Assert(t.Render(), NotNil)

	t.SetHeaders("id", "name", "price")
	t.SetSizes(12, 12, 12)
	t.SetAlignments(ALIGN_LEFT, ALIGN_CENTER, ALIGN_RIGHT)

	HeaderCapitalize = true

	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.Add(11, "{g}ABC{!}", 2.28, 400), NotNil)

	c.Assert(t.Render(), NotNil)

	HeaderCapitalize = false
}

func (s *TableSuite) TestPrintWithoutInit(c *C) {
	t := NewTable()
	t.Print("abcd", 1234)
}

func (s *TableSuite) TestAuxi(c *C) {
	t := &Table{Sizes: []int{1, 2, 3, 4}}

	c.Assert(getColumnsNum(t), Equals, len(t.Sizes))

	t = &Table{columnSizes: []int{4, 4}}
	renderRowData(t, []string{"ABCDABCDABCD", "ABCDABCDABCD"}, 2)

	MaxWidth = 80

	setColumnsSizes(t, 3)
	c.Assert(t.columnSizes[0], Equals, 23)
	c.Assert(t.columnSizes[1], Equals, 23)
	c.Assert(t.columnSizes[2], Equals, 25)

	MaxWidth = 0
}
