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

// Reader is CSV reader struct
type Reader struct {
	Header Header

	comma         string
	hasHeader     bool
	headerSkipped bool
	currentLine   int

	s *bufio.Scanner
}

// Row is CSV row
type Row []string

// Header is row with header data
type Header []string

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyDest is returned by the ReadTo method if empty destination slice was given
	ErrEmptyDest = errors.New("Destination slice length must be greater than 1")

	// ErrNilReader is returned when reader struct is nil
	ErrNilReader = errors.New("Reader is nil")

	// ErrEmptyHeader is returned when header has no data
	ErrEmptyHeader = errors.New("Header is empty")

	// ErrNilMap is returned when map is nil
	ErrNilMap = errors.New("Map is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader creates new CSV reader
func NewReader(r io.Reader, comma rune) *Reader {
	return &Reader{
		comma:     string(comma),
		hasHeader: false,
		s:         bufio.NewScanner(r),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads line from CSV file
func (r *Reader) Read() (Row, error) {
	if r == nil || r.s == nil {
		return nil, ErrNilReader
	}

	if r.hasHeader && !r.headerSkipped {
		err := r.readHeader()

		if err != nil {
			return nil, err
		}
	}

	if !r.s.Scan() {
		return nil, r.getError()
	}

	r.currentLine++

	return strings.Split(r.s.Text(), r.comma), nil
}

// ReadTo reads data to given slice
func (r *Reader) ReadTo(dst Row) error {
	switch {
	case r == nil, r.s == nil:
		return ErrNilReader
	case len(dst) == 0:
		return ErrEmptyDest
	}

	if r.hasHeader && !r.headerSkipped {
		err := r.readHeader()

		if err != nil {
			return err
		}
	}

	if !r.s.Scan() {
		return r.getError()
	}

	parseAndFill(r.s.Text(), dst, r.comma)
	r.currentLine++

	return nil
}

// Seq is an iterator over all CSV data
func (r *Reader) Seq(yield func(line int, row Row) bool) {
	if r == nil || r.s == nil {
		return
	}

	for {
		rr, err := r.Read()
		if err != nil || !yield(r.currentLine, rr) {
			return
		}
	}
}

// WithHeader sets header skip flag
func (r *Reader) WithHeader(flag bool) *Reader {
	if r == nil || r.s == nil {
		return nil
	}

	r.hasHeader = flag

	return r
}

// Line returns number of the last line read
func (r *Reader) Line() int {
	if r == nil || r.s == nil {
		return 0
	}

	return r.currentLine
}

// Error returns error from underlying scanner
func (r *Reader) Error() error {
	if r == nil || r.s == nil {
		return nil
	}

	return r.s.Err()
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

// Map maps row data using headers names
func (h Header) Map(m map[string]string, r Row) error {
	switch {
	case len(h) == 0:
		return ErrEmptyHeader
	case m == nil:
		return ErrNilMap
	}

	for i, name := range h {
		m[name] = r.Get(i)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getError returns error from scanner
func (r *Reader) getError() error {
	if r.s.Err() == nil {
		return io.EOF
	}

	return r.s.Err()
}

// readHeader reads header data
func (r *Reader) readHeader() error {
	if !r.s.Scan() {
		return r.getError()
	}

	r.headerSkipped = true
	r.currentLine++

	r.Header = strings.Split(r.s.Text(), r.comma)

	return nil
}

// parseAndFill parses CSV row and fills slice with data
func parseAndFill(src string, dst Row, sep string) {
	if src == "" {
		clean(dst, 0)
		return
	}

	l := len(dst)
	n := strings.Count(src, sep)
	i := 0

	if n == 0 {
		dst[0] = src
		clean(dst, 1)
		return
	}

	for i < n && i < l {
		dst[i], src, _ = strings.Cut(src, sep)
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
