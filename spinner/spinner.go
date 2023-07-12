// Package spinner provides methods for creating spinner animation for
// long-running tasks
package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_ACTION_DONE  uint8 = 0
	_ACTION_ERROR uint8 = 1
	_ACTION_SKIP  uint8 = 2
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

var spinnerFrames = []string{"⠸", "⠴", "⠤", "⠦", "⠇", "⠋", "⠉", "⠙"}

var framesDelay = []time.Duration{
	75 * time.Millisecond,
	55 * time.Millisecond,
	35 * time.Millisecond,
	55 * time.Millisecond,
	75 * time.Millisecond,
	75 * time.Millisecond,
	75 * time.Millisecond,
	75 * time.Millisecond,
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
	for {
		for i, frame := range spinnerFrames {
			mu.RLock()
			fmtc.Printf(
				SpinnerColorTag+"%s  {!}"+desc+"… "+TimeColorTag+"[%s]{!}",
				frame, timeutil.ShortDuration(time.Since(start)),
			)
			mu.RUnlock()

			time.Sleep(framesDelay[i])
			fmt.Print("\033[2K\r")

			if !isActive.Load() {
				isHidden.Store(true)
				return
			}
		}
	}
}

func stopSpinner(action uint8) {
	if !isActive.Load() {
		return
	}

	isActive.Store(false)

	for range time.NewTicker(time.Millisecond).C {
		if isHidden.Load() {
			break
		}
	}

	mu.RLock()
	switch action {
	case _ACTION_ERROR:
		fmtc.Printf(
			ErrColorTag+ErrSymbol+" {!}"+desc+" "+TimeColorTag+"(%s){!}\n",
			timeutil.ShortDuration(time.Since(start), true),
		)
	case _ACTION_SKIP:
		fmtc.Printf(
			SkipColorTag+SkipSymbol+" {!}"+desc+" "+TimeColorTag+"(%s){!}\n",
			timeutil.ShortDuration(time.Since(start), true),
		)
	default:
		fmtc.Printf(
			OkColorTag+OkSymbol+" {!}"+desc+" "+TimeColorTag+"(%s){!}\n",
			timeutil.ShortDuration(time.Since(start), true),
		)
	}

	mu.RUnlock()

	mu.Lock()
	desc, start = "", time.Time{}
	mu.Unlock()
}
