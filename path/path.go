//go:build !windows
// +build !windows

// Package path provides methods for working with paths (fully compatible with base path package)
package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	PATH "path"
	"path/filepath"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrBadPattern indicates a globbing pattern was malformed
var ErrBadPattern = errors.New("Syntax error in pattern")

// unsafePaths is slice with unsafe paths
var unsafePaths = []string{
	"/lost+found",
	"/bin",
	"/boot",
	"/etc",
	"/dev",
	"/lib",
	"/lib64",
	"/proc",
	"/root",
	"/sbin",
	"/selinux",
	"/sys",
	"/usr/bin",
	"/usr/lib",
	"/usr/lib64",
	"/usr/libexec",
	"/usr/sbin",
	"/usr/include",
	"/var/cache",
	"/var/db",
	"/var/lib",
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Base returns the last element of path
func Base(path string) string {
	return PATH.Base(path)
}

// Clean returns the shortest path name equivalent to path by purely lexical processing
func Clean(path string) string {
	path = evalHome(path)
	return PATH.Clean(path)
}

// Dir returns all but the last element of path, typically the path's directory
func Dir(path string) string {
	return PATH.Dir(path)
}

// DirN returns first N elements of path
func DirN(path string, n int) string {
	if len(path) <= 1 || n < 1 {
		return path
	}

	if path[0] == '/' {
		n++
	}

	var k int

	for i, r := range path {
		if r == '/' {
			k++
		}

		if k == n {
			return path[:i]
		}
	}

	return path
}

// Ext returns the file name extension used by path
func Ext(path string) string {
	return PATH.Ext(path)
}

// IsAbs reports whether the path is absolute
func IsAbs(path string) bool {
	return PATH.IsAbs(path)
}

// Join joins any number of path elements into a single path, adding a separating slash if necessary
func Join(elem ...string) string {
	return PATH.Join(elem...)
}

// Match reports whether name matches the shell file name pattern
func Match(pattern, name string) (matched bool, err error) {
	return PATH.Match(pattern, name)
}

// Split splits path immediately following the final slash, separating it into a directory and file name component
func Split(path string) (dir, file string) {
	return PATH.Split(path)
}

// Compact converts path to compact representation (e.g /some/random/directory/file.txt → /s/r/d/file.txt)
func Compact(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}

	pathSlice := strings.Split(path, "/")

	for i := 0; i < len(pathSlice)-1; i++ {
		if len(pathSlice[i]) > 1 {
			pathSlice[i] = pathSlice[i][0:1]
		}
	}

	return strings.Join(pathSlice, "/")
}

// IsSafe returns true is given path is safe to use (not points to system dirs)
func IsSafe(path string) bool {
	if path == "" {
		return false
	}

	absPath, err := filepath.Abs(Clean(path))

	if err != nil || absPath == "/" {
		return false
	}

	for _, up := range unsafePaths {
		if contains(absPath, up) {
			return false
		}
	}

	return true
}

// IsDotfile returns true if file name begins with a full stop
func IsDotfile(path string) bool {
	if path == "" {
		return false
	}

	if !strings.Contains(path, "/") {
		return path[0:1] == "."
	}

	pathBase := Base(path)

	return pathBase[0:1] == "."
}

// IsGlob returns true if given pattern is Unix-like glob
func IsGlob(pattern string) bool {
	if pattern == "" {
		return false
	}

	var rs bool

	for _, r := range pattern {
		switch r {
		case '?', '*':
			return true
		case '[':
			rs = true
		case ']':
			if rs {
				return true
			}
		}
	}

	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

func evalHome(path string) string {
	if path == "" || path[0:1] != "~" {
		return path
	}

	return os.Getenv("HOME") + path[1:]
}

func contains(path, subpath string) bool {
	spl := len(subpath)

	if len(path) < spl {
		return false
	}

	return path[:spl] == subpath
}
