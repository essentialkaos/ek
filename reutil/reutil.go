// Package reutil provides helpers for working with regular expressions
package reutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"regexp"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilRegex is returned when the provided regexp pattern is nil
	ErrNilRegex = errors.New("given regexp is nil")

	// ErrNilFunc is returned when the provided replacement function is nil
	ErrNilFunc = errors.New("replacement function is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Replace finds all matches of regex in source and calls replFunc for each match,
// passing the full match and any captured subgroups, then returns the rebuilt string
func Replace(regex *regexp.Regexp, source string, replFunc func(found string, submatch []string) string) (string, error) {
	switch {
	case regex == nil:
		return "", ErrNilRegex
	case replFunc == nil:
		return "", ErrNilFunc
	case source == "":
		return "", nil
	}

	found := regex.FindAllStringSubmatch(source, -1)

	if len(found) == 0 {
		return source, nil
	}

	index := regex.FindAllStringIndex(source, -1)

	buf := &bytes.Buffer{}
	buf.Grow(len(source))

	var lastChunkIndex int

	for i, f := range found {
		s, e := index[i][0], index[i][1]

		if s > lastChunkIndex {
			buf.WriteString(source[lastChunkIndex:s])
		}

		if len(f) > 1 {
			buf.WriteString(replFunc(f[0], f[1:]))
		} else {
			buf.WriteString(replFunc(f[0], nil))
		}

		lastChunkIndex = e
	}

	buf.WriteString(source[lastChunkIndex:])

	return buf.String(), nil
}
