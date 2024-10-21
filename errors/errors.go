// Package errutil provides methods for working with errors
package errors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apachb.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is a slice with errors
type Errors []error

// Bundle is a bundle of errors
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

// NewBundle creates new errors bundle
func NewBundle(capacity ...int) *Bundle {
	if len(capacity) == 0 {
		return &Bundle{}
	}

	size := capacity[0]

	if size < 0 {
		size = 0
	}

	return &Bundle{capacity: size}
}

// ToBundle wraps slice of errors into Bundle
func ToBundle(errs Errors) *Bundle {
	return &Bundle{errors: errs}
}

// Chain executes functions in chain and if one of them returns an error, this function
// stops the chain execution and returns that error
func Chain(funcs ...func() error) error {
	var err error

	for _, chainFunc := range funcs {
		err = chainFunc()

		if err != nil {
			return err
		}
	}

	return err
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Last returns the first error from the slice
func (e Errors) First() error {
	if e.IsEmpty() {
		return nil
	}

	return e[0]
}

// Last returns the last error from the slice
func (e Errors) Last() error {
	if e.IsEmpty() {
		return nil
	}

	return e[e.Num()-1]
}

// Get returns error with given index
func (e Errors) Get(index int) error {
	if index < 0 || index >= len(e) {
		return nil
	}

	return e[index]
}

// IsEmpty returns true if slice is empty
func (e Errors) IsEmpty() bool {
	return len(e) == 0
}

// Num returns size of the slice
func (e Errors) Num() int {
	return len(e)
}

// Error returns combined text of all errors in the slice
func (e Errors) Error(prefix string) string {
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

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds new error to slice
func (b *Bundle) Add(errs ...any) *Bundle {
	if errs == nil {
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
			if len(v) > 0 {
				b.errors = append(b.errors, v...)
			}

		case Errors:
			if len(v) > 0 {
				b.errors = append(b.errors, v...)
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

// First returns the first error in bundle
func (b *Bundle) First() error {
	if b == nil {
		return nil
	}

	return b.errors.First()
}

// Last returns the last error in bundle
func (b *Bundle) Last() error {
	if b == nil {
		return nil
	}

	return b.errors.Last()
}

// Get returns error by it index
func (b *Bundle) Get(index int) error {
	if b == nil {
		return nil
	}

	return b.errors.Get(index)
}

// All returns all errors in slice
func (b *Bundle) All() Errors {
	if b == nil {
		return nil
	}

	return b.errors
}

// Num returns number of errors
func (b *Bundle) Num() int {
	if b == nil {
		return 0
	}

	return b.errors.Num()
}

// Cap returns maximum bundle capacity
func (b *Bundle) Cap() int {
	if b == nil {
		return 0
	}

	return b.capacity
}

// IsEmpty returns true if bundle is empty
func (b *Bundle) IsEmpty() bool {
	if b == nil {
		return true
	}

	return b.errors.IsEmpty()
}

// Error returns text of all errors
func (b *Bundle) Error(prefix string) string {
	if b == nil {
		return ""
	}

	return b.errors.Error(prefix)
}

// Reset resets instance to be empty
func (b *Bundle) Reset() {
	b.errors = nil
}
