// +build windows

package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// GetSize returns window width and height
func GetSize() (int, int) {
	return -1, -1
}

// GetWidth returns window width
func GetWidth() int {
	return -1
}

// GetHeight returns window height
func GetHeight() int {
	return -1
}
