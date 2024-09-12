package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGet() {
	paramValue, err := Get("kernel.random.boot_id")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Boot ID: %s\n", paramValue)
}

func ExampleGetI() {
	paramValue, err := GetI("kernel.pty.max")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("PTY Max: %d\n", paramValue)
}

func ExampleGetI64() {
	paramValue, err := GetI64("fs.file-max")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File max: %d\n", paramValue)
}
