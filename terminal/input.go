package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Prompt is a prompt string
var Prompt = "> "

// TitleColorTag is an fmtc color tag used for input titles
var TitleColorTag = "{s}"

// AlwaysYes is a flag, if set ReadAnswer will always return true (useful for working
// with option for forced actions)
var AlwaysYes = false

// NewLine is a flag for extra new line after inputs
var NewLine = false

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrInvalidAnswer is error for wrong answer for yes/no question
var ErrInvalidAnswer = fmt.Errorf("Please enter Y or N")

// ////////////////////////////////////////////////////////////////////////////////// //

var dataInput io.Reader = os.Stdin

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads user input
func Read(title string) string {
	if NewLine {
		fmtc.NewLine()
		defer fmtc.NewLine()
	}

	if title != "" {
		fmtc.Println(strutil.B(TitleColorTag == "", title, TitleColorTag+title+"{!}"))
	}

	r := bufio.NewReader(dataInput)

	fmtc.Print(Prompt)

	input, _ := r.ReadString('\n')

	return input
}

// ReadAnswer reads user's answer to yes/no question
func ReadAnswer(title string, defaultAnswers ...string) bool {
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

		return true
	}

	r := bufio.NewReader(dataInput)

	for {
		fmtc.Print(Prompt)

		answer, _ := r.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strutil.Q(answer, defaultAnswer)

		switch strings.ToUpper(answer) {
		case "Y":
			return true
		case "N":
			return false
		default:
			fmtc.NewLine()
			Warn(ErrInvalidAnswer)
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
