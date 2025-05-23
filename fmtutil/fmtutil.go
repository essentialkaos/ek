// Package fmtutil provides methods for output formatting
package fmtutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/ansi"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Alignment uint8

const (
	LEFT   Alignment = 0
	CENTER Alignment = 1
	RIGHT  Alignment = 2
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

// OrderSeparator is a default order separator
var OrderSeparator = ","

// SizeSeparator is a default size separator
var SizeSeparator = ""

// ////////////////////////////////////////////////////////////////////////////////// //

var spaces = strings.Repeat(" ", 256)

// ////////////////////////////////////////////////////////////////////////////////// //

// PrettyNum formats number to "pretty" form (e.g 1234567 -> 1,234,567)
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

// PrettyDiff formats number to "pretty" form with + or - symbol at the beginning
func PrettyDiff(i int, separator ...string) string {
	if i > 0 {
		return "+" + PrettyNum(i, separator...)
	}

	return PrettyNum(i, separator...)
}

// PrettyNum formats float value to "pretty" percent form (e.g 12.3423 -> 12.3%)
func PrettyPerc(i float64) string {
	i = Float(i)

	if i < 0.01 {
		return "< 0.01%"
	}

	return PrettyNum(i) + "%"
}

// PrettySize formats value to "pretty" size (e.g 1478182 -> 1.34 Mb)
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

// PrettyBool formats boolean to "pretty" form (e.g true/false -> Y/N)
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

// ParseSize parses size and return it in bytes (e.g 1.34 Mb -> 1478182)
func ParseSize(size string) uint64 {
	v := strings.ToLower(strings.ReplaceAll(size, " ", ""))
	mod, suf := extractSizeInfo(v)

	if suf == "" {
		num, err := strconv.ParseUint(size, 10, 64)

		if err != nil {
			return 0
		}

		return num
	}

	v = strings.TrimRight(v, suf)
	numFlt, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0
	}

	return uint64(numFlt * mod)
}

// Float formats float numbers more nicely
func Float(f float64) float64 {
	if math.IsNaN(f) {
		return 0.0
	}

	if f < 10.0 {
		return mathutil.Round(f, 2)
	}

	return mathutil.Round(f, 1)
}

// Align can align text with ANSI control sequences (for example colors)
func Align(text string, alignment Alignment, size int) string {
	len := strutil.Len(ansi.RemoveCodes(text))

	if len >= size {
		return text
	}

	switch alignment {
	case RIGHT:
		return spaces[:size-len] + text
	case CENTER:
		pad := (size - len) / 2
		return spaces[:pad] +
			text + spaces[:size-(len+pad)]
	default:
		return text + spaces[:size-len]
	}
}

// Wrap wraps text using max line length
func Wrap(text, indent string, maxLineLength int) string {
	var word bytes.Buffer
	var isNewLine bool

	w := &wrapper{
		Indent:        indent,
		MaxLineLength: maxLineLength,
	}

	reader := strings.NewReader(text)

	for i := int64(0); i < reader.Size(); i++ {
		r, _, _ := reader.ReadRune()

		switch r {
		case ' ':
			// break
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
			continue
		}

		w.AddWord(word)
		word.Reset()
	}

	w.AddWord(word)
	word.Reset()

	return w.Result()
}

// ColorizePassword adds different fmtc color tags for numbers and letters
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

// CountDigits returns number of digits in integer
func CountDigits(i int) int {
	if i < 0 {
		return int(math.Log10(math.Abs(float64(i)))) + 2
	}

	return int(math.Log10(float64(i))) + 1
}

// ////////////////////////////////////////////////////////////////////////////////// //

func formatPrettyFloat(str, sep string) string {
	flt := strings.TrimRight(strutil.ReadField(str, 1, false, '.'), "0")

	if flt == "" {
		return appendPrettySymbol(strutil.ReadField(str, 0, false, '.'), sep)
	}

	return appendPrettySymbol(strutil.ReadField(str, 0, false, '.'), sep) + "." + flt
}

func appendPrettySymbol(str, sep string) string {
	if len(str) < 3 {
		return str
	}

	var b strings.Builder

	if str[0] == '-' {
		b.WriteRune('-')
		str = str[1:]
	}

	if len(str)%3 == 0 {
		b.Grow(len(str) + (len(str) / 3) - 1)
	} else {
		b.Grow(len(str) + len(str)/3)
		b.WriteString(str[:(len(str) % 3)])
		b.WriteString(sep)
	}

	for i := len(str) % 3; i < len(str); i++ {
		b.WriteByte(str[i])
		if (1+i-len(str)%3)%3 == 0 && i+1 != len(str) {
			b.WriteString(sep)
		}
	}

	return b.String()
}

// extractSizeInfo extracts size info
func extractSizeInfo(s string) (float64, string) {
	var mod float64

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
		mod = 1
	default:
		suf = ""
	}

	return mod, suf
}

// ////////////////////////////////////////////////////////////////////////////////// //

// AddWord appends word to the line
func (w *wrapper) AddWord(word bytes.Buffer) {
	if word.Len() == 0 {
		return
	}

	var wordLen int

	if ansi.HasCodesBytes(word.Bytes()) {
		wordLen = len(ansi.RemoveCodesBytes(word.Bytes()))
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
