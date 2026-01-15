// Package timeutil provides methods for working with time and date
package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/essentialkaos/ek/v13/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Format formats [time.Time] value using given format string
//
// Interpreted sequences:
//
//	'%%' a literal %
//	'%a' locale's abbreviated weekday name (e.g., Sun)
//	'%A' locale's full weekday name (e.g., Sunday)
//	'%b' locale's abbreviated month name (e.g., Jan)
//	'%B' locale's full month name (e.g., January)
//	'%c' locale's date and time (e.g., Thu Mar 3 23:05:25 2005)
//	'%C' century; like %Y, except omit last two digits (e.g., 20)
//	'%d' day of month (e.g, 01)
//	'%D' date; same as %m/%d/%y
//	'%e' day of month, space padded
//	'%F' full date; same as %Y-%m-%d
//	'%g' last two digits of year of ISO week number (see %G)
//	'%G' year of ISO week number (see %V); normally useful only with %V
//	'%h' same as %b
//	'%H' hour (00..23)
//	'%I' hour (01..12)
//	'%j' day of year (001..366)
//	'%k' hour ( 0..23)
//	'%K' milliseconds (000..999)
//	'%l' hour ( 1..12)
//	'%m' month (01..12)
//	'%M' minute (00..59)
//	'%n' a newline
//	'%N' nanoseconds (000000000..999999999)
//	'%p' AM or PM
//	'%P' like %p, but lower case
//	'%r' locale's 12-hour clock time (e.g., 11:11:04 PM)
//	'%R' 24-hour hour and minute; same as %H:%M
//	'%s' seconds since 1970-01-01 00:00:00 UTC
//	'%S' second (00..60)
//	'%t' a tab
//	'%T' time; same as %H:%M:%S
//	'%u' day of week (1..7); 1 is Monday
//	'%U' week number of year, with Sunday as first day of week (00..53)
//	'%V' ISO week number, with Monday as first day of week (01..53)
//	'%w' day of week (0..6); 0 is Sunday
//	'%W' week number of year, with Monday as first day of week (00..53)
//	'%x' locale's date representation (e.g., 12/31/99)
//	'%X' locale's time representation (e.g., 23:13:48)
//	'%y' last two digits of year (00..99)
//	'%Y' year
//	'%z' +hhmm numeric timezone (e.g., -0400)
//	'%:z' +hh:mm numeric timezone (e.g., -04:00)
//	'%Z' alphabetic time zone abbreviation (e.g., EDT)
func Format(d time.Time, f string) string {
	input := bytes.NewBufferString(f)
	output := bytes.NewBufferString("")

	for {
		r, _, err := input.ReadRune()

		if err != nil {
			break
		}

		switch r {
		case '%':
			replaceDateTag(d, input, output)
		default:
			output.WriteRune(r)
		}
	}

	return output.String()
}

// SecondsToDuration converts float64 to time.Duration
func SecondsToDuration(d float64) time.Duration {
	return time.Duration(1000000000.0 * d)
}

// ParseDuration parses duration in 1w2d3h5m6s format and return as seconds
func ParseDuration(dur string, defMod ...rune) (time.Duration, error) {
	if dur == "" {
		return 0, nil
	}

	var err error
	var result int64

	buf := &bytes.Buffer{}

	for _, sym := range dur {
		switch sym {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf.WriteRune(sym)
		case 'w', 'W':
			result, err = appendDur(result, buf, _WEEK)
		case 'd', 'D':
			result, err = appendDur(result, buf, _DAY)
		case 'h', 'H':
			result, err = appendDur(result, buf, _HOUR)
		case 'm', 'M':
			result, err = appendDur(result, buf, _MINUTE)
		case 's', 'S':
			result, err = appendDur(result, buf, _SECOND)
		default:
			return 0, fmt.Errorf("Unsupported symbol %q", string(sym))
		}

		if err != nil {
			return 0, err
		}
	}

	if buf.Len() != 0 {
		if result != 0 {
			return 0, fmt.Errorf("Misformatted duration %q", dur)
		}

		mod := 's'

		if len(defMod) != 0 {
			mod = defMod[0]
		}

		result, err = strconv.ParseInt(buf.String(), 10, 64)

		if err != nil {
			return 0, err
		}

		switch mod {
		case 'w', 'W':
			result *= _WEEK
		case 'd', 'D':
			result *= _DAY
		case 'h', 'H':
			result *= _HOUR
		case 'm', 'M':
			result *= _MINUTE
		case 's', 'S':
			result *= _SECOND
		}
	}

	return SECOND * time.Duration(result), nil
}

// StartOfHour returns start of the hour
func StartOfHour(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), 0, 0, 0, t.Location(),
	)
}

// EndOfHour returns end of the hour
func EndOfHour(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return StartOfHour(t).Add(HOUR).Add(-1 * time.Nanosecond)
}

// StartOfDay returns start of the day
func StartOfDay(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0, t.Location(),
	)
}

// EndOfDay returns end of the day
func EndOfDay(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return StartOfDay(t).Add(24 * HOUR).Add(-1 * time.Nanosecond)
}

// StartOfWeek returns the first day of the week
func StartOfWeek(t time.Time, firstDay time.Weekday) time.Time {
	if t.IsZero() {
		return t
	}

	for {
		if t.Weekday() == firstDay {
			return time.Date(
				t.Year(), t.Month(), t.Day(),
				0, 0, 0, 0, t.Location(),
			)
		}

		t = t.AddDate(0, 0, -1)
	}
}

// EndOfWeek returns the last day of the week
func EndOfWeek(t time.Time, firstDay time.Weekday) time.Time {
	if t.IsZero() {
		return t
	}

	return StartOfWeek(t, firstDay).AddDate(0, 0, 7).Add(-1 * time.Nanosecond)
}

// StartOfMonth returns the first day of the month
func StartOfMonth(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return time.Date(
		t.Year(), t.Month(), 1,
		0, 0, 0, 0, t.Location(),
	)
}

// EndOfMonth returns the last day of the month
func EndOfMonth(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return StartOfMonth(t).AddDate(0, 1, 0).Add(-1 * time.Nanosecond)
}

// StartOfYear returns the first day of the year
func StartOfYear(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return time.Date(
		t.Year(), time.January, 1,
		0, 0, 0, 0, t.Location(),
	)
}

// EndOfYear returns the last day of the year
func EndOfYear(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}

	return StartOfYear(t).AddDate(1, 0, 0).Add(-1 * time.Nanosecond)
}

// PrevDay returns previous day date
func PrevDay(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, -1))
}

// PrevWeek returns previous week date
func PrevWeek(t time.Time, firstDay time.Weekday) time.Time {
	return StartOfWeek(t.AddDate(0, 0, -7), firstDay)
}

// PrevMonth returns previous month date
func PrevMonth(t time.Time) time.Time {
	return StartOfMonth(t.AddDate(0, -1, 0))
}

// PrevYear returns previous year date
func PrevYear(t time.Time) time.Time {
	return StartOfYear(t.AddDate(-1, 0, 0))
}

// NextDay returns next day date
func NextDay(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, 1))
}

// NextWeek returns next week date
func NextWeek(t time.Time, firstDay time.Weekday) time.Time {
	return StartOfWeek(t.AddDate(0, 0, 7), firstDay)
}

// NextMonth returns next month date
func NextMonth(t time.Time) time.Time {
	return StartOfMonth(t.AddDate(0, 1, 0))
}

// NextYear returns next year date
func NextYear(t time.Time) time.Time {
	return StartOfYear(t.AddDate(1, 0, 0))
}

// NextWorkday returns previous workday
func PrevWorkday(t time.Time) time.Time {
	for {
		t = t.AddDate(0, 0, -1)

		switch t.Weekday() {
		case time.Saturday, time.Sunday:
			continue
		}

		return StartOfDay(t)
	}
}

// NextWeekend returns previous weekend
func PrevWeekend(t time.Time) time.Time {
	for {
		t = t.AddDate(0, 0, -1)

		switch t.Weekday() {
		case time.Saturday, time.Sunday:
			return StartOfDay(t)
		}
	}
}

// NextWorkday returns next workday
func NextWorkday(t time.Time) time.Time {
	for {
		t = t.AddDate(0, 0, 1)

		switch t.Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday,
			time.Thursday, time.Friday:
			return StartOfDay(t)
		}
	}
}

// NextWeekend returns next weekend
func NextWeekend(t time.Time) time.Time {
	for {
		t = t.AddDate(0, 0, 1)

		switch t.Weekday() {
		case time.Saturday, time.Sunday:
			return StartOfDay(t)
		}
	}
}

// AddWorkdays adds working days to a given date
func AddWorkdays(t time.Time, days int) time.Time {
	diff := 1

	if days < 0 {
		diff = -1
	}

	days = mathutil.Abs(days)

	for days > 0 {
		t = t.AddDate(0, 0, diff)

		switch t.Weekday() {
		case time.Saturday, time.Sunday:
			continue
		default:
			days--
		}
	}

	return t
}

// IsWeekend returns true if given day is weekend (saturday or sunday)
func IsWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

// Until returns time until given moment in given units
func Until(t time.Time, unit time.Duration) int {
	now := time.Now().In(t.Location())
	return int(t.Sub(now) / unit)
}

// Since returns time since given moment in given units
func Since(t time.Time, unit time.Duration) int {
	now := time.Now().In(t.Location())
	return int(now.Sub(t) / unit)
}

// DurationAs returns duration in given units
func DurationAs(t, unit time.Duration) int {
	return int(math.Round(float64(t) / float64(unit)))
}

// FromISOWeek returns [time.Time] from ISO week number and year
func FromISOWeek(week, year int, loc *time.Location) time.Time {
	week = mathutil.Between(week, 1, 53)

	if year <= 0 {
		year = time.Now().In(loc).Year()
	}

	return time.Date(year, 1, 1, 0, 0, 0, 0, loc).AddDate(0, 0, 7*(week-1))
}

// ParseWithAny tries to parse value using given layouts
func ParseWithAny(value string, layouts ...string) (time.Time, error) {
	if len(layouts) == 0 {
		return time.Time{}, fmt.Errorf("No layouts provided")
	}

	for _, l := range layouts {
		t, err := time.Parse(l, value)

		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("Value cannot be parsed using any of the provided layouts")
}

// UnixIn returns the time corresponding to the given Unix timestamp interpreted
// as local time in the specified timezone rather than UTC.
func UnixIn(sec int64, nsec int64, loc *time.Location) time.Time {
	return removeTZOffset(time.Unix(sec, nsec).In(loc))
}

// ToUnixIn returns the Unix timestamp in given timezone
func ToUnixIn(t time.Time, loc *time.Location) int64 {
	ts := t.Unix()
	_, offset := t.Zone()

	return ts + int64(offset)
}

// UnixMilliIn returns the time corresponding to the given Unix millisecond
// timestamp interpreted as local time in the specified timezone rather than UTC.
func UnixMilliIn(msec int64, loc *time.Location) time.Time {
	return removeTZOffset(time.UnixMilli(msec).In(loc))
}

// ToUnixMilliIn returns the Unix timestamp in given timezone
func ToUnixMilliIn(t time.Time, loc *time.Location) int64 {
	ts := t.UnixMilli()
	_, offset := t.Zone()

	return ts + (int64(offset) * 1_000)
}

// UnixMicroIn returns the time corresponding to the given Unix microsecond
// timestamp interpreted as local time in the specified timezone rather than UTC.
func UnixMicroIn(usec int64, loc *time.Location) time.Time {
	return removeTZOffset(time.UnixMicro(usec).In(loc))
}

// ToUnixMicroIn returns the Unix timestamp in given timezone
func ToUnixMicroIn(t time.Time, loc *time.Location) int64 {
	ts := t.UnixMicro()
	_, offset := t.Zone()

	return ts + (int64(offset) * 1_000_000)
}

// ToUnixNanoIn returns the Unix timestamp in given timezone
func ToUnixNanoIn(t time.Time, loc *time.Location) int64 {
	ts := t.UnixNano()
	_, offset := t.Zone()

	return ts + (int64(offset) * 1_000_000_000)
}

// ////////////////////////////////////////////////////////////////////////////////// //

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

// ////////////////////////////////////////////////////////////////////////////////// //

// replaceDateTag replaces date tag in format string
func replaceDateTag(d time.Time, input, output *bytes.Buffer) {
	r, _, err := input.ReadRune()

	if err != nil {
		return
	}

	switch r {
	case '%':
		fmt.Fprintf(output, "%%")
	case 'a':
		output.WriteString(getShortWeekday(d.Weekday()))
	case 'A':
		output.WriteString(getLongWeekday(d.Weekday()))
	case 'b', 'h':
		output.WriteString(getShortMonth(d.Month()))
	case 'B':
		output.WriteString(getLongMonth(d.Month()))
	case 'c':
		zn, _ := d.Zone()
		fmt.Fprintf(output, "%s %02d %s %d %02d:%02d:%02d %s %s",
			getShortWeekday(d.Weekday()),
			d.Day(),
			getShortMonth(d.Month()),
			d.Year(),
			getAMPMHour(d),
			d.Minute(),
			d.Second(),
			getAMPM(d, true),
			zn,
		)
	case 'C', 'g':
		output.WriteString(strconv.Itoa(d.Year())[0:2])
	case 'd':
		fmt.Fprintf(output, "%02d", d.Day())
	case 'D':
		fmt.Fprintf(output, "%02d/%02d/%s", d.Month(), d.Day(), strconv.Itoa(d.Year())[2:4])
	case 'e':
		fmt.Fprintf(output, "%2d", d.Day())
	case 'F':
		fmt.Fprintf(output, "%d-%02d-%02d", d.Year(), d.Month(), d.Day())
	case 'G':
		fmt.Fprintf(output, "%02d", d.Year())
	case 'H':
		fmt.Fprintf(output, "%02d", d.Hour())
	case 'I':
		fmt.Fprintf(output, "%02d", getAMPMHour(d))
	case 'j':
		fmt.Fprintf(output, "%03d", d.YearDay())
	case 'k':
		fmt.Fprintf(output, "%2d", d.Hour())
	case 'K':
		output.WriteString(fmt.Sprintf("%03d", d.Nanosecond())[:3])
	case 'l':
		output.WriteString(strconv.Itoa(getAMPMHour(d)))
	case 'm':
		fmt.Fprintf(output, "%02d", d.Month())
	case 'M':
		fmt.Fprintf(output, "%02d", d.Minute())
	case 'n':
		output.WriteString("\n")
	case 'N':
		fmt.Fprintf(output, "%09d", d.Nanosecond())
	case 'p':
		output.WriteString(getAMPM(d, false))
	case 'P':
		output.WriteString(getAMPM(d, true))
	case 'r':
		fmt.Fprintf(output, "%02d:%02d:%02d %s", getAMPMHour(d), d.Minute(), d.Second(), getAMPM(d, true))
	case 'R':
		fmt.Fprintf(output, "%02d:%02d", d.Hour(), d.Minute())
	case 's':
		output.WriteString(strconv.FormatInt(d.Unix(), 10))
	case 'S':
		fmt.Fprintf(output, "%02d", d.Second())
	case 'T':
		fmt.Fprintf(output, "%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second())
	case 'u':
		output.WriteString(strconv.Itoa(getWeekdayNum(d)))
	case 'V':
		_, wn := d.ISOWeek()
		fmt.Fprintf(output, "%02d", wn)
	case 'w':
		fmt.Fprintf(output, "%d", d.Weekday())
	case 'y':
		output.WriteString(strconv.Itoa(d.Year())[2:4])
	case 'Y':
		output.WriteString(strconv.Itoa(d.Year()))
	case 'z':
		output.WriteString(getTimezone(d, false))
	case ':':
		input.ReadRune()
		output.WriteString(getTimezone(d, true))
	case 'Z':
		zn, _ := d.Zone()
		output.WriteString(zn)
	default:
		output.WriteRune('%')
		output.WriteRune(r)
	}
}

// getShortWeekday returns short weekday name (e.g. Sun)
func getShortWeekday(d time.Weekday) string {
	long := getLongWeekday(d)

	if long == "" {
		return ""
	}

	return long[:3]
}

// getLongWeekday returns long weekday name (e.g. Sunday)
func getLongWeekday(d time.Weekday) string {
	switch int(d) {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	}

	return ""
}

// getShortMonth returns short month name (e.g. Jan)
func getShortMonth(m time.Month) string {
	long := getLongMonth(m)

	if long == "" {
		return ""
	}

	return long[:3]
}

// getLongMonth returns long month name (e.g. January)
func getLongMonth(m time.Month) string {
	switch int(m) {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	}

	return ""
}

// getAMPMHour returns hour in 12-hour format (e.g. 1..12)
func getAMPMHour(d time.Time) int {
	h := d.Hour()

	switch {
	case h == 0 || h == 12:
		return 12

	case h < 12:
		return h

	default:
		return h - 12
	}
}

// getAMPM returns AM or PM depending on the time and caps flag
func getAMPM(d time.Time, caps bool) string {
	if d.Hour() < 12 {
		switch caps {
		case true:
			return "AM"
		default:
			return "am"
		}
	} else {
		switch caps {
		case true:
			return "PM"
		default:
			return "pm"
		}
	}
}

// getWeekdayNum returns weekday number (1..7) where 1 is Monday and 7 is Sunday
func getWeekdayNum(d time.Time) int {
	r := int(d.Weekday())

	if r == 0 {
		r = 7
	}

	return r
}

// getTimezone returns timezone offset in +hh:mm or +hhmm format
func getTimezone(d time.Time, separator bool) string {
	negative := false
	_, tzofs := d.Zone()

	if tzofs < 0 {
		negative = true
		tzofs *= -1
	}

	hours := int64(tzofs) / _HOUR
	minutes := int64(tzofs) % _HOUR

	switch negative {
	case true:
		if separator {
			return fmt.Sprintf("-%02d:%02d", hours, minutes)
		}

		return fmt.Sprintf("-%02d%02d", hours, minutes)

	default:
		if separator {
			return fmt.Sprintf("+%02d:%02d", hours, minutes)
		}

		return fmt.Sprintf("+%02d%02d", hours, minutes)
	}
}

// appendDur appends duration value to the buffer and returns new value
func appendDur(value int64, buf *bytes.Buffer, mod int64) (int64, error) {
	v, err := strconv.ParseInt(buf.String(), 10, 64)

	if err != nil {
		return 0, err
	}

	buf.Reset()

	return value + (v * mod), nil
}

// removeTZOffset removes timezone offset to interpret timestamp as local time
func removeTZOffset(t time.Time) time.Time {
	_, offset := t.Zone()
	return t.Add(-time.Duration(offset) * time.Second)
}
