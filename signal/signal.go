// +build !windows

// Package signal provides methods for handling signals
package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"os/signal"
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

// Send send given signal to process
func Send(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Track catch signal and execute handler for this signal
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

// TrackAsync catch signal and execute async handler for this signal
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
