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
	num    int
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
	if errs == nil {
		return e
	}

	for _, err := range errs {
		if err != nil {
			e.errors = append(e.errors, err)
			e.num++
		}
	}

	return e
}

// Last return last error in slice
func (e *Errors) Last() error {
	if e.errors == nil || e.num == 0 {
		return nil
	}

	return e.errors[e.num-1]
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

	return e.num != 0
}

// Num return number of errors
func (e *Errors) Num() int {
	return e.num
}
