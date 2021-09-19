// Package lscolors provides methods for colorizing file names based on colors from dircolors
package lscolors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// colorMap is map ext -> ANSI color code
var colorMap map[string]string

// initialized is initialization flag
var initialized bool

// ////////////////////////////////////////////////////////////////////////////////// //

// Colorize return file name with ANSI color tags
func Colorize(file string) string {
	if !initialized {
		initialize()
	}

	if len(colorMap) == 0 {
		return file
	}

	for glob, color := range colorMap {
		isMatch, _ := path.Match(glob, file)

		if isMatch {
			return "\033[" + color + "m" + file + "\033[0m"
		}
	}

	return file
}

// ////////////////////////////////////////////////////////////////////////////////// //

// initialize builds color map
func initialize() {
	initialized = true

	lsColors := os.Getenv("LS_COLORS")

	if lsColors == "" {
		return
	}

	colorMap = make(map[string]string)

	for _, key := range strings.Split(lsColors, ":") {
		if !strings.HasPrefix(key, "*") || !strings.ContainsRune(key, '=') {
			continue
		}

		sepIndex := strings.IndexRune(key, '=')

		colorMap[key[:sepIndex]] = key[sepIndex+1:]
	}
}
