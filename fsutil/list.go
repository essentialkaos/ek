//go:build linux || darwin || freebsd

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	PATH "path"
	"path/filepath"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ListingFilter is struct with properties for filtering listing output
type ListingFilter struct {
	MatchPatterns    []string // Shell patterns the entry name must match
	NotMatchPatterns []string // Shell patterns the entry name must not match

	ATimeOlder   int64 // Entries with ATime ≤ this Unix timestamp (before date)
	ATimeYounger int64 // Entries with ATime ≥ this Unix timestamp (after date)
	CTimeOlder   int64 // Entries with CTime ≤ this Unix timestamp (before date)
	CTimeYounger int64 // Entries with CTime ≥ this Unix timestamp (after date)
	MTimeOlder   int64 // Entries with MTime ≤ this Unix timestamp (before date)
	MTimeYounger int64 // Entries with MTime ≥ this Unix timestamp (after date)

	SizeLess    int64 // Entries with size strictly less than this value in bytes
	SizeGreater int64 // Entries with size strictly greater than this value in bytes
	SizeEqual   int64 // Entries with size exactly equal to this value in bytes
	SizeZero    bool  // Entries with a size of zero bytes

	Perms    string // Entries that satisfy these permission checks (see [CheckPerms])
	NotPerms string // Entries that do not satisfy these permission checks (see [CheckPerms])
}

// ////////////////////////////////////////////////////////////////////////////////// //

// hasMatchPatterns checks if filter has match patterns
func (lf ListingFilter) hasMatchPatterns() bool {
	return len(lf.MatchPatterns) != 0
}

// hasNotMatchPatterns checks if filter has not-match patterns
func (lf ListingFilter) hasNotMatchPatterns() bool {
	return len(lf.NotMatchPatterns) != 0
}

// hasTimes checks if filter has time-related properties
func (lf ListingFilter) hasTimes() bool {
	switch {
	case lf.ATimeOlder != 0,
		lf.ATimeYounger != 0,
		lf.CTimeOlder != 0,
		lf.CTimeYounger != 0,
		lf.MTimeOlder != 0,
		lf.MTimeYounger != 0:
		return true
	}

	return false
}

// hasPerms checks if filter has permission-related properties
func (lf ListingFilter) hasPerms() bool {
	return lf.Perms != "" || lf.NotPerms != ""
}

// hasSize checks if filter has size-related properties
func (lf ListingFilter) hasSize() bool {
	return lf.SizeZero || lf.SizeGreater > 0 || lf.SizeLess > 0 || lf.SizeEqual > 0
}

// ////////////////////////////////////////////////////////////////////////////////// //

// List returns the names of all entries in dir.
// Hidden entries (starting with '.') are excluded when ignoreHidden is true.
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

// ListAll recursively returns the names of all files and directories under dir.
// Hidden entries are excluded when ignoreHidden is true.
func ListAll(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecAll(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecAll(dir, "", ignoreHidden, filters[0])
}

// ListAllDirs recursively returns the names of all directories under dir.
// Hidden entries are excluded when ignoreHidden is true.
func ListAllDirs(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecDirs(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecDirs(dir, "", ignoreHidden, filters[0])
}

// ListAllFiles recursively returns the names of all files under dir.
// Hidden entries are excluded when ignoreHidden is true.
func ListAllFiles(dir string, ignoreHidden bool, filters ...ListingFilter) []string {
	if len(filters) == 0 {
		return readDirRecFiles(dir, "", ignoreHidden, ListingFilter{})
	}

	return readDirRecFiles(dir, "", ignoreHidden, filters[0])
}

// ListToAbsolute prepends basePath to every entry in list, converting relative
// names to absolute paths in place
func ListToAbsolute(root string, list []string) {
	for i, t := range list {
		list[i] = filepath.Join(root, t)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// readDir reads directory and returns slice with names of files and directories
func readDir(dir string) []string {
	fd, err := syscall.Open(dir, syscall.O_CLOEXEC, 0)

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

// readDirRecAll reads directory recursively and returns slice with names of files
// and directories
func readDirRecAll(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		targetPath := filepath.Join(path, name)

		if !IsDir(targetPath) {
			if base == "" {
				if isMatch(name, targetPath, filter) {
					result = append(result, name)
				}
			} else {
				if isMatch(name, targetPath, filter) {
					result = append(result, filepath.Join(base, name))
				}
			}
		} else {
			if base == "" {
				if isMatch(name, targetPath, filter) {
					result = append(result, name)
					result = append(result, readDirRecAll(targetPath, name, ignoreHidden, filter)...)
				}
			} else {
				if isMatch(name, targetPath, filter) {
					basePath := filepath.Join(base, name)
					result = append(result, basePath)
					result = append(result, readDirRecAll(targetPath, basePath, ignoreHidden, filter)...)
				}
			}
		}
	}

	return result
}

// readDirRecAll reads directory recursively and returns slice with names of directories
func readDirRecDirs(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		targetPath := filepath.Join(path, name)

		if IsDir(targetPath) {
			if base == "" {
				if isMatch(name, targetPath, filter) {
					result = append(result, name)
					result = append(result, readDirRecDirs(targetPath, name, ignoreHidden, filter)...)
				}
			} else {
				if isMatch(name, targetPath, filter) {
					basePath := filepath.Join(base, name)
					result = append(result, basePath)
					result = append(result, readDirRecDirs(targetPath, basePath, ignoreHidden, filter)...)
				}
			}
		}
	}

	return result
}

// readDirRecDirs reads directory recursively and returns slice with names of files
func readDirRecFiles(path, base string, ignoreHidden bool, filter ListingFilter) []string {
	var result []string

	names := readDir(path)

	for _, name := range names {
		if name[0] == '.' && ignoreHidden {
			continue
		}

		targetPath := filepath.Join(path, name)

		if IsDir(targetPath) {
			if base == "" {
				result = append(result, readDirRecFiles(targetPath, name, ignoreHidden, filter)...)
			} else {
				result = append(result, readDirRecFiles(targetPath, filepath.Join(base, name), ignoreHidden, filter)...)
			}
		} else {
			if base == "" {
				if isMatch(name, targetPath, filter) {
					result = append(result, name)
				}
			} else {
				if isMatch(name, targetPath, filter) {
					result = append(result, filepath.Join(base, name))
				}
			}
		}
	}

	return result
}

// isMatch checks if file or directory matches filter
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
				return false
			}
		}
	}

	if hasMatchPatterns {
		match = false

		for _, pattern := range filter.MatchPatterns {
			matched, _ := PATH.Match(pattern, name)

			if matched {
				match = true
				break
			}
		}

		if !match {
			return false
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

// filterList filters slice with names of files and directories using given filter
func filterList(names []string, dir string, filter ListingFilter) []string {
	var filteredNames []string

	for _, name := range names {
		if isMatch(name, filepath.Join(dir, name), filter) {
			filteredNames = append(filteredNames, name)
		}
	}

	return filteredNames
}

// filterHidden filters out hidden files and directories (those that start with a dot)
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

// fixCount ensures that the count from [syscall.ReadDirent] is non-negative
func fixCount(n int, err error) (int, error) {
	if n < 0 {
		n = 0
	}

	return n, err
}
