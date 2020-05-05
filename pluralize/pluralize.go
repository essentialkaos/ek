// Package pluralize provides methods for pluralization
package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultPluralizer holds default pluralization function
var DefaultPluralizer = En

// ////////////////////////////////////////////////////////////////////////////////// //

// P pluralizes a word based on the passed number with custom format
func P(format string, n interface{}, data ...string) string {
	return PS(DefaultPluralizer, format, n, data...)
}

// PS pluralizes a word based on the passed number with custom pluralizer and format
func PS(p Pluralizer, format string, n interface{}, data ...string) string {
	nk, ok := convertNumber(n)

	if !ok {
		return format
	}

	if isNumberFirst(format) {
		return fmt.Sprintf(format, n, PluralizeSpecial(p, nk, data...))
	}

	return fmt.Sprintf(format, PluralizeSpecial(p, nk, data...), n)
}

// Pluralize pluralizes a word based on the passed number
func Pluralize(n int, data ...string) string {
	return PluralizeSpecial(DefaultPluralizer, n, data...)
}

// PluralizeSpecial pluralizes a word based on the passed number with custom pluralizer
func PluralizeSpecial(p Pluralizer, n int, data ...string) string {
	return safeSliceGet(data, p(n))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func safeSliceGet(data []string, index int) string {
	if len(data) < index {
		return ""
	}

	return data[index]
}

func convertNumber(n interface{}) (int, bool) {
	switch u := n.(type) {
	case int32:
		return int(u), true
	case int64:
		return int(u), true
	case uint:
		return int(u), true
	case uint32:
		return int(u), true
	case uint64:
		return int(u), true
	case float32:
		return int(u), true
	case float64:
		return int(u), true
	case int:
		return n.(int), true
	default:
		return 0, false
	}
}

func isNumberFirst(format string) bool {
	return strings.Index(format, "%s") == strings.LastIndex(format, "%")
}
