package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"

	"golang.org/x/term"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// IsTTY returns true if the current output device is a TTY
func IsTTY() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// IsFakeTTY returns true if a fake TTY is in use via the FAKETTY environment variable
func IsFakeTTY() bool {
	return false
}

// IsTMUX returns true if the process is running inside a tmux session
func IsTMUX() (bool, error) {
	return false, nil
}

// IsSystemd returns true if the current process was started by systemd
func IsSystemd() bool {
	return false
}
