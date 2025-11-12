//go:build !windows

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"syscall"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

// ////////////////////////////////////////////////////////////////////////////////// //

// tty is a path to TTY device file
var tty = "/dev/tty"

// ttyFile is a tty file descriptor
var ttyFile *os.File

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize returns window width (columns) and height (rows)
func GetSize() (int, int) {
	var err error

	if ttyFile == nil {
		ttyFile, err = os.Open(tty)

		if err != nil {
			return -1, -1
		}
	}

	var sz winsize

	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL, ttyFile.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)),
	)

	return int(sz.cols), int(sz.rows)
}

// GetWidth returns window width (columns)
func GetWidth() int {
	w, _ := GetSize()
	return w
}

// GetHeight returns window height (rows)
func GetHeight() int {
	_, h := GetSize()
	return h
}
