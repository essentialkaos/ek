// Package errutil provides methods for working with errors
package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is struct for handling many errors at once
type Errors struct {
	maxSize int
	errors  []error
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewErrors creates new struct
func NewErrors(maxSize ...int) *Errors {
	if len(maxSize) == 0 {
		return &Errors{}
	}

	size := maxSize[0]

	if size < 0 {
		size = 0
	}

	return &Errors{maxSize: size}
}

// Chain executes functions in chain and if one of them return error
// this function stop chain execution and return this error
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
func (e *Errors) Add(errs ...interface{}) *Errors {
	if errs == nil {
		return e
	}

	for _, err := range errs {
		switch v := err.(type) {
		case *Errors:
			if v != nil {
				e.errors = append(e.errors, v.errors...)
			}
		case []error:
			e.errors = append(e.errors, v...)
		case error:
			e.errors = append(e.errors, v)
		}
	}

	if e.maxSize > 0 && len(e.errors) > e.maxSize {
		e.errors = e.errors[len(e.errors)-e.maxSize:]
	}

	return e
}

// Last returns last error in slice
func (e *Errors) Last() error {
	if e == nil || e.errors == nil {
		return nil
	}

	return e.errors[len(e.errors)-1]
}

// All returns all errors in slice
func (e *Errors) All() []error {
	if e == nil || e.errors == nil {
		return nil
	}

	return e.errors
}

// HasErrors checks if slice contains errors
func (e *Errors) HasErrors() bool {
	if e == nil || e.errors == nil {
		return false
	}

	return len(e.errors) != 0
}

// Num returns number of errors
func (e *Errors) Num() int {
	if e == nil {
		return 0
	}

	return len(e.errors)
}

// Error returns text of all errors
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
