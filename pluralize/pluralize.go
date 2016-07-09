// Package pluralize provides methods for pluralization
package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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
