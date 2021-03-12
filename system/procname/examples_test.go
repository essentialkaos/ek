package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleSet() {
	args := os.Args

	// Mask first argument
	args[1] = strings.Repeat("*", len(args[1]))

	err := Set(args)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func ExampleReplace() {
	password := "mySuppaPass"

	// Replacing known password to asterisks
	err := Replace(password, "*****************")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
