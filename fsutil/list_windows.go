// +build windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

type ListingFilter struct {
	MatchPatterns    []string
	NotMatchPatterns []string

	ATimeOlder   int64
	ATimeYounger int64
	CTimeOlder   int64
	CTimeYounger int64
	MTimeOlder   int64
	MTimeYounger int64

	Perms    string
	NotPerms string

	hasMatchPatterns    bool
	hasNotMatchPatterns bool
	hasTimes            bool
	hasPerms            bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

func List(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	return []string{}
}

func ListAll(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	return []string{}
}

func ListAllDirs(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	return []string{}
}

func ListAllFiles(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	return []string{}
}

func ListToAbsolute(path string, list []string) {
	return
}
