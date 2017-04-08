// +build windows

package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize return window width and height
func GetSize() (int, int) {
	return -1, -1
}

// GetWidth return window width
func GetWidth() int {
	return -1
}

// GetHeight return window height
func GetHeight() int {
	return -1
}
