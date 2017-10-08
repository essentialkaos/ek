package procname

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
