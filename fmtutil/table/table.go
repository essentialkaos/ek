// Package table contains methods and structs for rendering data as a table
package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v9/fmtc"
	"pkg.re/essentialkaos/ek.v9/mathutil"
	"pkg.re/essentialkaos/ek.v9/terminal/window"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MIN_WINDOW_SIZE is minimal window size
const MIN_WINDOW_SIZE = 88

// ////////////////////////////////////////////////////////////////////////////////// //

type Table struct {
	Sizes   []int    // Custom columns sizes
	Headers []string // Slice with headers

	HeaderCapitalize bool   // Capitalize headers
	HeaderColorTag   string // Customize headers color

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

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTable create new table struct
func NewTable(headers ...string) *Table {
	return &Table{Headers: headers}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// SetHeaders set columns headers
func (t *Table) SetHeaders(headers ...string) *Table {
	if t == nil {
		return nil
	}

	t.Headers = headers

	return t
}

// SetSizes set column sizes
func (t *Table) SetSizes(sizes ...int) *Table {
	if t == nil {
		return nil
	}

	t.Sizes = sizes

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

	if len(data) == 0 || len(t.Headers) == 0 {
		return t
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
		t.separator = renderSeparator(t.columnSizes)
	}

	fmtc.Println("{s}" + t.separator + "{!}")

	return t
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
	t.Separator()

	t.headerShown = true

	if len(t.Headers) == 0 {
		return
	}

	colorTag := getHeaderColorTag(t)
	headerCapitalize := getHeaderCapitalize(t)
	totalHeaders := len(t.Headers)
	totalColumns := len(t.columnSizes)

	var headerText string

	for columnIndex, columnSize := range t.columnSizes {
		if columnIndex >= totalHeaders {
			headerText = strings.Repeat(" ", columnSize)
		} else {
			headerText = t.Headers[columnIndex]
		}

		if headerCapitalize {
			headerText = strings.ToUpper(headerText)
		}

		if columnIndex+1 == totalColumns {
			fmtc.Printf(" "+colorTag+"%s{!}\n", headerText)
		} else {
			headerSizeStr := strconv.Itoa(t.columnSizes[columnIndex])
			fmtc.Printf(" "+colorTag+"%"+headerSizeStr+"s{!} {s}|{!}", headerText)
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

		if columnIndex+1 == totalColumns {
			fmtc.Printf(" %s", fmtc.Sprintf(columnData))
		} else {
			fmtc.Printf(" " + alignData(columnData, t.columnSizes[columnIndex]) + " {s}|{!}")
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
func calculateColumnSizes(t *Table) []int {
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
				itemSizes := len(fmtc.Clean(item))

				if itemSizes > t.columnSizes[index] {
					t.columnSizes[index] = itemSizes
				}
			}
		}
	}

	if len(t.Headers) > 0 {
		for index, header := range t.Headers {
			headerSize := len(header)

			if headerSize > t.columnSizes[index] {
				t.columnSizes[index] = headerSize
			}
		}
	}

	return t.columnSizes
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

// renderSeparator render separator
func renderSeparator(sizes []int) string {
	var result string

	columns := len(sizes)

	for columnIndex, columnSize := range sizes {
		result += strings.Repeat("-", columnSize+2)

		if columnIndex+1 != columns {
			result += "+"
		}
	}

	winWidth := getWindowWidth()
	resultLen := len(result)

	if resultLen < winWidth {
		result += strings.Repeat("-", winWidth-resultLen)
	}

	return result
}

// alignRecord align records with color tags
func alignData(data string, size int) string {
	var dataSize int

	if strings.Contains(data, "{") {
		dataSize = len(fmtc.Clean(data))
	} else {
		dataSize = len(data)
	}

	if dataSize >= size {
		return data
	}

	return strings.Repeat(" ", size-dataSize) + data
}

// getWindowWidth return window width
func getWindowWidth() int {
	return mathutil.Between(window.GetWidth(), MIN_WINDOW_SIZE, 9999)
}

// getHeaderColorTag return header color tag
func getHeaderColorTag(t *Table) string {
	if t.HeaderColorTag != "" {
		return t.HeaderColorTag
	}

	return HeaderColorTag
}

// getHeaderCapitalize return header capitalization flag
func getHeaderCapitalize(t *Table) bool {
	if t.HeaderCapitalize {
		return true
	}

	return HeaderCapitalize
}
