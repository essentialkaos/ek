package mathutil

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

func ExampleBetween() {
	fmt.Println(Between(10, 1, 5))
	fmt.Println(Between(-3, 1, 5))
	fmt.Println(Between(4, 1, 5))
	// Output:
	// 5
	// 1
	// 4
}

func ExampleMin() {
	fmt.Println(Min(1, 10))
	fmt.Println(Min(3, -3))
	// Output:
	// 1
	// -3
}

func ExampleMax() {
	fmt.Println(Max(1, 10))
	fmt.Println(Max(3, -3))
	// Output:
	// 10
	// 3
}

func ExampleAbs() {
	fmt.Println(Abs(10))
	fmt.Println(Abs(-10))
	// Output:
	// 10
	// 10
}

func ExamplePerc() {
	fmt.Printf("%g%%\n", Perc(180, 600))
	// Output:
	// 30%
}

func ExampleRound() {
	fmt.Println(Round(3.14159, 2))
	// Output:
	// 3.14
}
