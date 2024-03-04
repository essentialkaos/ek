package value

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ParseInt64 parses value as Int64
func ParseInt64(v string, defvals ...int64) int64 {
	if v == "" {
		if len(defvals) == 0 {
			return 0
		}

		return defvals[0]
	}

	// HEX Parsing
	if len(v) >= 3 && v[0:2] == "0x" {
		vHex, err := strconv.ParseInt(v[2:], 16, 0)

		if err != nil {
			return 0
		}

		return vHex
	}

	vInt, err := strconv.ParseInt(v, 10, 0)

	if err != nil {
		return 0
	}

	return vInt
}

// ParseInt parses value as Int
func ParseInt(v string, defvals ...int) int {
	if len(defvals) != 0 {
		return int(ParseInt64(v, int64(defvals[0])))
	}

	return int(ParseInt64(v))
}

// ParseUint parses value as Uint
func ParseUint(v string, defvals ...uint) uint {
	if len(defvals) != 0 {
		return uint(ParseInt64(v, int64(defvals[0])))
	}

	return uint(ParseInt64(v))
}

// ParseUint64 parses value as Uint64
func ParseUint64(v string, defvals ...uint64) uint64 {
	if len(defvals) != 0 {
		return uint64(ParseInt64(v, int64(defvals[0])))
	}

	return uint64(ParseInt64(v))
}

// ParseFloat parses value as float
func ParseFloat(v string, defvals ...float64) float64 {
	if v == "" {
		if len(defvals) == 0 {
			return 0.0
		}

		return defvals[0]
	}

	vF, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0.0
	}

	return vF
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

	vM, err := strconv.ParseUint(v, 8, 32)

	if err != nil {
		return 0
	}

	return os.FileMode(vM)
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

// ParseTime parses value as time duration
func ParseTime(v string, defvals ...time.Duration) time.Duration {
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
