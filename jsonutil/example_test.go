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

func ExampleWrite() {
	var data = make(map[string]int)

	data["john"] = 100
	data["bob"] = 300

	err := Write("/path/to/file.json", data, 0600)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func ExampleRead() {
	var data = make(map[string]int)

	err := Read("/path/to/file.json", data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func ExampleWriteGz() {
	var data = make(map[string]int)

	data["john"] = 100
	data["bob"] = 300

	err := WriteGz("/path/to/file.gz", data, 0600)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func ExampleReadGz() {
	var data = make(map[string]int)

	err := ReadGz("/path/to/file.gz", data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
