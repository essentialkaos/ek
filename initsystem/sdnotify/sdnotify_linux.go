// Package sdnotify provides methods for sending notifications to systemd
package sdnotify

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"net"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNoSocket     = fmt.Errorf("NOTIFY_SOCKET is empty")
	ErrNotConnected = fmt.Errorf("Not connected to socket")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var conn net.Conn

// ////////////////////////////////////////////////////////////////////////////////// //

// Connect connects systemd to socket
func Connect() error {
	var err error

	socket := os.Getenv("NOTIFY_SOCKET")

	if err != nil {
		return ErrNoSocket
	}

	conn, err = net.Dial("unixgram", socket)

	if err != nil {
		return fmt.Errorf("Can't connect to socket: %w", err)
	}

	return nil
}

// Notify sends provided message to systemd
func Notify(msg string) error {
	if conn == nil {
		return ErrNotConnected
	}

	_, err := fmt.Fprint(conn, msg)

	return err
}

// Ready sends READY message to systemd
func Ready() error {
	return Notify("READY=1")
}

// Reloading sends RELOADING message to systemd
func Reloading() error {
	return Notify("RELOADING=1")
}

// Stopping sends STOPPING message to systemd
func Stopping() error {
	return Notify("STOPPING=1")
}

// MainPID sends MAINPID message with PID to systemd
func MainPID(pid int) error {
	return Notify(fmt.Sprintf("MAINPID=%d", pid))
}

// ExtendTimeout sends EXTEND_TIMEOUT_USEC message to systemd
func ExtendTimeout(sec float64) error {
	usec := uint(sec * 1_000_000)
	return Notify(fmt.Sprintf("EXTEND_TIMEOUT_USEC=%d", usec))
}

// Status sends status message to systemd
func Status(format string, a ...any) error {
	return Notify(fmt.Sprintf(format, a...))
}
