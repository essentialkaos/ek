// Package csv contains simple CSV parser
package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Reader is reader struct
type Reader struct {
	Comma rune
	br    *bufio.Reader
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrEmptyDest is returned by the ReadTo method if empty destintation slice was given
var ErrEmptyDest = errors.New("Destination slice length must be greater than 1")

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader create new CSV reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Comma: ';',
		br:    bufio.NewReader(r),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads line from CSV file
func (r *Reader) Read() ([]string, error) {
	str, _, err := r.br.ReadLine()

	if err != nil || len(str) == 0 {
		return []string{}, err
	}

	return strings.Split(string(str), string(r.Comma)), nil
}

// ReadTo reads data to given slice
func (r *Reader) ReadTo(dst []string) error {
	if len(dst) == 0 {
		return ErrEmptyDest
	}

	str, _, err := r.br.ReadLine()

	if err != nil {
		return err
	}

	parseAndFill(string(str), dst, string(r.Comma))

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseAndFill(src string, dst []string, sep string) {
	l := len(dst)

	if src == "" {
		clean(dst, 0, l)
		return
	}

	n := strings.Count(src, sep)
	i := 0

	if n == 0 {
		dst[0] = src
		clean(dst, 1, l)
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
		clean(dst, i, l)
	}
}

func clean(dst []string, since, to int) {
	for i := since; i < to; i++ {
		dst[i] = ""
	}
}
