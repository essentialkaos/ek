package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleEncodeToFile() {
	var data = make(map[string]int)

	data["john"] = 100
	data["bob"] = 300

	err := EncodeToFile("/path/to/file.json", data, 0600)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func ExampleDecodeFile() {
	var data = make(map[string]int)

	err := DecodeFile("/path/to/file.json", data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
