// Package csv contains simple CSV parser
package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
