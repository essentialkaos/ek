// Package errutil provides methods for working with errors
//
// Deprecated: Use package errors instead
package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is struct for handling many errors at once
type Errors struct {
	capacity int
	errors   []error
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewErrors creates new struct
//
// Deprecated: Use package errors instead
func NewErrors(capacity ...int) *Errors {
	if len(capacity) == 0 {
		return &Errors{}
	}

	size := capacity[0]

	if size < 0 {
		size = 0
	}

	return &Errors{capacity: size}
}

// Wrap wraps slice of errors into Errors struct
//
// Deprecated: Use package errors instead
func Wrap(errs []error) *Errors {
	return &Errors{errors: errs}
}

// Chain executes functions in chain and if one of them return error
// this function stop chain execution and return this error
//
// Deprecated: Use package errors instead
func Chain(funcs ...func() error) error {
	var err error

	for _, fc := range funcs {
		err = fc()

		if err != nil {
			return err
		}
	}

	return err
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds new error to slice
//
// Deprecated: Use package errors instead
func (e *Errors) Add(errs ...any) *Errors {
	if errs == nil {
		return e
	}

	for _, err := range errs {
		switch v := err.(type) {
		case *Errors:
			if v != nil {
				e.errors = append(e.errors, v.errors...)
			}
		case Errors:
			e.errors = append(e.errors, v.errors...)
		case []error:
			e.errors = append(e.errors, v...)
		case []string:
			for _, s := range v {
				e.errors = append(e.errors, errors.New(s))
			}
		case error:
			e.errors = append(e.errors, v)
		case string:
			e.errors = append(e.errors, errors.New(v))
		}
	}

	if e.capacity > 0 && len(e.errors) > e.capacity {
		e.errors = e.errors[len(e.errors)-e.capacity:]
	}

	return e
}

// First returns the first error
//
// Deprecated: Use package errors instead
func (e *Errors) First() error {
	if e == nil || len(e.errors) == 0 {
		return nil
	}

	return e.errors[0]
}

// Last returns the last error
//
// Deprecated: Use package errors instead
func (e *Errors) Last() error {
	if e == nil || len(e.errors) == 0 {
		return nil
	}

	return e.errors[len(e.errors)-1]
}

// Get returns error by it index
//
// Deprecated: Use package errors instead
func (e *Errors) Get(index int) error {
	if e == nil || len(e.errors) == 0 ||
		index < 0 || index >= len(e.errors) {
		return nil
	}

	return e.errors[index]
}

// All returns all errors in slice
//
// Deprecated: Use package errors instead
func (e *Errors) All() []error {
	if e == nil || e.errors == nil {
		return nil
	}

	return e.errors
}

// HasErrors checks if slice contains errors
//
// Deprecated: Use package errors instead
func (e *Errors) HasErrors() bool {
	if e == nil || e.errors == nil {
		return false
	}

	return len(e.errors) != 0
}

// Num returns number of errors
//
// Deprecated: Use package errors instead
func (e *Errors) Num() int {
	if e == nil {
		return 0
	}

	return len(e.errors)
}

// Cap returns max capacity
//
// Deprecated: Use package errors instead
func (e *Errors) Cap() int {
	if e == nil {
		return 0
	}

	return e.capacity
}

// Error returns text of all errors
//
// Deprecated: Use package errors instead
func (e *Errors) Error() string {
	if e == nil || len(e.errors) == 0 {
		return ""
	}

	var result string

	for _, err := range e.errors {
		result += "  " + err.Error() + "\n"
	}

	return result
}

// Reset resets Errors instance to be empty
//
// Deprecated: Use package errors instead
func (e *Errors) Reset() {
	e.errors = nil
}
