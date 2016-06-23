// Package timeutil with time utils
package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v2/fmtutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_MINUTE = 60
	_HOUR   = 3600
	_DAY    = 86400
	_WEEK   = 604800
)

// ////////////////////////////////////////////////////////////////////////////////// //

// PrettyDuration return pretty duration (e.g. 1 hour 45 seconds)
func PrettyDuration(d interface{}) string {
	var (
		r []string
		t int
	)

	switch d.(type) {
	case time.Duration:
		t = int(d.(time.Duration).Seconds())
	case int8:
		t = int(d.(int8))
	case int16:
		t = int(d.(int16))
	case int32:
		t = int(d.(int32))
	case int64:
		t = int(d.(int64))
	case int:
		t = d.(int)
	default:
		return "Wrong duration value"
	}

	for i := 0; i < 5; i++ {
		if t >= _WEEK {
			weeks := t / _WEEK
			t = t % _WEEK
			r = append(r, fmtutil.Pluralize(weeks, "week", "weeks"))
		} else if t >= _DAY {
			days := t / _DAY
			t = t % _DAY
			r = append(r, fmtutil.Pluralize(days, "day", "days"))
		} else if t >= _HOUR {
			hours := t / _HOUR
			t = t % _HOUR
			r = append(r, fmtutil.Pluralize(hours, "hour", "hours"))
		} else if t >= _MINUTE {
			minutes := t / _MINUTE
			t = t % _MINUTE
			r = append(r, fmtutil.Pluralize(minutes, "minute", "minutes"))
		} else if t >= 1 {
			if len(r) != 0 {
				r = append(r, "and")
			}
			r = append(r, fmtutil.Pluralize(t, "second", "seconds"))
			break
		} else if t <= 0 && len(r) == 0 {
			return "< 1 second"
		}
	}

	return strings.Join(r, " ")
}

// Format return formated date to string with linux date formating
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

// DurationToSeconds convert duration to seconds
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

func replaceDateTag(d time.Time, input, output *bytes.Buffer) {
	r, _, err := input.ReadRune()

	if err != nil {
		return
	}

	switch r {
	case '%':
		output.WriteString("%")
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
		output.WriteString(fmt.Sprintf("%s %02d %s %d %02d:%02d:%02d %s %s",
			getShortWeekday(d.Weekday()),
			d.Day(),
			getShortMonth(d.Month()),
			d.Year(),
			getAMPMHour(d),
			d.Minute(),
			d.Second(),
			getAMPM(d, true),
			zn))
	case 'C', 'g':
		output.WriteString(strconv.Itoa(d.Year())[0:2])
	case 'd':
		output.WriteString(fmt.Sprintf("%02d", d.Day()))
	case 'D':
		output.WriteString(fmt.Sprintf("%02d/%02d/%s", d.Month(), d.Day(), strconv.Itoa(d.Year())[2:4]))
	case 'e':
		if d.Day() >= 10 {
			output.WriteString(strconv.Itoa(d.Day()))
		} else {
			output.WriteString(" " + strconv.Itoa(d.Day()))
		}
	case 'F':
		output.WriteString(fmt.Sprintf("%d-%02d-%02d", d.Year(), d.Month(), d.Day()))
	case 'G':
		output.WriteString(strconv.Itoa(d.Year()))
	case 'H':
		output.WriteString(fmt.Sprintf("%02d", d.Hour()))
	case 'I':
		output.WriteString(fmt.Sprintf("%02d", getAMPMHour(d)))
	case 'j':
		output.WriteString(fmt.Sprintf("%03d", d.YearDay()))
	case 'k':
		output.WriteString(" " + strconv.Itoa(d.Hour()))
	case 'l':
		output.WriteString(strconv.Itoa(getAMPMHour(d)))
	case 'm':
		output.WriteString(fmt.Sprintf("%02d", d.Month()))
	case 'M':
		output.WriteString(fmt.Sprintf("%02d", d.Minute()))
	case 'N':
		output.WriteString(fmt.Sprintf("%09d", d.Nanosecond()))
	case 'p':
		output.WriteString(getAMPM(d, false))
	case 'P':
		output.WriteString(getAMPM(d, true))
	case 'r':
		output.WriteString(fmt.Sprintf("%02d:%02d:%02d %s", getAMPMHour(d), d.Minute(), d.Second(), getAMPM(d, true)))
	case 'R':
		output.WriteString(fmt.Sprintf("%02d:%02d", d.Hour(), d.Minute()))
	case 's':
		output.WriteString(strconv.FormatInt(d.Unix(), 10))
	case 'S':
		output.WriteString(fmt.Sprintf("%02d", d.Second()))
	case 'T':
		output.WriteString(fmt.Sprintf("%02d:%02d:%02d", d.Hour(), d.Minute(), d.Second()))
	case 'u':
		output.WriteString(strconv.Itoa(getWeekdayNum(d)))
	case 'V':
		_, wn := d.ISOWeek()
		output.WriteString(fmt.Sprintf("%02d", wn))
	case 'w':
		output.WriteString(strconv.Itoa(int(d.Weekday())))
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

	if h == 0 || h == 12 {
		return 12
	} else if h < 12 {
		return h
	} else {
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
