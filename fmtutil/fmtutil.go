// Package fmtutil provides methods for output formatting
package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v9/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_KILO = 1024
	_MEGA = 1048576
	_GIGA = 1073741824
	_TERA = 1099511627776
)

// OrderSeparator default order separator
var OrderSeparator = ","

// SizeSeparator default size separator
var SizeSeparator = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// PrettyNum show pretty num (e.g. 1234567 -> 1,234,567)
func PrettyNum(i interface{}) string {
	return getPrettyNum(i)
}

// PrettySize show pretty size (e.g. 1478182 -> 1.34 Mb)
func PrettySize(i interface{}) string {
	var f float64

	switch i.(type) {
	case int:
		f = float64(i.(int))
	case int32:
		f = float64(i.(int32))
	case int64:
		f = float64(i.(int64))
	case uint:
		f = float64(i.(uint))
	case uint32:
		f = float64(i.(uint32))
	case uint64:
		f = float64(i.(uint64))
	case float32:
		f = float64(i.(float32))
	case float64:
		f = i.(float64)
	}

	switch {
	case f >= _TERA:
		return fmt.Sprintf("%g", Float(f/_TERA)) + SizeSeparator + "TB"
	case f >= _GIGA:
		return fmt.Sprintf("%g", Float(f/_GIGA)) + SizeSeparator + "GB"
	case f >= _MEGA:
		return fmt.Sprintf("%g", Float(f/_MEGA)) + SizeSeparator + "MB"
	case f >= _KILO:
		return fmt.Sprintf("%g", Float(f/_KILO)) + SizeSeparator + "KB"
	default:
		return fmt.Sprintf("%d", i) + SizeSeparator + "B"
	}
}

// ParseSize parse pretty size and return size in bytes
func ParseSize(size string) uint64 {
	var (
		mlt uint64
		pfx string
	)

	ns := strings.ToLower(strings.Replace(size, " ", "", -1))

	switch {
	case strings.Contains(ns, "tb"):
		mlt = 1099511627776
		pfx = "tb"
	case strings.Contains(ns, "gb"):
		mlt = 1073741824
		pfx = "gb"
	case strings.Contains(ns, "mb"):
		mlt = 1048576
		pfx = "mb"
	case strings.Contains(ns, "kb"):
		mlt = 1024
		pfx = "kb"
	case strings.Contains(ns, "b"):
		mlt = 1
		pfx = "b"
	}

	if pfx == "" {
		num, err := strconv.ParseUint(size, 10, 64)

		if err != nil {
			return 0
		}

		return num
	}

	numFlt, err := strconv.ParseFloat(strings.TrimRight(ns, pfx), 64)

	if err != nil {
		return 0
	}

	return uint64(numFlt * float64(mlt))
}

// Float floating number pretty formating
func Float(f float64) float64 {
	if f < 10.0 {
		return mathutil.Round(f, 2)
	}

	return mathutil.Round(f, 1)
}

// Wrap wraps text using max line length
func Wrap(text, indent string, maxLineLength int) string {
	var (
		result = ""
		line   = ""
		words  = strings.Split(text, " ")
	)

	for _, word := range words {
		if strings.Contains(word, "\n") {
			wordSlice := strings.Split(word, "\n")

			if len(wordSlice) == 3 {
				if len(indent+line+wordSlice[0]) > maxLineLength {
					result += indent + line + "\n" + indent + wordSlice[0] + "\n\n"
				} else {
					result += indent + line + wordSlice[0] + "\n\n"
				}

				line = wordSlice[2] + " "

				continue
			} else {
				word = strings.Replace(word, "\n", "", -1)
			}
		}

		if len(indent+line+word) > maxLineLength {
			result += indent + line + "\n"
			line = word + " "

			continue
		}

		line += word + " "
	}

	if line != "" {
		// Append line without last space appended to
		// line in for loop
		result += indent + line[0:len(line)-1]
	}

	return result
}

// ColorizePassword add different fmtc color tags for numbers and letters
func ColorizePassword(password, letterTag, numTag, specialTag string) string {
	var result, curTag, prevTag string

	prevTag = "-"

	for _, r := range password {
		switch {
		case r >= 48 && r <= 57:
			curTag = numTag
		case r >= 91 && r <= 96:
			curTag = specialTag
		case r >= 65 && r <= 122:
			curTag = letterTag
		default:
			curTag = specialTag
		}

		if curTag != prevTag {
			if curTag == "" {
				result += "{!}" + string(r)
			} else {
				result += curTag + string(r)
			}
			prevTag = curTag
		} else {
			result += string(r)
		}
	}

	return result + "{!}"
}

// CountDigits return number of digits in integer
func CountDigits(i int) int {
	if i < 0 {
		return int(math.Log10(math.Abs(float64(i)))) + 2
	}

	return int(math.Log10(float64(i))) + 1
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getPrettyNum(i interface{}) string {
	var str string

	switch v := i.(type) {
	case int, int32, int64, uint, uint32, uint64:
		str = fmt.Sprintf("%d", v)

	case float32, float64:
		str = fmt.Sprintf("%.3f", v)

		return formatPrettyFloat(str)
	}

	return appendPrettySymbol(str)
}

func formatPrettyFloat(str string) string {
	slc := strings.Split(str, ".")
	flt := strings.TrimRight(slc[1], "0")

	if flt == "" {
		return appendPrettySymbol(slc[0])
	}

	return appendPrettySymbol(slc[0]) + "." + flt
}

func appendPrettySymbol(str string) string {
	l := len(str)

	if l <= 3 {
		return str
	}

	r := l % 3
	rs := ""

	for i := l - 3; i >= 0; i -= 3 {
		if i == 0 {
			rs = str[i:i+3] + rs
		} else {
			rs = OrderSeparator + str[i:i+3] + rs
		}
	}

	if r != 0 {
		rs = str[0:r] + rs
	}

	return rs
}
