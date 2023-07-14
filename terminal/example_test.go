package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRead() {
	// User must enter name
	input, err := ReadUI("Please enter user name", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User name: %s\v", input)

	// You can read user input without providing any title
	fmt.Println("Please enter user name")
	input, err = ReadUI("", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User name: %s\v", input)
}

func ExampleReadPassword() {
	Prompt = "› "
	MaskSymbol = "•"
	MaskSymbolColorTag = "{s}"

	// User must enter the password
	password, err := ReadPassword("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User password: %s\v", password)
}

func ExampleReadPasswordSecure() {
	Prompt = "› "
	MaskSymbol = "•"
	MaskSymbolColorTag = "{s}"

	// User must enter the password
	password, err := ReadPasswordSecure("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User password: %s\v", string(password.Data))

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

func ExamplePrintActionMessage() {
	statusOk := true

	PrintActionMessage("Starting service my-service")

	switch statusOk {
	case true:
		PrintActionStatus(0) // Print OK
	case false:
		PrintActionStatus(1) // Print ERROR
	}
}

func ExamplePrintActionStatus() {
	statusOk := true

	PrintActionMessage("Starting service my-service")

	switch statusOk {
	case true:
		PrintActionStatus(0) // Print OK
	case false:
		PrintActionStatus(1) // Print ERROR
	}
}

func ExampleError() {
	// Print red text to stderr
	Error("Error while sending data to %s", "https://example.com")
}

func ExampleWarn() {
	// Print yellow text to stderr
	Warn("Warning file %s is not found", "/home/john/test.txt")
}

func ExampleInfo() {
	// Print cyan text to stdout
	Warn("User %q will be created automatically", "bob")
}

func ExampleErrorPanel() {
	ErrorPanel(
		"Can't send data to remote server.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	)
}

func ExampleWarnPanel() {
	WarnPanel(
		"Can't find user bob on system.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	)
}

func ExampleInfoPanel() {
	InfoPanel(
		"Auto-saving is enabled - data will be saved after editing.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	)
}

func ExamplePanel() {
	Panel(
		"Yolo", "{m}",
		"Excepteur sint occaecat cupidatat non proident.",
		`{*}Lorem ipsum{!} dolor sit amet, {r*}consectetur{!} adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
	)
}
