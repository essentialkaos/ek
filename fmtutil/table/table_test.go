package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/essentialkaos/check"
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
	t := NewTable("abc")

	c.Assert(t.SetHeaders("abc", "def", "123"), NotNil)
	c.Assert(t.Headers, HasLen, 3)
}

func (s *TableSuite) TestRenderHeaders(c *C) {
	t := NewTable("abc", "def")
	t.RenderHeaders()
}

func (s *TableSuite) TestSetSizes(c *C) {
	t := NewTable()

	c.Assert(t.SetSizes(10, 20, 30), NotNil)
	c.Assert(t.Sizes, HasLen, 3)
}

func (s *TableSuite) TestSetAlignments(c *C) {
	t := NewTable()

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
	t := NewTable()

	c.Assert(t.Add(), NotNil)
	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.data, HasLen, 1)
}

func (s *TableSuite) TestPrint(c *C) {
	t := NewTable("ABCD", "ABCDEF", "ABCD")

	c.Assert(t.Print(10, "abc", 3.14), NotNil)
	c.Assert(t.Print(10, "abµ", 3.14, 400), NotNil)
	c.Assert(t.Border(), NotNil)
}

func (s *TableSuite) TestSeparator(c *C) {
	t := NewTable()
	c.Assert(t.Separator(), NotNil)

	t = NewTable()
	c.Assert(t.Add(1), NotNil)
	c.Assert(t.Add(2), NotNil)
	c.Assert(t.Separator(), NotNil)
	c.Assert(t.Add(3), NotNil)
	c.Assert(t.Add(4), NotNil)
	t.Render()
}

func (s *TableSuite) TestRender(c *C) {
	t := NewTable()

	c.Assert(t.Render(), NotNil)
	c.Assert(t.Print(), NotNil)

	t = NewTable()

	t.BorderSymbol = "—"
	t.SeparatorSymbol = "+"
	t.ColumnSeparatorSymbol = ":"
	t.HeaderColorTag = "{r}"
	t.BorderColorTag = "{b}"
	t.SeparatorColorTag = "{g}"
	t.HeaderCapitalize = true

	c.Assert(t.Render(), NotNil)

	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.Add(11, "簡単な例", 2.28), NotNil)
	c.Assert(t.Add(11, "ABC", strings.Repeat("ABC123_", 20)), NotNil)

	c.Assert(t.Render(), NotNil)

	fmt.Println()

	t.Width = 80

	t.SetHeaders("id", "name", "price")
	t.SetSizes(12, 12, 12)
	t.SetAlignments(ALIGN_LEFT, ALIGN_CENTER, ALIGN_RIGHT)

	c.Assert(t.Add(10, "abc", 3.14), NotNil)
	c.Assert(t.Add(11, "{g}ABC{!}", 2.28, 400), NotNil)

	c.Assert(t.Render(), NotNil)

	HeaderCapitalize = false
}

func (s *TableSuite) TestPrintWithoutInit(c *C) {
	t := NewTable()
	t.Print("abcd", 1234)
}

func (s *TableSuite) TestNil(c *C) {
	var t *Table

	c.Assert(t.SetHeaders("abc"), IsNil)
	c.Assert(t.SetSizes(10, 20, 30), IsNil)
	c.Assert(t.SetAlignments(0, 1, 2), IsNil)
	c.Assert(t.Add(10, "abc", 3.14), IsNil)
	c.Assert(t.Print(10, "abc", 3.14), IsNil)
	c.Assert(t.Render(), IsNil)
	c.Assert(t.Border(), IsNil)
	c.Assert(t.Separator(), IsNil)
	c.Assert(t.HasData(), Equals, false)

	c.Assert(func() { t.RenderHeaders() }, NotPanics)
}

func (s *TableSuite) TestAuxi(c *C) {
	t := &Table{Sizes: []int{1, 2, 3, 4}, Width: 88}

	c.Assert(getColumnsNum(t), Equals, len(t.Sizes))

	t = &Table{columnSizes: []int{4, 4}, Width: 88}
	renderRowData(t, []string{"ABCDABCDABCD", "ABCDABCDABCD"}, 2)

	setColumnsSizes(t, 3)

	t = &Table{columnSizes: []int{3, 3, 3}}

	c.Assert(getTableWidth(t), Equals, 0)
	c.Assert(getSeparatorSize(t), Equals, 16)
}
