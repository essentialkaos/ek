package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/pluralize"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	NS     = time.Nanosecond
	US     = time.Microsecond
	MS     = time.Millisecond
	SECOND = time.Second
	MINUTE = time.Minute
	HOUR   = time.Hour
	DAY    = 24 * HOUR
	WEEK   = 7 * DAY
	YEAR   = 365 * DAY
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SECOND int64 = 1
	_MINUTE int64 = 60
	_HOUR   int64 = 3600
	_DAY    int64 = 86400
	_WEEK   int64 = 604800
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Duration is a wrapper for time.Duration with additional methods for pretty formatting
type Duration struct {
	time.Duration
	isInvalid bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Pretty converts duration from any supported format (numbers or time.Duration)
// into pretty duration wrapper
func Pretty(d any) Duration {
	switch u := d.(type) {
	case time.Duration:
		return Duration{u, false}
	case int:
		return Duration{time.Duration(u) * SECOND, false}
	case int16:
		return Duration{time.Duration(u) * SECOND, false}
	case int32:
		return Duration{time.Duration(u) * SECOND, false}
	case uint:
		return Duration{time.Duration(u) * SECOND, false}
	case uint16:
		return Duration{time.Duration(u) * SECOND, false}
	case uint32:
		return Duration{time.Duration(u) * SECOND, false}
	case uint64:
		return Duration{time.Duration(u) * SECOND, false}
	case float32:
		return Duration{time.Duration(u) * SECOND, false}
	case float64:
		return Duration{time.Duration(u) * SECOND, false}
	case int64:
		return Duration{time.Duration(u) * SECOND, false}
	case string:
		dd, err := ParseDuration(u)

		if err != nil {
			return Duration{0, true}
		}

		return Duration{dd, false}
	}

	return Duration{0, true}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Short returns pretty short duration (e.g. 1:37)
func (d Duration) Short(highPrecision ...bool) string {
	switch {
	case d.isInvalid:
		return ""
	case d.Duration == 0:
		return "0:00"
	}

	if len(highPrecision) != 0 && highPrecision[0] {
		return getShortDuration(d.Duration, true)
	}

	return getShortDuration(d.Duration, false)
}

// Mini returns formatted value of duration (d/hr/m/s/ms/us/ns)
func (d Duration) Mini(separator ...string) string {
	if d.isInvalid {
		return ""
	}

	if len(separator) != 0 {
		return getPrettyMiniDuration(d.Duration, separator[0])
	}

	return getPrettyMiniDuration(d.Duration, " ")
}

// InDays returns pretty duration in days (e.g. 15 days)
func (d Duration) InDays() string {
	var days int

	switch {
	case d.isInvalid:
		return ""
	case d.Duration < 24*HOUR:
		days = 1
	default:
		days = int(d.Hours()) / 24
	}

	return pluralize.PS(pluralize.En, "%d %s", days, "day", "days")
}

// Simple returns simple pretty duration (seconds → minutes → hours → days)
func (d Duration) Simple() string {
	if d.isInvalid {
		return ""
	}

	sec := int64(d.Seconds())

	switch {
	case sec >= _DAY:
		return pluralize.PS(pluralize.En, "%d %s", sec/_DAY, "day", "days")
	case sec >= _HOUR:
		return pluralize.PS(pluralize.En, "%d %s", sec/_HOUR, "hour", "hours")
	case sec >= _MINUTE:
		return pluralize.PS(pluralize.En, "%d %s", sec/_MINUTE, "minute", "minutes")
	case sec >= 1:
		return pluralize.PS(pluralize.En, "%d %s", sec, "second", "seconds")
	}

	return "< 1 second"
}

// String returns pretty duration (e.g. 1 hour 45 seconds)
func (d Duration) String() string {
	if d.isInvalid {
		return ""
	}

	if d.Duration != 0 && d.Duration < SECOND {
		return getPrettyMiniDuration(d.Duration, " ")
	}

	return getPrettyLongDuration(d.Duration)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getShortDuration returns pretty short duration in format "h:mm:ss" or "m:ss" or "m:ss.sss"
func getShortDuration(dur time.Duration, highPrecision bool) string {
	var h, m int64

	s := dur.Seconds()
	d := int64(s)

	if d >= 3600 {
		h = d / 3600
		d %= 3600
	}

	if d >= 60 {
		m = d / 60
		d %= 60
	}

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, d)
	}

	if highPrecision && s < 10.0 {
		ms := fmt.Sprintf("%.3f", s-math.Floor(s))
		ms = strings.TrimPrefix(ms, "0.")
		return fmt.Sprintf("%d:%02d.%s", m, d, ms)
	}

	return fmt.Sprintf("%d:%02d", m, d)
}

// getPrettyMiniDuration returns pretty mini duration (e.g. 1d, 2h, 3m, 4s, 5ms, 6μs, 7ns)
func getPrettyMiniDuration(d time.Duration, separator string) string {
	switch {
	case d >= DAY:
		return fmt.Sprintf(
			"%.0f"+separator+"d",
			formatFloat(float64(d)/float64(DAY)),
		)

	case d >= HOUR:
		return fmt.Sprintf(
			"%.2g"+separator+"h",
			formatFloat(float64(d)/float64(HOUR)),
		)

	case d >= MINUTE:
		return fmt.Sprintf(
			"%.2g"+separator+"m",
			formatFloat(float64(d)/float64(MINUTE)),
		)

	case d >= SECOND:
		return fmt.Sprintf(
			"%.2g"+separator+"s",
			formatFloat(float64(d)/float64(SECOND)),
		)

	case d >= MS:
		return fmt.Sprintf(
			"%g"+separator+"ms",
			formatFloat(float64(d)/float64(MS)),
		)

	case d >= US:
		return fmt.Sprintf(
			"%g"+separator+"μs",
			formatFloat(float64(d)/float64(US)),
		)

	default:
		return fmt.Sprintf("%d"+separator+"ns", d.Nanoseconds())
	}
}

// getPrettyLongDuration returns pretty long duration (e.g. 1 week 2 days 3 hours)
func getPrettyLongDuration(dur time.Duration) string {
	d := int64(dur.Seconds())

	var result []string

MAINLOOP:
	for range 5 {
		switch {
		case d >= _WEEK:
			weeks := d / _WEEK
			d %= _WEEK
			result = append(result, pluralize.PS(pluralize.En, "%d %s", weeks, "week", "weeks"))
		case d >= _DAY:
			days := d / _DAY
			d %= _DAY
			result = append(result, pluralize.PS(pluralize.En, "%d %s", days, "day", "days"))
		case d >= _HOUR:
			hours := d / _HOUR
			d %= _HOUR
			result = append(result, pluralize.PS(pluralize.En, "%d %s", hours, "hour", "hours"))
		case d >= _MINUTE:
			minutes := d / _MINUTE
			d %= _MINUTE
			result = append(result, pluralize.PS(pluralize.En, "%d %s", minutes, "minute", "minutes"))
		case d >= 1:
			result = append(result, pluralize.PS(pluralize.En, "%d %s", d, "second", "seconds"))
			break MAINLOOP
		case d <= 0 && len(result) == 0:
			return "< 1 second"
		}
	}

	resultLen := len(result)

	if resultLen > 1 {
		return strings.Join(result[:resultLen-1], " ") + " and " + result[resultLen-1]
	}

	return result[0]
}

func formatFloat(f float64) float64 {
	if f < 10.0 {
		return mathutil.Round(f, 2)
	}

	return mathutil.Round(f, 1)
}
