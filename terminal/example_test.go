package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRead() {
	userName, err := Read("Please enter user name")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User name: %s\n", userName)
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
	// Add custom error message prefix
	ErrorPrefix = "▲ "

	// Print red text to stderr
	Error("Error while sending data to %s", "https://example.com")

	// Print message from error struct
	err := errors.New("My error")
	Error(err)
}

func ExampleWarn() {
	// Add custom warning message prefix
	WarnPrefix = "△ "

	// Print yellow text to stderr
	Warn("Warning file %s is not found", "/home/john/test.txt")

	// Print message from error struct
	err := errors.New("My warning")
	Warn(err)
}

func ExampleInfo() {
	// Add custom info message prefix
	InfoPrefix = "❕ "

	// Print cyan text to stdout
	Warn("User %q will be created automatically", "bob")
}
