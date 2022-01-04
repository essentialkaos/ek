package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ❗ ErrWrongSize is returned if given slice have the wrong size
	ErrWrongSize = errors.New("Given slice must have same size as os.Arg")

	// ❗ ErrWrongArguments is returned if one of given arguments is empty
	ErrWrongArguments = errors.New("Arguments can't be empty")
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
