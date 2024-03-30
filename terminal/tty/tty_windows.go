package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsTTY returns true if current output device is TTY
func IsTTY() bool {
	panic("UNSUPPORTED")
}

// ❗ IsFakeTTY returns true is fake TTY is used
func IsFakeTTY() bool {
	panic("UNSUPPORTED")
}

// ❗ IsTMUX returns true if we are currently working in tmux
func IsTMUX() (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ IsSystemd returns true if process started by systemd
func IsSystemd() bool {
	panic("UNSUPPORTED")
}

// ❗ GetSize returns window width and height
func GetSize() (int, int) {
	panic("UNSUPPORTED")
}

// ❗ GetWidth returns window width
func GetWidth() int {
	panic("UNSUPPORTED")
}

// ❗ GetHeight returns window height
func GetHeight() int {
	panic("UNSUPPORTED")
}
