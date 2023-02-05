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
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ErrBadPattern indicates a globbing pattern was malformed
var ErrBadPattern = errors.New("syntax error in pattern")

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Base returns the last element of path
func Base(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ Clean returns the shortest path name equivalent to path by purely lexical processing
func Clean(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ Dir returns all but the last element of path, typically the path's directory
func Dir(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ DirN returns first N elements of path
func DirN(path string, n int) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ Ext returns the file name extension used by path
func Ext(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ IsAbs reports whether the path is absolute
func IsAbs(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ Join joins any number of path elements into a single path, adding a separating slash if necessary
func Join(elem ...string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ Match reports whether name matches the shell file name pattern
func Match(pattern, name string) (matched bool, err error) {
	panic("UNSUPPORTED")
	return false, nil
}

// ❗ Split splits path immediately following the final slash, separating it into a directory and file name component
func Split(path string) (dir, file string) {
	panic("UNSUPPORTED")
	return "", ""
}

// ❗ Compact converts path to compact representation (e.g /some/random/directory/file.txt → /s/r/d/file.txt)
func Compact(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ IsSafe return true is given path is safe to use (not points to system dirs)
func IsSafe(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsDotfile return true if filename begins with a full stop
func IsDotfile(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsGlob returns true if given pattern is Unix-like glob
func IsGlob(pattern string) bool {
	panic("UNSUPPORTED")
	return false
}
