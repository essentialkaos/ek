// Package cron provides methods for working with cron expressions
package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"strconv"
	"strings"
	"time"
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
	_SYMBOL_PERIOD   = "-"
	_SYMBOL_INTERVAL = "/"
	_SYMBOL_ENUM     = ","
	_SYMBOL_ANY      = "*"
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
	minutes    *exprPart
	hours      *exprPart
	doms       *exprPart
	months     *exprPart
	dows       *exprPart
}

type exprPart struct {
	index  map[uint8]bool
	tokens []uint8
}

type exprInfo struct {
	min uint8
	max uint8
	nt  uint8 // Naming type
}

// ////////////////////////////////////////////////////////////////////////////////// //

var ErrMalformedExpression = errors.New("Expression must have 5 tokens")

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
func Parse(expr string) (*Expr, error) {
	result := &Expr{expression: expr}

	expr = strings.Replace(expr, "\t", " ", -1)
	expr = getAliasExpression(expr)

	exprAr := strings.Split(expr, " ")

	if len(exprAr) != 5 {
		return nil, ErrMalformedExpression
	}

	data := make([][]uint8, 5)

	for tn, ei := range info {
		token := exprAr[tn]

		switch {
		case isAnyToken(token):
			data[tn] = fillUintSlice(ei.min, ei.max, 1)
		case isEnumToken(token):
			data[tn] = getEnumFromToken(token, ei.nt)
		case isPeriodToken(token):
			ts, te := getPeriodFromToken(token, ei.nt)
			ts = between(ts, ei.min, ei.max)
			te = between(te, ei.min, ei.max)
			data[tn] = fillUintSlice(ts, te, 1)
		case isIntervalToken(token):
			data[tn] = fillUintSlice(ei.min, ei.max, getIntervalFromToken(token))
		default:
			data[tn] = []uint8{parseToken(token, ei.nt)}
		}
	}

	result.minutes = &exprPart{slice2map(data[0]), data[0]}
	result.hours = &exprPart{slice2map(data[1]), data[1]}
	result.doms = &exprPart{slice2map(data[2]), data[2]}
	result.months = &exprPart{slice2map(data[3]), data[3]}
	result.dows = &exprPart{slice2map(data[4]), data[4]}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsDue check if current moment is match fo expression
func (expr *Expr) IsDue(args ...time.Time) bool {
	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	if expr.minutes.index[uint8(t.Minute())] == false {
		return false
	}

	if expr.hours.index[uint8(t.Hour())] == false {
		return false
	}

	if expr.doms.index[uint8(t.Day())] == false {
		return false
	}

	if expr.months.index[uint8(t.Month())] == false {
		return false
	}

	if expr.dows.index[uint8(t.Weekday())] == false {
		return false
	}

	return true
}

// Next get time of next matched moment
func (expr *Expr) Next(args ...time.Time) time.Time {
	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	year := t.Year()

	mStart := getNearPrevIndex(expr.months.tokens, uint8(t.Month()))
	dStart := getNearPrevIndex(expr.doms.tokens, uint8(t.Day()))

	for y := year; y < year+5; y++ {
		for i := mStart; i < len(expr.months.tokens); i++ {
			for j := dStart; j < len(expr.doms.tokens); j++ {
				for k := 0; k < len(expr.hours.tokens); k++ {
					for l := 0; l < len(expr.minutes.tokens); l++ {
						d := time.Date(
							y,
							time.Month(expr.months.tokens[i]),
							int(expr.doms.tokens[j]),
							int(expr.hours.tokens[k]),
							int(expr.minutes.tokens[l]),
							0, 0, t.Location())

						if d.Unix() <= t.Unix() {
							continue
						}

						switch {
						case uint8(d.Month()) != expr.months.tokens[i],
							uint8(d.Day()) != expr.doms.tokens[j],
							uint8(d.Hour()) != expr.hours.tokens[k],
							uint8(d.Minute()) != expr.minutes.tokens[l],
							expr.dows.index[uint8(d.Weekday())] == false:
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
func (expr *Expr) Prev(args ...time.Time) time.Time {
	var t time.Time

	if len(args) >= 1 {
		t = args[0]
	} else {
		t = time.Now()
	}

	year := t.Year()

	mStart := getNearNextIndex(expr.months.tokens, uint8(t.Month()))
	dStart := getNearNextIndex(expr.doms.tokens, uint8(t.Day()))

	for y := year; y >= year-5; y-- {
		for i := mStart; i >= 0; i-- {
			for j := dStart; j >= 0; j-- {
				for k := len(expr.hours.tokens) - 1; k >= 0; k-- {
					for l := len(expr.minutes.tokens) - 1; l >= 0; l-- {
						d := time.Date(
							y,
							time.Month(expr.months.tokens[i]),
							int(expr.doms.tokens[j]),
							int(expr.hours.tokens[k]),
							int(expr.minutes.tokens[l]),
							0, 0, t.Location())

						if d.Unix() >= t.Unix() {
							continue
						}

						switch {
						case uint8(d.Month()) != expr.months.tokens[i],
							uint8(d.Day()) != expr.doms.tokens[j],
							uint8(d.Hour()) != expr.hours.tokens[k],
							uint8(d.Minute()) != expr.minutes.tokens[l],
							expr.dows.index[uint8(d.Weekday())] == false:
							continue
						}

						return d
					}
				}
			}

			dStart = len(expr.doms.tokens) - 1
		}

		mStart = len(expr.months.tokens) - 1
	}

	return time.Unix(0, 0)
}

// String return raw expression
func (expr *Expr) String() string {
	return expr.expression
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isAnyToken(t string) bool {
	return t == _SYMBOL_ANY
}

func isEnumToken(t string) bool {
	return strings.Contains(t, _SYMBOL_ENUM)
}

func isPeriodToken(t string) bool {
	return strings.Contains(t, _SYMBOL_PERIOD)
}

func isIntervalToken(t string) bool {
	return strings.Contains(t, _SYMBOL_INTERVAL)
}

func getEnumFromToken(t string, nt uint8) []uint8 {
	var result []uint8

	for _, tt := range strings.Split(t, _SYMBOL_ENUM) {
		switch {
		case isPeriodToken(tt):
			ts, te := getPeriodFromToken(tt, nt)
			result = append(result, fillUintSlice(ts, te, 1)...)
		default:
			result = append(result, parseToken(tt, nt))
		}
	}

	return result
}

func getPeriodFromToken(t string, nt uint8) (uint8, uint8) {
	ts := strings.Split(t, _SYMBOL_PERIOD)

	return parseToken(ts[0], nt), parseToken(ts[1], nt)
}

func getIntervalFromToken(t string) uint8 {
	ts := strings.Split(t, _SYMBOL_INTERVAL)

	return str2uint(ts[1])
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

func parseToken(t string, nt uint8) uint8 {
	switch nt {
	case _NAMES_DAYS:
		tu, ok := getDayNumByName(t)
		if ok {
			return tu
		}

	case _NAMES_MONTHS:
		tu, ok := getMonthNumByName(t)
		if ok {
			return tu
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

func slice2map(s []uint8) map[uint8]bool {
	result := make(map[uint8]bool)

	for _, u := range s {
		result[u] = true
	}

	return result
}

func str2uint(t string) uint8 {
	u, _ := strconv.ParseUint(t, 10, 8)
	return uint8(u)
}

func getNearNextIndex(items []uint8, item uint8) int {
	for i := 0; i < len(items); i++ {
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
