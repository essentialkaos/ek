// Package strutil provides methods for working with strings
package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

// Q is simple helper for working with default values
func Q(v ...string) string {
	for _, k := range v {
		if k != "" {
			return k
		}
	}

	return ""
}

// B is shorthand for choosing value by condition
func B(cond bool, positive, negative string) string {
	if cond {
		return positive
	}

	return negative
}

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

	var result strings.Builder

SOURCELOOP:
	for _, sourceSym := range source {
		for _, fromSym := range from {
			if fromSym == sourceSym {
				result.WriteString(to)
				continue SOURCELOOP
			}
		}

		result.WriteRune(sourceSym)
	}

	return result.String()
}

// Exclude excludes substring from given string
func Exclude(data, substr string) string {
	if len(data) == 0 || len(substr) == 0 {
		return data
	}

	return strings.ReplaceAll(data, substr, "")
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
	var result []string
	var item strings.Builder
	var waitChar rune

	for _, char := range data {
		switch char {
		case '"', '\'', '`', '“', '”', '‘', '’', '«', '»', '„':
			if waitChar == 0 {
				waitChar = getClosingChar(char)
			} else if waitChar != 0 && waitChar == char {
				result = appendField(result, item.String())
				item.Reset()
				waitChar = 0
			} else {
				item.WriteRune(char)
			}

		case ',', ';', ' ':
			if waitChar != 0 {
				item.WriteRune(char)
			} else {
				result = appendField(result, item.String())
				item.Reset()
			}

		default:
			item.WriteRune(char)
		}
	}

	if item.Len() != 0 {
		result = appendField(result, item.String())
	}

	return result
}

// Copy is method for force string copying
func Copy(v string) string {
	return (v + " ")[:len(v)]
}

// Before returns part of the string before given substring
func Before(s, substr string) string {
	if !strings.Contains(s, substr) {
		return s
	}

	return s[:strings.Index(s, substr)]
}

// After returns part of the string after given substring
func After(s, substr string) string {
	if !strings.Contains(s, substr) {
		return s
	}

	return s[strings.Index(s, substr)+len(substr):]
}

// HasPrefixAny tests whether the string s begins with one of given prefixes
func HasPrefixAny(s string, prefix ...string) bool {
	for _, prf := range prefix {
		if strings.HasPrefix(s, prf) {
			return true
		}
	}

	return false
}

// HasSuffixAny tests whether the string s ends with one of given suffixes
func HasSuffixAny(s string, suffix ...string) bool {
	for _, suf := range suffix {
		if strings.HasSuffix(s, suf) {
			return true
		}
	}

	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

func appendField(data []string, item string) []string {
	if strings.TrimSpace(item) == "" {
		return data
	}

	return append(data, item)
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
