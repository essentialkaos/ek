// Package pluralize provides methods for pluralization
package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DefaultPluralizer holds default pluralization function
var DefaultPluralizer = En

// ////////////////////////////////////////////////////////////////////////////////// //

// Pluralize is simple method for pluralization
func Pluralize(n int, data ...string) string {
	return PluralizeSpecial(DefaultPluralizer, n, data...)
}

// PluralizeSpecial is method which can be used for custom pluralization
func PluralizeSpecial(p Pluralizer, n int, data ...string) string {
	return strconv.Itoa(n) + " " + safeSliceGet(data, p(n))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func safeSliceGet(data []string, index int) string {
	if len(data) < index {
		return ""
	}

	return data[index]
}
