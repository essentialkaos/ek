// Package sortutil provides methods for sorting slices
package sortutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type versionSlice []string
type stringSlice []string

// ////////////////////////////////////////////////////////////////////////////////// //

func (s versionSlice) Len() int      { return len(s) }
func (s versionSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s versionSlice) Less(i, j int) bool {
	return VersionCompare(s[i], s[j])
}

func (s stringSlice) Len() int      { return len(s) }
func (s stringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s stringSlice) Less(i, j int) bool {
	return strings.ToLower(s[i]) < strings.ToLower(s[j])
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Versions sorts versions slice
func Versions(s []string) {
	sort.Sort(versionSlice(s))
}

// VersionCompare compares 2 versions and returns true if v1 less v2. This function
// can be used for version sorting with structs
func VersionCompare(v1, v2 string) bool {
	is := strings.Split(v1, ".")
	js := strings.Split(v2, ".")

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
		}
	}

	return true
}

// Strings sorts strings slice and support case insensitive mode
func Strings(s []string, caseInsensitive bool) {
	if caseInsensitive {
		sort.Sort(stringSlice(s))
	} else {
		sort.Strings(s)
	}
}
