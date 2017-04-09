// Package errutil provides methods for working with errors
package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Errors is struct for handling many errors at once
type Errors []error

// ////////////////////////////////////////////////////////////////////////////////// //

// NewErrors creates new struct
func NewErrors() *Errors {
	return &Errors{}
}

// Chain execute functions in chain and if one of them return error
// this function stop chain execution and return given error
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
	for _, err := range errs {
		if err != nil {
			*e = append(*e, err)
		}
	}

	return e
}

// Last return last error in slice
func (e *Errors) Last() error {
	if *e == nil {
		return nil
	}

	return (*e)[len(*e)-1]
}

// All return all errors in slice
func (e *Errors) All() []error {
	return *e
}

// HasErrors check if slice contains errors
func (e *Errors) HasErrors() bool {
	return len(*e) > 0
}

// Num return number of errors
func (e *Errors) Num() int {
	return len(*e)
}
