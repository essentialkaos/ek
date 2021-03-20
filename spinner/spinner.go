// Package spinner provides methods for creating spinner animation for
// long-running tasks
package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"sync"
	"time"

	"pkg.re/essentialkaos/ek.v12/fmtc"
	"pkg.re/essentialkaos/ek.v12/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var SpinnerColorTag = "{y}"
var OkColorTag = "{g}"
var ErrColorTag = "{r}"
var TimeColorTag = "{s-}"

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

var isActive = false
var isHidden = true

var mu = &sync.RWMutex{}

// ////////////////////////////////////////////////////////////////////////////////// //

// Show shows spinner with given task description
func Show(message string, args ...interface{}) {
	mu.RLock()
	if !isHidden {
		mu.RUnlock()
		return
	}
	mu.RUnlock()

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	isActive, isHidden = true, false
	start = time.Now()
	mu.Unlock()

	go showSpinner()
}

// Update updates task description
func Update(message string, args ...interface{}) {
	mu.RLock()
	if isHidden {
		mu.RUnlock()
		return
	}
	mu.RUnlock()

	mu.Lock()
	desc = fmt.Sprintf(message, args...)
	mu.Unlock()
}

// Done finishes spinner animation and shows task status
func Done(ok bool) {
	mu.RLock()

	if !isActive {
		mu.RUnlock()
		return
	}

	mu.RUnlock()

	mu.Lock()
	isActive = false
	mu.Unlock()

	for {
		mu.RLock()
		if isHidden {
			mu.RUnlock()
			break
		}
		mu.RUnlock()
	}

	mu.RLock()

	if ok {
		fmtc.Printf(
			OkColorTag+"✔  {!}%s "+TimeColorTag+"(%s){!}\n",
			desc, timeutil.ShortDuration(time.Since(start)),
		)
	} else {
		fmtc.Printf(
			ErrColorTag+"✖  {!}%s "+TimeColorTag+"(%s){!}\n",
			desc, timeutil.ShortDuration(time.Since(start)),
		)
	}

	mu.RUnlock()

	mu.Lock()
	desc, isActive, isHidden, start = "", false, true, time.Time{}
	mu.Unlock()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showSpinner() {
	for {
		for i, frame := range spinnerFrames {
			mu.RLock()
			fmtc.Printf(
				SpinnerColorTag+"%s  {!}%s… "+TimeColorTag+"[%s]{!}",
				frame, desc, timeutil.ShortDuration(time.Since(start)),
			)
			mu.RUnlock()
			time.Sleep(framesDelay[i])
			fmtc.Printf("\033[2K\r")

			if !isActive {
				mu.Lock()
				isHidden = true
				mu.Unlock()
				return
			}
		}
	}
}
