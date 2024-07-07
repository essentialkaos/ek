// Package spinner provides methods for creating spinner animation for
// long-running tasks
package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/timeutil"
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

// Show shows spinner with given task description
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

// Update updates task description
func Update(message string, args ...any) {
	if !isActive.Load() || isHidden.Load() {
		return
	}

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	mu.Unlock()
}

// Done finishes spinner animation and shows task status
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

// Skip finishes spinner animation and mark it as skipped
func Skip() {
	if !isActive.Load() {
		return
	}

	stopSpinner(_ACTION_SKIP)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showSpinner() {
	var i int

	spinnerColorTag := strutil.B(fmtc.IsTag(SpinnerColorTag), SpinnerColorTag, "{y}")
	timeColorTag := strutil.B(fmtc.IsTag(TimeColorTag), TimeColorTag, "{s-}")

	for {
		mu.RLock()
		fmtc.Printf(spinnerColorTag+"%s  {!}", spinnerFrames[i])
		fmtc.Print(desc + "… ")
		fmtc.Printf(timeColorTag+"[%s]{!}", timeutil.ShortDuration(time.Since(start)))
		mu.RUnlock()

		i++

		if i == 10 {
			i = 2
		}

		time.Sleep(framesDelay[i])
		fmt.Print("\033[2K\r")

		if !isActive.Load() {
			isHidden.Store(true)
			return
		}
	}
}

func stopSpinner(action uint8) {
	isActive.Store(false)

	for range time.NewTicker(time.Millisecond).C {
		if isHidden.Load() {
			break
		}
	}

	mu.RLock()

	timeColorTag := strutil.B(fmtc.IsTag(TimeColorTag), TimeColorTag, "{s-}")

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
	fmtc.Printf(timeColorTag+"(%s){!}\n", timeutil.ShortDuration(time.Since(start), true))

	mu.RUnlock()

	mu.Lock()
	desc, start = "", time.Time{}
	mu.Unlock()
}
