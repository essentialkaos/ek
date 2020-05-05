// Package timeutil provides methods for working with time
package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v12/mathutil"
	"pkg.re/essentialkaos/ek.v12/pluralize"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_MINUTE = 60
	_HOUR   = 3600
	_DAY    = 86400
	_WEEK   = 604800
)

// ////////////////////////////////////////////////////////////////////////////////// //

// It's ok to have so nested blocks in this method
// codebeat:disable[BLOCK_NESTING]

// PrettyDuration returns pretty duration (e.g. 1 hour 45 seconds)
func PrettyDuration(d interface{}) string {
	var (
		result   []string
		duration int
	)

	switch u := d.(type) {
	case time.Duration:
		if u < time.Second {
			return getShortDuration(u)
		}

		duration = int(u.Seconds())
	case int8:
		duration = int(u)
	case int16:
		duration = int(u)
	case int32:
		duration = int(u)
	case int64:
		duration = int(u)
	case int:
		duration = u
	default:
		return "Wrong duration value"
	}

MAINLOOP:
	for i := 0; i < 5; i++ {
		switch {
		case duration >= _WEEK:
			weeks := duration / _WEEK
			duration = duration % _WEEK
			result = append(result, pluralize.PS(pluralize.En, "%d %s", weeks, "week", "weeks"))
		case duration >= _DAY:
			days := duration / _DAY
			duration = duration % _DAY
			result = append(result, pluralize.PS(pluralize.En, "%d %s", days, "day", "days"))
		case duration >= _HOUR:
			hours := duration / _HOUR
			duration = duration % _HOUR
			result = append(result, pluralize.PS(pluralize.En, "%d %s", hours, "hour", "hours"))
		case duration >= _MINUTE:
			minutes := duration / _MINUTE
			duration = duration % _MINUTE
			result = append(result, pluralize.PS(pluralize.En, "%d %s", minutes, "minute", "minutes"))
		case duration >= 1:
			result = append(result, pluralize.PS(pluralize.En, "%d %s", duration, "second", "seconds"))
			break MAINLOOP
		case duration <= 0 && len(result) == 0:
			return "< 1 second"
		}
	}

	resultLen := len(result)

	if resultLen > 1 {
		return strings.Join(result[:resultLen-1], " ") + " and " + result[resultLen-1]
	}

	return result[0]
}

// codebeat:enable[BLOCK_NESTING]

// Format returns formatted date as a string
//
// Interpreted sequences:
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

// DurationToSeconds converts duration to seconds
func DurationToSeconds(d time.Duration) int64 {
	return int64(d / 1000000000)
}

// ParseDuration parses duration in 1w2d3h5m6s format and return as seconds
func ParseDuration(dur string) int64 {
	if dur == "" {
		return 0
	}

	var (
		result   int64
		value    string
		valueInt int64
	)

	for _, sym := range strings.ToLower(dur) {
		switch sym {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value += string(sym)

		case 'w':
			valueInt, _ = strconv.ParseInt(value, 10, 64)
			result += valueInt * _WEEK
			value = ""

		case 'd':
			valueInt, _ = strconv.ParseInt(value, 10, 64)
			result += valueInt * _DAY
			value = ""

		case 'h':
			valueInt, _ = strconv.ParseInt(value, 10, 64)
			result += valueInt * _HOUR
			value = ""

		case 'm':
			valueInt, _ = strconv.ParseInt(value, 10, 64)
			result += valueInt * _MINUTE
			value = ""

		case 's':
			valueInt, _ = strconv.ParseInt(value, 10, 64)
			result += valueInt
			value = ""
		}
	}

	if value != "" {
		valueInt, _ = strconv.ParseInt(value, 10, 64)
		result += valueInt
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// It's ok to have so long method here
// codebeat:disable[LOC,ABC]

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

// codebeat:enable[LOC,ABC]

func getShortWeekday(d time.Weekday) string {
	long := getLongWeekday(d)

	if long == "" {
		return ""
	}

	return long[:3]
}

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

func getShortMonth(m time.Month) string {
	long := getLongMonth(m)

	if long == "" {
		return ""
	}

	return long[:3]
}

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

func getWeekdayNum(d time.Time) int {
	r := int(d.Weekday())

	if r == 0 {
		r = 7
	}

	return r
}

func getTimezone(d time.Time, separator bool) string {
	negative := false
	_, tzofs := d.Zone()

	if tzofs < 0 {
		negative = true
		tzofs *= -1
	}

	hours := tzofs / _HOUR
	minutes := tzofs % _HOUR

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

func getShortDuration(d time.Duration) string {
	var duration float64

	switch {
	case d > time.Millisecond:
		duration = float64(d) / float64(time.Millisecond)
		return fmt.Sprintf("%g ms", formatFloat(duration))

	case d > time.Microsecond:
		duration = float64(d) / float64(time.Microsecond)
		return fmt.Sprintf("%g μs", formatFloat(duration))

	default:
		return fmt.Sprintf("%d ns", d.Nanoseconds())

	}
}

func formatFloat(f float64) float64 {
	if f < 10.0 {
		return mathutil.Round(f, 2)
	}

	return mathutil.Round(f, 1)
}
