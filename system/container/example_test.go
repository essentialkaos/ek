package container

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

func ExampleGetEngine() {
	fmt.Printf("Container engine: %s\n", GetEngine())
}

func ExampleIsContainer() {
	if !IsContainer() {
		fmt.Println("It's not containerized")
	} else {
		fmt.Printf("It's containerized using %s engine\n", GetEngine())
	}
}
