//go:build !windows
// +build !windows

package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "strings"

// ////////////////////////////////////////////////////////////////////////////////// //

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

// DirN returns first N elements of path
func DirN(path string, n int) string {
	if strings.Count(path, pathSeparator) < 2 || n == 0 {
		return path
	}

	if n > 0 {
		return dirNRight(path, n)
	}

	return dirNLeft(path, n*-1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// isSafePath checks if path is safe to use
func isSafePath(path string) bool {
	for _, p := range unsafePaths {
		if strings.HasPrefix(path, p) {
			return false
		}
	}

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //
