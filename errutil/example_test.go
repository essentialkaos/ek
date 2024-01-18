package errutil

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

func ExampleChain() {
	f1 := func() error { return nil }
	f2 := func() error { return nil }
	f3 := func() error { return fmt.Errorf("Error 3") }
	f4 := func() error { return fmt.Errorf("Error 4") }

	err := Chain(f1, f2, f3, f4)

	fmt.Println(err.Error())

	// Output:
	// Error 3
}

func ExampleErrors() {
	f1 := func() error { return nil }
	f2 := func() error { return nil }
	f3 := func() error { return fmt.Errorf("Error 3") }
	f4 := func() error { return fmt.Errorf("Error 4") }

	// An Errors needs no initialization
	var myErrs Errors

	myErrs.Add(f1())

	// Using NewErrors you can create Errors instance with limited capacity
	errs := NewErrors(10)

	errs.Add(f1())
	errs.Add(f2())
	errs.Add(f3())
	errs.Add(f4())

	fmt.Printf("Last error text: %v\n", errs.Last().Error())
	fmt.Printf("Number of errors: %d\n", errs.Num())
	fmt.Printf("Capacity: %d\n", errs.Cap())
	fmt.Printf("Has errors: %t\n", errs.HasErrors())

	// Output:
	// Last error text: Error 4
	// Number of errors: 2
	// Capacity: 10
	// Has errors: true
}

func ExampleErrors_Add() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("First: %v\n", myErrs.First())
	fmt.Printf("Last: %v\n", myErrs.Last())

	// Output:
	// First: Error 1
	// Last: Error 3
}

func ExampleErrors_First() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("First: %v\n", myErrs.First())

	// Output:
	// First: Error 1
}

func ExampleErrors_Last() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Last: %v\n", myErrs.Last())

	// Output:
	// Last: Error 3
}

func ExampleErrors_Get() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Index 1: %v\n", myErrs.Get(1))
	fmt.Printf("Index 99: %v\n", myErrs.Get(99))

	// Output:
	// Index 1: Error 2
	// Index 99: <nil>
}

func ExampleErrors_All() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors: %v\n", myErrs.All())

	// Output:
	// Errors: [Error 1 Error 2 Error 3]
}

func ExampleErrors_HasErrors() {
	var myErrs Errors

	fmt.Printf("Has errors: %t\n", myErrs.HasErrors())

	myErrs.Add(fmt.Errorf("Error"))

	fmt.Printf("Has errors: %t\n", myErrs.HasErrors())

	// Output:
	// Has errors: false
	// Has errors: true
}

func ExampleErrors_Num() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors num: %d\n", myErrs.Num())

	// Output:
	// Errors num: 3
}

func ExampleErrors_Cap() {
	myErrs := NewErrors(2)

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors cap: %d\n", myErrs.Cap())
	fmt.Printf("First: %v\n", myErrs.First())
	fmt.Printf("Last: %v\n", myErrs.Last())

	// Output:
	// Errors cap: 2
	// First: Error 2
	// Last: Error 3
}

func ExampleErrors_Error() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error 1"))
	myErrs.Add(fmt.Errorf("Error 2"))
	myErrs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors:\n%s\n", myErrs.Error())

	// Output:
	// Errors:
	//   Error 1
	//   Error 2
	//   Error 3
}

func ExampleErrors_Reset() {
	var myErrs Errors

	myErrs.Add(fmt.Errorf("Error"))
	myErrs.Reset()

	fmt.Printf("Has errors: %t\n", myErrs.HasErrors())

	// Output:
	// Has errors: false
}
