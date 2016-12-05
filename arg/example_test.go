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

func Example_parsing() {
	// Key is argument in format "short-name:long-name" or "long-name"
	// We highly recommend to define arguments names as constants
	argMap := Map{
		"s:string":   {},                                     // By default argument has string type
		"S:string2":  {Type: STRING, Value: "Default value"}, // You can predefine default values
		"int":        {Type: INT},                            // Integer without short name
		"I:int2":     {Type: INT, Min: 1, Max: 10},           // Integer with limits
		"f:float":    {Type: FLOAT, Value: 10.0},             // Float
		"b:boolean":  {Type: BOOL},                           // Boolean
		"r:required": {Type: INT, Required: true},            // Some arguments can be marked as required
		"m:merg":     {Type: STRING, Mergeble: true},         // Mergeble arguments can be defined more than one time
		"h:help":     {Type: BOOL, Alias: "u:usage about"},   // You can define argument aliases
		"e:example":  {Conflicts: "s:string S:string2"},      // Argument conflicts with string and string2 (arguments can't be set at same time)
		"E:example2": {Bound: "int I:int2"},                  // Argument bound with int and int2 (arguments must be set at same time)
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
