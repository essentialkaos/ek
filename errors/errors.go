// Package errors provides methods for working with errors
package errors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is a slice of errors
type Errors []error

// Bundle is a capacity-bounded, nil-safe collection of errors.
// Its zero value is ready to use without initialization.
type Bundle struct {
	capacity int
	errors   Errors
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New returns an error that formats as the given text. Each call to New returns
// a distinct error value even if the text is identical.
func New(text string) error {
	return errors.New(text)
}

// Is reports whether any error in err's tree matches target
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's tree that matches target, and if one is found,
// sets target to that error value and returns true. Otherwise, it returns false.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Join returns an error that wraps the given errors
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's type contains
// an Unwrap method returning error. Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewBundle creates a new empty [Bundle] with an optional maximum capacity.
// When capacity is exceeded, the oldest errors are dropped.
func NewBundle(capacity ...int) *Bundle {
	if len(capacity) == 0 {
		return &Bundle{}
	}

	size := max(capacity[0], 0)

	return &Bundle{capacity: size}
}

// ToBundle wraps an existing [Errors] slice into a [Bundle]
func ToBundle(errs Errors) *Bundle {
	b := &Bundle{}
	b.Add(errs)

	return b
}

// Chain calls each function in order and stops at the first non-nil error,
// returning it
func Chain(funcs ...func() error) error {
	var err error

	for _, chainFunc := range funcs {
		err = chainFunc()

		if err != nil {
			return err
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// First returns the first error in the slice, or nil if empty
func (e Errors) First() error {
	if e.IsEmpty() {
		return nil
	}

	return e[0]
}

// Last returns the last error in the slice, or nil if empty
func (e Errors) Last() error {
	if e.IsEmpty() {
		return nil
	}

	return e[e.Num()-1]
}

// Get returns the error at the given index, or nil if the index is out of range
func (e Errors) Get(index int) error {
	if index < 0 || index >= len(e) {
		return nil
	}

	return e[index]
}

// IsEmpty returns true if the slice contains no errors
func (e Errors) IsEmpty() bool {
	return len(e) == 0
}

// Num returns the number of errors in the slice
func (e Errors) Num() int {
	return len(e)
}

// ErrorWithPrefix returns combined text of all errors in the slice with prefix
// before every error
func (e Errors) ErrorWithPrefix(prefix string) string {
	var buf strings.Builder

	for i, err := range e {
		buf.WriteString(prefix)
		buf.WriteString(err.Error())

		if i+1 < len(e) {
			buf.WriteRune('\n')
		}
	}

	return buf.String()
}

// Error returns combined text of all errors in the slice
func (e Errors) Error() string {
	return e.ErrorWithPrefix("")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add appends one or more errors to the bundle.
// Accepts error, string, []error, []string, [Errors], [Bundle], and *Bundle values.
func (b *Bundle) Add(errs ...any) *Bundle {
	if len(errs) == 0 || b == nil {
		return b
	}

	for _, err := range errs {
		switch v := err.(type) {
		case *Bundle:
			if v != nil && len(v.errors) > 0 {
				b.errors = append(b.errors, v.errors...)
			}

		case Bundle:
			if len(v.errors) > 0 {
				b.errors = append(b.errors, v.errors...)
			}

		case []error:
			for _, e := range v {
				if e != nil {
					b.errors = append(b.errors, e)
				}
			}

		case Errors:
			for _, e := range v {
				if e != nil {
					b.errors = append(b.errors, e)
				}
			}

		case []string:
			for _, s := range v {
				b.errors = append(b.errors, errors.New(s))
			}

		case error:
			if v != nil {
				b.errors = append(b.errors, v)
			}

		case string:
			b.errors = append(b.errors, errors.New(v))
		}
	}

	if b.capacity > 0 && len(b.errors) > b.capacity {
		b.errors = b.errors[len(b.errors)-b.capacity:]
	}

	return b
}

// Addf formats an error using [fmt.Errorf] and appends it to the bundle
func (b *Bundle) Addf(format string, a ...any) *Bundle {
	b.Add(fmt.Errorf(format, a...))
	return b
}

// First returns the first error in the bundle, or nil if empty
func (b *Bundle) First() error {
	if b == nil {
		return nil
	}

	return b.errors.First()
}

// Last returns the last error in the bundle, or nil if empty
func (b *Bundle) Last() error {
	if b == nil {
		return nil
	}

	return b.errors.Last()
}

// Get returns the error at the given index, or nil if the index is out
// of range
func (b *Bundle) Get(index int) error {
	if b == nil {
		return nil
	}

	return b.errors.Get(index)
}

// All returns the full slice of collected errors
func (b *Bundle) All() Errors {
	if b == nil {
		return nil
	}

	return b.errors
}

// Num returns the number of errors currently in the bundle
func (b *Bundle) Num() int {
	if b == nil {
		return 0
	}

	return b.errors.Num()
}

// Cap returns the maximum number of errors the bundle will retain, or
// 0 if unbounded
func (b *Bundle) Cap() int {
	if b == nil {
		return 0
	}

	return b.capacity
}

// IsEmpty returns true if the bundle contains no errors
func (b *Bundle) IsEmpty() bool {
	if b == nil {
		return true
	}

	return b.errors.IsEmpty()
}

// ErrorWithPrefix returns all error messages joined by newlines, each prefixed
// with prefix
func (b *Bundle) ErrorWithPrefix(prefix string) string {
	if b == nil {
		return ""
	}

	return b.errors.ErrorWithPrefix(prefix)
}

// Error returns all error messages joined by newlines
func (b *Bundle) Error() string {
	if b == nil {
		return ""
	}

	return b.errors.Error()
}

// Reset removes all errors from the bundle without changing its capacity
func (b *Bundle) Reset() {
	if b != nil {
		b.errors = nil
	}
}
