//go:build linux || darwin

// Package procname provides methods for changing process name in the process tree
package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"runtime"
	"strings"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrWrongSize is returned if given slice have the wrong size
	ErrWrongSize = errors.New("invalid args length (must have the same length as os.Arg)")

	// ErrEmptyFrom is returned if the "from" argument is empty
	ErrEmptyFrom = errors.New("from value is empty")

	// ErrEmptyTo is returned if the "to" argument is empty
	ErrEmptyTo = errors.New("to value is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Set changes the current process command in the process tree.
// The given slice must have the same number of elements as os.Args.
func Set(args []string) error {
	if len(args) != len(os.Args) {
		return ErrWrongSize
	}

	for i, arg := range args {
		if arg != os.Args[i] {
			changeArgument(i, arg)
		}
	}

	return nil
}

// Replace replaces one argument in the process command.
//
// WARNING: Be careful with using os.Args or options.Parse result
// as 'from' argument. After using this method given variable content
// will be replaced. Use strutil.Copy method in this case.
func Replace(from, to string) error {
	switch {
	case from == "":
		return ErrEmptyFrom
	case to == "":
		return ErrEmptyTo
	}

	for i, arg := range os.Args {
		if arg == from {
			changeArgument(i, to)
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// changeArgument overwrites the argument at the given index in the os.Args
// backing array
func changeArgument(index int, newArg string) {
	curArg := os.Args[index]
	curArgLen := len(curArg)
	newArgLen := len(newArg)

	switch {
	case curArgLen > newArgLen:
		newArg += strings.Repeat(" ", curArgLen-newArgLen)
	case curArgLen < newArgLen:
		newArg = newArg[:curArgLen]
	}

	arg := unsafe.Slice(unsafe.StringData(curArg), curArgLen)
	n := copy(arg, newArg)

	if n < len(arg) {
		arg[n] = 0
	}

	runtime.KeepAlive(curArg)
}
