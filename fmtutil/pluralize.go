package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Pluralize is simple method for pluralization
func Pluralize(num int, one, many string) string {
	if num == 1 || num == -1 {
		return strconv.Itoa(num) + " " + one
	}

	return strconv.Itoa(num) + " " + many
}
