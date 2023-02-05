package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ListingFilter is struct with properties for filtering listing output
type ListingFilter struct {
	MatchPatterns    []string // ❗ Slice with shell file name patterns
	NotMatchPatterns []string // ❗ Slice with shell file name patterns

	ATimeOlder   int64 // ❗ Files with ATime less or equal to defined timestamp (BEFORE date)
	ATimeYounger int64 // ❗ Files with ATime greater or equal to defined timestamp (AFTER date)
	CTimeOlder   int64 // ❗ Files with CTime less or equal to defined timestamp (BEFORE date)
	CTimeYounger int64 // ❗ Files with CTime greater or equal to defined timestamp (AFTER date)
	MTimeOlder   int64 // ❗ Files with MTime less or equal to defined timestamp (BEFORE date)
	MTimeYounger int64 // ❗ Files with MTime greater or equal to defined timestamp (AFTER date)

	SizeLess    int64 // ❗ Files with size less than defined
	SizeGreater int64 // ❗ Files with size greater than defined
	SizeEqual   int64 // ❗ Files with size equals to defined
	SizeZero    bool  // ❗ Empty files

	Perms    string // ❗ Permission (see fsutil.CheckPerms for more info)
	NotPerms string // ❗ Permission (see fsutil.CheckPerms for more info)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ List is lightweight method for listing directory
func List(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ ListAll is lightweight method for listing all files and directories
func ListAll(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ ListAllDirs is lightweight method for listing all directories
func ListAllDirs(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ ListAllFiles is lightweight method for listing all files
func ListAllFiles(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	panic("UNSUPPORTED")
	return nil
}

// ❗ ListToAbsolute converts slice with relative paths to slice with absolute paths
func ListToAbsolute(path string, list []string) {
	panic("UNSUPPORTED")
}
