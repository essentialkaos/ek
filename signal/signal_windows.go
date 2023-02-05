package signal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

const (
	ABRT   = 0  // ❗ ABRT signal code
	ALRM   = 1  // ❗ ALRM signal code
	BUS    = 2  // ❗ BUS signal code
	CHLD   = 3  // ❗ CHLD signal code
	CONT   = 4  // ❗ CONT signal code
	FPE    = 5  // ❗ FPE signal code
	HUP    = 6  // ❗ HUP signal code
	ILL    = 7  // ❗ ILL signal code
	INT    = 8  // ❗ INT signal code
	IO     = 9  // ❗ IO signal code
	IOT    = 10 // ❗ IOT signal code
	KILL   = 11 // ❗ KILL signal code
	PIPE   = 12 // ❗ PIPE signal code
	PROF   = 13 // ❗ PROF signal code
	QUIT   = 14 // ❗ QUIT signal code
	SEGV   = 15 // ❗ SEGV signal code
	STOP   = 16 // ❗ STOP signal code
	SYS    = 17 // ❗ SYS signal code
	TERM   = 18 // ❗ TERM signal code
	TRAP   = 19 // ❗ TRAP signal code
	TSTP   = 20 // ❗ TSTP signal code
	TTIN   = 21 // ❗ TTIN signal code
	TTOU   = 22 // ❗ TTOU signal code
	URG    = 23 // ❗ URG signal code
	USR1   = 24 // ❗ USR1 signal code
	USR2   = 25 // ❗ USR2 signal code
	VTALRM = 26 // ❗ VTALRM signal code
	WINCH  = 27 // ❗ WINCH signal code
	XCPU   = 28 // ❗ XCPU signal code
	XFSZ   = 29 // ❗ XFSZ signal code
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Send send given signal to process
func Send(pid int, signal int) error {
	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Handlers is map signal → handler
type Handlers map[int]func()

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Track catch signal and execute handler for this signal
func (h Handlers) Track() {
	panic("UNSUPPORTED")
}

// ❗ TrackAsync catch signal and execute async handler for this signal
func (h Handlers) TrackAsync() {
	panic("UNSUPPORTED")
}
