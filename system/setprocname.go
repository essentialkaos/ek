// +build linux darwin

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

// SetProcName change current process name
// New process name must have same length or less.
func SetProcName(args []string) error {
	if len(args) != len(os.Args) {
		return errors.New("Wrong arguments slice size")
	}

	for i := 0; i < len(args); i++ {
		if args[i] == os.Args[i] {
			continue
		}

		changeArgument(i, args[i])
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
