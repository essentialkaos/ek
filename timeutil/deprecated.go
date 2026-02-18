//nolint:all
package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "time"

// ////////////////////////////////////////////////////////////////////////////////// //

// SecondsToDuration converts float64 to time.Duration
//
// Deprecated: Use [ToSeconds] instead
func SecondsToDuration(d float64) time.Duration {
	return ToSeconds(d)
}

// PrettyDuration returns pretty duration (e.g. 1 hour 45 seconds)
//
// Deprecated: Use [Duration.String] instead
func PrettyDuration(d any) string {
	return Pretty(d).String()
}

// PrettyDurationSimple returns simple pretty duration (seconds → minutes → hours → days)
//
// Deprecated: Use [Duration.Simple] instead
func PrettyDurationSimple(d any) string {
	return Pretty(d).Simple()
}

// PrettyDurationInDays returns pretty duration in days (e.g. 15 days)
//
// Deprecated: Use [Duration.InDays] instead
func PrettyDurationInDays(d any) string {
	return Pretty(d).InDays()
}

// ShortDuration returns pretty short duration (e.g. 1:37)
//
// Deprecated: Use [Duration.Short] instead
func ShortDuration(d any, highPrecision ...bool) string {
	return Pretty(d).Short(highPrecision...)
}

// MiniDuration returns formatted value of duration (d/hr/m/s/ms/us/ns)
//
// Deprecated: Use [Duration.Mini] instead
func MiniDuration(d time.Duration, separator ...string) string {
	return Pretty(d).Mini(separator...)
}
