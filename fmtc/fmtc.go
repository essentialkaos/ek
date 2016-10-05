// Package fmtc provides methods similar to fmt for colored output
package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_CODE_RESET     = "\033[0m"
	_CODE_BACKSPACE = "\b"
	_CODE_BELL      = "\a"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// T is struct can be used for printing temporary messages
type T struct {
	size int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var codes = map[int]int{
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
	's': 37, // Grey (Smokey)
	'w': 97, // White

	// Background
	'D': 40,  // Black (Dark)
	'R': 41,  // Red
	'G': 42,  // Green
	'Y': 43,  // Yellow
	'B': 44,  // Blue
	'M': 45,  // Magenta
	'C': 46,  // Cyan
	'S': 47,  // Grey (Smokey)
	'W': 107, // White
}

// ////////////////////////////////////////////////////////////////////////////////// //

// DisableColors disable all colors and modificators in output
var DisableColors = false

// ////////////////////////////////////////////////////////////////////////////////// //

// Println formats using the default formats for its operands and writes to standard
// output. Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (int, error) {
	applyColors(&a, DisableColors)
	return fmt.Println(a...)
}

// Printf formats according to a format specifier and writes to standard output. It
// returns the number of bytes written and any write error encountered.
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

// Printf remove previous message (if printed) and print new message
func (t *T) Printf(f string, a ...interface{}) (int, error) {
	if t.size != 0 {
		fmt.Printf(getSymbols(_CODE_BACKSPACE, t.size) + "\033[0K")
	}

	t.size = len(fmt.Sprintf(searchColors(f, true), a...))

	return fmt.Printf(searchColors(f, DisableColors), a...)
}

// Println remove previous message (if printed) and print new message
func (t *T) Println(a ...interface{}) (int, error) {
	if t.size != 0 {
		fmt.Printf(getSymbols(_CODE_BACKSPACE, t.size) + "\033[0K")
	}

	t.size = 0

	return Println(a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func tag2ANSI(tag string, clean bool) string {
	if clean {
		return ""
	}

	var (
		modificator = 0
		charColor   = 39
		bgColor     = 49
		light       = false
	)

	for _, key := range tag {
		code, ok := codes[int(key)]

		if !ok {
			return fmt.Sprint(tag)
		}

		switch key {
		case '-':
			light = true
		case '!', '*', '^', '_', '~', '@':
			modificator = code
		case 'd', 'r', 'g', 'y', 'b', 'm', 'c', 's', 'w':
			charColor = code
		case 'D', 'R', 'G', 'Y', 'B', 'M', 'C', 'S', 'W':
			bgColor = code
		}
	}

	if light {
		switch charColor {
		case 97:
			break
		case 37:
			charColor = 90
		default:
			charColor += 60
		}
	}

	return fmt.Sprintf("\033[%d;%d;%dm", modificator, charColor, bgColor)
}

func replaceColorTags(input, output *bytes.Buffer, clean bool) bool {
	tag := bytes.NewBufferString("")

LOOP:
	for {
		i, _, err := input.ReadRune()

		if err != nil {
			output.WriteString("{")
			output.WriteString(tag.String())
			return true
		}

		switch i {
		default:
			tag.WriteRune(i)
		case '{':
			output.WriteString("{")
			output.WriteString(tag.String())
			tag = bytes.NewBufferString("")
		case '}':
			break LOOP
		}
	}

	tagStr := tag.String()

	if tagStr == "!" {
		if !clean {
			output.WriteString(_CODE_RESET)
		}

		return true
	}

	colorCode := tag2ANSI(tagStr, clean)

	if colorCode == tagStr {
		output.WriteString("{")
		output.WriteString(colorCode)
		output.WriteString("}")

		return true
	}

	output.WriteString(colorCode)

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

// ////////////////////////////////////////////////////////////////////////////////// //
