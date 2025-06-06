// Package cron provides methods for working with cron expressions
package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Aliases
const (
	YEARLY   = "0 0 1 1 *"
	ANNUALLY = "0 0 1 1 *"
	MONTHLY  = "0 0 1 * *"
	WEEKLY   = "0 0 * * 0"
	DAILY    = "0 0 * * *"
	HOURLY   = "0 * * * *"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_SYMBOL_PERIOD   = '-'
	_SYMBOL_INTERVAL = '/'
	_SYMBOL_ENUM     = ','
	_SYMBOL_ANY      = '*'
)

const (
	_NAMES_NONE   uint8 = 0
	_NAMES_DAYS   uint8 = 1
	_NAMES_MONTHS uint8 = 2
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Expr cron expression struct
type Expr struct {
	expression string
	minutes    []uint8
	hours      []uint8
	doms       []uint8
	months     []uint8
	dows       []uint8
}

// ////////////////////////////////////////////////////////////////////////////////// //

type exprInfo struct {
	min uint8
	max uint8
	nt  uint8 // Naming type
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrMalformedExpression is returned by the Parse method if given cron expression has
	// wrong or unsupported format
	ErrMalformedExpression = errors.New("Expression must have 5 tokens")

	// ErrZeroInterval is returned if interval part of expression is empty
	ErrZeroInterval = errors.New("Interval can't be less or equals 0")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var info = []exprInfo{
	{0, 59, _NAMES_NONE},
	{0, 23, _NAMES_NONE},
	{1, 31, _NAMES_NONE},
	{1, 12, _NAMES_MONTHS},
	{0, 6, _NAMES_DAYS},
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Parse parse cron expression
// https://en.wikipedia.org/wiki/Cron
func Parse(expr string) (*Expr, error) {
	expr = strings.Replace(expr, "\t", " ", -1)
	expr = getAliasExpression(expr)

	if strings.Count(expr, " ") < 4 {
		return nil, ErrMalformedExpression
	}

	result := &Expr{expression: expr}

	for tn, ei := range info {
		var data []uint8
		var err error

		token := strutil.ReadField(expr, tn, true, ' ')

		switch {
		case isAnyToken(token):
			data = fillUintSlice(ei.min, ei.max, 1)
		case isEnumToken(token):
			data, err = parseEnumToken(token, ei)
		case isPeriodToken(token):
			data, err = parsePeriodToken(token, ei)
		case isIntervalToken(token):
			data, err = parseIntervalToken(token, ei)
		default:
			data, err = parseSimpleToken(token, ei)
		}

		if err != nil {
			return nil, fmt.Errorf("Can't parse token %q: %w", token, err)
		}

		switch tn {
		case 0:
			result.minutes = data
		case 1:
			result.hours = data
		case 2:
			result.doms = data
		case 3:
			result.months = data
		case 4:
			result.dows = data
		}
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsDue check if current moment is match for expression
func (e *Expr) IsDue(args ...time.Time) bool {
	if e == nil {
		return false
	}

	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	if !contains(e.minutes, uint8(t.Minute())) {
		return false
	}

	if !contains(e.hours, uint8(t.Hour())) {
		return false
	}

	if !contains(e.doms, uint8(t.Day())) {
		return false
	}

	if !contains(e.months, uint8(t.Month())) {
		return false
	}

	if !contains(e.dows, uint8(t.Weekday())) {
		return false
	}

	return true
}

// Next get time of next matched moment
func (e *Expr) Next(args ...time.Time) time.Time {
	if e == nil {
		return time.Time{}
	}

	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	year := t.Year()

	mStart := getNearPrevIndex(e.months, uint8(t.Month()))
	dStart := getNearPrevIndex(e.doms, uint8(t.Day()))

	for y := year; y < year+5; y++ {
		for i := mStart; i < len(e.months); i++ {
			for j := dStart; j < len(e.doms); j++ {
				for k := 0; k < len(e.hours); k++ {
					for l := 0; l < len(e.minutes); l++ {
						d := time.Date(
							y,
							time.Month(e.months[i]),
							int(e.doms[j]),
							int(e.hours[k]),
							int(e.minutes[l]),
							0, 0, t.Location(),
						)

						if d.Unix() <= t.Unix() {
							continue
						}

						switch {
						case uint8(d.Month()) != e.months[i],
							uint8(d.Day()) != e.doms[j],
							uint8(d.Hour()) != e.hours[k],
							uint8(d.Minute()) != e.minutes[l],
							!contains(e.dows, uint8(d.Weekday())):
							continue
						}

						return d
					}
				}
			}

			dStart = 0
		}

		mStart = 0
	}

	return time.Unix(0, 0)
}

// Prev get time of prev matched moment
func (e *Expr) Prev(args ...time.Time) time.Time {
	if e == nil {
		return time.Time{}
	}

	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	year := t.Year()

	mStart := getNearNextIndex(e.months, uint8(t.Month()))
	dStart := getNearNextIndex(e.doms, uint8(t.Day()))

	for y := year; y >= year-5; y-- {
		for i := mStart; i >= 0; i-- {
			for j := dStart; j >= 0; j-- {
				for k := len(e.hours) - 1; k >= 0; k-- {
					for l := len(e.minutes) - 1; l >= 0; l-- {
						d := time.Date(
							y,
							time.Month(e.months[i]),
							int(e.doms[j]),
							int(e.hours[k]),
							int(e.minutes[l]),
							0, 0, t.Location(),
						)

						if d.Unix() >= t.Unix() {
							continue
						}

						switch {
						case uint8(d.Month()) != e.months[i],
							uint8(d.Day()) != e.doms[j],
							uint8(d.Hour()) != e.hours[k],
							uint8(d.Minute()) != e.minutes[l],
							!contains(e.dows, uint8(d.Weekday())):
							continue
						}

						return d
					}
				}
			}

			dStart = len(e.doms) - 1
		}

		mStart = len(e.months) - 1
	}

	return time.Unix(0, 0)
}

// String return raw expression
func (e *Expr) String() string {
	if e == nil {
		return "Expr{nil}"
	}

	return e.expression
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isAnyToken(t string) bool {
	return t == string(_SYMBOL_ANY)
}

func isEnumToken(t string) bool {
	return strings.ContainsRune(t, _SYMBOL_ENUM)
}

func isPeriodToken(t string) bool {
	return strings.ContainsRune(t, _SYMBOL_PERIOD)
}

func isIntervalToken(t string) bool {
	return strings.ContainsRune(t, _SYMBOL_INTERVAL)
}

func parseEnumToken(t string, ei exprInfo) ([]uint8, error) {
	var result []uint8

	for i := 0; i <= strings.Count(t, string(_SYMBOL_ENUM)); i++ {
		tt := strutil.ReadField(t, i, false, _SYMBOL_ENUM)

		switch {
		case isPeriodToken(tt):
			d, err := parsePeriodToken(tt, ei)

			if err != nil {
				return nil, err
			}

			result = append(result, d...)

		default:
			t, err := parseToken(tt, ei.nt)

			if err != nil {
				return nil, err
			}

			result = append(result, t)
		}
	}

	return result, nil
}

func parsePeriodToken(t string, ei exprInfo) ([]uint8, error) {
	t1, err := parseToken(strutil.ReadField(t, 0, false, _SYMBOL_PERIOD), ei.nt)

	if err != nil {
		return nil, err
	}

	t2, err := parseToken(strutil.ReadField(t, 1, false, _SYMBOL_PERIOD), ei.nt)

	if err != nil {
		return nil, err
	}

	return fillUintSlice(
		between(t1, ei.min, ei.max),
		between(t2, ei.min, ei.max),
		1,
	), nil
}

func parseIntervalToken(t string, ei exprInfo) ([]uint8, error) {
	i, err := str2uint(strutil.ReadField(t, 1, false, _SYMBOL_INTERVAL))

	if err != nil {
		return nil, err
	}

	if i == 0 {
		return nil, ErrZeroInterval
	}

	return fillUintSlice(ei.min, ei.max, i), nil
}

func parseSimpleToken(t string, ei exprInfo) ([]uint8, error) {
	v, err := parseToken(t, ei.nt)

	if err != nil {
		return nil, err
	}

	return []uint8{v}, nil
}

func getAliasExpression(expr string) string {
	switch expr {
	case "@yearly":
		return YEARLY
	case "@annually":
		return ANNUALLY
	case "@monthly":
		return MONTHLY
	case "@weekly":
		return WEEKLY
	case "@daily":
		return DAILY
	case "@hourly":
		return HOURLY
	}

	return expr
}

func parseToken(t string, nt uint8) (uint8, error) {
	switch nt {
	case _NAMES_DAYS:
		tu, ok := getDayNumByName(t)
		if ok {
			return tu, nil
		}

	case _NAMES_MONTHS:
		tu, ok := getMonthNumByName(t)
		if ok {
			return tu, nil
		}
	}

	return str2uint(t)
}

func getDayNumByName(token string) (uint8, bool) {
	switch strings.ToLower(token) {
	case "sun":
		return 0, true
	case "mon":
		return 1, true
	case "tue":
		return 2, true
	case "wed":
		return 3, true
	case "thu":
		return 4, true
	case "fri":
		return 5, true
	case "sat":
		return 6, true
	}

	return 0, false
}

func getMonthNumByName(token string) (uint8, bool) {
	switch strings.ToLower(token) {
	case "jan":
		return 1, true
	case "feb":
		return 2, true
	case "mar":
		return 3, true
	case "apr":
		return 4, true
	case "may":
		return 5, true
	case "jun":
		return 6, true
	case "jul":
		return 7, true
	case "aug":
		return 8, true
	case "sep":
		return 9, true
	case "oct":
		return 10, true
	case "nov":
		return 11, true
	case "dec":
		return 12, true
	}

	return 0, false
}

func fillUintSlice(start, end, interval uint8) []uint8 {
	var result []uint8

	for i := start; i <= end; i += interval {
		result = append(result, i)
	}

	return result
}

func str2uint(t string) (uint8, error) {
	u, err := strconv.ParseUint(t, 10, 8)

	if err != nil {
		return 0, err
	}

	return uint8(u), nil
}

func getNearNextIndex(items []uint8, item uint8) int {
	for i := range len(items) {
		if items[i] >= item {
			return i
		}
	}

	return 0
}

func getNearPrevIndex(items []uint8, item uint8) int {
	for i := len(items) - 1; i >= 0; i-- {
		if items[i] <= item {
			return i
		}
	}

	return len(items) - 1
}

func between(val, min, max uint8) uint8 {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

func contains(data []uint8, item uint8) bool {
	for _, v := range data {
		if item == v {
			return true
		}
	}

	return false
}
