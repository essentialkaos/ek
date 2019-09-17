package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
