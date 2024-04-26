package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "fmt"

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
