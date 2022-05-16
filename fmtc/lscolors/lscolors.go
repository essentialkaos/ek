// Package lscolors provides methods for colorizing file names based on colors from dircolors
package lscolors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// RESET_SEQ is ANSI reset sequence
const RESET_SEQ = "\033[0m"

// ////////////////////////////////////////////////////////////////////////////////// //

// colorMap is map ext -> ANSI color code
var colorMap map[string]string

// initialized is initialization flag
var initialized bool

// ////////////////////////////////////////////////////////////////////////////////// //

// GetColor returns ANSI control sequence with color for given file
func GetColor(file string) string {
	if !initialized {
		initialize()
	}

	if len(colorMap) == 0 {
		return ""
	}

	for glob, color := range colorMap {
		isMatch, _ := path.Match(glob, file)

		if isMatch {
			return "\033[" + color + "m"
		}
	}

	return ""
}

// Colorize return file name with ANSI control sequences
func Colorize(file string) string {
	colorSeq := GetColor(file)

	if colorSeq == "" {
		return file
	}

	return colorSeq + file + RESET_SEQ
}

// Colorize return path with ANSI control sequences
func ColorizePath(fullPath string) string {
	file := path.Base(fullPath)
	colorSeq := GetColor(file)

	if colorSeq == "" {
		return fullPath
	}

	return colorSeq + fullPath + RESET_SEQ
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
