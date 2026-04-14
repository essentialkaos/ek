package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StringsNatural sorts a slice of strings using natural (human-friendly) order
func StringsNatural(s []string) {
	if len(s) <= 1 {
		return
	}

	sort.Slice(s, func(i, j int) bool {
		return NaturalLess(s[i], s[j])
	})
}

// NaturalLess returns true if s1 is naturally less than s2, so that e.g.
// "abc2" < "abc12". Only ASCII digits (0–9) are considered numeric.
// This code based on sortorder package created by @fvbommel
func NaturalLess(s1, s2 string) bool {
	i1, i2 := 0, 0
	l1, l2 := len(s1), len(s2)

	for i1 < l1 && i2 < l2 {
		c1, c2 := s1[i1], s2[i2]
		d1, d2 := isDigit(c1), isDigit(c2)

		if d1 != d2 {
			return d1
		}

		if !d1 {
			if c1 != c2 {
				return c1 < c2
			}

			i1++
			i2++

			continue
		}

		for i1 < l1 && s1[i1] == '0' {
			i1++
		}

		for i2 < l2 && s2[i2] == '0' {
			i2++
		}

		n1, n2 := i1, i2

		for i1 < l1 && isDigit(s1[i1]) {
			i1++
		}

		for i2 < l2 && isDigit(s2[i2]) {
			i2++
		}

		ln1, ln2 := i1-n1, i2-n2

		if ln1 != ln2 {
			return ln1 < ln2
		}

		nl1, nl2 := s1[n1:i1], s2[n2:i2]

		if nl1 != nl2 {
			return nl1 < nl2
		}
	}

	return l1 < l2
}

// ////////////////////////////////////////////////////////////////////////////////// //

// isDigit checks if a byte is a digit (0-9)
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
