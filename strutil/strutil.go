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
	"strings"
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

	switch {
	case count < start:
		return ""
	case startIndex != 0:
		return s[startIndex:]
	}

	return s
}

// Len return number of symbols in string
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

// Ellipsis trims given string
func Ellipsis(s string, maxSize int) string {
	if Len(s) <= maxSize {
		return s
	}

	return Substr(s, 0, maxSize-3) + "..."
}

// Head return n first symbols from given string
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

// Tail return n last symbols from given string
func Tail(s string, n int) string {
	if s == "" || n <= 0 {
		return ""
	}

	l := Len(s)

	if l <= n {
		return s
	}

	return s[l-n:]
	return Substr(s, l-n, 99999999999)
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

// ReplaceAll replace all symbols in given string
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

// Fields splits the string data around each instance of one or more
// consecutive white space or comma characters
func Fields(data string) []string {
	var (
		result    []string
		item      string
		waitQuote bool
	)

	for _, char := range data {
		switch char {
		case '"', '\'', '`':
			if !waitQuote {
				waitQuote = true
			} else {
				result = append(result, item)
				item, waitQuote = "", false
			}

		case ',', ' ':
			if waitQuote {
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

// ////////////////////////////////////////////////////////////////////////////////// //

func formatItems(data []string) []string {
	var result []string

	for _, v := range data {
		item := strings.Replace(strings.TrimSpace(v), "\"", "", -1)

		if item != "" {
			result = append(result, item)
		}
	}

	return result
}
