package errors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleBundle() {
	f1 := func() error { return nil }
	f2 := func() error { return nil }
	f3 := func() error { return fmt.Errorf("Error 3") }
	f4 := func() error { return fmt.Errorf("Error 4") }

	// An Bundle needs no initialization
	var myErrs Bundle

	myErrs.Add(f1())

	// Using NewBundle you can create Bundle instance with limited capacity
	errs := NewBundle(10)

	errs.Add(f1())
	errs.Add(f2())
	errs.Add(f3())
	errs.Add(f4())

	fmt.Printf("Last error text: %v\n", errs.Last().Error())
	fmt.Printf("Number of errors: %d\n", errs.Num())
	fmt.Printf("Capacity: %d\n", errs.Cap())
	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	// Output:
	// Last error text: Error 4
	// Number of errors: 2
	// Capacity: 10
	// Has errors: true
}

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

func ExampleToBundle() {
	e := []error{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	errs := ToBundle(e)

	fmt.Printf("Last error text: %v\n", errs.Last().Error())
	fmt.Printf("Number of errors: %d\n", errs.Num())
	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	// Output:
	// Last error text: Error 3
	// Number of errors: 3
	// Has errors: true
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleErrors_First() {
	errs := Errors{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	fmt.Printf("First: %v\n", errs.First())

	// Output:
	// First: Error 1
}

func ExampleErrors_Last() {
	errs := Errors{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	fmt.Printf("Last: %v\n", errs.Last())

	// Output:
	// Last: Error 3
}

func ExampleErrors_Get() {
	errs := Errors{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	fmt.Printf("Index 1: %v\n", errs.Get(1))
	fmt.Printf("Index 99: %v\n", errs.Get(99))

	// Output:
	// Index 1: Error 2
	// Index 99: <nil>
}

func ExampleErrors_IsEmpty() {
	var errs Errors

	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	errs = Errors{fmt.Errorf("Error 1")}

	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	// Output:
	// Has errors: false
	// Has errors: true
}

func ExampleErrors_Num() {
	errs := Errors{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	fmt.Printf("Errors num: %d\n", errs.Num())

	// Output:
	// Errors num: 3
}

func ExampleErrors_Error() {
	errs := Errors{
		fmt.Errorf("Error 1"),
		fmt.Errorf("Error 2"),
		fmt.Errorf("Error 3"),
	}

	fmt.Printf("Errors:\n%s\n", errs.Error(" - "))

	// Output:
	// Errors:
	//  - Error 1
	//  - Error 2
	//  - Error 3
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleBundle_Add() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("First: %v\n", errs.First())
	fmt.Printf("Last: %v\n", errs.Last())

	// Output:
	// First: Error 1
	// Last: Error 3
}

func ExampleBundle_First() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("First: %v\n", errs.First())

	// Output:
	// First: Error 1
}

func ExampleBundle_Last() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Last: %v\n", errs.Last())

	// Output:
	// Last: Error 3
}

func ExampleBundle_Get() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Index 1: %v\n", errs.Get(1))
	fmt.Printf("Index 99: %v\n", errs.Get(99))

	// Output:
	// Index 1: Error 2
	// Index 99: <nil>
}

func ExampleBundle_All() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors: %v\n", errs.All())

	// Output:
	// Errors: [Error 1 Error 2 Error 3]
}

func ExampleBundle_IsEmpty() {
	var errs Bundle

	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	errs.Add(fmt.Errorf("Error"))

	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	// Output:
	// Has errors: false
	// Has errors: true
}

func ExampleBundle_Num() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors num: %d\n", errs.Num())

	// Output:
	// Errors num: 3
}

func ExampleBundle_Cap() {
	errs := NewBundle(2)

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors cap: %d\n", errs.Cap())
	fmt.Printf("First: %v\n", errs.First())
	fmt.Printf("Last: %v\n", errs.Last())

	// Output:
	// Errors cap: 2
	// First: Error 2
	// Last: Error 3
}

func ExampleBundle_Error() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error 1"))
	errs.Add(fmt.Errorf("Error 2"))
	errs.Add(fmt.Errorf("Error 3"))

	fmt.Printf("Errors:\n%s\n", errs.Error(" - "))

	// Output:
	// Errors:
	//  - Error 1
	//  - Error 2
	//  - Error 3
}

func ExampleBundle_Reset() {
	var errs Bundle

	errs.Add(fmt.Errorf("Error"))
	errs.Reset()

	fmt.Printf("Has errors: %t\n", !errs.IsEmpty())

	// Output:
	// Has errors: false
}
