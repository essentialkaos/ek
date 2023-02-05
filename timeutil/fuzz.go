//go:build gofuzz
// +build gofuzz

package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
