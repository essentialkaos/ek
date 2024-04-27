package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRead() {
	// User must enter name
	input, err := Read("Please enter user name", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User name: %s\n", input)

	// You can read user input without providing any title
	fmt.Println("Please enter user name")
	input, err = Read("", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User name: %s\n", input)
}

func ExampleReadPassword() {
	Prompt = "› "
	MaskSymbol = "•"
	MaskSymbolColorTag = "{s}"

	// User must enter the password
	password, err := ReadPassword("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User password: %s\n", password)
}

func ExampleReadPasswordSecure() {
	Prompt = "› "
	MaskSymbol = "•"
	MaskSymbolColorTag = "{s}"

	// User must enter the password
	password, err := ReadPasswordSecure("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User password: %s\n", string(password.Data))

	password.Destroy()
}

func ExampleReadAnswer() {
	// If the user doesn't enter any value, we will use the default
	// value (Y in this case)
	ok, err := ReadAnswer("Remove this file?", "Y")

	if !ok || err != nil {
		return
	}

	if ok {
		fmt.Println("File removed")
	}
}

func ExampleAddHistory() {
	input, err := Read("Please enter user name", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Save entered value to the input history
	AddHistory(input)

	fmt.Printf("User name: %s\n", input)
}

func ExampleSetHistoryCapacity() {
	input, err := Read("Please enter user name", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Limit history size to last 3 entries
	SetHistoryCapacity(3)

	// Save entered value to the input history
	AddHistory(input)

	fmt.Printf("User name: %s\n", input)
}

func ExampleSetCompletionHandler() {
	commands := []string{"add", "delete", "search", "help", "quit"}

	input.SetCompletionHandler(func(input string) []string {
		var result []string

		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				result = append(result, c)
			}
		}

		return result
	})

	input.SetHintHandler(func(input string) string {
		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				return c[len(input):]
			}
		}

		return ""
	})

	input, err := input.Read("Please enter command", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Command: %s\n", input)
}

func ExampleSetHintHandler() {
	commands := []string{"add", "delete", "search", "help", "quit"}

	input.SetCompletionHandler(func(input string) []string {
		var result []string

		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				result = append(result, c)
			}
		}

		return result
	})

	input.SetHintHandler(func(input string) string {
		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				return c[len(input):]
			}
		}

		return ""
	})

	input, err := input.Read("Please enter command", true)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Command: %s\n", input)
}
