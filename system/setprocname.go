// +build linux, darwin
package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"reflect"
	"unsafe"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// SetProcName change current process name
// New process name must have same length or less.
func SetProcName(name string) {
	titleStr := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	title := (*[1 << 30]byte)(unsafe.Pointer(titleStr.Data))[:titleStr.Len]

	newTitle := name
	curTitle := os.Args[0]

	if len(curTitle) > len(newTitle) {
		spaces := "                                                                "
		newTitle += spaces[:len(curTitle)-len(newTitle)]
	}

	n := copy(title, newTitle)
	if n < len(title) {
		title[n] = 0
	}
}
