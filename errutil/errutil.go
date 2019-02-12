// Package errutil provides methods for working with errors
package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

// Chain execute functions in chain and if one of them return error
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
	if e == nil || e.num == 0 || e.errors == nil {
		return nil
	}

	return e.errors[e.num-1]
}

// All return all errors in slice
func (e *Errors) All() []error {
	if e == nil || e.errors == nil {
		return nil
	}

	return e.errors
}

// HasErrors check if slice contains errors
func (e *Errors) HasErrors() bool {
	if e == nil || e.errors == nil {
		return false
	}

	return e.num != 0
}

// Num return number of errors
func (e *Errors) Num() int {
	if e == nil {
		return 0
	}

	return e.num
}

// Error return text of all errors
func (e *Errors) Error() string {
	if e == nil || e.num == 0 {
		return ""
	}

	var result string

	for _, err := range e.errors {
		result += "  " + err.Error() + "\n"
	}

	return result
}
