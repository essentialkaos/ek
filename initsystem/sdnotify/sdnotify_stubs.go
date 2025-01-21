//go:build !linux
// +build !linux

// Package sdnotify provides methods for sending notifications to systemd
package sdnotify

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Connect connects systemd to socket
func Connect() error {
	panic("UNSUPPORTED")
}

// ❗ Notify sends provided message to systemd
func Notify(msg string) error {
	panic("UNSUPPORTED")
}

// ❗ Ready sends READY message to systemd
func Ready() error {
	panic("UNSUPPORTED")
}

// ❗ Reloading sends RELOADING message to systemd
func Reloading() error {
	panic("UNSUPPORTED")
}

// ❗ Stopping sends STOPPING message to systemd
func Stopping() error {
	panic("UNSUPPORTED")
}

// ❗ MainPID sends MAINPID message with PID to systemd
func MainPID(pid int) error {
	panic("UNSUPPORTED")
}

// ❗ ExtendTimeout sends EXTEND_TIMEOUT_USEC message to systemd
func ExtendTimeout(sec float64) error {
	panic("UNSUPPORTED")
}

// ❗ Status sends status message to systemd
func Status(format string, a ...any) error {
	panic("UNSUPPORTED")
}
