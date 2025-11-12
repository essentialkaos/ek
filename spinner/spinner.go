// Package spinner provides methods for creating spinner animation for
// long-running tasks
package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/terminal/tty"
	"github.com/essentialkaos/ek/v13/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DURATION_SHORT  uint8 = iota // e.g. 1:34
	DURATION_MINI                // e.g. 15ms
	DURATION_SIMPLE              // e.g. 10 seconds
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_ACTION_DONE uint8 = iota
	_ACTION_ERROR
	_ACTION_SKIP
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SpinnerColorTag is spinner animation color tag (see fmtc package)
var SpinnerColorTag = "{y}"

// OkColorTag is check color tag (see fmtc package)
var OkColorTag = "{g}"

// ErrColorTag is cross color tag (see fmtc package)
var ErrColorTag = "{r}"

// SkipColorTag is skipped action color tag (see fmtc package)
var SkipColorTag = "{s-}"

// TimeColorTag is time color tag (see fmtc package)
var TimeColorTag = "{s-}"

// OkSymbol contains symbol for action with no problems
var OkSymbol = "✔ "

// ErrSymbol contains symbol for action with problems
var ErrSymbol = "✖ "

// SkipSymbol contains symbol for skipped action
var SkipSymbol = "✔ "

// DisableAnimation is global animation off switch flag
var DisableAnimation = false

// DurationFormat is format used for printing result action duration
var DurationFormat = DURATION_SHORT

// ////////////////////////////////////////////////////////////////////////////////// //

var spinnerFrames = []string{"⠒", "⠲", "⠴", "⠤", "⠦", "⠇", "⠋", "⠉", "⠙", "⠸"}

var framesDelay = []time.Duration{
	105 * time.Millisecond,
	95 * time.Millisecond,
	75 * time.Millisecond,
	55 * time.Millisecond,
	35 * time.Millisecond,
	55 * time.Millisecond,
	75 * time.Millisecond,
	75 * time.Millisecond,
	75 * time.Millisecond,
	95 * time.Millisecond,
}

var desc string
var start time.Time

var isActive = &atomic.Bool{}
var isHidden = &atomic.Bool{}

var mu = &sync.RWMutex{}

// ////////////////////////////////////////////////////////////////////////////////// //

// Show starts spinner animation and shows task description
func Show(message string, args ...any) {
	if isActive.Load() {
		return
	}

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	start = time.Now()

	isActive.Store(true)
	isHidden.Store(false)

	if DisableAnimation {
		isHidden.Store(true)
	} else {
		go showSpinner()
	}
	mu.Unlock()
}

// Update updates spinner description
func Update(message string, args ...any) {
	if !isActive.Load() || isHidden.Load() {
		return
	}

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	mu.Unlock()
}

// Done finishes spinner animation and marks it as done
func Done(ok bool) {
	if !isActive.Load() {
		return
	}

	if ok {
		stopSpinner(_ACTION_DONE)
	} else {
		stopSpinner(_ACTION_ERROR)
	}
}

// Skip finishes spinner animation and marks it as skipped
func Skip() {
	if !isActive.Load() {
		return
	}

	stopSpinner(_ACTION_SKIP)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showSpinner starts spinner animation in a separate goroutine
func showSpinner() {
	var dur time.Duration

	spinnerColorTag := strutil.B(fmtc.IsTag(SpinnerColorTag), SpinnerColorTag, "{y}")
	timeColorTag := strutil.B(fmtc.IsTag(TimeColorTag), TimeColorTag, "{s-}")

	mu.RLock()
	fmtc.Printf(spinnerColorTag+"%s  {!}", spinnerFrames[0])
	fmtc.LPrintf(getMaxDescSize(), desc)
	fmtc.Print("… " + timeColorTag + "[0:00]{!}")
	mu.RUnlock()

	frame := 1

	for {
		dur += framesDelay[frame]

		mu.RLock()

		fmtc.Printf("\033[1G"+spinnerColorTag+"%s  {!}", spinnerFrames[frame])

		if dur >= time.Second/2 {
			fmtc.LPrintf(getMaxDescSize(), desc)
			fmtc.Printf("… "+timeColorTag+"[%s]{!}\033[K", timeutil.Pretty(time.Since(start)).Short(false))

			dur = 0
		}

		mu.RUnlock()

		fmt.Print("\033[999G")

		frame++

		if frame == 10 {
			frame = 2
		}

		time.Sleep(framesDelay[frame])

		if !isActive.Load() {
			isHidden.Store(true)
			return
		}
	}
}

// stopSpinner stops spinner animation and prints final message
func stopSpinner(action uint8) {
	isActive.Store(false)

	for range time.NewTicker(time.Millisecond).C {
		if isHidden.Load() {
			break
		}
	}

	mu.RLock()

	timeColorTag := strutil.B(fmtc.IsTag(TimeColorTag), TimeColorTag, "{s-}")

	fmt.Print("\033[1G")

	switch action {
	case _ACTION_ERROR:
		errColorTag := strutil.B(fmtc.IsTag(ErrColorTag), ErrColorTag, "{r}")
		fmtc.Printf(errColorTag + ErrSymbol + " {!}")
	case _ACTION_SKIP:
		skipColorTag := strutil.B(fmtc.IsTag(SkipColorTag), SkipColorTag, "{s-}")
		fmtc.Printf(skipColorTag + SkipSymbol + " {!}")
	default:
		okColorTag := strutil.B(fmtc.IsTag(OkColorTag), OkColorTag, "{g}")
		fmtc.Printf(okColorTag + OkSymbol + " {!}")
	}

	fmtc.Print(desc + " ")
	fmtc.Printfn(timeColorTag+"(%s){!}\033[K", formatDuration(time.Since(start)))

	mu.RUnlock()

	mu.Lock()
	desc, start = "", time.Time{}
	mu.Unlock()
}

// formatDuration formats duration based on the global DurationFormat setting
func formatDuration(d time.Duration) string {
	switch DurationFormat {
	case DURATION_MINI:
		return timeutil.Pretty(d).Mini("")
	case DURATION_SIMPLE:
		return timeutil.Pretty(d).Simple()
	}

	return timeutil.Pretty(d).Short(true)
}

// getMaxDescSize returns size of the current line for the description
func getMaxDescSize() int {
	w := tty.GetWidth()
	return mathutil.B(w < 20, 9999, w-14)
}
