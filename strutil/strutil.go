// Package strutil provides methods for working with strings
package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// EllipsisSuffix is ellipsis suffix
var EllipsisSuffix = "..."

// ////////////////////////////////////////////////////////////////////////////////// //

var defaultFieldsSeparators = []string{" ", "\t"}

// ////////////////////////////////////////////////////////////////////////////////// //

// Concat is method for fast string concatenation
func Concat(s ...string) string {
	var buffer bytes.Buffer

	for _, v := range s {
		buffer.WriteString(v)
	}

	return buffer.String()
}

// Substr returns the part of a string between the start index and a number
// of characters after it (unicode supported)
func Substr(s string, start, length int) string {
	if s == "" || length <= 0 || start >= Len(s) {
		return ""
	}

	if start < 0 {
		start = 0
	}

	var count int
	var startIndex int

	for i := range s {
		if count == start {
			startIndex = i
		}

		if count-start == length {
			return s[startIndex:i]
		}

		count++
	}

	if startIndex != 0 {
		return s[startIndex:]
	}

	return s
}

// Substring returns the part of the string between the start and end (unicode supported)
func Substring(s string, start, end int) string {
	if s == "" || start == end || start >= Len(s) {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = start
		start = 0
	}

	var count int
	var startIndex int

	for i := range s {
		if count == start {
			startIndex = i
		}

		if count == end {
			return s[startIndex:i]
		}

		count++
	}

	if startIndex != 0 {
		return s[startIndex:]
	}

	return s
}

// Extract extracts a substring safely (unicode NOT supported)
func Extract(s string, start, end int) string {
	if s == "" || end < start {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end > len(s) {
		end = len(s)
	}

	return s[start:end]
}

// Len returns number of symbols in string (unicode supported)
func Len(s string) int {
	if s == "" {
		return 0
	}

	var count int

	for range s {
		count++
	}

	return count
}

// Ellipsis trims given string (unicode supported)
func Ellipsis(s string, maxSize int) string {
	if Len(s) <= maxSize {
		return s
	}

	return Substr(s, 0, maxSize-Len(EllipsisSuffix)) + EllipsisSuffix
}

// Head returns n first symbols from given string (unicode supported)
func Head(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}

	l := Len(s)

	if l <= n {
		return s
	}

	return Substr(s, 0, n)
}

// Tail returns n last symbols from given string (unicode supported)
func Tail(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}

	l := Len(s)

	if l <= n {
		return s
	}

	return Substr(s, l-n, l)
}

// PrefixSize returns prefix size
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

// SuffixSize returns suffix size
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

// ReplaceAll replaces all symbols in given string
func ReplaceAll(source, from, to string) string {
	if source == "" {
		return ""
	}

	var result string

SOURCELOOP:
	for _, sourceSym := range source {
		for _, fromSym := range from {
			if fromSym == sourceSym {
				result += to
				continue SOURCELOOP
			}
		}

		result += string(sourceSym)
	}

	return result
}

// Exclude excludes substring from given string
// It little bit faster than strings.ReplaceAll
func Exclude(data, substr string) string {
	if len(data) == 0 || len(substr) == 0 {
		return data
	}

	k := strings.Count(data, substr)

	if k == 0 {
		return data
	}

	b := make([]byte, len(data)-(len(substr)*k))
	p, w := 0, 0

	for i := 0; i < k; i++ {
		j := p
		j += strings.Index(data[p:], substr)
		w += copy(b[w:], data[p:j])
		p = j + len(substr)
	}

	w += copy(b[w:], data[p:])

	return string(b[0:w])
}

// ReadField reads field with given index from data
func ReadField(data string, index int, multiSep bool, separators ...string) string {
	if data == "" || index < 0 {
		return ""
	}

	if len(separators) == 0 {
		separators = defaultFieldsSeparators
	}

	curIndex, startPointer := -1, -1

MAINLOOP:
	for i, r := range data {
		for _, s := range separators {
			if r == rune(s[0]) {
				if curIndex == index {
					return data[startPointer:i]
				}

				if !multiSep {
					startPointer = i + 1
					curIndex++
					continue MAINLOOP
				}

				startPointer = -1
				continue MAINLOOP
			}
		}

		if startPointer == -1 {
			startPointer = i
			curIndex++
		}
	}

	if index > curIndex {
		return ""
	}

	return data[startPointer:]
}

// Fields splits the string data around each instance of one or more
// consecutive white space or comma characters
func Fields(data string) []string {
	var (
		result   []string
		item     string
		waitChar rune
	)

	for _, char := range data {
		switch char {
		case '"', '\'', '`', '“', '”', '‘', '’', '«', '»', '„':
			if waitChar == 0 {
				waitChar = getClosingChar(char)
			} else if waitChar != 0 && waitChar == char {
				result = append(result, item)
				item, waitChar = "", 0
			} else {
				item += string(char)
			}

		case ',', ';', ' ':
			if waitChar != 0 {
				item += string(char)
			} else {
				result = append(result, item)
				item = ""
			}

		default:
			item += string(char)
		}
	}

	if item != "" {
		result = append(result, item)
	}

	return formatItems(result)
}

// Copy is method for force string copying
func Copy(v string) string {
	return (v + " ")[:len(v)]
}

// ////////////////////////////////////////////////////////////////////////////////// //

func formatItems(data []string) []string {
	var result []string

	for _, v := range data {
		item := strings.TrimSpace(v)

		if item != "" {
			result = append(result, item)
		}
	}

	return result
}

func getClosingChar(r rune) rune {
	switch r {
	case '“':
		return '”'
	case '„':
		return '“'
	case '‘':
		return '’'
	case '«':
		return '»'
	default:
		return r
	}
}
