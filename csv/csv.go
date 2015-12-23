// Package csv contains simple csv parser
package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
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

// NewReader create new reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Comma: ';',
		br:    bufio.NewReader(r),
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads line from csv file
func (r *Reader) Read() ([]string, error) {
	str, _, err := r.br.ReadLine()

	if err != nil || len(str) == 0 {
		return []string{}, err
	}

	return strings.Split(string(str), string(r.Comma)), nil
}
