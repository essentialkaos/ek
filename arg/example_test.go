package arg

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleParsing() {
	// Key is argument in format "short-name:long-name" or "long-name"
	// We highly recommend to define arguments names as constants
	argMap := Map{
		"s:string":   &V{},                                     // By default argument has string type
		"S:string2":  &V{Type: STRING, Value: "Default value"}, // You can predefine default values
		"int":        &V{Type: INT},                            // Integer without short name
		"I:int2":     &V{Type: INT, Min: 1, Max: 10},           // Integer with limits
		"f:float":    &V{Type: FLOAT, Value: 10.0},             // Float
		"b:boolean":  &V{Type: BOOL},                           // Boolean
		"r:required": &V{Type: INT, Required: true},            // Some arguments can be marked as required
		"m:merg":     &V{Type: STRING, Mergeble: true},         // Mergeble arguments can be defined more than one time
		"h:help":     &V{Type: BOOL, Alias: "u:usage about"},   // You can define argument aliases
	}

	// args contains unparsed values
	args, errs := Parse(argMap)

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	if Has("s:string") {
		fmt.Println("\"--string/-s\" is set")
	}

	fmt.Printf("Unparsed: %v\n\n", args)

	fmt.Printf("string → %s\n", GetS("string"))
	fmt.Printf("int → %d\n", GetI("int"))
	fmt.Printf("float → %f\n", GetF("f:float"))
	fmt.Printf("boolean → %t\n", GetB("b:boolean"))
}
