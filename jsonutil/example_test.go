package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
