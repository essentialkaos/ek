// +build windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ListingFilter is struct with properties for filtering listing output
type ListingFilter struct {
	MatchPatterns    []string
	NotMatchPatterns []string

	ATimeOlder   int64
	ATimeYounger int64
	CTimeOlder   int64
	CTimeYounger int64
	MTimeOlder   int64
	MTimeYounger int64

	SizeLess    int64
	SizeGreater int64
	SizeEqual   int64
	SizeZero    bool

	Perms    string
	NotPerms string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func List(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	return []string{}
}

func ListAll(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	return []string{}
}

func ListAllDirs(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	return []string{}
}

func ListAllFiles(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	return []string{}
}

func ListToAbsolute(path string, list []string) {
	return
}
