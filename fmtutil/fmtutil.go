// Package fmtutil provides methods for output formatting
package fmtutil

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
	"strings"

	"github.com/essentialkaos/ek/v14/ansi"
	"github.com/essentialkaos/ek/v14/mathutil"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Alignment defines the horizontal text alignment strategy
type Alignment uint8

const (
	LEFT   Alignment = 0 // Align text to the left, padding on the right
	CENTER Alignment = 1 // Center text, distributing padding on both sides
	RIGHT  Alignment = 2 // Align text to the right, padding on the left
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_KILO float64 = 1024
	_MEGA float64 = 1048576
	_GIGA float64 = 1073741824
	_TERA float64 = 1099511627776
)

// ////////////////////////////////////////////////////////////////////////////////// //

type wrapper struct {
	Indent        string
	MaxLineLength int

	result  bytes.Buffer
	line    bytes.Buffer
	lineLen int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// OrderSeparator is the thousands separator used by [PrettyNum] and [PrettyPerc]
var OrderSeparator = ","

// SizeSeparator is placed between the numeric value and unit in [PrettySize]
var SizeSeparator = ""

// ////////////////////////////////////////////////////////////////////////////////// //

var spaces string

// ////////////////////////////////////////////////////////////////////////////////// //

// PrettyNum formats any numeric value with thousands separators
// (e.g. 1234567 → "1,234,567").
// An optional separator overrides OrderSeparator for this call only.
func PrettyNum(i any, separator ...string) string {
	var str string

	sep := OrderSeparator

	if len(separator) > 0 {
		sep = separator[0]
	}

	switch v := i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		str = fmt.Sprintf("%d", v)

		return appendPrettySymbol(str, sep)

	case float32, float64:
		str = fmt.Sprintf("%.3f", v)

		if str == "NaN" {
			return "0"
		}

		return formatPrettyFloat(str, sep)
	}

	// Return value for unsupported types as is
	return fmt.Sprintf("%v", i)
}

// PrettyDiff formats an integer like [PrettyNum] but prepends "+" for positive values
func PrettyDiff(i int, separator ...string) string {
	if i > 0 {
		return "+" + PrettyNum(i, separator...)
	}

	return PrettyNum(i, separator...)
}

// PrettyPerc formats a float value as a human-readable percentage
// (e.g. 12.3423 → "12.3%"). Values below 0.01 are rendered as "< 0.01%".
func PrettyPerc(i float64) string {
	i = Float(i)

	if i < 0.01 {
		return "< 0.01%"
	}

	return PrettyNum(i) + "%"
}

// PrettySize formats a numeric byte count as a human-readable size string
// (e.g. 184713 → "180.4KB"). An optional separator overrides SizeSeparator for
// this call.
func PrettySize[N mathutil.Numeric](i N, separator ...string) string {
	sep := SizeSeparator

	if len(separator) > 0 {
		sep = separator[0]
	}

	f := float64(i)

	if math.IsNaN(f) {
		return "0" + sep + "B"
	}

	switch {
	case math.Abs(f) >= _TERA:
		return fmt.Sprintf("%g", Float(f/_TERA)) + sep + "TB"
	case math.Abs(f) >= _GIGA:
		return fmt.Sprintf("%g", Float(f/_GIGA)) + sep + "GB"
	case math.Abs(f) >= _MEGA:
		return fmt.Sprintf("%g", Float(f/_MEGA)) + sep + "MB"
	case math.Abs(f) >= _KILO:
		return fmt.Sprintf("%g", Float(f/_KILO)) + sep + "KB"
	default:
		return fmt.Sprintf("%g", mathutil.Round(f, 0)) + sep + "B"
	}
}

// PrettyBool formats a boolean as "Y"/"N" by default.
// Provide two optional strings to use custom true/false labels (e.g. "Yep",
// "Nope").
func PrettyBool(b bool, vals ...string) string {
	switch {
	case b && len(vals) >= 1:
		return vals[0]
	case !b && len(vals) >= 2:
		return vals[1]
	case b:
		return "Y"
	}

	return "N"
}

// ParseSize parses a human-readable size string and returns the equivalent byte count
// (e.g. "2.2 GB" → 2362232012).
func ParseSize(size string) (uint64, error) {
	v := strings.ToLower(strings.ReplaceAll(size, " ", ""))
	mod, suf := extractSizeInfo(v)

	v = strings.TrimSuffix(v, suf)

	if v == "" {
		return 0, fmt.Errorf("size has no digits")
	}

	numFlt, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0, err
	}

	return uint64(numFlt * mod), nil
}

// Float rounds f to 2 decimal places below 10, or 1 decimal place above. Returns 0.0
// for NaN.
func Float(f float64) float64 {
	if math.IsNaN(f) {
		return 0.0
	}

	if f < 10.0 {
		return mathutil.Round(f, 2)
	}

	return mathutil.Round(f, 1)
}

// Align pads text to size columns, respecting ANSI escape codes in width calculation
func Align(text string, alignment Alignment, size int) string {
	vlen := strutil.Len(ansi.Remove(text))

	if vlen >= size {
		return text
	}

	switch alignment {
	case RIGHT:
		return padding(size-vlen) + text
	case CENTER:
		pad := (size - vlen) / 2
		return padding(pad) +
			text + padding(size-(vlen+pad))
	default:
		return text + padding(size-vlen)
	}
}

// Wrap breaks text into lines no longer than maxLineLength columns, prefixing each
// line with indent. Existing newlines are preserved; double newlines become blank lines.
func Wrap(text, indent string, maxLineLength int) string {
	var word bytes.Buffer
	var isNewLine bool

	w := &wrapper{
		Indent:        indent,
		MaxLineLength: maxLineLength,
	}

	for _, r := range text {
		switch r {
		case ' ':
			w.AddWord(word)
			word.Reset()

		case '\n', '\r':
			if !isNewLine {
				isNewLine = true
				w.AddWord(word)
				word.Reset()
			} else {
				w.NewLine()
			}

		default:
			isNewLine = false
			word.WriteRune(r)
		}
	}

	w.AddWord(word)
	word.Reset()

	return w.Result()
}

// ColorizePassword wraps character groups in fmtc colour tags. letterTag is applied
// to letters, numTag to digits, and specialTag to everything else.
func ColorizePassword(password, letterTag, numTag, specialTag string) string {
	var curTag string

	prevTag := "-"

	var b strings.Builder
	b.Grow(len(password) * 2)

	for _, r := range password {
		switch {
		case r >= '0' && r <= '9':
			curTag = numTag
		case r >= '[' && r <= '`': // punctuation between upper/lower ASCII
			curTag = specialTag
		case r >= 'A' && r <= 'z':
			curTag = letterTag
		default:
			curTag = specialTag
		}

		if curTag != prevTag {
			if curTag == "" {
				b.WriteString("{!}")
				b.WriteRune(r)
			} else {
				b.WriteString(curTag)
				b.WriteRune(r)
			}
			prevTag = curTag
		} else {
			b.WriteRune(r)
		}
	}

	return b.String() + "{!}"
}

// CountDigits returns the number of decimal digits in i, including the minus sign
// for negatives.
func CountDigits(i int) int {
	switch {
	case i == 0:
		return 1
	case i < 0:
		return int(math.Log10(math.Abs(float64(i)))) + 2
	}

	return int(math.Log10(float64(i))) + 1
}

// ////////////////////////////////////////////////////////////////////////////////// //

// formatPrettyFloat inserts thousands separators into the integer part of a decimal
// string and strips trailing fractional zeros
func formatPrettyFloat(str, sep string) string {
	flt := strings.TrimRight(strutil.ReadField(str, 1, false, '.'), "0")

	if flt == "" {
		return appendPrettySymbol(strutil.ReadField(str, 0, false, '.'), sep)
	}

	return appendPrettySymbol(strutil.ReadField(str, 0, false, '.'), sep) + "." + flt
}

// appendPrettySymbol inserts sep every three digits from the right
// (e.g. "1234567" → "1,234,567")
func appendPrettySymbol(str, sep string) string {
	if len(str) < 3 {
		return str
	}

	var b strings.Builder

	if str[0] == '-' {
		b.WriteRune('-')
		str = str[1:]
	}

	offset := len(str) % 3

	for i, ch := range str {
		if i > 0 && (i-offset)%3 == 0 {
			b.WriteString(sep)
		}

		b.WriteRune(ch)
	}

	return b.String()
}

// extractSizeInfo returns the byte multiplier and suffix for a size string
func extractSizeInfo(s string) (float64, string) {
	mod := 1.0
	suf := strings.TrimLeft(s, "0123456789. ")

	switch suf {
	case "tb", "tib":
		mod = 1024 * 1024 * 1024 * 1024
	case "t":
		mod = 1000 * 1000 * 1000 * 1000
	case "gb", "gib":
		mod = 1024 * 1024 * 1024
	case "g", "kkk":
		mod = 1000 * 1000 * 1000
	case "mb", "mib":
		mod = 1024 * 1024
	case "m", "kk":
		mod = 1000 * 1000
	case "kb":
		mod = 1024
	case "k":
		mod = 1000
	case "b":
		// okay
	default:
		suf = ""
	}

	return mod, suf
}

// padding generates string with spaces
func padding(n int) string {
	if n <= 0 {
		return ""
	}

	if spaces == "" {
		spaces = strings.Repeat(" ", 256)
	}

	return spaces[:min(n, 256)]
}

// ////////////////////////////////////////////////////////////////////////////////// //

// AddWord appends word to the line
func (w *wrapper) AddWord(word bytes.Buffer) {
	if word.Len() == 0 {
		return
	}

	var wordLen int

	if ansi.HasBytes(word.Bytes()) {
		wordLen = len(ansi.RemoveBytes(word.Bytes()))
	} else {
		wordLen = len(word.String())
	}

	if w.line.Len() != 0 && len(w.Indent)+w.lineLen+wordLen > w.MaxLineLength {
		w.result.WriteString(w.Indent)
		w.result.Write(w.line.Bytes())
		w.result.WriteRune('\n')
		w.line.Reset()
		w.lineLen = wordLen + 1
	} else {
		w.lineLen += wordLen + 1
		if w.line.Len() != 0 {
			w.line.WriteRune(' ')
		}
	}

	w.line.Write(word.Bytes())
}

// NewLine adds new line
func (w *wrapper) NewLine() {
	w.result.WriteString(w.Indent)
	w.result.Write(w.line.Bytes())
	w.result.WriteRune('\n')
	w.result.WriteRune('\n')
	w.line.Reset()
	w.lineLen = 0
}

// Result returns result as a string
func (w *wrapper) Result() string {
	// If line buffer isn't empty append it to the result
	if w.line.Len() != 0 {
		w.result.WriteString(w.Indent)
		w.result.Write(w.line.Bytes())
	}

	w.line.Reset()

	return w.result.String()
}
