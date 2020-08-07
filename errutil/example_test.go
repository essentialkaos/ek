package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
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

	errs := NewErrors()

	errs.Add(f1())
	errs.Add(f2())
	errs.Add(f3())
	errs.Add(f4())

	fmt.Printf("Last error text: %v\n", errs.Last().Error())
	fmt.Printf("Number of errors: %d\n", errs.Num())
	fmt.Printf("Has errors: %t\n", errs.HasErrors())

	// Output:
	// Last error text: Error 4
	// Number of errors: 2
	// Has errors: true
}
