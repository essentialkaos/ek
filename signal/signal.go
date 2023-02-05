//go:build !windows
// +build !windows

// Package signal provides methods for handling POSIX signals
package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Signal codes
const (
	ABRT   = syscall.SIGABRT   // ABRT signal code
	ALRM   = syscall.SIGALRM   // ALRM signal code
	BUS    = syscall.SIGBUS    // BUS signal code
	CHLD   = syscall.SIGCHLD   // CHLD signal code
	CONT   = syscall.SIGCONT   // CONT signal code
	FPE    = syscall.SIGFPE    // FPE signal code
	HUP    = syscall.SIGHUP    // HUP signal code
	ILL    = syscall.SIGILL    // ILL signal code
	INT    = syscall.SIGINT    // INT signal code
	IO     = syscall.SIGIO     // IO signal code
	IOT    = syscall.SIGIOT    // IOT signal code
	KILL   = syscall.SIGKILL   // KILL signal code
	PIPE   = syscall.SIGPIPE   // PIPE signal code
	PROF   = syscall.SIGPROF   // PROF signal code
	QUIT   = syscall.SIGQUIT   // QUIT signal code
	SEGV   = syscall.SIGSEGV   // SEGV signal code
	STOP   = syscall.SIGSTOP   // STOP signal code
	SYS    = syscall.SIGSYS    // SYS signal code
	TERM   = syscall.SIGTERM   // TERM signal code
	TRAP   = syscall.SIGTRAP   // TRAP signal code
	TSTP   = syscall.SIGTSTP   // TSTP signal code
	TTIN   = syscall.SIGTTIN   // TTIN signal code
	TTOU   = syscall.SIGTTOU   // TTOU signal code
	URG    = syscall.SIGURG    // URG signal code
	USR1   = syscall.SIGUSR1   // USR1 signal code
	USR2   = syscall.SIGUSR2   // USR2 signal code
	VTALRM = syscall.SIGVTALRM // VTALRM signal code
	WINCH  = syscall.SIGWINCH  // WINCH signal code
	XCPU   = syscall.SIGXCPU   // XCPU signal code
	XFSZ   = syscall.SIGXFSZ   // XFSZ signal code
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Handlers is map signal → handler
type Handlers map[os.Signal]func()

// ////////////////////////////////////////////////////////////////////////////////// //

// Send sends given signal to process
func Send(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Track catches signal and executes handler for this signal
func (h Handlers) Track() {
	c := make(chan os.Signal, 2)

	for s := range h {
		signal.Notify(c, s)
	}

	go func() {
		for {
			sig := <-c

			handler := h[sig]

			if handler != nil {
				handler()
			}
		}
	}()
}

// TrackAsync catches signal and executes async handler for this signal
func (h Handlers) TrackAsync() {
	c := make(chan os.Signal, 2)

	for s := range h {
		signal.Notify(c, s)
	}

	go func() {
		for {
			sig := <-c

			handler := h[sig]

			if handler != nil {
				go handler()
			}
		}
	}()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,ABC]

// GetByName returns signal with given name
func GetByName(name string) (syscall.Signal, error) {
	switch strings.ToUpper(name) {
	case "SIGABRT", "ABRT":
		return ABRT, nil
	case "SIGALRM", "ALRM":
		return ALRM, nil
	case "SIGBUS", "BUS":
		return BUS, nil
	case "SIGCHLD", "CHLD":
		return CHLD, nil
	case "SIGCONT", "CONT":
		return CONT, nil
	case "SIGFPE", "FPE":
		return FPE, nil
	case "SIGHUP", "HUP":
		return HUP, nil
	case "SIGILL", "ILL":
		return ILL, nil
	case "SIGINT", "INT":
		return INT, nil
	case "SIGIO", "IO":
		return IO, nil
	case "SIGIOT", "IOT":
		return IOT, nil
	case "SIGKILL", "KILL":
		return KILL, nil
	case "SIGPIPE", "PIPE":
		return PIPE, nil
	case "SIGPROF", "PROF":
		return PROF, nil
	case "SIGQUIT", "QUIT":
		return QUIT, nil
	case "SIGSEGV", "SEGV":
		return SEGV, nil
	case "SIGSTOP", "STOP":
		return STOP, nil
	case "SIGSYS", "SYS":
		return SYS, nil
	case "SIGTERM", "TERM":
		return TERM, nil
	case "SIGTRAP", "TRAP":
		return TRAP, nil
	case "SIGTSTP", "TSTP":
		return TSTP, nil
	case "SIGTTIN", "TTIN":
		return TTIN, nil
	case "SIGTTOU", "TTOU":
		return TTOU, nil
	case "SIGURG", "URG":
		return URG, nil
	case "SIGUSR1", "USR1":
		return USR1, nil
	case "SIGUSR2", "USR2":
		return USR2, nil
	case "SIGVTALRM", "VTALRM":
		return VTALRM, nil
	case "SIGWINCH", "WINCH":
		return WINCH, nil
	case "SIGXCPU", "XCPU":
		return XCPU, nil
	case "SIGXFSZ", "XFSZ":
		return XFSZ, nil
	}

	return syscall.Signal(-1), fmt.Errorf("Unknown signal name %s", name)
}

// GetByCode returns signal with given code
func GetByCode(code int) (syscall.Signal, error) {
	switch code {
	case 1:
		return HUP, nil
	case 2:
		return INT, nil
	case 3:
		return QUIT, nil
	case 4:
		return ILL, nil
	case 5:
		return TRAP, nil
	case 6:
		return ABRT, nil
	case 8:
		return FPE, nil
	case 9:
		return KILL, nil
	case 10:
		return BUS, nil
	case 11:
		return SEGV, nil
	case 12:
		return SYS, nil
	case 13:
		return PIPE, nil
	case 14:
		return ALRM, nil
	case 15:
		return TERM, nil
	case 16:
		return USR1, nil
	case 17:
		return USR2, nil
	case 18:
		return CHLD, nil
	case 20:
		return TSTP, nil
	case 21:
		return URG, nil
	case 23:
		return STOP, nil
	case 25:
		return CONT, nil
	case 26:
		return TTIN, nil
	case 27:
		return TTOU, nil
	case 28:
		return VTALRM, nil
	case 29:
		return PROF, nil
	case 30:
		return XCPU, nil
	case 31:
		return XFSZ, nil
	}

	return syscall.Signal(-1), fmt.Errorf("Unknown signal code %d", code)
}

// codebeat:enable[LOC,ABC]

// ////////////////////////////////////////////////////////////////////////////////// //
