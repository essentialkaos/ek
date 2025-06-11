package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StringsNatural sorts a slice of strings in natural order
// Limitation: only ASCII digits (0-9) are considered.
func StringsNatural(s []string) {
	if len(s) <= 1 {
		return
	}

	sort.Slice(s, func(i, j int) bool {
		return NaturalLess(s[i], s[j])
	})
}

// NaturalLess compares two strings using natural ordering. This means that e.g.
// "abc2" < "abc12"
// This code based on sortorder package created by @fvbommel
func NaturalLess(s1, s2 string) bool {
	i1, i2 := 0, 0
	l1, l2 := len(s1), len(s2)

	for i1 < l1 && i2 < l2 {
		c1, c2 := s1[i1], s2[i2]
		d1, d2 := isDigit(c1), isDigit(c2)

		if d1 != d2 {
			return d1
		} else if !d1 {
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

func isDigit(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}
