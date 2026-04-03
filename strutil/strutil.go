// Package strutil provides methods for working with strings
package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// EllipsisSuffix is the string appended when truncating with Ellipsis
var EllipsisSuffix = "..."

// ////////////////////////////////////////////////////////////////////////////////// //

var defaultFieldsSeparators = []rune{' ', '\t'}

// ////////////////////////////////////////////////////////////////////////////////// //

// Q returns the first non-empty string from the given arguments
func Q(v ...string) string {
	for _, k := range v {
		if k != "" {
			return k
		}
	}

	return ""
}

// B returns positive if cond is true, otherwise returns negative
func B(cond bool, positive, negative string) string {
	if cond {
		return positive
	}

	return negative
}

// Concat concatenates all given strings into a single string
func Concat(s ...string) string {
	var buffer bytes.Buffer

	for _, v := range s {
		buffer.WriteString(v)
	}

	return buffer.String()
}

// Substr returns the part of s starting at rune index start with the given
// length. Negative start is treated as 0. Unicode is supported.
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

// Substring returns the part of s between rune indices start and end.
// Negative end is counted from the end of the string. Unicode is supported.
func Substring(s string, start, end int) string {
	if s == "" || start == end || start >= Len(s) {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = Len(s) + end
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

// Extract returns a byte-indexed substring of s between start and end.
// Unlike [Substr] and [Substring], Unicode is not supported.
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

// Len returns the number of Unicode code points in s
func Len(s string) int {
	return utf8.RuneCountInString(s)
}

// LenVisual returns the number of columns required to render s using a
// monospaced font, counting wide (CJK) characters as 2 columns each.
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

// Ellipsis truncates s to maxSize runes, appending EllipsisSuffix if trimmed.
// Unicode is supported.
func Ellipsis(s string, maxSize int) string {
	if Len(s) <= maxSize {
		return s
	}

	return Substr(s, 0, maxSize-Len(EllipsisSuffix)) + EllipsisSuffix
}

// Head returns the first n runes of s. Unicode is supported.
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

// Tail returns the last n runes of s. Unicode is supported.
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

// PrefixSize returns the number of leading runes in str equal to prefix
func PrefixSize(str string, prefix rune) int {
	if str == "" {
		return 0
	}

	var result int

	for _, r := range str {
		if r != prefix {
			return result
		}

		result++
	}

	return result
}

// SuffixSize returns the number of trailing runes in str equal to suffix
func SuffixSize(str string, suffix rune) int {
	if str == "" {
		return 0
	}

	var result int

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)

		if r != suffix {
			return result
		}

		str = str[:len(str)-size]

		result++
	}

	return result
}

// ReplaceAll replaces every rune found in replset with the string to
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

// ReplaceIgnoreCase replaces all case-insensitive occurrences of from with to
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

// Exclude removes all occurrences of substr from data
func Exclude(data, substr string) string {
	if len(data) == 0 || len(substr) == 0 {
		return data
	}

	return strings.ReplaceAll(data, substr, "")
}

// ReadField returns the field at the given zero-based index from data,
// splitting on separators (default: space and tab). If multiSep is true,
// consecutive separators are treated as one.
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

// Fields splits data on whitespace, commas, and semicolons, respecting
// quoted substrings (", ', `, and Unicode quote pairs) and backslash escapes.
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
			switch {
			case waitChar == 0 && !escaped:
				waitChar = getClosingChar(char)
			case waitChar != 0 && waitChar == char && !escaped:
				result = appendField(result, buf.String())
				buf.Reset()
				waitChar = 0
			default:
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

// Copy returns an independent copy of v, preventing retention of a larger
// underlying array
func Copy(v string) string {
	b := make([]byte, len(v))

	copy(b, v)

	return string(b)
}

// Before returns the part of s before the first occurrence of substr.
// Returns s unchanged if substr is not found.
func Before(s, substr string) string {
	i := strings.Index(s, substr)

	if i == -1 {
		return s
	}

	return s[:i]
}

// After returns the part of s after the first occurrence of substr.
// Returns s unchanged if substr is not found.
func After(s, substr string) string {
	i := strings.Index(s, substr)

	if i == -1 {
		return s
	}

	return s[i+len(substr):]
}

// HasPrefixAny reports whether s begins with any of the given prefixes
func HasPrefixAny(s string, prefix ...string) bool {
	for _, prf := range prefix {
		if strings.HasPrefix(s, prf) {
			return true
		}
	}

	return false
}

// HasSuffixAny reports whether s ends with any of the given suffixes
func HasSuffixAny(s string, suffix ...string) bool {
	for _, suf := range suffix {
		if strings.HasSuffix(s, suf) {
			return true
		}
	}

	return false
}

// IndexByteSkip returns the byte index of c in s after skipping the first
// skip occurrences. A negative skip searches from the end, skipping the last
// abs(skip) occurrences.
func IndexByteSkip(s string, c byte, skip int) int {
	if skip == 0 {
		return strings.IndexByte(s, c)
	}

	counter := 0

	if skip > 0 {
		for i := range len(s) {
			if s[i] == c {
				counter++
			}

			if counter > skip {
				return i
			}
		}
	} else {
		skip *= -1

		for i := len(s) - 1; i >= 0; i-- {
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

// SqueezeRepeats replaces each run of a repeated character listed in set
// with a single instance of that character
func SqueezeRepeats(s string, set string) string {
	if s == "" || set == "" {
		return s
	}

	var buf strings.Builder
	buf.Grow(len(s))

	runes := []rune(s)

	for i, r := range runes {
		if strings.ContainsRune(set, r) && i > 0 && runes[i-1] == r {
			continue
		}
		buf.WriteRune(r)
	}

	return buf.String()
}

// Mask replaces runes between byte positions start and end with maskingRune.
// A negative end is counted from the end of s.
//
// start - the first masked symbol
// end - the first non-masked symbol
func Mask(s string, start, end int, maskingRune rune) string {
	runes := []rune(s)
	n := len(runes)

	if end < 0 {
		end = n + end
	}

	var buf bytes.Buffer

	for i, r := range runes {
		if i >= start && i < end {
			buf.WriteRune(maskingRune)
		} else {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

// JoinFunc concatenates elems into a single string separated by sep,
// applying f to each element before joining.
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

// Wrap splits s into lines of at most n runes by inserting newline characters.
// Word boundaries are not considered; splitting is purely by rune count.
func Wrap(s string, n int) string {
	if s == "" || n <= 0 || len(s) <= n {
		return s
	}

	var buf strings.Builder
	var count int

	for _, r := range s {
		if count == n {
			buf.WriteByte('\n')
			count = 0
		}

		buf.WriteRune(r)

		count++
	}

	return buf.String()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// appendField appends item to data if item is not empty
func appendField(data []string, item string) []string {
	if strings.TrimSpace(item) == "" {
		return data
	}

	return append(data, item)
}

// getClosingChar returns the closing character for the given opening character
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
