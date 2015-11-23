// Package provides methods for parsing semver version info
package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Version struct {
	raw          string
	versionSlice []int
	preRelease   string
	build        string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Parse(ver string) *Version {
	if ver == "" {
		return &Version{}
	}

	result := &Version{raw: ver}

	if strings.Contains(ver, "+") {
		bs := strings.Split(ver, "+")

		if bs[1] == "" {
			ver = bs[0]
		} else {
			ver = bs[0]
			result.build = bs[1]
		}
	}

	if strings.Contains(ver, "-") {
		ps := strings.Split(ver, "-")

		if ps[1] == "" {
			ver = ps[0]
		} else {
			ver = ps[0]
			result.preRelease = ps[1]
		}
	}

	for _, version := range strings.Split(ver, ".") {
		iv, err := strconv.Atoi(version)

		if err != nil {
			break
		}

		result.versionSlice = append(result.versionSlice, iv)
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Major return major version
func (v *Version) Major() int {
	if v == nil || v.raw == "" || len(v.versionSlice) == 0 {
		return -1
	}

	return v.versionSlice[0]
}

// Minor return minor version
func (v *Version) Minor() int {
	if v == nil || v.raw == "" || len(v.versionSlice) == 0 {
		return -1
	}

	if len(v.versionSlice) == 1 {
		return 0
	}

	return v.versionSlice[1]
}

// Patch return patch version
func (v *Version) Patch() int {
	if v == nil || v.raw == "" || len(v.versionSlice) == 0 {
		return -1
	}

	if len(v.versionSlice) <= 2 {
		return 0
	}

	return v.versionSlice[2]
}

// PreRelease return prerelease version
func (v *Version) PreRelease() string {
	if v == nil || v.raw == "" {
		return ""
	}

	return v.preRelease
}

// Build return build
func (v *Version) Build() string {
	if v == nil || v.raw == "" {
		return ""
	}

	return v.build
}

// String return version as string
func (v *Version) String() string {
	if v == nil {
		return ""
	}

	return v.raw
}
