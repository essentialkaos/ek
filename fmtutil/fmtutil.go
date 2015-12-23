// Package fmtutil provides methods for output formating
package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/mathutil"
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

// ////////////////////////////////////////////////////////////////////////////////// //

func getPrettyNum(i interface{}) string {
	var nStr string
	var isFloat bool

	switch v := i.(type) {
	case int, int32, int64, uint, uint32, uint64:
		nStr = fmt.Sprintf("%d", v)

		if len(nStr) <= 3 {
			return nStr
		}
	case float32, float64:
		isFloat = true
		nStr = fmt.Sprintf("%.2f", v)

		if len(nStr) <= 6 {
			return nStr
		}
	}

	if isFloat {
		aStr := strings.Split(nStr, ".")
		cStr := appendPrettySymbol(aStr[0])
		nStr = cStr + "." + aStr[1]
	} else {
		nStr = appendPrettySymbol(nStr)
	}

	return nStr
}

func appendPrettySymbol(s string) string {
	l := len(s)
	r := l % 3

	rs := ""

	for i := l - 3; i >= 0; i -= 3 {
		if i == 0 {
			rs = s[i:i+3] + rs
		} else {
			rs = OrderSeparator + s[i:i+3] + rs
		}
	}

	if r != 0 {
		rs = s[0:r] + rs
	}

	return rs
}
