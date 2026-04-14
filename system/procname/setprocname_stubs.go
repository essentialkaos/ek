//go:build !linux && !darwin

package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ❗ ErrWrongSize is returned if given slice have the wrong size
	ErrWrongSize = errors.New("invalid args length (must have the same length as os.Arg)")

	// ❗ ErrEmptyFrom is returned if the "from" argument is empty
	ErrEmptyFrom = errors.New("from value is empty")

	// ❗ ErrEmptyTo is returned if the "to" argument is empty
	ErrEmptyTo = errors.New("to value is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ Set changes current process command in process tree
func Set(args []string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ Replace replaces one argument in process command
//
// WARNING: Be careful with using os.Args or options.Parse result
// as 'from' argument. After using this method given variable content
// will be replaced. Use strutil.Copy method in this case.
func Replace(from, to string) error {
	panic("UNSUPPORTED")
	return nil
}
