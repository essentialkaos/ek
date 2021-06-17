package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReadUI() {
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

	// User must enter password
	input, err := ReadUI("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User password: %s\v", input)
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

func ExamplePrintErrorMessage() {
	// Print red text to stderr
	PrintErrorMessage("Error while sending data")
}

func ExamplePrintWarnMessage() {
	// Print yellow text to stderr
	PrintWarnMessage("Warning file is not found")
}
