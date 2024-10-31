// Package table contains methods and structs for rendering data in tabular format
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

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/terminal/tty"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Column alignment flags
const (
	ALIGN_LEFT   uint8 = 0
	ALIGN_CENTER uint8 = 1
	ALIGN_RIGHT  uint8 = 2

	AL uint8 = 0 // Short form of ALIGN_LEFT
	AC uint8 = 1 // Short form of ALIGN_CENTER
	AR uint8 = 2 // Short form of ALIGN_RIGHT
)

// ////////////////////////////////////////////////////////////////////////////////// //

// _SEPARATOR_TAG is unique value used as placeholder for data separator
const _SEPARATOR_TAG = "--[@SEPARATOR@]--"

// ////////////////////////////////////////////////////////////////////////////////// //

// Table is struct which can be used for table rendering
type Table struct {
	Sizes     []int    // Custom columns sizes
	Headers   []string // Slice with headers
	Alignment []uint8  // Columns alignment

	// Width is table maximum width
	Width int

	// Breaks is an interval for separators between given number of rows
	Breaks int

	// SeparatorSymbol is symbol used for borders rendering
	BorderSymbol string

	// SeparatorSymbol is symbol used for separator rendering
	SeparatorSymbol string

	// ColumnSeparatorSymbol is column separator symbol
	ColumnSeparatorSymbol string

	// HeaderColorTag is fmtc tag used for headers
	HeaderColorTag string

	// BorderColorTag is fmtc tag used for separator
	BorderColorTag string

	// SeparatorColorTag is fmtc tag used for separator
	SeparatorColorTag string

	// HeaderCapitalize is a flag for capitalizing headers
	HeaderCapitalize bool

	// HideTopBorder is a flag for disabling bottom border rendering
	HideTopBorder bool

	// HideBottomBorder is a flag for disabling bottom border rendering
	HideBottomBorder bool

	// FullScreen is a flag for full-screen table
	FullScreen bool

	// Processor is function used for processing and formatting input data
	Processor func(data []any) []string

	// Slice with data
	data [][]string

	// Separator cache
	separator string

	// Flag will be set if header is rendered
	headerShown bool

	// Slice with auto calculated sizes
	columnSizes []int

	// Cursor is number of the latest record
	cursor int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// HeaderCapitalize is a flag for capitalizing headers by default
var HeaderCapitalize = false

// HeaderColorTag is default fmtc tag used for headers
var HeaderColorTag = "{*}"

// SeparatorSymbol is default symbol used for borders rendering
var BorderSymbol = "-"

// BorderColorTag is default fmtc tag used for separator
var BorderColorTag = "{s}"

// SeparatorSymbol is default symbol used for separator rendering
var SeparatorSymbol = "-"

// SeparatorColorTag is default fmtc tag used for separator
var SeparatorColorTag = "{s}"

// ColumnSeparatorSymbol is default column separator symbol
var ColumnSeparatorSymbol = "|"

// Breaks is an interval for separators between given number of rows
var Breaks = -1

// FullScreen is a flag for full-screen table by default
var FullScreen = true

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTable creates new table struct
func NewTable(headers ...string) *Table {
	return &Table{
		HeaderCapitalize: HeaderCapitalize,
		FullScreen:       FullScreen,
		Breaks:           Breaks,
		Headers:          headers,
		Processor:        convertSlice,
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// SetHeaders sets headers
func (t *Table) SetHeaders(headers ...string) *Table {
	if t == nil {
		return nil
	}

	t.Headers = headers

	return t
}

// SetSizes sets size of columns
func (t *Table) SetSizes(sizes ...int) *Table {
	if t == nil {
		return nil
	}

	t.Sizes = sizes

	return t
}

// SetAlignments sets alignment of columns
func (t *Table) SetAlignments(align ...uint8) *Table {
	if t == nil {
		return nil
	}

	t.Alignment = align

	return t
}

// Add adds given data to stack
func (t *Table) Add(data ...any) *Table {
	if t == nil {
		return nil
	}

	if len(data) == 0 {
		return t
	}

	t.data = append(t.data, t.Processor(data))

	return t
}

// Print immediately prints given data
func (t *Table) Print(data ...any) *Table {
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
	renderRowData(t, t.Processor(data), len(t.columnSizes))

	return t
}

// HasData returns true if table stack has some data
func (t *Table) HasData() bool {
	return t != nil && len(t.data) != 0
}

// Separator renders separator
func (t *Table) Separator() *Table {
	if t == nil {
		return nil
	}

	if !t.HasData() {
		renderSeparator(t)
	} else {
		t.Add(_SEPARATOR_TAG)
	}

	return t
}

// Border renders table border
func (t *Table) Border() *Table {
	if t == nil {
		return nil
	}

	renderBorder(t)

	return t
}

// RenderHeaders renders headers
func (t *Table) RenderHeaders() {
	if t == nil {
		return
	}

	if len(t.columnSizes) == 0 {
		calculateColumnSizes(t)
	}

	renderHeaders(t)
}

// Render renders data
func (t *Table) Render() *Table {
	if t == nil {
		return nil
	}

	// Nothing to render
	if len(t.Headers) == 0 && len(t.data) == 0 {
		return t
	}

	prepareRender(t)

	if len(t.Headers) == 0 && !t.HideTopBorder {
		renderBorder(t)
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

	if !t.HideTopBorder {
		renderBorder(t)
	}

	totalHeaders := len(t.Headers)
	totalColumns := len(t.columnSizes)

	var headerText string

	for columnIndex, columnSize := range t.columnSizes {
		if columnIndex >= totalHeaders {
			headerText = strings.Repeat(" ", columnSize)
		} else {
			headerText = t.Headers[columnIndex]
		}

		if t.HeaderCapitalize {
			headerText = strings.ToUpper(headerText)
		}

		fmtc.Print(
			" " + strutil.Q(t.HeaderColorTag, HeaderColorTag) + formatText(headerText, t.columnSizes[columnIndex],
				getAlignment(t, columnIndex)) + "{!} ",
		)

		if columnIndex+1 != totalColumns {
			fmtc.Printf(
				strutil.Q(t.SeparatorColorTag, SeparatorColorTag)+"%s{!}",
				strutil.Q(t.ColumnSeparatorSymbol, ColumnSeparatorSymbol),
			)
		} else {
			fmtc.NewLine()
		}
	}

	renderBorder(t)
}

// renderData render table data
func renderData(t *Table) {
	totalColumns := len(t.columnSizes)

	for _, rowData := range t.data {
		if rowData[0] == _SEPARATOR_TAG {
			renderSeparator(t)
			continue
		}

		renderRowData(t, rowData, totalColumns)
	}

	if !t.HideBottomBorder {
		renderBorder(t)
	}
}

// renderRowData render data in row
func renderRowData(t *Table, data []string, totalColumns int) {
	if t.Breaks > 0 && t.cursor > 0 && t.cursor%t.Breaks == 0 {
		renderSeparator(t)
	}

	for columnIndex, columnData := range data {
		if columnIndex == totalColumns {
			break
		}

		if strutil.LenVisual(fmtc.Clean(columnData)) > t.columnSizes[columnIndex] {
			fmtc.Print(" " + strutil.Ellipsis(columnData, t.columnSizes[columnIndex]) + " ")
		} else {
			if columnIndex+1 == totalColumns && getAlignment(t, columnIndex) == ALIGN_LEFT {
				fmtc.Print(" " + formatText(columnData, -1, ALIGN_LEFT))
			} else {
				fmtc.Print(" " + formatText(columnData, t.columnSizes[columnIndex], getAlignment(t, columnIndex)) + " ")
			}
		}

		if columnIndex+1 != totalColumns {
			fmtc.Printf(
				strutil.Q(t.SeparatorColorTag, SeparatorColorTag)+"%s{!}",
				strutil.Q(t.ColumnSeparatorSymbol, ColumnSeparatorSymbol),
			)
		}
	}

	t.cursor++

	fmtc.NewLine()
}

// renderSeparator prints separator
func renderSeparator(t *Table) {
	if t.separator == "" {
		t.separator = strings.Repeat(strutil.Q(t.SeparatorSymbol, SeparatorSymbol), getSeparatorSize(t))
	}

	fmtc.Println(strutil.Q(t.SeparatorColorTag, SeparatorColorTag) + t.separator + "{!}")
}

// renderBorder renders table border
func renderBorder(t *Table) {
	border := strings.Repeat(strutil.Q(t.BorderSymbol, BorderSymbol), getSeparatorSize(t))

	fmtc.Println(strutil.Q(t.BorderColorTag, BorderColorTag) + border + "{!}")
}

// convertSlice convert slice with any to slice with strings
func convertSlice(data []any) []string {
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
				if item == _SEPARATOR_TAG {
					continue
				}

				itemSizes := strutil.LenVisual(fmtc.Clean(item))

				if itemSizes > t.columnSizes[index] {
					t.columnSizes[index] = itemSizes
				}
			}
		}
	}

	if len(t.Headers) > 0 {
		for index, header := range t.Headers {
			headerSize := strutil.LenVisual(header)

			if headerSize > t.columnSizes[index] {
				t.columnSizes[index] = headerSize
			}
		}
	}

	tableWidth := getTableWidth(t)

	if tableWidth > 0 {
		var fullSize int

		for columnIndex, columnSize := range t.columnSizes {
			if columnIndex+1 == totalColumns {
				t.columnSizes[columnIndex] = ((tableWidth - fullSize) - (totalColumns * 3)) + 1
			}

			fullSize += columnSize
		}
	}
}

// setColumnsSizes set columns sizes by number of columns
func setColumnsSizes(t *Table, columns int) {
	tableWidth := getTableWidth(t)
	t.columnSizes = make([]int, columns)

	totalSize := 0
	columnSize := (tableWidth / columns) - 3

	for index := range t.columnSizes {
		t.columnSizes[index] = columnSize
		totalSize += columnSize

		if index+1 == columns {
			if totalSize+(columns*3) < tableWidth {
				fmt.Println(1001)
				t.columnSizes[index]++
			}

			t.columnSizes[index]++
		}
	}
}

// getColumnsNum returns number of columns
func getColumnsNum(t *Table) int {
	var columns int

	if len(t.data) > 0 {
		for _, row := range t.data {
			rowColumns := len(row)

			if rowColumns > columns {
				columns = rowColumns
			}
		}
	}

	if len(t.Headers) > columns {
		columns = len(t.Headers)
	}

	if len(t.Sizes) > columns {
		columns = len(t.Sizes)
	}

	return columns
}

// formatText align text with color tags
func formatText(data string, size int, align uint8) string {
	var dataSize int

	if strings.Contains(data, "{") {
		dataSize = strutil.LenVisual(fmtc.Clean(data))
	} else {
		dataSize = strutil.LenVisual(data)
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
	tableWidth := getTableWidth(t)

	if tableWidth > 0 {
		return tableWidth
	}

	var size int

	for _, columnSize := range t.columnSizes {
		size += columnSize
	}

	return size + (len(t.columnSizes) * 3) - 2
}

// getTableWidth returns maximum width of table
func getTableWidth(t *Table) int {
	if t.Width > 0 {
		return mathutil.Between(t.Width, 80, 9999)
	}

	if t.FullScreen || len(t.columnSizes) == 0 {
		return mathutil.Between(tty.GetWidth(), 80, 9999)
	}

	return 0
}
