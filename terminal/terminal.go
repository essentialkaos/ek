//go:build !windows
// +build !windows

// Package terminal provides methods for working with user input
package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/essentialkaos/ek/v12/ansi"
	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/secstr"
	"github.com/essentialkaos/ek/v12/sliceutil"

	"github.com/essentialkaos/go-linenoise/v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type PanelOption uint8

const (
	// PANEL_WRAP is panel rendering option for automatic text wrapping
	PANEL_WRAP PanelOption = iota + 1

	// PANEL_INDENT is panel rendering option for indent using new lines
	// before and after panel
	PANEL_INDENT

	// PANEL_BOTTOM_LINE is panel rendering option for drawing bottom line of panel
	PANEL_BOTTOM_LINE

	// PANEL_LABEL_POWERLINE is panel rendering option for using powerline symbols
	PANEL_LABEL_POWERLINE
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = linenoise.ErrKillSignal

// Prompt is prompt string
var Prompt = "> "

// MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// HideLength is flag for hiding password length
var HideLength = false

// HidePassword is flag for hiding password while typing
// Because of using the low-level linenoise method for this feature, we can not use a
// custom masking symbol, so it always will be an asterisk (*).
var HidePassword = false

var (
	// MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
	MaskSymbolColorTag = ""

	// TitleColorTag is fmtc color tag used for input titles
	TitleColorTag = "{s}"

	// ErrorColorTag is fmtc color tag used for error messages
	ErrorColorTag = "{r}"

	// WarnColorTag is fmtc color tag used for warning messages
	WarnColorTag = "{y}"

	// InfoColorTag is fmtc color tag used for info messages
	InfoColorTag = "{c-}"
)

var (
	// ErrorPrefix is prefix for error messages
	ErrorPrefix = ""

	// WarnPrefix is prefix for warning messages
	WarnPrefix = ""

	// InfoPrefix is prefix for info messages
	InfoPrefix = ""
)

// PanelWidth is panel width (≥ 40)
var PanelWidth = 88

// ////////////////////////////////////////////////////////////////////////////////// //

var tmux int8

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads user's input
func Read(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, false)
}

// ReadAnswer reads user's answer for yes/no question
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	var defaultAnswer string

	if len(defaultAnswers) != 0 {
		defaultAnswer = defaultAnswers[0]
	}

	for {
		answer, err := readUserInput(
			getAnswerTitle(title, defaultAnswer), false, false,
		)

		if err != nil {
			return false, err
		}

		if answer == "" {
			answer = defaultAnswer
		}

		switch strings.ToUpper(answer) {
		case "Y":
			return true, nil
		case "N":
			return false, nil
		default:
			PrintWarnMessage("\nPlease enter Y or N\n")
		}
	}
}

// ReadPassword reads password or some private input which will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, true)
}

// ReadPasswordSecure reads password or some private input which will be hidden
// after pressing Enter
func ReadPasswordSecure(title string, nonEmpty bool) (*secstr.String, error) {
	password, err := readUserInput(title, nonEmpty, true)

	if err != nil {
		return nil, err
	}

	return secstr.NewSecureString(&password)
}

// PrintActionMessage prints message about action currently in progress
func PrintActionMessage(message string) {
	fmtc.Printf("{*}%s:{!} ", message)
}

// PrintActionStatus prints message with action execution status
func PrintActionStatus(status int) {
	switch status {
	case 0:
		fmtc.Println("{g}OK{!}")
	case 1:
		fmtc.Println("{r}ERROR{!}")
	}
}

// Error prints error message
func Error(message string, args ...any) {
	if len(args) == 0 {
		fmtc.Fprintf(os.Stderr, ErrorPrefix+ErrorColorTag+"%s{!}\n", message)
	} else {
		fmtc.Fprintf(os.Stderr, ErrorPrefix+ErrorColorTag+"%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// Warn prints warning message
func Warn(message string, args ...any) {
	if len(args) == 0 {
		fmtc.Fprintf(os.Stderr, WarnPrefix+WarnColorTag+"%s{!}\n", message)
	} else {
		fmtc.Fprintf(os.Stderr, WarnPrefix+WarnColorTag+"%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// Info prints info message
func Info(message string, args ...any) {
	if len(args) == 0 {
		fmtc.Fprintf(os.Stdout, InfoPrefix+InfoColorTag+"%s{!}\n", message)
	} else {
		fmtc.Fprintf(os.Stdout, InfoPrefix+InfoColorTag+"%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// ErrorPanel shows panel with error message
func ErrorPanel(title, message string, options ...PanelOption) {
	Panel("ERROR", ErrorColorTag, title, message, options...)
}

// WarnPanel shows panel with warning message
func WarnPanel(title, message string, options ...PanelOption) {
	Panel("WARNING", WarnColorTag, title, message, options...)
}

// InfoPanel shows panel with warning message
func InfoPanel(title, message string, options ...PanelOption) {
	Panel("INFO", InfoColorTag, title, message, options...)
}

// Panel show panel with given label, title, and message
func Panel(label, colorTag, title, message string, options ...PanelOption) {
	var buf *bytes.Buffer

	if sliceutil.Contains(options, PANEL_INDENT) {
		fmtc.NewLine()
	}

	if sliceutil.Contains(options, PANEL_LABEL_POWERLINE) {
		fmtc.Printf(colorTag+"{@*} %s {!}"+colorTag+"{!} "+colorTag+"%s{!}\n", label, title)
	} else {
		fmtc.Printf(colorTag+"{@*} %s {!} "+colorTag+"%s{!}\n", label, title)
	}

	switch {
	case sliceutil.Contains(options, PANEL_WRAP):
		buf = bytes.NewBufferString(
			fmtutil.Wrap(fmtc.Sprint(message), "", mathutil.Max(38, PanelWidth-2)) + "\n",
		)
	default:
		buf = bytes.NewBufferString(message)
		buf.WriteRune('\n')
	}

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		fmtc.Print(colorTag + "┃{!} " + line)
	}

	if sliceutil.Contains(options, PANEL_BOTTOM_LINE) {
		fmtc.Println(colorTag + "┖" + strings.Repeat("─", mathutil.Max(38, PanelWidth-2)))
	}

	if sliceutil.Contains(options, PANEL_INDENT) {
		fmtc.NewLine()
	}
}

// AddHistory adds line to input history
func AddHistory(data string) {
	linenoise.AddHistory(data)
}

// SetCompletionHandler adds function for autocompletion
func SetCompletionHandler(h func(input string) []string) {
	linenoise.SetCompletionHandler(h)
}

// SetHintHandler adds function for input hints
func SetHintHandler(h func(input string) string) {
	linenoise.SetHintHandler(h)
}

// DEPRECATED /////////////////////////////////////////////////////////////////////// //

// ReadUI reads user's input
//
// Deprecated: Use method Read instead
func ReadUI(title string, nonEmpty bool) (string, error) {
	return Read(title, nonEmpty)
}

// PrintErrorMessage prints error message
//
// Deprecated: Use method Error instead
func PrintErrorMessage(message string, args ...any) {
	Error(message, args...)
}

// PrintWarnMessage prints warning message
//
// Deprecated: Use method Warn instead
func PrintWarnMessage(message string, args ...any) {
	Warn(message, args...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getMask returns mask for password
func getMask(message string) string {
	var masking string

	// Remove fmtc color tags and ANSI escape codes
	prompt := fmtc.Clean(ansi.RemoveCodes(Prompt))
	prefix := strings.Repeat(" ", utf8.RuneCountInString(prompt))
	length := utf8.RuneCountInString(message)

	if !HideLength {
		if isTmuxSession() {
			masking = strings.Repeat("*", length)
		} else {
			masking = strings.Repeat(MaskSymbol, length)
		}
	} else {
		masking = "[hidden]" + strings.Repeat(" ", mathutil.Max(0, length-8))
	}

	if !fsutil.IsCharacterDevice("/dev/stdin") && os.Getenv("FAKETTY") == "" {
		return fmtc.Sprintf(Prompt) + masking
	}

	return fmt.Sprintf("%s\033[1A%s", prefix, masking)
}

// getAnswerTitle returns title with info about default answer
func getAnswerTitle(title, defaultAnswer string) string {
	if title == "" {
		return ""
	}

	switch strings.ToUpper(defaultAnswer) {
	case "Y":
		return fmt.Sprintf(TitleColorTag+"%s ({*}Y{!*}/n){!}", title)
	case "N":
		return fmt.Sprintf(TitleColorTag+"%s (y/{*}N{!*}){!}", title)
	default:
		return fmt.Sprintf(TitleColorTag+"%s (y/n){!}", title)
	}
}

// readUserInput reads user input
func readUserInput(title string, nonEmpty, private bool) (string, error) {
	if title != "" {
		fmtc.Println(TitleColorTag + title + "{!}")
	}

	var input string
	var err error

	if private && HidePassword {
		linenoise.SetMaskMode(true)
	}

	for {
		input, err = linenoise.Line(fmtc.Sprintf(Prompt))

		if private && HidePassword {
			linenoise.SetMaskMode(false)
		}

		if err != nil {
			return "", err
		}

		if nonEmpty && strings.TrimSpace(input) == "" {
			PrintWarnMessage("\nYou must enter non-empty value\n")
			continue
		}

		if private && input != "" {
			if !HidePassword {
				if MaskSymbolColorTag == "" {
					fmt.Println(getMask(input))
				} else {
					fmtc.Println(MaskSymbolColorTag + getMask(input) + "{!}")
				}
			}
		} else {
			if !fsutil.IsCharacterDevice("/dev/stdin") && os.Getenv("FAKETTY") == "" {
				fmt.Println(fmtc.Sprintf(Prompt) + input)
			}
		}

		break
	}

	return input, err
}

// isTmuxSession returns true if we work in tmux session
func isTmuxSession() bool {
	if tmux == 0 {
		if os.Getenv("TMUX") == "" {
			tmux = -1
		} else {
			tmux = 1
		}
	}

	return tmux == 1
}
