package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"

	"golang.org/x/term"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize returns the terminal width (columns) and height (rows)
func GetSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		return -1, -1
	}

	return w, h
}

// GetWidth returns the terminal width in columns
func GetWidth() int {
	w, _ := GetSize()
	return w
}

// GetHeight returns the terminal height in rows
func GetHeight() int {
	_, h := GetSize()
	return h
}
