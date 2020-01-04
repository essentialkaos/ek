// +build !windows

// Package signal provides methods for handling signals
package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Signal codes
const (
	ABRT   = syscall.SIGABRT
	ALRM   = syscall.SIGALRM
	BUS    = syscall.SIGBUS
	CHLD   = syscall.SIGCHLD
	CONT   = syscall.SIGCONT
	FPE    = syscall.SIGFPE
	HUP    = syscall.SIGHUP
	ILL    = syscall.SIGILL
	INT    = syscall.SIGINT
	IO     = syscall.SIGIO
	IOT    = syscall.SIGIOT
	KILL   = syscall.SIGKILL
	PIPE   = syscall.SIGPIPE
	PROF   = syscall.SIGPROF
	QUIT   = syscall.SIGQUIT
	SEGV   = syscall.SIGSEGV
	STOP   = syscall.SIGSTOP
	SYS    = syscall.SIGSYS
	TERM   = syscall.SIGTERM
	TRAP   = syscall.SIGTRAP
	TSTP   = syscall.SIGTSTP
	TTIN   = syscall.SIGTTIN
	TTOU   = syscall.SIGTTOU
	URG    = syscall.SIGURG
	USR1   = syscall.SIGUSR1
	USR2   = syscall.SIGUSR2
	VTALRM = syscall.SIGVTALRM
	WINCH  = syscall.SIGWINCH
	XCPU   = syscall.SIGXCPU
	XFSZ   = syscall.SIGXFSZ
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Handlers is map signal->handler
type Handlers map[os.Signal]func()

// ////////////////////////////////////////////////////////////////////////////////// //

// Send sends given signal to process
func Send(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Track catches signal and executes handler for this signal
func (h Handlers) Track() {
	c := make(chan os.Signal)

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
	c := make(chan os.Signal)

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

// GetByName returns signal with given name
func GetByName(name string) os.Signal {
	switch strings.ToUpper(name) {
	case "SIGABRT", "ABRT":
		return ABRT
	case "SIGALRM", "ALRM":
		return ALRM
	case "SIGBUS", "BUS":
		return BUS
	case "SIGCHLD", "CHLD":
		return CHLD
	case "SIGCONT", "CONT":
		return CONT
	case "SIGFPE", "FPE":
		return FPE
	case "SIGHUP", "HUP":
		return HUP
	case "SIGILL", "ILL":
		return ILL
	case "SIGINT", "INT":
		return INT
	case "SIGIO", "IO":
		return IO
	case "SIGIOT", "IOT":
		return IOT
	case "SIGKILL", "KILL":
		return KILL
	case "SIGPIPE", "PIPE":
		return PIPE
	case "SIGPROF", "PROF":
		return PROF
	case "SIGQUIT", "QUIT":
		return QUIT
	case "SIGSEGV", "SEGV":
		return SEGV
	case "SIGSTOP", "STOP":
		return STOP
	case "SIGSYS", "SYS":
		return SYS
	case "SIGTERM", "TERM":
		return TERM
	case "SIGTRAP", "TRAP":
		return TRAP
	case "SIGTSTP", "TSTP":
		return TSTP
	case "SIGTTIN", "TTIN":
		return TTIN
	case "SIGTTOU", "TTOU":
		return TTOU
	case "SIGURG", "URG":
		return URG
	case "SIGUSR1", "USR1":
		return USR1
	case "SIGUSR2", "USR2":
		return USR2
	case "SIGVTALRM", "VTALRM":
		return VTALRM
	case "SIGWINCH", "WINCH":
		return WINCH
	case "SIGXCPU", "XCPU":
		return XCPU
	case "SIGXFSZ", "XFSZ":
		return XFSZ
	}

	return -1
}

// GetByCode returns signal with given code
func GetByCode(code int) os.Signal {
	switch code {
	case 1:
		return HUP
	case 2:
		return INT
	case 3:
		return QUIT
	case 4:
		return ILL
	case 5:
		return TRAP
	case 6:
		return ABRT
	case 8:
		return FPE
	case 9:
		return KILL
	case 10:
		return BUS
	case 11:
		return SEGV
	case 12:
		return SYS
	case 13:
		return PIPE
	case 14:
		return ALRM
	case 15:
		return TERM
	case 16:
		return USR1
	case 17:
		return USR2
	case 18:
		return CHLD
	case 20:
		return TSTP
	case 21:
		return URG
	case 23:
		return STOP
	case 25:
		return CONT
	case 26:
		return TTIN
	case 27:
		return TTOU
	case 28:
		return VTALRM
	case 29:
		return PROF
	case 30:
		return XCPU
	case 31:
		return XFSZ
	}

	return -1
}

// ////////////////////////////////////////////////////////////////////////////////// //
