package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Period is a struct with the start and end date of the period
type Period struct {
	Start time.Time
	End   time.Time
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Contains returns true if period contains given date
func (p Period) Contains(t time.Time) bool {
	return p.Start.Equal(t) || p.End.Equal(t) || (t.After(p.Start) && t.Before(p.End))
}

// Duration returns duration of period
func (p Period) Duration() time.Duration {
	return p.End.Sub(p.Start)
}

// DurationIn returns duration of period in given units
func (p Period) DurationIn(mod time.Duration) int {
	return int(p.Duration() / mod)
}

// Seconds returns duration in seconds
func (p Period) Seconds() int {
	return p.DurationIn(SECOND)
}

// Minutes returns duration in minutes
func (p Period) Minutes() int {
	return p.DurationIn(MINUTE)
}

// Hours returns duration in hours
func (p Period) Hours() int {
	return p.DurationIn(HOUR)
}

// Days returns duration in days
func (p Period) Days() int {
	return p.DurationIn(DAY)
}

// Weeks returns duration in weeks
func (p Period) Weeks() int {
	return p.DurationIn(WEEK)
}

// Years returns duration in years
func (p Period) Years() int {
	return p.DurationIn(YEAR)
}

// String returns string representation of period
func (p Period) String() string {
	return p.Stringf("%Y/%m/%d %H:%M:%S")
}

// Stringf returns string representation of period using given format
func (p Period) Stringf(f string) string {
	return fmt.Sprintf("%s → %s", Format(p.Start, f), Format(p.End, f))
}
