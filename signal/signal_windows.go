package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

const (
	ABRT   = 0
	ALRM   = 1
	BUS    = 2
	CHLD   = 3
	CONT   = 4
	FPE    = 5
	HUP    = 6
	ILL    = 7
	INT    = 8
	IO     = 9
	IOT    = 10
	KILL   = 11
	PIPE   = 12
	PROF   = 13
	QUIT   = 14
	SEGV   = 15
	STOP   = 16
	SYS    = 17
	TERM   = 18
	TRAP   = 19
	TSTP   = 20
	TTIN   = 21
	TTOU   = 22
	URG    = 23
	USR1   = 24
	USR2   = 25
	VTALRM = 26
	WINCH  = 27
	XCPU   = 28
	XFSZ   = 29
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
func (h Handlers) Track() {
	panic("UNSUPPORTED")
}

// TrackAsync catch signal and execute async handler for this signal
func (h Handlers) TrackAsync() {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //
