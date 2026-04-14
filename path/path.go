// Package path provides methods for working with paths (fully compatible with base path package)
package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrBadPattern indicates that a globbing pattern is malformed
var ErrBadPattern = errors.New("syntax error in pattern")

// ////////////////////////////////////////////////////////////////////////////////// //

// pathSeparator is a string representation of the filepath separator
var pathSeparator = string(filepath.Separator)

// homeDirFunc is function to get user home dir
var homeDirFunc = os.UserHomeDir

// ////////////////////////////////////////////////////////////////////////////////// //

// Base returns the last element of path
func Base(path string) string {
	return filepath.Base(path)
}

// Clean returns the shortest path name equivalent to path by purely lexical
// processing, expanding a leading ~ to the user's home directory
func Clean(path string) string {
	path = evalHome(path)
	return filepath.Clean(path)
}

// Dir returns all but the last element of path, typically the path's directory
func Dir(path string) string {
	return filepath.Dir(path)
}

// Ext returns the file name extension used by path
func Ext(path string) string {
	return filepath.Ext(path)
}

// IsAbs reports whether the path is absolute
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// Join joins any number of path elements into a single path, adding a
// separating slash if necessary
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// JoinSecure joins all path elements under root, resolving symlinks at each
// step and returning an error if the final path escapes the root directory
func JoinSecure(root string, elem ...string) (string, error) {
	resolved, err := filepath.EvalSymlinks(root)

	if err == nil {
		root = resolved
	}

	result := root

	for _, e := range elem {
		result = Clean(result + pathSeparator + e)
		resultSym, err := filepath.EvalSymlinks(result)

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				resultSym = result
			} else {
				return "", fmt.Errorf("can't eval symlinks: %w", err)
			}
		}

		result = resultSym
	}

	rootWithSep := root

	if !strings.HasSuffix(rootWithSep, pathSeparator) {
		rootWithSep += pathSeparator
	}

	if result != root && !strings.HasPrefix(result, rootWithSep) {
		return "", fmt.Errorf("final destination (%s) is outside root (%s)", result, root)
	}

	return result, nil
}

// Match reports whether name matches the shell file name pattern
func Match(pattern, name string) (matched bool, err error) {
	return filepath.Match(pattern, name)
}

// Split splits path immediately following the final slash, separating it into
// a directory and file name component
func Split(path string) (dir, file string) {
	return filepath.Split(path)
}

// Compact converts path to compact representation (e.g /some/random/directory/file.txt
// → /s/r/d/file.txt)
func Compact(path string) string {
	if !strings.ContainsRune(path, filepath.Separator) {
		return path
	}

	pathSlice := strings.Split(path, pathSeparator)

	for i := range len(pathSlice) - 1 {
		if len(pathSlice[i]) > 1 && !strings.HasSuffix(pathSlice[i], ":") {
			pathSlice[i] = pathSlice[i][:1]
		}
	}

	return strings.Join(pathSlice, pathSeparator)
}

// IsSafe reports whether the given path is safe to use and does not point to
// a protected system directory
func IsSafe(path string) bool {
	if path == "" {
		return false
	}

	absPath, err := filepath.Abs(Clean(path))

	if err != nil || absPath == pathSeparator {
		return false
	}

	return isSafePath(absPath)
}

// IsDotfile reports whether the base name of the given path begins with a dot
func IsDotfile(path string) bool {
	if path == "" {
		return false
	}

	if !strings.ContainsRune(path, filepath.Separator) {
		return strings.HasPrefix(path, ".")
	}

	pathBase := Base(path)

	return strings.HasPrefix(pathBase, ".")
}

// IsGlob reports whether the given pattern contains Unix-style glob
// metacharacters
func IsGlob(pattern string) bool {
	if pattern == "" {
		return false
	}

	var openBracket bool

	for _, r := range pattern {
		switch r {
		case '?', '*':
			return true
		case '[':
			openBracket = true
		case ']':
			if openBracket {
				return true
			}
		}
	}

	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

// dirNRight returns the path to the Nth directory from the right
func dirNRight(path string, n int) string {
	if path[0] == filepath.Separator {
		n++
	}

	var k int

	for i, r := range path {
		if r == filepath.Separator {
			k++
		}

		if k == n {
			return path[:i]
		}
	}

	return path
}

// dirNLeft returns the path to the Nth directory from the left
func dirNLeft(path string, n int) string {
	if path[len(path)-1] == filepath.Separator {
		n++
	}

	var k int

	for i := len(path) - 1; i > 0; i-- {
		if path[i] == filepath.Separator {
			k++
		}

		if k == n {
			return path[:i]
		}
	}

	return path
}

// evalHome evaluates the home directory in the given path
func evalHome(path string) string {
	if path == "" || path[:1] != "~" {
		return path
	}

	homeDir, err := homeDirFunc()

	if err != nil {
		return path
	}

	return homeDir + path[1:]
}
