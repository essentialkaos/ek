// Package csv contains simple CSV parser
package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// COMMA is default separator for CSV cells
const COMMA = ';'

// ////////////////////////////////////////////////////////////////////////////////// //

// Reader is CSV reader struct
type Reader struct {
	Comma      rune
	SkipHeader bool

	headerSkipped bool
	currentLine   int

	br *bufio.Reader
}

// Row is CSV row
type Row []string

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrEmptyDest is returned by the ReadTo method if empty destination slice was given
var ErrEmptyDest = errors.New("Destination slice length must be greater than 1")

// ErrNilReader is returned when reader struct is nil
var ErrNilReader = errors.New("Reader is nil")

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader create new CSV reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Comma:      COMMA,
		SkipHeader: false,

		br: bufio.NewReader(r),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads line from CSV file
func (r *Reader) Read() (Row, error) {
	if r == nil {
		return nil, ErrNilReader
	}

	if r.SkipHeader && !r.headerSkipped {
		r.br.ReadLine()
		r.headerSkipped = true
		r.currentLine++
	}

	str, _, err := r.br.ReadLine()

	if err != nil || len(str) == 0 {
		return nil, err
	}

	r.currentLine++

	return strings.Split(string(str), string(r.Comma)), nil
}

// ReadTo reads data to given slice
func (r *Reader) ReadTo(dst Row) error {
	if r == nil {
		return ErrNilReader
	}

	if len(dst) == 0 {
		return ErrEmptyDest
	}

	if r.SkipHeader && !r.headerSkipped {
		r.br.ReadLine()
		r.headerSkipped = true
		r.currentLine++
	}

	str, _, err := r.br.ReadLine()

	if err != nil {
		return err
	}

	parseAndFill(string(str), dst, string(r.Comma))
	r.currentLine++

	return nil
}

// WithComma sets comma (fields delimiter) for CSV reader
func (r *Reader) WithComma(comma rune) *Reader {
	if r == nil {
		return nil
	}

	r.Comma = comma

	return r
}

// WithHeaderSkip sets header skip flag
func (r *Reader) WithHeaderSkip(flag bool) *Reader {
	if r == nil {
		return nil
	}

	r.SkipHeader = flag

	return r
}

// Line returns number of the last line read
func (r *Reader) Line() int {
	if r == nil {
		return 0
	}

	return r.currentLine
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Size returns size of the row
func (r Row) Size() int {
	return len(r)
}

// Cells returns number of cells filled with data
func (r Row) Cells() int {
	var size int

	for _, c := range r {
		if len(c) > 0 {
			size++
		}
	}

	return size
}

// IsEmpty returns true if all cells are empty
func (r Row) IsEmpty() bool {
	return r.Cells() == 0
}

// Has returns true if row contains cell with given index filled with data
func (r Row) Has(index int) bool {
	return index < len(r) && r[index] != ""
}

// Get returns value of the cell with given index
func (r Row) Get(index int) string {
	if index >= len(r) {
		return ""
	}

	return r[index]
}

// GetB returns cell value as boolean
func (r Row) GetB(index int) bool {
	switch strings.ToLower(r.Get(index)) {
	case "1", "true", "t", "yes", "y", "+":
		return true
	}

	return false
}

// GetI returns cell value as int
func (r Row) GetI(index int) (int, error) {
	return strconv.Atoi(r.Get(index))
}

// GetI8 returns cell value as int8
func (r Row) GetI8(index int) (int8, error) {
	i, err := strconv.ParseInt(r.Get(index), 10, 8)

	if err != nil {
		return 0, err
	}

	return int8(i), nil
}

// GetI16 returns cell value as int16
func (r Row) GetI16(index int) (int16, error) {
	i, err := strconv.ParseInt(r.Get(index), 10, 16)

	if err != nil {
		return 0, err
	}

	return int16(i), nil
}

// GetI32 returns cell value as int32
func (r Row) GetI32(index int) (int32, error) {
	i, err := strconv.ParseInt(r.Get(index), 10, 32)

	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

// GetI64 returns cell value as int64
func (r Row) GetI64(index int) (int64, error) {
	i, err := strconv.ParseInt(r.Get(index), 10, 64)

	if err != nil {
		return 0, err
	}

	return int64(i), nil
}

// GetF returns cell value as float
func (r Row) GetF(index int) (float64, error) {
	return strconv.ParseFloat(r.Get(index), 64)
}

// GetF32 returns cell value as float32
func (r Row) GetF32(index int) (float32, error) {
	f, err := strconv.ParseFloat(r.Get(index), 32)

	if err != nil {
		return 0, err
	}

	return float32(f), nil
}

// GetU returns cell value as uint
func (r Row) GetU(index int) (uint, error) {
	u, err := strconv.ParseUint(r.Get(index), 10, 32)

	if err != nil {
		return 0, err
	}

	return uint(u), nil
}

// GetU8 returns cell value as uint8
func (r Row) GetU8(index int) (uint8, error) {
	u, err := strconv.ParseUint(r.Get(index), 10, 8)

	if err != nil {
		return 0, err
	}

	return uint8(u), nil
}

// GetU16 returns cell value as uint16
func (r Row) GetU16(index int) (uint16, error) {
	u, err := strconv.ParseUint(r.Get(index), 10, 16)

	if err != nil {
		return 0, err
	}

	return uint16(u), nil
}

// GetU32 returns cell value as uint32
func (r Row) GetU32(index int) (uint32, error) {
	u, err := strconv.ParseUint(r.Get(index), 10, 32)

	if err != nil {
		return 0, err
	}

	return uint32(u), nil
}

// GetU64 returns cell value as uint64
func (r Row) GetU64(index int) (uint64, error) {
	u, err := strconv.ParseUint(r.Get(index), 10, 64)

	if err != nil {
		return 0, err
	}

	return uint64(u), nil
}

// ForEach executes given function for every cell in a row
func (r Row) ForEach(fn func(index int, value string) error) error {
	for i, c := range r {
		err := fn(i, c)

		if err != nil {
			return err
		}
	}

	return nil
}

// ToString returns string representation of row
func (r Row) ToString(comma rune) string {
	var buf bytes.Buffer

	for i, c := range r {
		buf.WriteString(c)

		if i+1 != len(r) {
			buf.WriteRune(comma)
		}
	}

	return buf.String()
}

// ToBytes returns representation of row as a byte slice
func (r Row) ToBytes(comma rune) []byte {
	var buf bytes.Buffer

	for i, c := range r {
		buf.WriteString(c)

		if i+1 != len(r) {
			buf.WriteRune(comma)
		}
	}

	return buf.Bytes()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseAndFill parses CSV row and fills slice with data
func parseAndFill(src string, dst Row, sep string) {
	l := len(dst)

	if src == "" {
		clean(dst, 0)
		return
	}

	n := strings.Count(src, sep)
	i := 0

	if n == 0 {
		dst[0] = src
		clean(dst, 1)
		return
	}

	for i < n && i < l {
		m := strings.Index(src, sep)

		dst[i] = src[:m]
		src = src[m+1:]

		i++
	}

	if src != "" && i != l {
		dst[i] = src
	}

	if i < l-1 {
		clean(dst, i+1)
	}
}

// clean cleans destination slice
func clean(dst Row, from int) {
	for i := from; i < len(dst); i++ {
		dst[i] = ""
	}
}
