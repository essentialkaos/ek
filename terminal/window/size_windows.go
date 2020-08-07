// +build windows

package window

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
