// Package sdnotify provides methods for sending notifications to systemd
package sdnotify

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"net"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrNoSocket     = errors.New("NOTIFY_SOCKET environment variable is empty")
	ErrNotConnected = errors.New("not connected to socket")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var conn net.Conn

// ////////////////////////////////////////////////////////////////////////////////// //

// Connect establishes a connection to the systemd notification socket
func Connect() error {
	if conn != nil {
		return nil // Already connected
	}

	var err error

	socket := os.Getenv("NOTIFY_SOCKET")

	if socket == "" {
		return ErrNoSocket
	}

	conn, err = net.Dial("unixgram", socket)

	if err != nil {
		return fmt.Errorf("can't connect to systemd notifications socket: %w", err)
	}

	return nil
}

// Disconnect closes connection to the systemd notification socket
func Disconnect() error {
	if conn == nil {
		return ErrNotConnected
	}

	err := conn.Close()

	if err != nil {
		return fmt.Errorf("can't close connection to systemd notifications socket: %w", err)
	}

	return nil
}

// Notify sends provided message to systemd
func Notify(msg string) error {
	if conn == nil {
		return ErrNotConnected
	}

	_, err := fmt.Fprint(conn, msg)

	if err != nil {
		return fmt.Errorf("can't send notification to socket: %w", err)
	}

	return nil
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
func MainPID() error {
	return Notify(fmt.Sprintf("MAINPID=%d", os.Getpid()))
}

// ExtendTimeout sends EXTEND_TIMEOUT_USEC to systemd with the given duration
// in seconds
func ExtendTimeout(sec float64) error {
	if sec <= 0 {
		return fmt.Errorf("invalid timeout %g", sec)
	}

	return Notify(fmt.Sprintf("EXTEND_TIMEOUT_USEC=%d", uint64(sec*1_000_000)))
}

// Status sends a printf-formatted status message to systemd
func Status(format string, a ...any) error {
	return Notify(fmt.Sprintf(format, a...))
}
