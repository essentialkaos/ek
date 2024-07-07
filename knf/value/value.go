package value

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek.v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var maxCheckFail = false

// ////////////////////////////////////////////////////////////////////////////////// //

// ParseInt64 parses value as 64-bit int
func ParseInt64(v string, defvals ...int64) int64 {
	if v == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	if isHex(v) {
		h, err := strconv.ParseInt(v[2:], 16, 0)

		if err != nil {
			return 0
		}

		return h
	}

	i, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return 0
	}

	return i
}

// ParseInt parses value as int
func ParseInt(v string, defvals ...int) int {
	var i int64

	if len(defvals) != 0 {
		i = ParseInt64(v, int64(defvals[0]))
	} else {
		i = ParseInt64(v)
	}

	if maxCheckFail || i > math.MaxInt {
		return 0
	}

	return int(i)
}

// ParseUint64 parses value as 64-bit uint
func ParseUint64(v string, defvals ...uint64) uint64 {
	if v == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	if isHex(v) {
		h, err := strconv.ParseUint(v[2:], 16, 0)

		if err != nil {
			return 0
		}

		return h
	}

	i, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		return 0
	}

	return i
}

// ParseUint parses value as uint
func ParseUint(v string, defvals ...uint) uint {
	var u uint64

	if len(defvals) != 0 {
		u = ParseUint64(v, uint64(defvals[0]))
	} else {
		u = ParseUint64(v)
	}

	if maxCheckFail || u > math.MaxUint {
		return 0
	}

	return uint(u)
}

// ParseFloat parses value as float
func ParseFloat(v string, defvals ...float64) float64 {
	if v == "" {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	f, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0.0
	}

	return f
}

// ParseFloat parses value as boolean
func ParseBool(v string, defvals ...bool) bool {
	if v == "" {
		if len(defvals) == 0 {
			return false
		}

		return defvals[0]
	}

	switch strings.ToLower(v) {
	case "", "0", "false", "no":
		return false
	default:
		return true
	}
}

// ParseMode parses value as file mode
func ParseMode(v string, defvals ...os.FileMode) os.FileMode {
	if v == "" {
		if len(defvals) == 0 {
			return os.FileMode(0)
		}

		return defvals[0]
	}

	m, err := strconv.ParseUint(v, 8, 32)

	if err != nil {
		return 0
	}

	return os.FileMode(m)
}

// ParseDuration parses value as duration
func ParseDuration(v string, mod time.Duration, defvals ...time.Duration) time.Duration {
	if v == "" {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	return time.Duration(ParseInt64(v)) * mod
}

// ParseTimeDuration parses value as time duration
func ParseTimeDuration(v string, defvals ...time.Duration) time.Duration {
	if v == "" {
		if len(defvals) == 0 {
			return time.Duration(0)
		}

		return defvals[0]
	}

	v = strings.ToLower(v)
	m := 1

	switch v[len(v)-1:] {
	case "s":
		v, m = v[:len(v)-1], 1
	case "m":
		v, m = v[:len(v)-1], 60
	case "h":
		v, m = v[:len(v)-1], 3600
	case "d":
		v, m = v[:len(v)-1], 24*3600
	case "w":
		v, m = v[:len(v)-1], 7*24*3600
	}

	i, err := strconv.Atoi(v)

	if err != nil {
		return time.Duration(0)
	}

	return time.Duration(i*m) * time.Second
}

// ParseTimestamp parses value as Unix timestamp
func ParseTimestamp(v string, defvals ...time.Time) time.Time {
	if v == "" {
		if len(defvals) == 0 {
			return time.Time{}
		}

		return defvals[0]
	}

	i, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		return time.Time{}
	}

	return time.Unix(i, 0)
}

// ParseTimezone parses value as timezone
func ParseTimezone(v string, defvals ...*time.Location) *time.Location {
	if v == "" {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	l, _ := time.LoadLocation(v)

	return l
}

// ParseList parses value as list
func ParseList(v string, defvals ...[]string) []string {
	if v == "" {
		if len(defvals) == 0 {
			return nil
		}

		return defvals[0]
	}

	return strutil.Fields(v)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isHex(v string) bool {
	return len(v) >= 3 && v[0:2] == "0x"
}
