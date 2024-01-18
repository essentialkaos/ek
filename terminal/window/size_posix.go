//go:build !windows
// +build !windows

// Package window provides methods for working terminal window
package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v12/terminal/tty"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize returns window width and height
//
// Deprecated: Use method package tty instead
func GetSize() (int, int) {
	return tty.GetSize()
}

// GetWidth returns window width
//
// Deprecated: Use method package tty instead
func GetWidth() int {
	w, _ := GetSize()
	return w
}

// GetHeight returns window height
//
// Deprecated: Use method package tty instead
func GetHeight() int {
	_, h := GetSize()
	return h
}
