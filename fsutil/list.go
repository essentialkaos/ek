// +build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"path"
	"syscall"
)

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

	Perms    string
	NotPerms string

	hasMatchPatterns    bool
	hasNotMatchPatterns bool
	hasTimes            bool
	hasPerms            bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (lf *ListingFilter) init() *ListingFilter {
	if len(lf.MatchPatterns) != 0 {
		lf.hasMatchPatterns = true
	}

	if len(lf.NotMatchPatterns) != 0 {
		lf.hasNotMatchPatterns = true
	}

	switch {
	case lf.ATimeOlder != 0,
		lf.ATimeOlder != 0,
		lf.ATimeYounger != 0,
		lf.CTimeOlder != 0,
		lf.CTimeYounger != 0,
		lf.MTimeOlder != 0,
		lf.MTimeYounger != 0:
		lf.hasTimes = true
	}

	if lf.Perms != "" || lf.NotPerms != "" {
		lf.hasPerms = true
	}

	return lf
}

// ////////////////////////////////////////////////////////////////////////////////// //

// List is lightweight method for listing directory
func List(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	var names = readDir(dir)

	if ignoreHidden {
		names = filterHidden(names)
	}

	if len(filters) != 0 {
		names = filterList(names, filters[0].init())
	}

	return names
}

// ListAll is lightweight method for listing all files and directories
func ListAll(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecAll(dir, "", ignoreHidden, nil)
	}

	return readDirRecAll(dir, "", ignoreHidden, filters[0].init())
}

// ListAllDirs is lightweight method for listing all directories
func ListAllDirs(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecDirs(dir, "", ignoreHidden, nil)
	}

	return readDirRecDirs(dir, "", ignoreHidden, filters[0].init())
}

// ListAllFiles is lightweight method for listing all files
func ListAllFiles(dir string, ignoreHidden bool, filters ...*ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecFiles(dir, "", ignoreHidden, nil)
	}

	return readDirRecFiles(dir, "", ignoreHidden, filters[0].init())
}

// ListAbsolute convert slice with relative paths to slice with absolute paths
func ListAbsolute(path string, list []string) {
	for i, t := range list {
		list[i] = path + "/" + t
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readDir(dir string) []string {
	fd, err := syscall.Open(dir, syscall.O_CLOEXEC, 0644)

	if err != nil {
		return []string{}
	}

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

func readDirRecAll(path, base string, ignoreHidden bool, filter *ListingFilter) []string {
	var result = make([]string, 0)

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

func readDirRecDirs(path, base string, ignoreHidden bool, filter *ListingFilter) []string {
	var result = make([]string, 0)

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

func readDirRecFiles(path, base string, ignoreHidden bool, filter *ListingFilter) []string {
	var result = make([]string, 0)

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

func isMatch(name, fullPath string, filter *ListingFilter) bool {
	if filter == nil {
		return true
	}

	var match = true

	if filter.hasNotMatchPatterns {
		for _, pattern := range filter.NotMatchPatterns {
			matched, _ := path.Match(pattern, name)

			if matched {
				match = false
				break
			}
		}
	} else if filter.hasMatchPatterns {
		for _, pattern := range filter.MatchPatterns {
			matched, _ := path.Match(pattern, name)

			if matched {
				match = true
				break
			}

			match = false
		}
	}

	if !filter.hasTimes && !filter.hasPerms {
		return match
	}

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

	if !filter.hasPerms {
		return match
	}

	if filter.Perms != "" {
		match = match && CheckPerms(filter.Perms, fullPath) == true
	}

	if filter.NotPerms != "" {
		match = match && CheckPerms(filter.Perms, fullPath) == false
	}

	return match
}

func filterList(names []string, filter *ListingFilter) []string {
	var filteredNames []string

	for _, name := range names {
		if isMatch(name, name, filter) {
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
