package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Date is tyny date struct
type Date uint32

// ////////////////////////////////////////////////////////////////////////////////// //

// StartDate contains start date timestamp
var StartDate uint32 = 1322611200

// ////////////////////////////////////////////////////////////////////////////////// //

// TinyDate create tiny date struct by timestamp
func TinyDate(t int64) Date {
	return Date(uint32(t) - StartDate)
}

// Unix return linux timestamp
func (d Date) Unix() int64 {
	return int64(StartDate + uint32(d))
}

// Time return time struct
func (d Date) Time() time.Time {
	return time.Unix(int64(uint32(d)+StartDate), 0)
}
