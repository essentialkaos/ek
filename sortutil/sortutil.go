// Package with utils for sorting slices
package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type versionSlice []string

// ////////////////////////////////////////////////////////////////////////////////// //

func (s versionSlice) Len() int      { return len(s) }
func (s versionSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s versionSlice) Less(i, j int) bool {
	is := strings.Split(s[i], ".")
	js := strings.Split(s[j], ".")

	il, jl := len(is), len(js)

	l := il

	if jl > l {
		l = jl
	}

	for k := 0; k < l; k++ {
		switch {
		case il-1 < k:
			return true
		case jl-1 < k:
			return false
		}

		if is[k] == js[k] {
			continue
		}

		ii, err1 := strconv.ParseInt(is[k], 10, 64)
		ji, err2 := strconv.ParseInt(js[k], 10, 64)

		if err1 != nil || err2 != nil {
			return is[k] < js[k]
		}

		switch {
		case ii < ji:
			return true
		case ii > ji:
			return false
		default:
			continue
		}
	}

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Versions sort versions slice
func Versions(s []string) {
	sort.Sort(versionSlice(s))
}
