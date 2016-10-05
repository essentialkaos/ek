package env

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

func Example_Get() {
	env := Get()

	// Print PATH environment variable
	fmt.Println(env["PATH"])

	// Path return PATH variable as slice
	for i, p := range env.Path() {
		fmt.Printf("%d %s\n", i, p)
	}

	// You can use getters for different value formats
	fmt.Printf("Integer value %s = %d\n", "INT_VALUE", env.GetI("INT_VALUE"))
	fmt.Printf("Float value %s = %f\n", "FLOAT_VALUE", env.GetI("FLOAT_VALUE"))
	fmt.Printf("String value %s = %d\n", "STR_VALUE", env.GetS("STR_VALUE"))
}

func Example_Which() {
	echoPath := Which("echo")

	fmt.Printf("Full path to echo binary is %s\n", echoPath)

	// Output:
	// Full path to echo binary is /bin/echo
}
