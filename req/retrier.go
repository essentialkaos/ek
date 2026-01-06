package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"time"

	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/timeutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Retrier is retrier struct
type Retrier struct {
	e *Engine
}

// Retry contains retry configuration
type Retry struct {
	Num              int           // Number of tries (1 or more)
	Pause            time.Duration // Pause between tries
	Status           int           // Required HTTP status (100-599)
	MinStatus        int           // Minimal HTTP status number (100-599)
	IgnoreRetryAfter bool          // Ignore Retry-After header
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilEngine is returned if retrier struct is nil
	ErrNilRetrier = fmt.Errorf("Retrier is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewRetrier creates new retrier instance
func NewRetrier(e ...*Engine) *Retrier {
	engine := Global

	if len(e) != 0 {
		engine = e[0]
	}

	return &Retrier{e: engine}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Do tries to send given request
func (rt *Retrier) Do(r Request, rr Retry) (*Response, error) {
	return rt.doRequest("", r, rr)
}

// Get tries to send GET request
func (rt *Retrier) Get(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(GET, r, rr)
}

// Post tries to send POST request
func (rt *Retrier) Post(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(POST, r, rr)
}

// Put tries to send PUT request
func (rt *Retrier) Put(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(PUT, r, rr)
}

// Head tries to send HEAD request
func (rt *Retrier) Head(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(HEAD, r, rr)
}

// Patch tries to send PATCH request
func (rt *Retrier) Patch(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(PATCH, r, rr)
}

// Delete tries to send DELETE request
func (rt *Retrier) Delete(r Request, rr Retry) (*Response, error) {
	return rt.doRequest(DELETE, r, rr)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate validates retry configuration
func (r Retry) Validate() error {
	switch {
	case r.Num < 1:
		return fmt.Errorf("Number of tries must be equal or greater that 1 (%d < 1)", r.Num)
	case r.Status != 0 && (r.Status < 100 || r.Status > 599):
		return fmt.Errorf("Invalid HTTP status code %d", r.Status)
	case r.MinStatus != 0 && (r.MinStatus < 100 || r.MinStatus > 599):
		return fmt.Errorf("Invalid minimal HTTP status code %d", r.MinStatus)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// doRequest tries to send request with retry logic
func (rt *Retrier) doRequest(method string, r Request, rr Retry) (*Response, error) {
	switch {
	case rt == nil:
		return nil, ErrNilRetrier
	case rt.e == nil:
		return nil, ErrNilEngine
	}

	var lastErr error

	for range rr.Num {
		resp, err := rt.e.doRequest(r, method)

		if err != nil {
			lastErr = err
		} else {
			switch {
			case rr.Status != 0 && resp.StatusCode != rr.Status:
				lastErr = fmt.Errorf(
					"All requests completed with non-ok status code (%d is required)",
					rr.Status,
				)
			case rr.MinStatus != 0 && resp.StatusCode > rr.MinStatus:
				lastErr = fmt.Errorf(
					"All requests completed with non-ok status code (status code must be greater than %d)",
					rr.Status,
				)
			}
		}

		if lastErr == nil {
			return resp, nil
		}

		retryPause := getRetryPause(rr, resp)

		if retryPause > 0 {
			time.Sleep(rr.Pause)
		}
	}

	return nil, lastErr
}

// getRetryPause returns pause between requests
func getRetryPause(rr Retry, resp *Response) time.Duration {
	if rr.Pause > 0 {
		return rr.Pause
	}

	if resp == nil {
		return 0
	}

	retryAfter := resp.Header.Get("Retry-After")

	if retryAfter == "" || rr.IgnoreRetryAfter {
		return 0
	}

	var pause time.Duration

	if mathutil.IsNumber(retryAfter) {
		pauseFloat, err := strconv.ParseFloat(retryAfter, 64)

		if err != nil {
			return 0
		}

		pause = timeutil.SecondsToDuration(pauseFloat)
	} else {
		pauseUntil, err := time.Parse(time.RFC1123, retryAfter)

		if err != nil {
			return 0
		}

		return time.Until(pauseUntil)
	}

	return pause
}
