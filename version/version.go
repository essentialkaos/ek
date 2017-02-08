// Package version provides methods for parsing semver version info
package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Version contains version data
type Version struct {
	raw          string
	versionSlice []int
	preRelease   string
	build        string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var preRegExp = regexp.MustCompile(`([a-zA-Z-.]{1,})([0-9]{0,})`)

// ////////////////////////////////////////////////////////////////////////////////// //

// Parse parse version string and return version struct
func Parse(v string) (*Version, error) {
	if v == "" {
		return nil, errors.New("Version can't be empty")
	}

	result := &Version{raw: v}

	if strings.Contains(v, "+") {
		bs := strings.Split(v, "+")

		if bs[1] == "" {
			return nil, errors.New("Build number is empty")
		} else {
			v = bs[0]
			result.build = bs[1]
		}
	}

	if strings.Contains(v, "-") {
		ps := strings.Split(v, "-")

		if ps[1] == "" {
			return nil, errors.New("Prerelease number is empty")
		} else {
			v = ps[0]
			result.preRelease = ps[1]
		}
	}

	for _, version := range strings.Split(v, ".") {
		iv, err := strconv.Atoi(version)

		if err != nil {
			return nil, err
		}

		result.versionSlice = append(result.versionSlice, iv)
	}

	return result, nil
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

// Equal return true if version are equal to given
func (v *Version) Equal(version *Version) bool {
	if v.Major() != version.Major() {
		return false
	}

	if v.Minor() != version.Minor() {
		return false
	}

	if v.Patch() != version.Patch() {
		return false
	}

	if v.PreRelease() != version.PreRelease() {
		return false
	}

	if v.Build() != version.Build() {
		return false
	}

	return true
}

// Less return true if given version is greater
func (v *Version) Less(version *Version) bool {
	if v.Major() < version.Major() {
		return true
	}

	if v.Minor() < version.Minor() {
		return true
	}

	if v.Patch() < version.Patch() {
		return true
	}

	pr1, pr2 := v.PreRelease(), version.PreRelease()

	if pr1 != pr2 {
		return prereleaseLess(pr1, pr2)
	}

	return false
}

// Greater return true if given version is less
func (v *Version) Greater(version *Version) bool {
	if v.Major() > version.Major() {
		return true
	}

	if v.Minor() > version.Minor() {
		return true
	}

	if v.Patch() > version.Patch() {
		return true
	}

	pr1, pr2 := v.PreRelease(), version.PreRelease()

	if pr1 != pr2 {
		return !prereleaseLess(pr1, pr2)
	}

	return false
}

// Contains check is current version contains given version
func (v *Version) Contains(version *Version) bool {
	if v.Major() != version.Major() {
		return false
	}

	if len(v.versionSlice) == 1 {
		return true
	}

	if v.Minor() != version.Minor() {
		return false
	}

	if len(v.versionSlice) == 2 {
		return true
	}

	if v.Patch() != version.Patch() {
		return false
	}

	return false
}

// String return version as string
func (v *Version) String() string {
	if v == nil {
		return ""
	}

	return v.raw
}

// ////////////////////////////////////////////////////////////////////////////////// //

// prereleaseLess
func prereleaseLess(pr1, pr2 string) bool {
	// Current version is release and given is prerelease
	if pr1 == "" && pr2 != "" {
		return false
	}

	// Current version is prerelease and given is release
	if pr1 != "" && pr2 == "" {
		return true
	}

	// Parse prerelease
	pr1Re := preRegExp.FindStringSubmatch(pr1)
	pr2Re := preRegExp.FindStringSubmatch(pr2)

	pr1Name := pr1Re[1]
	pr2Name := pr2Re[1]

	if pr1Name > pr2Name {
		return false
	}

	if pr1Name < pr2Name {
		return true
	}

	// Errors not important, because if subver is empty
	// Atoi return 0
	pr1Ver, _ := strconv.Atoi(pr1Re[2])
	pr2Ver, _ := strconv.Atoi(pr2Re[2])

	return pr1Ver < pr2Ver
}
