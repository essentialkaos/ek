//go:build !windows
// +build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	PATH "path"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ListingFilter is struct with properties for filtering listing output
type ListingFilter struct {
	MatchPatterns    []string // Slice with shell file name patterns
	NotMatchPatterns []string // Slice with shell file name patterns

	ATimeOlder   int64 // Files with ATime less or equal to defined timestamp (BEFORE date)
	ATimeYounger int64 // Files with ATime greater or equal to defined timestamp (AFTER date)
	CTimeOlder   int64 // Files with CTime less or equal to defined timestamp (BEFORE date)
	CTimeYounger int64 // Files with CTime greater or equal to defined timestamp (AFTER date)
	MTimeOlder   int64 // Files with MTime less or equal to defined timestamp (BEFORE date)
	MTimeYounger int64 // Files with MTime greater or equal to defined timestamp (AFTER date)

	SizeLess    int64 // Files with size less than defined
	SizeGreater int64 // Files with size greater than defined
	SizeEqual   int64 // Files with size equals to defined
	SizeZero    bool  // Empty files

	Perms    string // Permission (see fsutil.CheckPerms for more info)
	NotPerms string // Permission (see fsutil.CheckPerms for more info)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (lf ListingFilter) hasMatchPatterns() bool {
	return len(lf.MatchPatterns) != 0
}

func (lf ListingFilter) hasNotMatchPatterns() bool {
	return len(lf.NotMatchPatterns) != 0
}

func (lf ListingFilter) hasTimes() bool {
	switch {
	case lf.ATimeOlder != 0,
		lf.ATimeOlder != 0,
		lf.ATimeYounger != 0,
		lf.CTimeOlder != 0,
		lf.CTimeYounger != 0,
		lf.MTimeOlder != 0,
		lf.MTimeYounger != 0:
		return true
	}

	return false
}

func (lf ListingFilter) hasPerms() bool {
	return lf.Perms != "" || lf.NotPerms != ""
}

func (lf ListingFilter) hasSize() bool {
	return lf.SizeZero || lf.SizeGreater > 0 || lf.SizeLess > 0 || lf.SizeEqual > 0
}

// ////////////////////////////////////////////////////////////////////////////////// //

// List is lightweight method for listing directory
func List(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	var names = readDir(dir)

	if ignoreHidden {
		names = filterHidden(names)
	}

	if len(filters) != 0 {
		names = filterList(names, dir, filters[0])
	}

	return names
}

// ListAll is lightweight method for listing all files and directories
func ListAll(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecAll(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecAll(dir, "", ignoreHidden, filters[0])
}

// ListAllDirs is lightweight method for listing all directories
func ListAllDirs(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecDirs(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecDirs(dir, "", ignoreHidden, filters[0])
}

// ListAllFiles is lightweight method for listing all files
func ListAllFiles(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecFiles(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecFiles(dir, "", ignoreHidden, filters[0])
}

// ListToAbsolute converts slice with relative paths to slice with absolute paths
func ListToAbsolute(path string, list []string) {
	for i, t := range list {
		list[i] = path + "/" + t
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readDir(dir string) []string {
	fd, err := syscall.Open(dir, syscall.O_CLOEXEC, 0644)

	if err != nil {
		return nil
	}

	defer syscall.Close(fd)

	var size = 100
	var n = -1

	var nbuf int
	var bufp int

	var buf = make([]byte, 4096)
	var names = make([]string, 0, size)

	for n != 0 {
		if bufp >= nbuf {
			bufp = 0

			var errno error

			nbuf, errno = fixCount(syscall.ReadDirent(fd, buf))

			if errno != nil {
				return names
			}

			if nbuf <= 0 {
				break
			}
		}

		var nb, nc int
		nb, nc, names = syscall.ParseDirent(buf[bufp:nbuf], n, names)
		bufp += nb
		n -= nc
	}

	return names
}

func readDirRecAll(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		if !IsDir(path + "/" + name) {
			if base == "" {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, name)
				}
			} else {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, base+"/"+name)
				}
			}
		} else {
			if base == "" {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, name)
					result = append(result, readDirRecAll(path+"/"+name, name, ignoreHidden, filter)...)
				}
			} else {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, base+"/"+name)
					result = append(result, readDirRecAll(path+"/"+name, base+"/"+name, ignoreHidden, filter)...)
				}
			}
		}
	}

	return result
}

func readDirRecDirs(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		if IsDir(path + "/" + name) {
			if base == "" {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, name)
					result = append(result, readDirRecDirs(path+"/"+name, name, ignoreHidden, filter)...)
				}
			} else {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, base+"/"+name)
					result = append(result, readDirRecDirs(path+"/"+name, base+"/"+name, ignoreHidden, filter)...)
				}
			}
		}
	}

	return result
}

func readDirRecFiles(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		if IsDir(path + "/" + name) {
			if base == "" {
				result = append(result, readDirRecFiles(path+"/"+name, name, ignoreHidden, filter)...)
			} else {
				result = append(result, readDirRecFiles(path+"/"+name, base+"/"+name, ignoreHidden, filter)...)
			}
		} else {
			if base == "" {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, name)
				}
			} else {
				if isMatch(name, path+"/"+name, filter) {
					result = append(result, base+"/"+name)
				}
			}
		}
	}

	return result
}

// It's ok to have long function with many conditions to filter some entities
// codebeat:disable[LOC,ABC,CYCLO]

func isMatch(name, fullPath string, filter ListingFilter) bool {
	var (
		hasNotMatchPatterns = filter.hasNotMatchPatterns()
		hasMatchPatterns    = filter.hasMatchPatterns()
		hasTimes            = filter.hasTimes()
		hasPerms            = filter.hasPerms()
		hasSize             = filter.hasSize()
	)

	if !hasNotMatchPatterns && !hasMatchPatterns && !hasTimes && !hasPerms && !hasSize {
		return true
	}

	var match = true

	if hasNotMatchPatterns {
		for _, pattern := range filter.NotMatchPatterns {
			matched, _ := PATH.Match(pattern, name)

			if matched {
				match = false
				break
			}
		}
	}

	if hasMatchPatterns {
		for _, pattern := range filter.MatchPatterns {
			matched, _ := PATH.Match(pattern, name)

			if matched {
				match = true
				break
			}

			match = false
		}
	}

	if !hasTimes && !hasPerms && !hasSize {
		return match
	}

	if hasTimes {
		atime, mtime, ctime, err := GetTimestamps(fullPath)

		if err != nil {
			return match
		}

		if filter.MTimeYounger != 0 {
			match = match && mtime >= filter.MTimeYounger
		}

		if filter.MTimeOlder != 0 {
			match = match && mtime <= filter.MTimeOlder
		}

		if filter.CTimeYounger != 0 {
			match = match && ctime >= filter.CTimeYounger
		}

		if filter.CTimeOlder != 0 {
			match = match && ctime <= filter.CTimeOlder
		}

		if filter.ATimeYounger != 0 {
			match = match && atime >= filter.ATimeYounger
		}

		if filter.ATimeOlder != 0 {
			match = match && atime <= filter.ATimeOlder
		}
	}

	if !hasPerms && !hasSize {
		return match
	}

	if hasPerms {
		if filter.Perms != "" {
			match = match && CheckPerms(filter.Perms, fullPath)
		}

		if filter.NotPerms != "" {
			match = match && !CheckPerms(filter.NotPerms, fullPath)
		}
	}

	if hasSize {
		if filter.SizeZero {
			match = match && GetSize(fullPath) == 0
		} else {
			if filter.SizeEqual > 0 {
				match = match && GetSize(fullPath) == filter.SizeEqual
			}

			if filter.SizeGreater > 0 {
				match = match && GetSize(fullPath) > filter.SizeGreater
			}

			if filter.SizeLess > 0 {
				match = match && GetSize(fullPath) < filter.SizeLess
			}
		}
	}

	return match
}

// codebeat:enable[LOC,ABC,CYCLO]

func filterList(names []string, dir string, filter ListingFilter) []string {
	var filteredNames []string

	for _, name := range names {
		if isMatch(name, dir+"/"+name, filter) {
			filteredNames = append(filteredNames, name)
		}
	}

	return filteredNames
}

func filterHidden(names []string) []string {
	var filteredNames []string

	for _, name := range names {
		if name[0] == '.' {
			continue
		}

		filteredNames = append(filteredNames, name)
	}

	return filteredNames
}

func fixCount(n int, err error) (int, error) {
	if n < 0 {
		n = 0
	}

	return n, err
}
