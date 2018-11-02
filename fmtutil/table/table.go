// Package table contains methods and structs for rendering data as a table
package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"pkg.re/essentialkaos/ek.v10/fmtc"
	"pkg.re/essentialkaos/ek.v10/mathutil"
	"pkg.re/essentialkaos/ek.v10/strutil"
	"pkg.re/essentialkaos/ek.v10/terminal/window"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Column alignment flags
const (
	ALIGN_LEFT   uint8 = 0
	ALIGN_CENTER       = 1
	ALIGN_RIGHT        = 2
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Table is struct which can be used for table rendering
type Table struct {
	Sizes     []int    // Custom columns sizes
	Headers   []string // Slice with headers
	Alignment []uint8  // Columns alignment

	// Slice with data
	data [][]string

	// Separator cache
	separator string

	// Flag will be set if header is rendered
	headerShown bool

	// Slice with auto calculated sizes
	columnSizes []int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// HeaderCapitalize is flag for capitalizing headers by default
var HeaderCapitalize = false

// HeaderColorTag is fmtc tag used for headers by default for all tables
var HeaderColorTag = "{*}"

// SeparatorSymbol used for separator generation
var SeparatorSymbol = "-"

// ColumnSeparatorSymbol is column separator symbol
var ColumnSeparatorSymbol = "|"

// MaxWidth is a maximum table width
var MaxWidth = 0

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTable create new table struct
func NewTable(headers ...string) *Table {
	return &Table{Headers: headers}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// SetHeaders allow to set columns headers
func (t *Table) SetHeaders(headers ...string) *Table {
	if t == nil {
		return nil
	}

	t.Headers = headers

	return t
}

// SetSizes allow to set column sizes
func (t *Table) SetSizes(sizes ...int) *Table {
	if t == nil {
		return nil
	}

	t.Sizes = sizes

	return t
}

// SetAlignments allow to set column alignment
func (t *Table) SetAlignments(align ...uint8) *Table {
	if t == nil {
		return nil
	}

	t.Alignment = align

	return t
}

// Add add given data to stack
func (t *Table) Add(data ...interface{}) *Table {
	if t == nil {
		return nil
	}

	if len(data) == 0 {
		return t
	}

	t.data = append(t.data, convertSlice(data))

	return t
}

// Print print given data
func (t *Table) Print(data ...interface{}) *Table {
	if t == nil {
		return nil
	}

	if len(data) == 0 {
		return t
	}

	if len(t.Headers) == 0 && len(t.Sizes) == 0 {
		setColumnsSizes(t, len(data))
	}

	prepareRender(t)
	renderRowData(t, convertSlice(data), len(t.columnSizes))

	return t
}

// HasData return true if table have some data
func (t *Table) HasData() bool {
	return t != nil && len(t.data) != 0
}

// Separator print separator
func (t *Table) Separator() *Table {
	if t == nil {
		return nil
	}

	if t.separator == "" {
		t.separator = strings.Repeat(SeparatorSymbol, getSeparatorSize(t))
	}

	fmtc.Println("{s}" + t.separator + "{!}")

	return t
}

// RenderHeaders force headers rendering
func (t *Table) RenderHeaders() {
	if len(t.columnSizes) == 0 {
		calculateColumnSizes(t)
	}

	renderHeaders(t)
}

// Render render data
func (t *Table) Render() *Table {
	if t == nil {
		return nil
	}

	// Nothing to render
	if len(t.Headers) == 0 && len(t.data) == 0 {
		return t
	}

	prepareRender(t)

	if len(t.Headers) == 0 {
		t.Separator()
	}

	if t.data != nil {
		renderData(t)
	}

	// Remove data after rendering
	t.separator = ""
	t.data = nil
	t.columnSizes = nil
	t.headerShown = false

	return t
}

// ////////////////////////////////////////////////////////////////////////////////// //

// prepareRender prepare table for render
func prepareRender(t *Table) {
	if len(t.columnSizes) == 0 {
		calculateColumnSizes(t)
	}

	if !t.headerShown {
		renderHeaders(t)
	}
}

// renderHeaders render headers
func renderHeaders(t *Table) {
	t.headerShown = true

	if len(t.Headers) == 0 {
		return
	}

	t.Separator()

	totalHeaders := len(t.Headers)
	totalColumns := len(t.columnSizes)

	var headerText string

	for columnIndex, columnSize := range t.columnSizes {
		if columnIndex >= totalHeaders {
			headerText = strings.Repeat(" ", columnSize)
		} else {
			headerText = t.Headers[columnIndex]
		}

		if HeaderCapitalize {
			headerText = strings.ToUpper(headerText)
		}

		fmtc.Printf(" " + HeaderColorTag + formatText(headerText, t.columnSizes[columnIndex], getAlignment(t, columnIndex)) + "{!} ")

		if columnIndex+1 != totalColumns {
			fmtc.Printf("{s}%s{!}", ColumnSeparatorSymbol)
		} else {
			fmtc.NewLine()
		}
	}

	t.Separator()
}

// renderData render table data
func renderData(t *Table) {
	totalColumns := len(t.columnSizes)

	for _, rowData := range t.data {
		renderRowData(t, rowData, totalColumns)
	}

	t.Separator()
}

// renderRowData render data in row
func renderRowData(t *Table, rowData []string, totalColumns int) {
	for columnIndex, columnData := range rowData {
		if columnIndex == totalColumns {
			break
		}

		if strutil.Len(fmtc.Clean(columnData)) > t.columnSizes[columnIndex] {
			fmtc.Print(" " + strutil.Ellipsis(columnData, t.columnSizes[columnIndex]) + " ")
		} else {
			fmtc.Print(" " + formatText(columnData, t.columnSizes[columnIndex], getAlignment(t, columnIndex)) + " ")
		}

		if columnIndex+1 != totalColumns {
			fmtc.Printf("{s}%s{!}", ColumnSeparatorSymbol)
		}
	}

	fmtc.NewLine()
}

// convertSlice convert slice with interface{} to slice with strings
func convertSlice(data []interface{}) []string {
	var result []string

	for _, item := range data {
		result = append(result, fmt.Sprintf("%v", item))
	}

	return result
}

// calculateColumnSizes calculate size for each column
func calculateColumnSizes(t *Table) {
	totalColumns := getColumnsNum(t)
	t.columnSizes = make([]int, totalColumns)

	if len(t.Sizes) != 0 {
		for columnIndex := range t.Sizes {
			if columnIndex < totalColumns {
				t.columnSizes[columnIndex] = t.Sizes[columnIndex]
			}
		}
	}

	if len(t.data) > 0 {
		for _, row := range t.data {
			for index, item := range row {
				itemSizes := strutil.Len(fmtc.Clean(item))

				if itemSizes > t.columnSizes[index] {
					t.columnSizes[index] = itemSizes
				}
			}
		}
	}

	if len(t.Headers) > 0 {
		for index, header := range t.Headers {
			headerSize := strutil.Len(header)

			if headerSize > t.columnSizes[index] {
				t.columnSizes[index] = headerSize
			}
		}
	}

	var fullSize int

	windowWidth := getWindowWidth()

	for columnIndex, columnSize := range t.columnSizes {
		if columnIndex+1 == totalColumns {
			t.columnSizes[columnIndex] = ((windowWidth - fullSize) - (totalColumns * 3)) + 1
		}

		fullSize += columnSize
	}
}

// setColumnsSizes set columns sizes by number of columns
func setColumnsSizes(t *Table, columns int) {
	windowWidth := getWindowWidth()
	t.columnSizes = make([]int, columns)

	totalSize := 0
	columnSize := (windowWidth / columns) - 3

	for index := range t.columnSizes {
		t.columnSizes[index] = columnSize
		totalSize += columnSize

		if index+1 == columns {
			if totalSize+(columns*3) < windowWidth {
				t.columnSizes[index]++
			}

			t.columnSizes[index]++
		}
	}
}

// getColumnsNum return number of columns
func getColumnsNum(t *Table) int {
	if len(t.data) > 0 {
		var columns int

		for _, row := range t.data {
			rowColumns := len(row)

			if rowColumns > columns {
				columns = rowColumns
			}
		}

		return columns
	}

	if len(t.Headers) > 0 {
		return len(t.Headers)
	}

	return len(t.Sizes)
}

// formatText align text with color tags
func formatText(data string, size int, align uint8) string {
	var dataSize int

	if strings.Contains(data, "{") {
		dataSize = strutil.Len(fmtc.Clean(data))
	} else {
		dataSize = strutil.Len(data)
	}

	if dataSize >= size {
		return data
	}

	switch align {
	case ALIGN_RIGHT:
		return strings.Repeat(" ", size-dataSize) + data

	case ALIGN_CENTER:
		prefixSize := (size - dataSize) / 2
		suffixSize := size - (prefixSize + dataSize)
		return strings.Repeat(" ", prefixSize) + data + strings.Repeat(" ", suffixSize)
	}

	return data + strings.Repeat(" ", size-dataSize)
}

// getAlignment return align for given column
func getAlignment(t *Table, columnIndex int) uint8 {
	l := len(t.Alignment)

	if l == 0 || columnIndex >= l {
		return 0
	}

	return t.Alignment[columnIndex]
}

// getSeparatorSize return separator size based on size of all columns
func getSeparatorSize(t *Table) int {
	if len(t.columnSizes) == 0 {
		return getWindowWidth()
	}

	var size int

	for _, columnSize := range t.columnSizes {
		size += columnSize
	}

	return size + (len(t.columnSizes) * 3) - 1
}

// getWindowWidth return window width
func getWindowWidth() int {
	if MaxWidth > 0 {
		return mathutil.Between(MaxWidth, 80, 9999)
	}

	return mathutil.Between(window.GetWidth(), 80, 9999)
}
