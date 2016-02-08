// Package strutil provides methods for working with strings
package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Concat fast string concatenation
func Concat(s ...string) string {
	var buffer bytes.Buffer

	for _, v := range s {
		buffer.WriteString(v)
	}

	return buffer.String()
}

// Substr return substring from given string
func Substr(s string, start int, end int) string {
	if s == "" {
		return ""
	}

	sl := len(s)

	if start > sl {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end > sl || end == 0 {
		end = sl
	}

	input := bytes.NewBufferString(s)
	output := bytes.NewBufferString("")

	input.Next(start)

	current := 0
	max := end - start

	for {
		r, _, _ := input.ReadRune()

		if r == rune(65533) {
			continue
		}

		output.WriteRune(r)

		current++

		if current == max {
			break
		}
	}

	return output.String()
}

// Ellipsis trims given string
func Ellipsis(s string, maxSize int) string {
	if len([]rune(s)) > maxSize {
		return Substr(s, 0, maxSize-3) + "..."
	}

	return s
}

// Head return n first symbols from given string
func Head(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}

	l := len(s)

	if l <= n {
		return s
	}

	return s[:n]
}

// Tail return n last symbols from given string
func Tail(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}

	l := len(s)

	if l <= n {
		return s
	}

	return s[l-n:]
}

// PrefixSize return prefix size
func PrefixSize(str string, prefix rune) int {
	if str == "" {
		return 0
	}

	var result int

	for i := 0; i < len(str); i++ {
		if rune(str[i]) != prefix {
			return result
		}

		result++
	}

	return result
}

// SuffixSize return suffix size
func SuffixSize(str string, suffix rune) int {
	if str == "" {
		return 0
	}

	var result int

	for i := len(str) - 1; i >= 0; i-- {
		if rune(str[i]) != suffix {
			return result
		}

		result++
	}

	return result
}
