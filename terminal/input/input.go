//go:build !windows

// Package input provides methods for reading user input
package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
	"github.com/essentialkaos/ek/v13/secstr"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/terminal"
	"github.com/essentialkaos/ek/v13/terminal/tty"

	linenoise "github.com/essentialkaos/go-linenoise/v3"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CompletionHandler is the function signature for Tab-completion callbacks
type CompletionHandler = func(input string) []string

// HintHandler is the function signature for inline input hint callbacks
type HintHandler = func(input string) string

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = linenoise.ErrKillSignal

// ////////////////////////////////////////////////////////////////////////////////// //

// Prompt is the string displayed before each user input line
var Prompt = "> "

// MaskSymbol is the character used to replace each typed character when masking
// passwords
var MaskSymbol = "*"

// HideLength conceals the password length by replacing the mask with a fixed
// "[hidden]" label
var HideLength = false

// HidePassword hides typed characters in real time using linenoise's built-in
// mask mode. When enabled, the masking symbol is always an asterisk regardless
// of [MaskSymbol].
var HidePassword = false

// MaskSymbolColorTag is the fmtc color tag applied when rendering the mask symbol
var MaskSymbolColorTag = ""

// TitleColorTag is the fmtc color tag applied to input title lines
var TitleColorTag = "{s}"

// AlwaysYes causes [ReadAnswer] to return true without prompting the user,
// useful for non-interactive or forced-action modes
var AlwaysYes = false

// NewLine controls whether an extra blank line is printed before and after each input
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

// InvalidAnswerMessage is returned when the user's answer is neither Y nor N
var InvalidAnswerMessage = "Please enter Y or N"

// ////////////////////////////////////////////////////////////////////////////////// //

var oldTMUXFlag int8

var isStdinSet = hasStdinData()

// ////////////////////////////////////////////////////////////////////////////////// //

// Read displays an optional title, prompts for a line of text, and applies
// any provided validators before returning the input
func Read(title string, validators ...Validator) (string, error) {
	return readUserInput(title, false, validators)
}

// ReadAnswer prompts the user with a yes/no question and returns their answer
// as a bool. An optional defaultAnswer ("Y" or "N") is used when the user submits an
// empty response.
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	var defaultAnswer string

	if len(defaultAnswers) != 0 {
		defaultAnswer = defaultAnswers[0]
	}

	if AlwaysYes {
		if title != "" {
			fmtc.Println(TitleColorTag + getAnswerTitle(title, defaultAnswer) + "{!}")
		}

		fmtc.Println(Prompt + "{s}Y{!}")

		if NewLine {
			fmtc.NewLine()
		}

		return true, nil
	}

	for {
		answer, err := readUserInput(
			getAnswerTitle(title, defaultAnswer), false, nil,
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
			terminal.Warn(InvalidAnswerMessage)
			fmtc.NewLine()
		}
	}
}

// ReadPassword prompts for a line of text that is masked on screen after entry
func ReadPassword(title string, validators ...Validator) (string, error) {
	return readUserInput(title, true, validators)
}

// ReadPasswordSecure works like ReadPassword but returns the value wrapped in a
// secstr.String to prevent the secret from lingering in plain memory
func ReadPasswordSecure(title string, validators ...Validator) (*secstr.String, error) {
	password, err := readUserInput(title, true, validators)

	if err != nil {
		return nil, err
	}

	return secstr.NewSecureString(&password)
}

// AddHistory appends a line to the linenoise input history
func AddHistory(data string) {
	linenoise.AddHistory(data)
}

// SetHistoryCapacity sets the maximum number of entries retained in input history.
// Returns an error if capacity is less than 1.
func SetHistoryCapacity(capacity int) error {
	return linenoise.SetHistoryCapacity(capacity)
}

// SetCompletionHandler registers a callback invoked on Tab key press to provide
// completions
func SetCompletionHandler(h CompletionHandler) {
	linenoise.SetCompletionHandler(h)
}

// SetHintHandler registers a callback that provides an inline hint shown while the
// user types
func SetHintHandler(h HintHandler) {
	linenoise.SetHintHandler(h)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getMask returns mask for password
func getMask(message string) string {
	var masking string

	// Remove fmtc color tags and ANSI escape codes
	prompt := fmtc.Clean(ansi.Remove(Prompt))
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
		masking = "[hidden]" + strings.Repeat(" ", max(0, length-8))
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
		return fmt.Sprintf(TitleColorTag+"%s{!} {s}(Y/n){!}", title)
	case "N":
		return fmt.Sprintf(TitleColorTag+"%s{!} {s}(y/N){!}", title)
	default:
		return fmt.Sprintf(TitleColorTag+"%s{!} {s}(y/n){!}", title)
	}
}

// readUserInput reads user input
func readUserInput(title string, private bool, validators []Validator) (string, error) {
	if title != "" {
		fmtc.Println(strutil.B(TitleColorTag == "", title, TitleColorTag+title+"{!}"))
	}

	var input string
	var err error

	if NewLine {
		defer fmtc.NewLine()
	}

	if private && HidePassword {
		linenoise.SetMaskMode(true)
		defer linenoise.SetMaskMode(false)
	}

INPUT_LOOP:
	for {
		input, err = linenoise.Line(fmtc.Sprintf(Prompt))

		if err != nil {
			return "", err
		}

		if isStdinSet {
			fmtc.Println(fmtc.Sprintf(Prompt))
			isStdinSet = false
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

		if len(validators) != 0 {
			for _, validator := range validators {
				input, err = validator.Validate(input)

				if err != nil {
					fmtc.NewLine()
					terminal.Warn(err.Error())
					fmtc.NewLine()
					continue INPUT_LOOP
				}
			}
		}

		break
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

// hasStdinData returns true if there is some data in stdin
func hasStdinData() bool {
	stdin, err := os.Stdin.Stat()

	if err != nil {
		return false
	}

	if stdin.Mode()&os.ModeCharDevice != 0 {
		return false
	}

	return true
}
