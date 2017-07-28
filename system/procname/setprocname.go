// +build linux darwin

// Package procname provides methods for changing process name in the process tree
package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrWrongSize is returned if given slice have the wrong size
var ErrWrongSize = errors.New("Given slice must have same size as os.Arg")

// ErrWrongArguments is returned if one of given arguments is empty
var ErrWrongArguments = errors.New("Arguments can't be empty")

// ////////////////////////////////////////////////////////////////////////////////// //

// Set change current process command in process tree
func Set(args []string) error {
	if len(args) != len(os.Args) {
		return ErrWrongSize
	}

	for i := 0; i < len(args); i++ {
		if args[i] == os.Args[i] {
			continue
		}

		changeArgument(i, args[i])
	}

	return nil
}

// Replace replace one argument in process command
func Replace(from, to string) error {
	if from == "" || to == "" {
		return ErrWrongArguments
	}

	for i, arg := range os.Args {
		if arg == from {
			changeArgument(i, to)
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func changeArgument(index int, newArg string) {
	argStrHr := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[index]))
	arg := (*[1 << 30]byte)(unsafe.Pointer(argStrHr.Data))[:argStrHr.Len]

	curArg := os.Args[index]
	curArgLen := len(curArg)
	newArgLen := len(newArg)

	switch {
	case curArgLen > newArgLen:
		newArg = newArg + strings.Repeat(" ", curArgLen-newArgLen)
	case curArgLen < newArgLen:
		newArg = newArg[:curArgLen]
	}

	n := copy(arg, newArg)

	if n < len(arg) {
		arg[n] = 0
	}
}
