// +build !windows

// Package window provides methods for working terminal window
package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

// tty is path to tty device file
var tty = "/dev/tty"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize return window width and height
func GetSize() (int, int) {
	t, err := os.OpenFile(tty, syscall.O_RDONLY, 0)

	if err != nil {
		return -1, -1
	}

	var sz winsize

	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL, t.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)),
	)

	return int(sz.cols), int(sz.rows)
}

// GetWidth return window width
func GetWidth() int {
	w, _ := GetSize()
	return w
}

// GetHeight return window height
func GetHeight() int {
	_, h := GetSize()
	return h
}
