// +build !windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var dirStack []string

// ////////////////////////////////////////////////////////////////////////////////// //

// Push changes current working directory and add previous working directory to stack
func Push(dir string) string {
	var wd string

	if dirStack == nil {
		wd, _ = os.Getwd()
		dirStack = append(dirStack, wd)
	}

	err := os.Chdir(dir)

	if err != nil {
		return ""
	}

	wd, _ = os.Getwd()

	dirStack = append(dirStack, wd)

	return wd
}

// Pop changes current working directory to previous in stack
func Pop() string {
	var wd string

	if dirStack == nil {
		wd, _ = os.Getwd()
		dirStack = append(dirStack, wd)
	}

	dl := len(dirStack)

	switch dl {

	case 0, 1:
		// nop

	default:
		err := os.Chdir(dirStack[dl-2])

		if err != nil {
			return ""
		}

		dirStack = dirStack[0 : dl-1]
	}

	wd, _ = os.Getwd()

	return wd
}
