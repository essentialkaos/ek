package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "strings"

// ////////////////////////////////////////////////////////////////////////////////// //

// unsafePaths is slice with unsafe paths
var unsafePaths = []string{
	`\Recovery`,
	`\Windows`,
	`\ProgramData`,
	`\Program Files (x86)`,
	`\Program Files`,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DirN returns first N elements of path
func DirN(path string, n int) string {
	if strings.Count(path, pathSeparator) < 2 || n == 0 {
		return path
	}

	disk, p, ok := strings.Cut(path, ":")

	if !ok {
		p = path
		disk = ""
	} else {
		disk += ":"
	}

	if n > 0 {
		return disk + dirNRight(p, n)
	}

	return disk + dirNLeft(p, n*-1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isSafePath(path string) bool {
	for _, p := range unsafePaths {
		if strings.ContainsRune(path, ':') {
			_, path, _ = strings.Cut(path, ":")
		}

		if strings.HasPrefix(path, p) {
			return false
		}
	}

	return true
}
