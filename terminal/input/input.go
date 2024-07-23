//go:build !windows
// +build !windows

// Package input provides methods for reading user input
package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/essentialkaos/ek/v13/ansi"
	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/secstr"
	"github.com/essentialkaos/ek/v13/terminal"
	"github.com/essentialkaos/ek/v13/terminal/tty"

	linenoise "github.com/essentialkaos/go-linenoise/v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CompletionHandler is completion handler
type CompletionHandler = func(input string) []string

// HintHandler is hint handler
type HintHandler = func(input string) string

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = linenoise.ErrKillSignal

// ////////////////////////////////////////////////////////////////////////////////// //

// Prompt is a prompt string
var Prompt = "> "

// MaskSymbol is a symbol used to maks passwords
var MaskSymbol = "*"

// HideLength is a flag to hide the password length
var HideLength = false

// HidePassword is a flag to hide the password while typing.
// Because we are using the low-level linenoise method for this feature, we cannot
// use a custom masking symbol, so it will always be an asterisk (*).
var HidePassword = false

// MaskSymbolColorTag is an fmtc color tag used for MaskSymbol output
var MaskSymbolColorTag = ""

// TitleColorTag is an fmtc color tag used for input titles
var TitleColorTag = "{s}"

// AlwaysYes is a flag, if set ReadAnswer will always return true (useful for working
// with option for forced actions)
var AlwaysYes = false

// NewLine is a flag for extra new line after inputs
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

var oldTMUXFlag int8

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads user input
func Read(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, false)
}

// ReadAnswer reads user's answer to yes/no question
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	var defaultAnswer string

	if len(defaultAnswers) != 0 {
		defaultAnswer = defaultAnswers[0]
	}

	if AlwaysYes {
		if title != "" {
			fmtc.Println(TitleColorTag + getAnswerTitle(title, defaultAnswer) + "{!}")
		}

		fmtc.Println(Prompt + "y")

		if NewLine {
			fmtc.NewLine()
		}

		return true, nil
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
			terminal.Warn("Please enter Y or N")
			fmtc.NewLine()
		}
	}
}

// ReadPassword reads password or some private input that will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, true)
}

// ReadPasswordSecure reads password or some private input that will be hidden
// after pressing Enter
func ReadPasswordSecure(title string, nonEmpty bool) (*secstr.String, error) {
	password, err := readUserInput(title, nonEmpty, true)

	if err != nil {
		return nil, err
	}

	return secstr.NewSecureString(&password)
}

// AddHistory adds line to input history
func AddHistory(data string) {
	linenoise.AddHistory(data)
}

// SetHistoryCapacity sets maximum capacity of history
func SetHistoryCapacity(capacity int) error {
	return linenoise.SetHistoryCapacity(capacity)
}

// SetCompletionHandler adds autocompletion function (using Tab key)
func SetCompletionHandler(h CompletionHandler) {
	linenoise.SetCompletionHandler(h)
}

// SetHintHandler adds function for input hints
func SetHintHandler(h HintHandler) {
	linenoise.SetHintHandler(h)
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
		// Check for old versions of TMUX with rendering problems
		if isOldTMUXSession() {
			masking = strings.Repeat("*", length)
		} else {
			masking = strings.Repeat(MaskSymbol, length)
		}
	} else {
		masking = "[hidden]" + strings.Repeat(" ", mathutil.Max(0, length-8))
	}

	if !tty.IsTTY() {
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
			terminal.Warn("\nYou must enter non-empty value\n")
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
			if !tty.IsTTY() {
				fmt.Println(fmtc.Sprintf(Prompt) + input)
			}
		}

		break
	}

	if NewLine {
		fmtc.NewLine()
	}

	return input, err
}

// isOldTMUXSession returns true if we work in tmux session and version is older than 3.x
func isOldTMUXSession() bool {
	if oldTMUXFlag == 0 {
		if os.Getenv("TMUX") == "" {
			oldTMUXFlag = -1
		} else {
			version := os.Getenv("TERM_PROGRAM_VERSION")

			if len(version) > 0 && (version[0] == '1' || version[0] == '2') {
				oldTMUXFlag = 1
			} else {
				oldTMUXFlag = -1
			}
		}
	}

	return oldTMUXFlag == 1
}
