// Package errutil provides methods for working with errors
package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is struct for handling many errors at once
type Errors struct {
	errors []error
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewErrors creates new struct
func NewErrors() *Errors {
	return &Errors{}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds new error to slice
func (e *Errors) Add(errs ...error) *Errors {
	if errs == nil || len(errs) == 0 {
		return e
	}

	e.errors = append(e.errors, errs...)

	return e
}

// Last return last error in slice
func (e *Errors) Last() error {
	if e.errors == nil || len(e.errors) == 0 {
		return nil
	}

	return e.errors[len(e.errors)-1]
}

// All return all errors in slice
func (e *Errors) All() []error {
	if e.errors == nil {
		return make([]error, 0)
	}

	return e.errors
}

// HasErrors check if slice contains errors
func (e *Errors) HasErrors() bool {
	if e.errors == nil {
		return false
	}

	return len(e.errors) != 0
}
