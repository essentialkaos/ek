package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReadUI() {

	// user must enter name
	input, err := ReadUI("Please enter user name", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User name: %s\v", input)
}

func ExampleReadPassword() {

	// user must enter password
	input, err := ReadUI("Please enter password", true)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("User password: %s\v", input)
}

func ExampleReadAnswer() {

	// is user doesn't enter any value, we use default value (Y in this case)
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
	// print this text with red color
	PrintErrorMessage("Error while sending data")
}

func ExamplePrintWarnMessage() {
	// print this text with yellow color
	PrintWarnMessage("Warning file is not found")
}
