// Package fmtc provides methods similar to fmt for colored output
package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v12/color"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_CODE_RESET      = "\033[0m"
	_CODE_CLEAN_LINE = "\033[2K\r"
	_CODE_BACKSPACE  = "\b"
	_CODE_BELL       = "\a"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// codes map tag -> escape code
var codes = map[rune]int{
	// Special
	'-': -1, // Light colors
	'!': 0,  // Default
	'*': 1,  // Bold
	'^': 2,  // Dim
	'_': 4,  // Underline
	'~': 5,  // Blink
	'@': 7,  // Reverse

	// Text
	'd': 30, // Black (Dark)
	'r': 31, // Red
	'g': 32, // Green
	'y': 33, // Yellow
	'b': 34, // Blue
	'm': 35, // Magenta
	'c': 36, // Cyan
	's': 37, // Gray (Smokey)
	'w': 97, // White

	// Background
	'D': 40,  // Black (Dark)
	'R': 41,  // Red
	'G': 42,  // Green
	'Y': 43,  // Yellow
	'B': 44,  // Blue
	'M': 45,  // Magenta
	'C': 46,  // Cyan
	'S': 47,  // Gray (Smokey)
	'W': 107, // White
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DisableColors disables all colors and modificators in output
var DisableColors = os.Getenv("NO_COLOR") != ""

// ////////////////////////////////////////////////////////////////////////////////// //

var colors256Supported bool
var colorsTCSupported bool
var colorsSupportChecked bool

var term = os.Getenv("TERM")
var colorTerm = os.Getenv("COLORTERM")

// ////////////////////////////////////////////////////////////////////////////////// //

// Print formats using the default formats for its operands and writes to standard
// output. Spaces are added between operands when neither is a string. It returns
// the number of bytes written and any write error encountered.
//
// Supported color codes:
//
//    Modificators:
//     - Light colors
//     ! Default
//     * Bold
//     ^ Dim
//     _ Underline
//     ~ Blink
//     @ Reverse
//
//    Foreground colors:
//     d Black (Dark)
//     r Red
//     g Green
//     y Yellow
//     b Blue
//     m Magenta
//     c Cyan
//     s Gray (Smokey)
//     w White
//
//    Background colors:
//     D Black (Dark)
//     R Red
//     G Green
//     Y Yellow
//     B Blue
//     M Magenta
//     C Cyan
//     S Gray (Smokey)
//     W White
//
//    256 colors:
//     #code foreground color
//     %code background color
//
//    24-bit colors (TrueColor):
//      #hex foreground color
//      %hex background color
//
func Print(a ...interface{}) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Print(a...)
}

// Println formats using the default formats for its operands and writes to standard
// output. Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Println(a...)
}

// Printf formats according to a format specifier and writes to standard output. It
// returns the number of bytes written and any write error encountered.
func Printf(f string, a ...interface{}) (int, error) {
	return fmt.Printf(searchColors(f, -1, DisableColors), a...)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. It returns the
// number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Fprint(w, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It returns
// the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...interface{}) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Fprintln(w, a...)
}

// Fprintf formats according to a format specifier and writes to w. It returns
// the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, searchColors(f, -1, DisableColors), a...)
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string {
	applyColors(&a, -1, DisableColors)
	return fmt.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting
// string.
func Sprintf(f string, a ...interface{}) string {
	return fmt.Sprintf(searchColors(f, -1, DisableColors), a...)
}

// Sprintln formats using the default formats for its operands and returns the
// resulting string. Spaces are always added between operands and a newline is
// appended.
func Sprintln(a ...interface{}) string {
	applyColors(&a, -1, DisableColors)
	return fmt.Sprintln(a...)
}

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
func Errorf(f string, a ...interface{}) error {
	return errors.New(Sprintf(f, a...))
}

// TPrintf removes all content on the current line and prints the new message
func TPrintf(f string, a ...interface{}) (int, error) {
	fmt.Printf(_CODE_CLEAN_LINE)
	return Printf(f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// with a new line symbol at the end
func TPrintln(a ...interface{}) (int, error) {
	fmt.Printf(_CODE_CLEAN_LINE)
	return Println(a...)
}

// LPrintf formats according to a format specifier and writes to standard output
// limited by the text size
func LPrintf(maxSize int, f string, a ...interface{}) (int, error) {
	s := fmt.Sprintf(f, a...)
	return fmt.Printf(searchColors(s, maxSize, DisableColors))
}

// LPrintln formats using the default formats for its operands and writes to standard
// output limited by the text size
func LPrintln(maxSize int, a ...interface{}) (int, error) {
	applyColors(&a, maxSize, DisableColors)
	return fmt.Println(a...)
}

// TLPrintf removes all content on the current line and prints the new message
// limited by the text size
func TLPrintf(maxSize int, f string, a ...interface{}) (int, error) {
	fmt.Printf(_CODE_CLEAN_LINE)
	return LPrintf(maxSize, f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// limited by the text size with a new line symbol at the end
func TLPrintln(maxSize int, a ...interface{}) (int, error) {
	fmt.Printf(_CODE_CLEAN_LINE)
	return LPrintln(maxSize, a...)
}

// NewLine prints a newline to standard output
func NewLine() (int, error) {
	return fmt.Println("")
}

// Clean returns string without color tags
func Clean(s string) string {
	return searchColors(s, -1, true)
}

// Bell prints alert (bell) symbol
func Bell() {
	fmt.Printf(_CODE_BELL)
}

// Is256ColorsSupported returns true if 256 colors is supported by terminal
func Is256ColorsSupported() bool {
	if colorsSupportChecked {
		return colors256Supported
	}

	checkForColorsSupport()

	return colors256Supported
}

// IsTrueColorSupported returns true if TrueColor (24-bit colors) is supported by terminal
func IsTrueColorSupported() bool {
	if colorsSupportChecked {
		return colorsTCSupported
	}

	checkForColorsSupport()

	return colorsTCSupported
}

// ////////////////////////////////////////////////////////////////////////////////// //

// codebeat:disable[LOC,BLOCK_NESTING]

func tag2ANSI(tag string, clean bool) string {
	if clean {
		return ""
	}

	if isExtendedColorTag(tag) {
		return parseExtendedColor(tag)
	}

	light := strings.Contains(tag, "-")
	reset := strings.Contains(tag, "!")

	var chars string

	for _, key := range tag {
		code := codes[key]

		switch {
		case light && code == 37: // Light gray = Dark gray
			chars += "90;"
			continue
		case light && code == 97: // Light gray = Dark gray
			chars += "97;"
			continue
		}

		switch key {
		case '-', '!':
			continue

		case '*', '^', '_', '~', '@':
			if reset {
				chars += getResetCode(code)
			} else {
				chars += strconv.Itoa(code)
			}

		case 'D', 'R', 'G', 'Y', 'B', 'M', 'C', 'S', 'W':
			chars += strconv.Itoa(code)

		case 'd', 'r', 'g', 'y', 'b', 'm', 'c', 's', 'w':
			if light {
				chars += strconv.Itoa(code + 60)
			} else {
				chars += strconv.Itoa(code)
			}
		}

		chars += ";"
	}

	if chars == "" {
		return ""
	}

	return fmt.Sprintf("\033[" + chars[:len(chars)-1] + "m")
}

// codebeat:enable[LOC,BLOCK_NESTING]

func parseExtendedColor(tag string) string {
	if len(tag) == 7 {
		hex := strings.TrimLeft(tag, "#%")
		h, _ := color.Parse("#" + hex)
		c := h.ToRGB()

		if strings.HasPrefix(tag, "#") {
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
		}

		return fmt.Sprintf("\033[48;2;%d;%d;%dm", c.R, c.G, c.B)
	}

	if strings.HasPrefix(tag, "#") {
		return "\033[38;5;" + tag[1:] + "m"
	}

	return "\033[48;5;" + tag[1:] + "m"
}

func getResetCode(code int) string {
	if code == codes['*'] {
		code++
	}

	return "2" + strconv.Itoa(code)
}

func replaceColorTags(input, output *bytes.Buffer, clean bool) bool {
	tag := bytes.NewBufferString("")

LOOP:
	for {
		i, _, err := input.ReadRune()

		if err != nil {
			output.WriteString("{" + tag.String())
			return true
		}

		switch i {
		default:
			tag.WriteRune(i)
		case '{':
			output.WriteString("{" + tag.String())
			tag = bytes.NewBufferString("")
		case '}':
			break LOOP
		}
	}

	tagStr := tag.String()

	if !isValidTag(tagStr) {
		output.WriteString("{" + tagStr + "}")
		return true
	}

	if tagStr == "!" {
		if !clean {
			output.WriteString(_CODE_RESET)
		}

		return true
	}

	output.WriteString(tag2ANSI(tagStr, clean))

	return false
}

func searchColors(text string, limit int, clean bool) string {
	if text == "" {
		return ""
	}

	closed := true
	counter := 0
	input := bytes.NewBufferString(text)
	output := bytes.NewBufferString("")

	for {
		i, _, err := input.ReadRune()

		if err != nil {
			break
		}

		switch i {
		case '{':
			closed = replaceColorTags(input, output, clean)
		case rune(65533):
			continue
		default:
			output.WriteRune(i)
			counter++
		}

		if counter == limit {
			break
		}
	}

	if !closed {
		output.WriteString(_CODE_RESET)
	}

	return output.String()
}

func applyColors(a *[]interface{}, limit int, clean bool) {
	for i, x := range *a {
		if s, ok := x.(string); ok {
			(*a)[i] = searchColors(s, limit, clean)
		}
	}
}

func isValidTag(tag string) bool {
	switch {
	case tag == "",
		strings.Trim(tag, "-") == "",
		strings.Count(tag, "!") > 1,
		strings.Contains(tag, "!") && strings.Contains(tag, "-"):
		return false
	}

	if isValidExtendedTag(tag) {
		return true
	}

	for _, r := range tag {
		_, hasCode := codes[r]

		if !hasCode {
			return false
		}
	}

	return true
}

func isExtendedColorTag(tag string) bool {
	return strings.HasPrefix(tag, "#") || strings.HasPrefix(tag, "%")
}

func isValidExtendedTag(tag string) bool {
	if !isExtendedColorTag(tag) {
		return false
	}

	tag = strings.TrimLeft(tag, "#%")

	switch len(tag) {
	case 6:
		hex, err := strconv.ParseInt(tag, 16, 64)
		if err != nil || hex < 0x000000 || hex > 0xffffff {
			return false
		}
	default:
		code, err := strconv.Atoi(tag)
		if err != nil || code < 0 || code > 256 {
			return false
		}
	}

	return true
}

func checkForColorsSupport() {
	if strings.Contains(term, "256color") {
		colors256Supported = true
	}

	if term == "iterm" || colorTerm == "truecolor" ||
		strings.Contains(term, "truecolor") ||
		strings.HasPrefix(term, "vte") {
		colors256Supported, colorsTCSupported = true, true
	}

	colorsSupportChecked = true
}

// ////////////////////////////////////////////////////////////////////////////////// //
