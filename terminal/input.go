package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Prompt is the string displayed before each user input line
var Prompt = "> "

// TitleColorTag is the fmtc color tag applied to input title lines
var TitleColorTag = "{s}"

// AlwaysYes causes [ReadAnswer] to return true without prompting the user,
// useful for non-interactive or forced-action modes
var AlwaysYes = false

// NewLine controls whether an extra blank line is printed before and after each input
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

// InvalidAnswerMesssage is returned when the user's answer is neither Y nor N
var InvalidAnswerMesssage = "Please enter Y or N"

// ////////////////////////////////////////////////////////////////////////////////// //

var dataInput io.Reader = os.Stdin

// ////////////////////////////////////////////////////////////////////////////////// //

// Read prints an optional title and returns the raw text entered by the user
func Read(title string) (string, error) {
	if NewLine {
		fmtc.NewLine()
		defer fmtc.NewLine()
	}

	if title != "" {
		fmtc.Println(strutil.B(TitleColorTag == "", title, TitleColorTag+title+"{!}"))
	}

	r := bufio.NewReader(dataInput)

	fmtc.Print(Prompt)

	input, err := r.ReadString('\n')

	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("can't read user input: %w", err)
	}

	return input, nil
}

// ReadAnswer prompts the user with a yes/no question and returns their answer
// as a bool. An optional defaultAnswer ("Y" or "N") is used when the user submits
// an empty response.
func ReadAnswer(title string, defaultAnswers ...string) (bool, error) {
	var defaultAnswer string

	if len(defaultAnswers) != 0 {
		defaultAnswer = defaultAnswers[0]
	}

	if NewLine {
		fmtc.NewLine()
		defer fmtc.NewLine()
	}

	if title != "" {
		fmtc.Println(TitleColorTag + getAnswerTitle(title, defaultAnswer) + "{!}")
	}

	if AlwaysYes {
		fmtc.Println(Prompt + "{s}Y{!}")

		if NewLine {
			fmtc.NewLine()
		}

		return true, nil
	}

	r := bufio.NewReader(dataInput)

	for {
		fmtc.Print(Prompt)

		answer, err := r.ReadString('\n')

		if err != nil && !errors.Is(err, io.EOF) {
			return false, fmt.Errorf("can't read user input: %w", err)
		}

		answer = strings.TrimSpace(answer)
		answer = strutil.Q(answer, defaultAnswer)

		switch strings.ToUpper(answer) {
		case "Y":
			return true, nil
		case "N":
			return false, nil
		default:
			fmtc.NewLine()
			Warn(InvalidAnswerMesssage)
			fmtc.NewLine()
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

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
