// Package strutil provides methods for working with strings
package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

var defaultFieldsSeparators = []rune{' ', '\t'}

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

// LenVisual returns number of space required for rendering given string using
// monospaced font.
//
// Warning: This method can be inaccurate in some cases, use with care
func LenVisual(s string) int {
	if s == "" {
		return 0
	}

	var count int

	for _, r := range s {
		if r > 11263 {
			count += 2
		} else {
			count++
		}
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

// ReplaceAll replaces all symbols from replset in string
func ReplaceAll(source, replset, to string) string {
	if source == "" {
		return ""
	}

	var result strings.Builder

SOURCELOOP:
	for _, sourceSym := range source {
		for _, fromSym := range replset {
			if fromSym == sourceSym {
				result.WriteString(to)
				continue SOURCELOOP
			}
		}

		result.WriteRune(sourceSym)
	}

	return result.String()
}

// ReplaceIgnoreCase replaces part of the string ignoring case
func ReplaceIgnoreCase(source, from, to string) string {
	if source == "" || from == "" {
		return source
	}

	var result strings.Builder

	from = strings.ToLower(from)
	lowSource := strings.ToLower(source)

	for {
		index := strings.Index(lowSource, from)

		if index == -1 {
			result.WriteString(source)
			break
		}

		result.WriteString(source[:index])
		result.WriteString(to)

		source = source[index+len(from):]
		lowSource = lowSource[index+len(from):]
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
func ReadField(data string, index int, multiSep bool, separators ...rune) string {
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
			if r == s {
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
	var buf bytes.Buffer
	var waitChar rune
	var escaped bool

	for _, char := range data {
		switch char {
		case '\\':
			buf.WriteRune(char)
			escaped = true
		case '"', '\'', '`', '“', '”', '‘', '’', '«', '»', '„':
			if waitChar == 0 && !escaped {
				waitChar = getClosingChar(char)
			} else if waitChar != 0 && waitChar == char && !escaped {
				result = appendField(result, buf.String())
				buf.Reset()
				waitChar = 0
			} else {
				if escaped && buf.Len() > 0 {
					buf.Truncate(buf.Len() - 1)
				}
				buf.WriteRune(char)
				escaped = false
			}

		case ',', ';', ' ':
			if waitChar != 0 {
				buf.WriteRune(char)
				escaped = false
			} else {
				result = appendField(result, buf.String())
				buf.Reset()
			}

		default:
			buf.WriteRune(char)
			escaped = false
		}
	}

	if buf.Len() != 0 {
		result = appendField(result, buf.String())
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

// IndexByteSkip returns the index of the given byte in the string after skipping
// the first or the last N occurrences
func IndexByteSkip(s string, c byte, skip int) int {
	if skip == 0 {
		return strings.IndexByte(s, c)
	}

	counter := 0

	if skip > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				counter++
			}

			if counter > skip {
				return i
			}
		}
	} else {
		skip *= -1

		for i := len(s) - 1; i > 0; i-- {
			if s[i] == c {
				counter++
			}

			if counter > skip {
				return i
			}
		}
	}

	return -1
}

// SqueezeRepeats replaces each sequence of a repeated character that is listed in
// the specified set
func SqueezeRepeats(s string, set string) string {
	if s == "" || set == "" {
		return s
	}

	for _, r := range set {
		l, rs := len(s), string(r)

		for {
			s = strings.ReplaceAll(s, rs+rs, rs)

			if len(s) == l {
				break
			}

			l = len(s)
		}
	}

	return s
}

// Mask masks part of the given string using given symbol
//
// start - the first masked symbol
// end - the first non-masked symbol
func Mask(s string, start, end int, maskingRune rune) string {
	var buffer bytes.Buffer

	if end < 0 {
		end = len(s) + end
	}

	for i, r := range s {
		if i >= start && i < end {
			buffer.WriteRune(maskingRune)
		} else {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
}

// JoinFunc concatenates the elements of its first argument into a single string,
// and modifies each element using the given function
func JoinFunc(elems []string, sep string, f func(s string) string) string {
	if len(elems) == 0 {
		return ""
	}

	var buf bytes.Buffer

	buf.WriteString(f(elems[0]))

	for _, s := range elems[1:] {
		buf.WriteString(sep)
		buf.WriteString(f(s))
	}

	return buf.String()
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
