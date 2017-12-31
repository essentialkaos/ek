// Package fmtc provides methods similar to fmt for colored output
package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_CODE_RESET     = "\033[0m"
	_CODE_BACKSPACE = "\b"
	_CODE_BELL      = "\a"
)

// ////////////////////////////////////////////////////////////////////////////////// //

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

// DisableColors disable all colors and modificators in output
var DisableColors = false

// ////////////////////////////////////////////////////////////////////////////////// //

// tmpSize is size of temporary output
var tmpSize int

// ////////////////////////////////////////////////////////////////////////////////// //

// Println formats using the default formats for its operands and writes to standard
// output. Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
//
// Supported color codes.
// Modificators:
//  - Light colors
//  ! Default
//  * Bold
//  ^ Dim
//  _ Underline
//  ~ Blink
//  @ Reverse
//
// Foreground colors:
//  d Black (Dark)
//  r Red
//  g Green
//  y Yellow
//  b Blue
//  m Magenta
//  c Cyan
//  s Gray (Smokey)
//  w White
//
// Background colors:
//  D Black (Dark)
//  R Red
//  G Green
//  Y Yellow
//  B Blue
//  M Magenta
//  C Cyan
//  S Gray (Smokey)
//  W White
//
func Println(a ...interface{}) (int, error) {
	applyColors(&a, DisableColors)
	return fmt.Println(a...)
}

// Printf formats according to a format specifier and writes to standard output. It
// returns the number of bytes written and any write error encountered.
//
// Supported color codes.
// Modificators:
//  - Light colors
//  ! Default
//  * Bold
//  ^ Dim
//  _ Underline
//  ~ Blink
//  @ Reverse
//
// Foreground colors:
//  d Black (Dark)
//  r Red
//  g Green
//  y Yellow
//  b Blue
//  m Magenta
//  c Cyan
//  s Gray (Smokey)
//  w White
//
// Background colors:
//  D Black (Dark)
//  R Red
//  G Green
//  Y Yellow
//  B Blue
//  M Magenta
//  C Cyan
//  S Gray (Smokey)
//  W White
//
func Printf(f string, a ...interface{}) (int, error) {
	return fmt.Printf(searchColors(f, DisableColors), a...)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. It returns the
// number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	applyColors(&a, DisableColors)
	return fmt.Fprint(w, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It returns
// the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...interface{}) (int, error) {
	applyColors(&a, DisableColors)
	return fmt.Fprintln(w, a...)
}

// Fprintf formats according to a format specifier and writes to w. It returns
// the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, searchColors(f, DisableColors), a...)
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string {
	applyColors(&a, DisableColors)
	return fmt.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting
// string.
func Sprintf(f string, a ...interface{}) string {
	return fmt.Sprintf(searchColors(f, DisableColors), a...)
}

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
func Errorf(f string, a ...interface{}) error {
	return errors.New(Sprintf(f, a...))
}

// NewLine prints a newline to standard output
func NewLine() (int, error) {
	return fmt.Println("")
}

// Clean return string without color tags
func Clean(s string) string {
	return searchColors(s, true)
}

// Bell print alert symbol
func Bell() {
	fmt.Printf(_CODE_BELL)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// TPrintf remove the previous message (if printed) and print new message
func TPrintf(f string, a ...interface{}) (int, error) {
	if tmpSize != 0 {
		fmt.Printf(getSymbols(_CODE_BACKSPACE, tmpSize) + "\033[0K")
	}

	if f == "" && len(a) == 0 {
		tmpSize = 0
		return 0, nil
	}

	tmpSize = strutil.Len(fmt.Sprintf(searchColors(f, true), a...))

	return fmt.Printf(searchColors(f, DisableColors), a...)
}

// TPrintln remove the previous message (if printed) and print new message
func TPrintln(a ...interface{}) (int, error) {
	if tmpSize != 0 {
		fmt.Printf(getSymbols(_CODE_BACKSPACE, tmpSize) + "\033[0K")
	}

	tmpSize = 0

	return Println(a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func tag2ANSI(tag string, clean bool) string {
	if clean {
		return ""
	}

	light := strings.Contains(tag, "-")
	reset := strings.Contains(tag, "!")

	var chars string

	for _, key := range tag {
		code := codes[key]

		// Light gray = Dark gray
		if light && code == 37 {
			chars += "90;"
			continue
		}

		// Light white = white
		if light && code == 97 {
			chars += "97;"
			continue
		}

		switch key {
		case '-':
		case '!':
		case '*':
			if reset {
				chars += "2" + strconv.Itoa(code+1) + ";"
			} else {
				chars += strconv.Itoa(code) + ";"
			}

		case '^', '_', '~', '@':
			if reset {
				chars += "2" + strconv.Itoa(code) + ";"
			} else {
				chars += strconv.Itoa(code) + ";"
			}

		case 'D', 'R', 'G', 'Y', 'B', 'M', 'C', 'S', 'W':
			chars += strconv.Itoa(code) + ";"

		case 'd', 'r', 'g', 'y', 'b', 'm', 'c', 's', 'w':
			if light {
				chars += strconv.Itoa(code+60) + ";"
			} else {
				chars += strconv.Itoa(code) + ";"
			}
		}
	}

	return fmt.Sprintf("\033[" + chars[:len(chars)-1] + "m")
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

func searchColors(text string, clean bool) string {
	if text == "" {
		return ""
	}

	closed := true
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
		}
	}

	if !closed {
		output.WriteString(_CODE_RESET)
	}

	return output.String()
}

func applyColors(a *[]interface{}, clean bool) {
	for i, x := range *a {
		if s, ok := x.(string); ok {
			(*a)[i] = searchColors(s, clean)
		}
	}
}

func getSymbols(symbol string, count int) string {
	result := ""

	for i := 0; i < count; i++ {
		result += symbol
	}

	return result
}

func isValidTag(tag string) bool {
	for _, r := range tag {
		_, hasCode := codes[r]

		if !hasCode {
			return false
		}
	}

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //
