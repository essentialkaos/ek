// Package pluralize provides methods for pluralization
package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v12/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultPluralizer holds default pluralization function
var DefaultPluralizer = En

// ////////////////////////////////////////////////////////////////////////////////// //

// P pluralizes a word based on the passed number with custom format
func P[N mathutil.Numeric](format string, n N, data ...string) string {
	return PS(DefaultPluralizer, format, n, data...)
}

// PS pluralizes a word based on the passed number with custom pluralizer and format
func PS[N mathutil.Numeric](p Pluralizer, format string, n N, data ...string) string {
	if isNumberFirst(format) {
		return fmt.Sprintf(format, n, PluralizeSpecial(p, convertNumber(n), data...))
	}

	return fmt.Sprintf(format, PluralizeSpecial(p, convertNumber(n), data...), n)
}

// Pluralize pluralizes a word based on the passed number
func Pluralize[N mathutil.Numeric](n N, data ...string) string {
	return PluralizeSpecial(DefaultPluralizer, n, data...)
}

// PluralizeSpecial pluralizes a word based on the passed number with custom pluralizer
func PluralizeSpecial[N mathutil.Numeric](p Pluralizer, n N, data ...string) string {
	return safeSliceGet(data, p(convertNumber(n)))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// safeSliceGet returns value from slice with given index
func safeSliceGet(data []string, index int) string {
	if len(data) < index {
		return ""
	}

	return data[index]
}

// convertNumber converts numeric to uint64
func convertNumber(n any) uint64 {
	switch u := n.(type) {
	case int:
		return uint64(mathutil.Abs(u))
	case int8:
		return uint64(mathutil.Abs(u))
	case int16:
		return uint64(mathutil.Abs(u))
	case int32:
		return uint64(mathutil.Abs(u))
	case int64:
		return uint64(mathutil.Abs(u))
	case uint:
		return uint64(u)
	case uint8:
		return uint64(u)
	case uint16:
		return uint64(u)
	case uint32:
		return uint64(u)
	case uint64:
		return uint64(u)
	case float32:
		return uint64(mathutil.Abs(u))
	case float64:
		return uint64(mathutil.Abs(u))
	}

	return 0
}

// isNumberFirst returns true if number is first in format string
func isNumberFirst(format string) bool {
	return strings.Index(format, "%s") == strings.LastIndex(format, "%")
}
