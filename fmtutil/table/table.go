// Package table contains methods and structs for rendering data in tabular format
package table

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v13/ansi"
	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/terminal/tty"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	ALIGN_LEFT   uint8 = 0 // Aligns column content to the left
	ALIGN_CENTER uint8 = 1 // Centers column content
	ALIGN_RIGHT  uint8 = 2 // Aligns column content to the right

	AL uint8 = 0 // Short form of [ALIGN_LEFT]
	AC uint8 = 1 // Short form of [ALIGN_CENTER]
	AR uint8 = 2 // Short form of [ALIGN_RIGHT]
)

// ////////////////////////////////////////////////////////////////////////////////// //

// _SEPARATOR_TAG is unique value used as placeholder for data separator
const _SEPARATOR_TAG = "--[@SEPARATOR@]--"

// ////////////////////////////////////////////////////////////////////////////////// //

// Table is struct which can be used for table rendering
type Table struct {
	// Sizes defines fixed widths for each column; overrides auto-calculation
	Sizes []int

	// Headers contains the column header labels
	Headers []string

	// Alignment defines per-column text alignment using [ALIGN_LEFT], [ALIGN_CENTER],
	// or [ALIGN_RIGHT]
	Alignment []uint8

	// Width sets the maximum table width in characters; clamped to [80, 9999]
	Width int

	// Breaks defines the row interval at which automatic separators are inserted
	Breaks int

	// BorderSymbol is the character used to render top and bottom borders
	BorderSymbol string

	// SeparatorSymbol is the character used to render horizontal separators between rows
	SeparatorSymbol string

	// ColumnSeparatorSymbol is the character used to divide adjacent columns
	ColumnSeparatorSymbol string

	// HeaderColorTag is the fmtc color tag applied to header text
	HeaderColorTag string

	// BorderColorTag is the fmtc color tag applied to border lines
	BorderColorTag string

	// SeparatorColorTag is the fmtc color tag applied to separator lines
	SeparatorColorTag string

	// HeaderCapitalize controls whether header labels are uppercased before rendering
	HeaderCapitalize bool

	// HideTopBorder disables rendering of the top border line
	HideTopBorder bool

	// HideBottomBorder disables rendering of the bottom border line
	HideBottomBorder bool

	// FullScreen stretches the table to the full terminal width
	FullScreen bool

	// Processor is the function used to convert a row of any values into strings.
	// It defaults to fmt.Sprintf("%v", …) conversion for each element.
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

// HeaderCapitalize controls whether headers are uppercased by default for new tables
var HeaderCapitalize = false

// HeaderColorTag is the default fmtc color tag applied to header text
var HeaderColorTag = "{*}"

// BorderSymbol is the default character used for top and bottom border lines
var BorderSymbol = "-"

// BorderColorTag is the default fmtc color tag applied to border lines
var BorderColorTag = "{s}"

// SeparatorSymbol is the default character used for horizontal separator lines
var SeparatorSymbol = "-"

// SeparatorColorTag is the default fmtc color tag applied to separator lines
var SeparatorColorTag = "{s}"

// ColumnSeparatorSymbol is the default character used to divide adjacent columns
var ColumnSeparatorSymbol = "|"

// Breaks is the default row interval for automatic separator insertion
var Breaks = 0

// FullScreen controls whether new tables expand to the full terminal width by default
var FullScreen = true

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTable creates and returns a new Table with optional column headers
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

// SetHeaders sets the column headers for the table
func (t *Table) SetHeaders(headers ...string) *Table {
	if t == nil {
		return nil
	}

	t.Headers = headers

	return t
}

// SetSizes sets fixed widths for each column in order
func (t *Table) SetSizes(sizes ...int) *Table {
	if t == nil {
		return nil
	}

	t.Sizes = sizes

	return t
}

// SetAlignments sets the alignment for each column using [ALIGN_LEFT], [ALIGN_CENTER],
// or [ALIGN_RIGHT]
func (t *Table) SetAlignments(align ...uint8) *Table {
	if t == nil {
		return nil
	}

	t.Alignment = align

	return t
}

// Add appends a row of data to the table's internal buffer for deferred rendering
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

// Print immediately renders a single row without buffering it
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

// HasData reports whether the table's buffer contains any pending rows
func (t *Table) HasData() bool {
	return t != nil && len(t.data) != 0
}

// Separator adds a horizontal separator at the current position in the buffer,
// or renders it immediately if no data is buffered
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

// Border immediately renders a single horizontal border line
func (t *Table) Border() *Table {
	if t == nil {
		return nil
	}

	renderBorder(t)

	return t
}

// RenderHeaders calculates column sizes and renders the header row with borders
func (t *Table) RenderHeaders() {
	if t == nil {
		return
	}

	if len(t.columnSizes) == 0 {
		calculateColumnSizes(t)
	}

	renderHeaders(t)
}

// Render flushes the buffered data to output, rendering headers, rows, and borders,
// then resets the table's internal state
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
	t.cursor = 0

	return t
}

// ////////////////////////////////////////////////////////////////////////////////// //

// prepareRender calculates column sizes and renders the header if not yet shown
func prepareRender(t *Table) {
	if len(t.columnSizes) == 0 {
		calculateColumnSizes(t)
	}

	if !t.headerShown {
		renderHeaders(t)
	}
}

// renderHeaders renders the header row surrounded by border lines
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

// renderData renders all buffered rows and the bottom border
func renderData(t *Table) {
	totalColumns := len(t.columnSizes)

	for _, rowData := range t.data {
		if len(rowData) == 0 {
			continue
		}

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

// renderRowData renders a single row, inserting automatic separators based on [Breaks]
func renderRowData(t *Table, data []string, totalColumns int) {
	if t.Breaks > 0 && t.cursor > 0 && t.cursor%t.Breaks == 0 {
		renderSeparator(t)
	}

	for columnIndex, columnData := range data {
		if columnIndex == totalColumns {
			break
		}

		dataLen := getDataLen(columnData)

		if dataLen > t.columnSizes[columnIndex] {
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

// renderSeparator prints a full-width separator line using [SeparatorSymbol]
func renderSeparator(t *Table) {
	if t.separator == "" {
		t.separator = strings.Repeat(strutil.Q(t.SeparatorSymbol, SeparatorSymbol), getSeparatorSize(t))
	}

	fmtc.Println(strutil.Q(t.SeparatorColorTag, SeparatorColorTag) + t.separator + "{!}")
}

// renderBorder prints a full-width border line using [BorderSymbol]
func renderBorder(t *Table) {
	border := strings.Repeat(strutil.Q(t.BorderSymbol, BorderSymbol), getSeparatorSize(t))

	fmtc.Println(strutil.Q(t.BorderColorTag, BorderColorTag) + border + "{!}")
}

// convertSlice converts a slice of any values to a slice of their string
// representations
func convertSlice(data []any) []string {
	result := make([]string, len(data))

	for i, item := range data {
		result[i] = fmt.Sprintf("%v", item)
	}

	return result
}

// calculateColumnSizes computes optimal column widths based on headers,
// data content, and table width
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

				itemLen := getDataLen(item)

				if itemLen > t.columnSizes[index] {
					t.columnSizes[index] = itemLen
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

// setColumnsSizes distributes table width evenly across the given number of columns
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
				t.columnSizes[index]++
			}

			t.columnSizes[index]++
		}
	}
}

// getColumnsNum returns the maximum number of columns across headers, sizes,
// and data rows
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

// formatText pads or centers the given string to the target size according to
// the alignment flag
func formatText(data string, size int, align uint8) string {
	dataLen := getDataLen(data)

	if dataLen >= size {
		return data
	}

	switch align {
	case ALIGN_RIGHT:
		return strings.Repeat(" ", size-dataLen) + data

	case ALIGN_CENTER:
		prefixSize := (size - dataLen) / 2
		suffixSize := size - (prefixSize + dataLen)
		return strings.Repeat(" ", prefixSize) + data + strings.Repeat(" ", suffixSize)
	}

	return data + strings.Repeat(" ", size-dataLen)
}

// getAlignment returns the alignment setting for the given column index, defaulting
// to [ALIGN_LEFT]
func getAlignment(t *Table, columnIndex int) uint8 {
	l := len(t.Alignment)

	if l == 0 || columnIndex >= l {
		return 0
	}

	return t.Alignment[columnIndex]
}

// getSeparatorSize returns the total rendered width for separators and borders
func getSeparatorSize(t *Table) int {
	tableWidth := getTableWidth(t)

	if tableWidth > 0 {
		return tableWidth
	}

	var size int

	for _, columnSize := range t.columnSizes {
		size += columnSize
	}

	return size + (len(t.columnSizes) * 3) - 1
}

// getTableWidth returns the effective table width, clamped to [80, 9999], or 0 for
// natural sizing
func getTableWidth(t *Table) int {
	if t.Width > 0 {
		return mathutil.Between(t.Width, 80, 9999)
	}

	if t.FullScreen || len(t.columnSizes) == 0 {
		return mathutil.Between(tty.GetWidth(), 80, 9999)
	}

	return 0
}

// getDataLen returns the visible character length of a string, stripping ANSI codes
// and fmtc tags
func getDataLen(data string) int {
	return strutil.LenVisual(ansi.Remove(fmtc.Clean(data)))
}
