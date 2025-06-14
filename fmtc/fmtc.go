// Package fmtc provides methods similar to fmt for colored output
package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
	"sync"

	"github.com/essentialkaos/ek/v13/color"
	"github.com/essentialkaos/ek/v13/env"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_CODE_RESET      = "\033[0m"
	_CODE_CLEAN_LINE = "\033[2K\r"
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
	'&': 3,  // Italic
	'_': 4,  // Underline
	'~': 5,  // Blink
	'@': 7,  // Reverse
	'+': 8,  // Hidden
	'=': 9,  // Strikethrough

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

var isColorsSupported bool    // 16 colors support
var isColors256Supported bool // 256 colors support
var isColorsTCSupported bool  // 24bit (TrueColor) colors support
var isColorsSupportChecked bool

var boldDisableEnvVar = env.Var("FMTC_NO_BOLD")
var italicDisableEnvVar = env.Var("FMTC_NO_ITALIC")
var blinkDisableEnvVar = env.Var("FMTC_NO_BLINK")

var colorsMap *sync.Map

var termEnvVar = env.Var("TERM")
var colorTermEnvVar = env.Var("COLORTERM")

// ////////////////////////////////////////////////////////////////////////////////// //

// AddColor defines or redifines named color
func AddColor(name, tag string) error {
	if colorsMap == nil {
		colorsMap = &sync.Map{}
	}

	switch {
	case name == "":
		return errors.New("Can't add named color: name can't be empty")
	case tag == "":
		return errors.New("Can't add named color: tag can't be empty")
	case !IsTag(tag):
		return fmt.Errorf("Can't add named color: %q is not valid color tag", tag)
	case !isValidNamedTag("?" + name):
		return fmt.Errorf("Can't add named color: %q is not valid name", name)
	}

	colorsMap.Store(name, searchColors(tag, -1, false, false))

	return nil
}

// RemoveColor removes named color
func RemoveColor(name string) {
	if colorsMap == nil || name == "" {
		return
	}

	colorsMap.Delete(name)
}

// Print formats using the default formats for its operands and writes to standard
// output. Spaces are added between operands when neither is a string. It returns
// the number of bytes written and any write error encountered.
//
// Supported color codes:
//
//	Modificators:
//	 - Light colors
//	 ! Default
//	 * Bold
//	 & Italic
//	 ^ Dim
//	 _ Underline
//	 = Strikethrough
//	 ~ Blink
//	 @ Reverse
//	 + Hidden
//
//	Foreground colors:
//	 d Black (Dark)
//	 r Red
//	 g Green
//	 y Yellow
//	 b Blue
//	 m Magenta
//	 c Cyan
//	 s Gray (Smokey)
//	 w White
//
//	Background colors:
//	 D Black (Dark)
//	 R Red
//	 G Green
//	 Y Yellow
//	 B Blue
//	 M Magenta
//	 C Cyan
//	 S Gray (Smokey)
//	 W White
//
//	256 colors:
//	 #code foreground color
//	 %code background color
//
//	24-bit colors (TrueColor):
//	  #hex foreground color
//	  %hex background color
//
//	Named colors:
//	  ?name
func Print(a ...any) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Print(a...)
}

// Println formats using the default formats for its operands and writes to standard
// output. Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
func Println(a ...any) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Println(a...)
}

// Printf formats according to a format specifier and writes to standard output. It
// returns the number of bytes written and any write error encountered.
func Printf(f string, a ...any) (int, error) {
	return fmt.Printf(searchColors(f, -1, DisableColors, true), a...)
}

// Printfn formats according to a format specifier and writes to standard output with
// the new line at the end. It returns the number of bytes written and any write
// error encountered.
func Printfn(f string, a ...any) (int, error) {
	return fmt.Printf(searchColors(f, -1, DisableColors, true)+"\n", a...)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. It returns the
// number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...any) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Fprint(w, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It returns
// the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...any) (int, error) {
	applyColors(&a, -1, DisableColors)
	return fmt.Fprintln(w, a...)
}

// Fprintf formats according to a format specifier and writes to w. It returns
// the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, f string, a ...any) (int, error) {
	return fmt.Fprintf(w, searchColors(f, -1, DisableColors, true), a...)
}

// Fprintfn formats according to a format specifier and writes to w with the newline
// at the end. It returns the number of bytes written and any write error encountered.
func Fprintfn(w io.Writer, f string, a ...any) (int, error) {
	return fmt.Fprintf(w, searchColors(f, -1, DisableColors, true)+"\n", a...)
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a string.
func Sprint(a ...any) string {
	applyColors(&a, -1, DisableColors)
	return fmt.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting
// string.
func Sprintf(f string, a ...any) string {
	return fmt.Sprintf(searchColors(f, -1, DisableColors, true), a...)
}

// Sprintfn formats according to a format specifier and returns the resulting
// string with the newline at the end.
func Sprintfn(f string, a ...any) string {
	return fmt.Sprintf(searchColors(f, -1, DisableColors, true)+"\n", a...)
}

// Sprintln formats using the default formats for its operands and returns the
// resulting string. Spaces are always added between operands and a newline is
// appended.
func Sprintln(a ...any) string {
	applyColors(&a, -1, DisableColors)
	return fmt.Sprintln(a...)
}

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
func Errorf(f string, a ...any) error {
	return errors.New(Sprintf(f, a...))
}

// TPrint removes all content on the current line and prints the new message
func TPrint(a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return Print(a...)
}

// TPrintf removes all content on the current line and prints the new message
func TPrintf(f string, a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return Printf(f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// with a new line symbol at the end
func TPrintln(a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return Println(a...)
}

// LPrint formats using the default formats for its operands and writes to standard
// output limited by the text size
func LPrint(maxSize int, a ...any) (int, error) {
	s := fmt.Sprint(a...)
	return fmt.Print(searchColors(s, maxSize, DisableColors, true))
}

// LPrintf formats according to a format specifier and writes to standard output
// limited by the text size
func LPrintf(maxSize int, f string, a ...any) (int, error) {
	s := fmt.Sprintf(f, a...)
	return fmt.Print(searchColors(s, maxSize, DisableColors, true))
}

// LPrintfn formats according to a format specifier and writes to standard output
// limited by the text size and with the newline at the end
func LPrintfn(maxSize int, f string, a ...any) (int, error) {
	s := fmt.Sprintf(f, a...)
	return fmt.Print(searchColors(s, maxSize, DisableColors, true) + "\n")
}

// LPrintln formats using the default formats for its operands and writes to standard
// output limited by the text size
func LPrintln(maxSize int, a ...any) (int, error) {
	applyColors(&a, maxSize, DisableColors)
	return fmt.Println(a...)
}

// TLPrint removes all content on the current line and prints the new message
// limited by the text size
func TLPrint(maxSize int, a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return LPrint(maxSize, a...)
}

// TLPrintf removes all content on the current line and prints the new message
// limited by the text size
func TLPrintf(maxSize int, f string, a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return LPrintf(maxSize, f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// limited by the text size with a new line symbol at the end
func TLPrintln(maxSize int, a ...any) (int, error) {
	fmt.Print(_CODE_CLEAN_LINE)
	return LPrintln(maxSize, a...)
}

// NewLine prints a newline to standard output
func NewLine(num ...int) (int, error) {
	if len(num) == 0 {
		return fmt.Print("\n")
	}

	lineNum := num[0]

	if lineNum <= 1 {
		lineNum = 1
	}

	return fmt.Print(strings.Repeat("\n", lineNum))
}

// Clean returns string without color tags
func Clean(s string) string {
	return searchColors(s, -1, true, true)
}

// Render converts color tags to ANSI escape codes
func Render(s string) string {
	return searchColors(s, -1, false, true)
}

// Bell prints alert (bell) symbol
func Bell() {
	fmt.Print(_CODE_BELL)
}

// IsColorsSupported returns true if 16 colors is supported by terminal
func IsColorsSupported() bool {
	if isColorsSupportChecked {
		return isColorsSupported
	}

	checkForColorsSupport()

	return isColorsSupported
}

// Is256ColorsSupported returns true if 256 colors is supported by terminal
func Is256ColorsSupported() bool {
	if isColorsSupportChecked {
		return isColors256Supported
	}

	checkForColorsSupport()

	return isColors256Supported
}

// IsTrueColorSupported returns true if TrueColor (24-bit colors) is supported by terminal
func IsTrueColorSupported() bool {
	if isColorsSupportChecked {
		return isColorsTCSupported
	}

	checkForColorsSupport()

	return isColorsTCSupported
}

// IsTag tests whether the given value is a valid color tag (or sequence
// of tags) and can be encoded into an escape sequence
func IsTag(tag string) bool {
	if tag == "" {
		return true // Empty value is valid color tag ¯\_(ツ)_/¯
	}

	if !strings.HasPrefix(tag, "{") || !strings.HasSuffix(tag, "}") {
		return false
	}

	return Clean(tag) == ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// tag2ANSI converts color tag to ANSI escape code
func tag2ANSI(tag string, clean bool) string {
	switch {
	case clean:
		return ""
	case isExtendedColorTag(tag):
		return parseExtendedColor(tag)
	case isNamedColorTag(tag):
		return parseNamedColor(tag, clean)
	}

	light := strings.Contains(tag, "-")
	reset := strings.Contains(tag, "!")

	var chars string

	for _, key := range tag {
		switch {
		case key == '*' && !boldDisableEnvVar.Is(""),
			key == '&' && !italicDisableEnvVar.Is(""),
			key == '~' && !blinkDisableEnvVar.Is(""):
			continue
		}

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

		case '*', '&', '^', '_', '=', '~', '@', '+':
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

	return fmt.Sprintf("\033[" + strutil.Substring(chars, 0, -1) + "m")
}

// parseExtendedColor parses extended color tag and returns ANSI code for it
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

// parseNamedColor parses named color tag and returns ANSI code for it
func parseNamedColor(tag string, clean bool) string {
	if colorsMap == nil || clean {
		return ""
	}

	tag = strings.TrimLeft(tag, "?")
	t, ok := colorsMap.Load(tag)

	if !ok {
		return ""
	}

	return t.(string)
}

// getResetCode returns reset code for the given code
func getResetCode(code int) string {
	if code == codes['*'] {
		code++
	}

	return "2" + strconv.Itoa(code)
}

// replaceColorTags reads color tags from input and writes ANSI codes to output buffer
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

// searchColors searches for color tags in the input text and replaces them with
// ANSI escape codes. It limits the number of processed characters to 'limit'.
func searchColors(text string, limit int, clean, close bool) string {
	if text == "" {
		return ""
	}

	closed, counter := true, 0
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

	if !closed && !clean && close {
		output.WriteString(_CODE_RESET)
	}

	return output.String()
}

// applyColors applies color tags to the slice of data
func applyColors(a *[]any, limit int, clean bool) {
	for i, x := range *a {
		if s, ok := x.(string); ok {
			(*a)[i] = searchColors(s, limit, clean, true)
		}
	}
}

// isValidSimpleTag checks if the tag is a valid simple color tag
func isValidSimpleTag(tag string) bool {
	switch {
	case tag == "",
		strings.Trim(tag, "-") == "",
		strings.Count(tag, "!") > 1,
		strings.Contains(tag, "!") && strings.Contains(tag, "-"):
		return false
	}

	for _, r := range tag {
		_, hasCode := codes[r]

		if !hasCode {
			return false
		}
	}

	return true
}

// isExtendedColorTag checks if the tag is an extended color tag
func isExtendedColorTag(tag string) bool {
	switch {
	case len(tag) < 2,
		!strings.HasPrefix(tag, "#") &&
			!strings.HasPrefix(tag, "%"):
		return false
	}

	return true
}

// isValidTag checks if the given data is a valid color tag
func isValidTag(tag string) bool {
	return isValidSimpleTag(tag) || isValidExtendedTag(tag) || isValidNamedTag(tag)
}

// isValidExtendedTag checks if the given data is a valid extended color tag
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

// isNamedColorTag checks if the tag is a named color tag
func isNamedColorTag(tag string) bool {
	return len(tag) >= 2 && strings.HasPrefix(tag, "?")
}

// isValidNamedTag checks if the tag is a valid named color tag
func isValidNamedTag(tag string) bool {
	if !isNamedColorTag(tag) {
		return false
	}

	for _, r := range strings.TrimLeft(tag, "?") {
		switch {
		case r == 95,
			r >= 48 && r <= 57,
			r >= 65 && r <= 90,
			r >= 97 && r <= 122:
			continue
		default:
			return false
		}
	}

	return true
}

// checkForColorsSupport checks if terminal supports colored output
func checkForColorsSupport() {
	term := termEnvVar.Get()
	colorTerm := colorTermEnvVar.Get()

	switch {
	case strings.Contains(term, "xterm"),
		strings.Contains(term, "color"),
		term == "screen":
		isColorsSupported = true
	}

	if strings.Contains(term, "256color") {
		isColors256Supported = true
	}

	if term == "iterm" || colorTerm == "truecolor" ||
		strings.Contains(term, "truecolor") ||
		strings.HasPrefix(term, "vte") {
		isColors256Supported, isColorsTCSupported = true, true
	}

	isColorsSupportChecked = true
}

// ////////////////////////////////////////////////////////////////////////////////// //
