// Package reutil provides helpers for working with regular expressions
package reutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
	// ErrNilRegex is returned if regex pattern struct is nil
	ErrNilRegex = errors.New("Given regexp is nil")

	// ErrNilFunc is returned if replacement function is nil
	ErrNilFunc = errors.New("Replacement function is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Replace replaces parts of found strings using given method
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

	var lastChunkIndex int

	for n := 0; n < len(found); n++ {
		f, s, e := found[n], index[n][0], index[n][1]

		if s > 0 {
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
