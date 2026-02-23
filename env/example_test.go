package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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
	fmt.Printf("String value %s = %s\n", "STR_VALUE", env.Get("STR_VALUE"))
}

func ExampleWhich() {
	echoPath := Which("echo")

	fmt.Printf("Full path to echo binary is %s\n", echoPath)
}

func ExampleVar() {
	v := Var("PATH")

	fmt.Println(v.Get())
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleEnv_Path() {
	env := Get()

	for i, p := range env.Path() {
		fmt.Printf("%d %s\n", i, p)
	}
}

func ExampleEnv_Get() {
	env := Get()

	fmt.Printf("String value %s = %s\n", "STR_VALUE", env.Get("STR_VALUE"))
}

func ExampleEnv_GetI() {
	env := Get()

	fmt.Printf("Integer value %s = %d\n", "INT_VALUE", env.GetI("INT_VALUE"))
}

func ExampleEnv_GetF() {
	env := Get()

	fmt.Printf("Float value %s = %g\n", "FLOAT_VALUE", env.GetF("FLOAT_VALUE"))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleVariable_Get() {
	v := Var("STR_VALUE")

	fmt.Printf("String value %s = %s\n", "STR_VALUE", v.Get())
}

func ExampleVariable_String() {
	v := Var("STR_VALUE")

	fmt.Printf("String value %s = %s\n", "STR_VALUE", v.Get())
}

func ExampleVariable_Is() {
	v := Var("STR_VALUE")

	fmt.Printf("%s == %s: %t\n", "STR_VALUE", "test", v.Is("test"))
}

func ExampleVariable_Reset() {
	v := Var("STR_VALUE")

	fmt.Printf("String value %s = %s\n", "STR_VALUE", v.Get())

	// reset read variable
	v.Reset()
}
