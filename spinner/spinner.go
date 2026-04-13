// Package spinner provides methods for creating spinner animation for
// long-running tasks
package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v14/fmtc"
	"github.com/essentialkaos/ek/v14/mathutil"
	"github.com/essentialkaos/ek/v14/strutil"
	"github.com/essentialkaos/ek/v14/terminal/tty"
	"github.com/essentialkaos/ek/v14/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Duration format constants control how elapsed time is rendered in the spinner output
const (
	DURATION_SHORT  uint8 = iota // Short format, e.g. 1:34
	DURATION_MINI                // Minimal format, e.g. 15ms
	DURATION_SIMPLE              // Human-readable format, e.g. 10 seconds
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_ACTION_DONE uint8 = iota
	_ACTION_ERROR
	_ACTION_SKIP
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SpinnerColorTag is the fmtc color tag applied to the spinner animation frame
var SpinnerColorTag = "{y}"

// OkColorTag is the fmtc color tag applied to the success symbol
var OkColorTag = "{g}"

// ErrColorTag is the fmtc color tag applied to the error symbol
var ErrColorTag = "{r}"

// SkipColorTag is the fmtc color tag applied to the skip symbol
var SkipColorTag = "{s-}"

// TimeColorTag is the fmtc color tag applied to the elapsed time value
var TimeColorTag = "{s-}"

// OkSymbol is the terminal symbol printed when an action completes successfully
var OkSymbol = "✔ "

// ErrSymbol is the terminal symbol printed when an action fails
var ErrSymbol = "✖ "

// SkipSymbol is the terminal symbol printed when an action is skipped
var SkipSymbol = "✔ "

// DisableAnimation disables the animated spinner and runs in static output mode instead
var DisableAnimation = false

// DurationFormat controls the format used for printing the elapsed duration on completion
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

// Show starts the spinner animation with the given task description.
// Accepts fmt-style format arguments. No-op if the spinner is already active.
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

// Update replaces the spinner's description text while it is running.
// Accepts fmt-style format arguments. No-op if the spinner is not active or is hidden.
func Update(message string, args ...any) {
	if !isActive.Load() || isHidden.Load() {
		return
	}

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	mu.Unlock()
}

// Done stops the spinner and prints the final status line.
// Pass true to mark the action as successful, false to mark it as failed.
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

// Skip stops the spinner and marks the current action as skipped
func Skip() {
	if !isActive.Load() {
		return
	}

	stopSpinner(_ACTION_SKIP)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showSpinner is the goroutine that renders spinner animation frames to the terminal
func showSpinner() {
	var dur time.Duration

	spinnerColorTag := strutil.B(fmtc.IsTag(SpinnerColorTag), SpinnerColorTag, "{y}")
	timeColorTag := strutil.B(fmtc.IsTag(TimeColorTag), TimeColorTag, "{s-}")

	mu.RLock()
	fmtc.Printf(spinnerColorTag+"%s  {!}", spinnerFrames[0])
	fmtc.LPrint(getMaxDescSize(), desc)
	fmtc.Print("… " + timeColorTag + "[0:00]{!}")
	mu.RUnlock()

	frame := 1

	for {
		dur += framesDelay[frame]

		mu.RLock()

		fmtc.Printf("\033[1G"+spinnerColorTag+"%s  {!}", spinnerFrames[frame])

		if dur >= time.Second/2 {
			fmtc.LPrint(getMaxDescSize(), desc)
			fmtc.Printf("… "+timeColorTag+"[%s]{!}\033[K", timeutil.Pretty(time.Since(start)).Short(false))

			dur = 0
		}

		mu.RUnlock()

		fmt.Print("\033[999G")

		frame++

		// Frames 0 and 1 are used only for the initial display and first loop iteration.
		// After the first cycle, the spinner loops from frame 2 onward to maintain
		// a consistent cadence.
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

// stopSpinner signals the animation goroutine to stop and prints the final result line
func stopSpinner(action uint8) {
	isActive.Store(false)

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if isHidden.Load() {
			break
		}
	}

	mu.Lock()
	defer mu.Unlock()

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

	desc, start = "", time.Time{}
}

// formatDuration formats a duration value using the format selected by [DurationFormat]
func formatDuration(d time.Duration) string {
	switch DurationFormat {
	case DURATION_MINI:
		return timeutil.Pretty(d).Mini("")
	case DURATION_SIMPLE:
		return timeutil.Pretty(d).Simple()
	}

	return timeutil.Pretty(d).Short(true)
}

// getMaxDescSize returns the maximum character width available for the description text
func getMaxDescSize() int {
	w := tty.GetWidth()
	return mathutil.B(w < 20, 9999, w-14)
}
