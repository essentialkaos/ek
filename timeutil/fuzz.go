// +build gofuzz

package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var fuzzTestDate = time.Now()

// ////////////////////////////////////////////////////////////////////////////////// //

func Fuzz(data []byte) int {
	f := Format(fuzzTestDate, string(data))

	if f != "" {
		return 0
	}

	return 1
}
