//go:build !linux && !darwin

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsTTY returns true if the current output device is a TTY
func IsTTY() bool {
	panic("UNSUPPORTED")
}

// ❗ IsFakeTTY returns true if a fake TTY is in use via the FAKETTY environment variable
func IsFakeTTY() bool {
	panic("UNSUPPORTED")
}

// ❗ IsTMUX returns true if the process is running inside a tmux session
func IsTMUX() (bool, error) {
	panic("UNSUPPORTED")
}

// ❗ IsSystemd returns true if the current process was started by systemd
func IsSystemd() bool {
	panic("UNSUPPORTED")
}

// ❗ GetSize returns the terminal width (columns) and height (rows)
func GetSize() (int, int) {
	panic("UNSUPPORTED")
}

// ❗ GetWidth returns the terminal width in columns
func GetWidth() int {
	panic("UNSUPPORTED")
}

// ❗ GetHeight returns the terminal height in rows
func GetHeight() int {
	panic("UNSUPPORTED")
}
