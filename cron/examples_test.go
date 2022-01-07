package cron

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleParse() {
	expr, err := Parse("0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	if err != nil {
		return
	}

	m1 := time.Date(2020, 1, 1, 18, 15, 0, 0, time.Local)
	m2 := time.Date(2020, 1, 1, 18, 20, 0, 0, time.Local)

	fmt.Printf("Execute1: %t\n", expr.IsDue(m1))
	fmt.Printf("Execute2: %t\n", expr.IsDue(m2))

	// Output:
	// Execute1: true
	// Execute2: false
}

func ExampleExpr_Next() {
	expr, err := Parse("0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	if err != nil {
		return
	}

	m := time.Date(2020, 1, 1, 18, 15, 0, 0, time.UTC)

	fmt.Printf("%v\n", expr.Next(m))

	// Output:
	// 2020-01-01 18:30:00 +0000 UTC
}

func ExampleExpr_Prev() {
	expr, err := Parse("0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	if err != nil {
		return
	}

	m := time.Date(2020, 1, 1, 18, 15, 0, 0, time.UTC)

	fmt.Printf("%v\n", expr.Prev(m))

	// Output:
	// 2020-01-01 18:00:00 +0000 UTC
}

func ExampleExpr_String() {
	expr, err := Parse("0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri")

	if err != nil {
		return
	}

	fmt.Printf("%s\n", expr.String())

	// Output:
	// 0,15,30,45 0,6,12,18 1-10,15,31 * Mon-Fri
}
