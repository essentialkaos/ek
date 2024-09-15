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

func ExampleAll() {
	params, err := All()

	if err != nil {
		panic(err.Error())
	}

	for n, v := range params {
		fmt.Printf("%s → %s\n", n, v)
	}
}

func ExampleParams_Get() {
	params, err := All()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Boot ID: %s\n", params.Get("kernel.random.boot_id"))
}

func ExampleParams_GetI() {
	params, err := All()

	if err != nil {
		panic(err.Error())
	}

	paramValue, err := params.GetI("kernel.pty.max")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("PTY Max: %d\n", paramValue)
}

func ExampleParams_GetI64() {
	params, err := All()

	if err != nil {
		panic(err.Error())
	}

	paramValue, err := params.GetI64("fs.file-max")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File max: %d\n", paramValue)
}
