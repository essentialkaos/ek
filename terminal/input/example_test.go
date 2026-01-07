package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
	userName, err := Read("Please enter user name", NotEmpty)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User name: %s\n", userName)

	// You can read user input without providing any title
	fmt.Println("Please enter user name")
	userName, err = Read("")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User name: %s\n", userName)

	// You can define many validators at once
	userEmail, err := Read("Please enter user email", NotEmpty, IsEmail)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User email: %s\n", userEmail)
}

func ExampleReadPassword() {
	Prompt = "› "
	MaskSymbol = "•"
	MaskSymbolColorTag = "{s}"
	NewLine = true

	// User must enter the password
	password, err := ReadPassword("Please enter password", NotEmpty)

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
	password, err := ReadPasswordSecure("Please enter password", NotEmpty)

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
	input, err := Read("Please enter user name", NotEmpty)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Save entered value to the input history
	AddHistory(input)

	fmt.Printf("User name: %s\n", input)
}

func ExampleSetHistoryCapacity() {
	input, err := Read("Please enter user name", NotEmpty)

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

	SetCompletionHandler(func(input string) []string {
		var result []string

		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				result = append(result, c)
			}
		}

		return result
	})

	SetHintHandler(func(input string) string {
		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				return c[len(input):]
			}
		}

		return ""
	})

	input, err := Read("Please enter command", NotEmpty)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Command: %s\n", input)
}

func ExampleSetHintHandler() {
	commands := []string{"add", "delete", "search", "help", "quit"}

	SetCompletionHandler(func(input string) []string {
		var result []string

		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				result = append(result, c)
			}
		}

		return result
	})

	SetHintHandler(func(input string) string {
		for _, c := range commands {
			if strings.HasPrefix(c, input) {
				return c[len(input):]
			}
		}

		return ""
	})

	input, err := Read("Please enter command", NotEmpty)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Command: %s\n", input)
}
