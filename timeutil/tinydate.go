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

// Date is tiny date
type Date uint32

// ////////////////////////////////////////////////////////////////////////////////// //

// StartDate contains start date timestamp
var StartDate = uint32(1322611200)

// ////////////////////////////////////////////////////////////////////////////////// //

// TinyDate creates tiny date struct by timestamp
func TinyDate(t int64) Date {
	return Date(uint32(t) - StartDate)
}

// Unix returns unix timestamp
func (d Date) Unix() int64 {
	return int64(StartDate + uint32(d))
}

// Time returns time struct
func (d Date) Time() time.Time {
	return time.Unix(int64(uint32(d)+StartDate), 0)
}
