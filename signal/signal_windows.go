// +build !linux, !darwin, windows

package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

const (
	ABRT   = 0
	ALRM   = 0
	BUS    = 0
	CHLD   = 0
	CONT   = 0
	FPE    = 0
	HUP    = 0
	ILL    = 0
	INT    = 0
	IO     = 0
	IOT    = 0
	KILL   = 0
	PIPE   = 0
	PROF   = 0
	QUIT   = 0
	SEGV   = 0
	STOP   = 0
	SYS    = 0
	TERM   = 0
	TRAP   = 0
	TSTP   = 0
	TTIN   = 0
	TTOU   = 0
	URG    = 0
	USR1   = 0
	USR2   = 0
	VTALRM = 0
	WINCH  = 0
	XCPU   = 0
	XFSZ   = 0
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Send send given signal to process
func Send(pid int, signal int) error {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Handlers is map signal->handler
type Handlers map[int]func()

// ////////////////////////////////////////////////////////////////////////////////// //

// Track catch signal and execute handler for this signal
func (h Handlers) Track() {}

// TrackAsync catch signal and execute async handler for this signal
func (h Handlers) TrackAsync() {}

// ////////////////////////////////////////////////////////////////////////////////// //
