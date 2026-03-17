// Package lscolors provides methods for colorizing file names based on colors from dircolors
package lscolors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"path"
	"strings"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	RESET  = "rs" // Reset
	DIR    = "di" // Directory
	LINK   = "ln" // Symbolic link
	FIFO   = "pi" // Pipe
	SOCK   = "so" // Socket
	BLK    = "bd" // Block device driver
	CHR    = "cd" // Character device driver
	STICKY = "st" // Dir with the sticky bit set (+t) and not other-writable
	EXEC   = "ex" // Executable files
)

// ////////////////////////////////////////////////////////////////////////////////// //

// colorMap is map ext -> ANSI color code
var colorMap map[string]string

// once is object for lazy colors initialization
var once sync.Once

// ////////////////////////////////////////////////////////////////////////////////// //

// DisableColors disables all colors in output
var DisableColors = os.Getenv("NO_COLOR") != ""

// ////////////////////////////////////////////////////////////////////////////////// //

// GetColor returns ANSI control sequence with color for given file
func GetColor(file string) string {
	once.Do(initialize)

	if DisableColors || len(colorMap) == 0 {
		return ""
	}

	if colorMap[file] != "" {
		return "\033[" + colorMap[file] + "m"
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

	return colorSeq + file + "\033[0m"
}

// ColorizePath returns the full path with ANSI control sequences applied
// to the basename
func ColorizePath(fullPath string) string {
	file := path.Base(fullPath)
	colorSeq := GetColor(file)

	if colorSeq == "" {
		return fullPath
	}

	return colorSeq + fullPath + "\033[0m"
}

// ////////////////////////////////////////////////////////////////////////////////// //

// initialize builds color map
func initialize() {
	if DisableColors {
		return
	}

	lsColors := os.Getenv("LS_COLORS")

	if lsColors == "" {
		return
	}

	colorMap = map[string]string{RESET: "0"}

	for key := range strings.SplitSeq(lsColors, ":") {
		if !strings.ContainsRune(key, '=') || !strings.ContainsRune(key, ';') {
			continue
		}

		sepIndex := strings.IndexRune(key, '=')

		colorMap[key[:sepIndex]] = key[sepIndex+1:]
	}
}
