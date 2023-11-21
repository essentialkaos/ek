package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsTTY returns true if current output device is TTY
func IsTTY() bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsFakeTTY returns true is fake TTY is used
func IsFakeTTY() bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ GetSize returns window width and height
func GetSize() (int, int) {
	panic("UNSUPPORTED")
	return -1, -1
}

// ❗ GetWidth returns window width
func GetWidth() int {
	panic("UNSUPPORTED")
	return -1
}

// ❗ GetHeight returns window height
func GetHeight() int {
	panic("UNSUPPORTED")
	return -1
}