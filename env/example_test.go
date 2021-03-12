package env

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

func ExampleGet() {
	env := Get()

	// Print PATH environment variable
	fmt.Println(env["PATH"])

	// Path return PATH variable as slice
	for i, p := range env.Path() {
		fmt.Printf("%d %s\n", i, p)
	}

	// You can use getters for different value formats
	fmt.Printf("Integer value %s = %d\n", "INT_VALUE", env.GetI("INT_VALUE"))
	fmt.Printf("Float value %s = %g\n", "FLOAT_VALUE", env.GetF("FLOAT_VALUE"))
	fmt.Printf("String value %s = %s\n", "STR_VALUE", env.GetS("STR_VALUE"))
}

func ExampleWhich() {
	echoPath := Which("echo")

	fmt.Printf("Full path to echo binary is %s\n", echoPath)
}
