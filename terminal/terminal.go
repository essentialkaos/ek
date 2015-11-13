// +build linux, darwin, !windows

package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fmtc"
	"fmt"
	"github.com/essentialkaos/go.linenoise"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _PROMPT = "> "

// ////////////////////////////////////////////////////////////////////////////////// //

// KillSignalError is error type when user cancel input
var KillSignalError = linenoise.KillSignalError

// ////////////////////////////////////////////////////////////////////////////////// //

// ReadUI read user input
func ReadUI(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, false)
}

// ReadAnswer read user answer for Y/n question
func ReadAnswer(title, defaultAnswer string) bool {
	for {
		answer, err := ReadUI(title, false)

		if err != nil {
			return false
		}

		if answer == "" {
			answer = defaultAnswer
		}

		answer = strings.ToUpper(answer)

		switch answer {
		case "Y":
			return true
		case "N":
			return false
		default:
			fmtc.Println("\n{y}Please enter Y or N{!}\n")
		}
	}
}

// ReadPassword read password or some private input which will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, true)
}

// PrintErrorMessage print error message
func PrintErrorMessage(message string, args ...interface{}) {
	if len(args) == 0 {
		fmtc.Printf("{r}%s{!}\n", message)
	} else {
		fmtc.Printf("{r}%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// PrintWarnMessage print warning message
func PrintWarnMessage(message string, args ...interface{}) {
	if len(args) == 0 {
		fmtc.Printf("{y}%s{!}\n", message)
	} else {
		fmtc.Printf("{y}%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// PrintActionMessage print message about action currently in progress
func PrintActionMessage(message string) {
	fmtc.Printf("{*}%s:{!} ", message)
}

// PrintActionStatus print message with action execution status
func PrintActionStatus(status int) {
	switch status {
	case 0:
		fmtc.Println("{g}OK{!}")
	case 1:
		fmtc.Println("{r}ERROR{!}")
	}
}

// AddHstory add line to input history
func AddHstory(ui string) {
	linenoise.AddHistory(ui)
}

// SetCompletionHandler add function for autocompletion
func SetCompletionHandler(compfunc func(in string) []string) {
	linenoise.SetCompletionHandler(compfunc)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getPrivateHider(data string) string {
	prefix := ""
	result := ""

	for i := 0; i < len(_PROMPT); i++ {
		prefix += " "
	}

	for i := 0; i < len(data); i++ {
		result += "*"
	}

	return fmt.Sprintf("%s\033[1A%s", prefix, result)
}

func readUserInput(title string, nonEmpty bool, private bool) (string, error) {
	if title != "" {
		fmtc.Printf("{c}%s:{!}\n", title)
	}

	var ui string
	var err error

	for {
		ui, err = linenoise.Line(_PROMPT)

		if err != nil {
			return "", err
		}

		if nonEmpty && ui == "" {
			PrintWarnMessage("\nYou must enter non empty value\n")
			continue
		}

		if private && ui != "" {
			fmt.Println(getPrivateHider(ui))
		}

		break
	}

	return ui, err
}
