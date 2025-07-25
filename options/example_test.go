package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewOptions() {
	opts := NewOptions()

	// Add options
	opts.Add("u:user", &V{Type: STRING, Value: "john"})
	opts.Add("l:lines", &V{Type: INT, Min: 1, Max: 100})

	// args contains unparsed values
	input := "-u bob -l 12 file.txt"
	args, errs := opts.Parse(strings.Split(input, " "))

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

func ExampleNewArguments() {
	args := NewArguments("head", "file.txt", "10")

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Command: %s\n", args.Get(0).String())
	fmt.Printf("File: %s\n", args.Get(1).String())

	lines, _ := args.Get(2).Int()
	fmt.Printf("Lines: %d\n", lines)
	// Output:
	// Arguments: [head file.txt 10]
	// Command: head
	// File: file.txt
	// Lines: 10
}

func ExampleParse() {
	// Key is option in format "short-name:long-name" or "long-name"
	// We highly recommend defining options names as constants
	optMap := Map{
		"s:string":   {},                                             // By default, argument has string type
		"S:string2":  {Type: STRING, Value: "Default value"},         // You can predefine default values
		"int":        {Type: INT},                                    // Integer without short name
		"I:int2":     {Type: INT, Min: 1, Max: 10},                   // Integer with limits
		"f:float":    {Type: FLOAT, Value: 10.0},                     // Float
		"b:boolean":  {Type: BOOL},                                   // Boolean
		"m:merg":     {Type: STRING, Mergeble: true},                 // Mergeble options can be defined more than one time
		"h:help":     {Type: BOOL, Alias: "u:usage about"},           // You can define argument aliases
		"e:example":  {Conflicts: []string{"s:string", "S:string2"}}, // Option conflicts with string and string2 (options can't be set at same time)
		"E:example2": {Bound: "int I:int2"},                          // Option bound with int and int2 (options must be set at same time)
	}

	// args is a slice with command arguments
	args, errs := Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	if Has("s:string") {
		fmt.Println("\"--string/-s\" is set")
	} else {
		fmt.Println("\"--string/-s\" is not set")
	}

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("First argument: %s\n\n", args.Get(0).String())

	fmt.Printf("string → %s\n", GetS("string"))
	fmt.Printf("int → %d\n", GetI("int"))
	fmt.Printf("float → %f\n", GetF("f:float"))
	fmt.Printf("boolean → %t\n", GetB("b:boolean"))
}

func ExampleAdd() {
	// Add options
	Add("u:user", &V{Type: STRING, Value: "john"})
	Add("l:lines", &V{Type: INT, Min: 1, Max: 100})

	args, _ := Parse()

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Lines: %d\n", GetI("l:lines"))
}

func ExampleAddMap() {
	// Add options
	AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	args, _ := Parse()

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Lines: %d\n", GetI("l:lines"))
}

func ExampleGetS() {
	args, _ := Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Lines: %d\n", GetI("l:lines"))
}

func ExampleGetI() {
	args, _ := Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Lines: %d\n", GetI("l:lines"))
}

func ExampleGetB() {
	args, _ := Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Force: %t\n", GetB("f:force"))
}

func ExampleGetF() {
	args, _ := Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"r:ratio": {Type: FLOAT},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", GetS("u:user"))
	fmt.Printf("Ratio: %g\n", GetF("r:ratio"))
}

func ExampleSplit() {
	// Use null-terminated string instead of default whitespace for merge.
	// Note that the merge symbol must be set before the parsing options.
	MergeSymbol = "\x00"

	args, _ := Parse(Map{
		"u:user":  {Mergeble: true},
		"r:ratio": {Type: FLOAT},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Users: %s\n", Split("u:user"))
	fmt.Printf("Ratio: %g\n", GetF("r:ratio"))
}

func ExampleIs() {
	Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	if Is("u:user", "bob") && Is("lines", 10) {
		fmt.Println("User is bob and lines number is 10")
	}
}

func ExampleHas() {
	args, _ := Parse(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Has user option: %t\n", Has("u:user"))
	fmt.Printf("Has lines option: %t\n", Has("l:lines"))
}

func ExampleDelete() {
	args, _ := Parse(Map{
		"u:user":     {Type: STRING, Value: "john"},
		"l:lines":    {Type: INT, Min: 1, Max: 100},
		"p:password": {},
	})

	// Remove password option after usage
	Delete("p:password")

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Has password option: %t\n", Has("p:password"))
}

func ExampleFormat() {
	o := "t:test"

	fmt.Printf("Option: %s\n", Format(o))
	// Output:
	// Option: -t/--test
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleV_String() {
	v := &V{
		Type:      INT,
		Value:     25,
		Alias:     "items",
		Conflicts: "E:empty",
		Bound:     "c:create",
		Min:       10,
		Max:       1000,
	}

	fmt.Println(v.String())

	// Output:
	// Int{Value:25 Min:10 Max:1000 Alias:--items Conflicts:-E/--empty Bound:-c/--create}
}

func ExampleMap_String() {
	m := Map{
		"s:size": &V{
			Type:      INT,
			Value:     25,
			Alias:     []string{"items", "objects"},
			Conflicts: "E:empty",
			Bound:     "c:create",
			Min:       10,
			Max:       1000,
		},
	}

	fmt.Println(m.String())

	// Output:
	// options.Map[size:Int{Value:25 Min:10 Max:1000 Alias:[--items --objects] Conflicts:-E/--empty Bound:-c/--create}]
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleOptions_Add() {
	opts := NewOptions()

	// Add options
	opts.Add("u:user", &V{Type: STRING, Value: "john"})
	opts.Add("l:lines", &V{Type: INT, Min: 1, Max: 100})

	args, _ := opts.Parse([]string{"-u", "bob", "-l", "12", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

func ExampleOptions_AddMap() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	args, _ := opts.Parse([]string{"-u", "bob", "-l", "12", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

func ExampleOptions_GetS() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	args, _ := opts.Parse([]string{"-u", "bob", "-l", "12", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

func ExampleOptions_GetI() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	args, _ := opts.Parse([]string{"-u", "bob", "-l", "12", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

func ExampleOptions_GetB() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"f:force": {Type: BOOL},
	})

	args, _ := opts.Parse([]string{"-u", "bob", "-f", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Force: %t\n", opts.GetB("f:force"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Force: true
}

func ExampleOptions_GetF() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"r:ratio": {Type: FLOAT},
	})

	args, _ := opts.Parse([]string{"-u", "bob", "-r", "2.35", "file.txt"})

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Ratio: %g\n", opts.GetF("r:ratio"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Ratio: 2.35
}

func ExampleOptions_Split() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Mergeble: true},
		"r:ratio": {Type: FLOAT},
	})

	// Use null-terminated string instead of default whitespace for merge.
	// Note that the merge symbol must be set before the parsing options.
	MergeSymbol = "\x00"

	input := "-u bob -u john -u dave -r 3.14 file.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Users: %s\n", opts.Split("u:user"))
	fmt.Printf("Ratio: %g\n", opts.GetF("r:ratio"))
	// Output:
	// Arguments: [file.txt]
	// Users: [bob john dave]
	// Ratio: 3.14
}

func ExampleOptions_Has() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":  {Type: STRING, Value: "john"},
		"l:lines": {Type: INT, Min: 1, Max: 100},
	})

	input := "-u bob file.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Has user option: %t\n", opts.Has("u:user"))
	fmt.Printf("Has lines option: %t\n", opts.Has("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// Has user option: true
	// Has lines option: false
}

func ExampleOptions_Delete() {
	opts := NewOptions()

	// Add options
	opts.AddMap(Map{
		"u:user":     {Type: STRING, Value: "john"},
		"l:lines":    {Type: INT, Min: 1, Max: 100},
		"p:password": {},
	})

	input := "-u bob -p Test1234 file.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	opts.Delete("p:password")

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Has user option: %t\n", opts.Has("u:user"))
	fmt.Printf("Has lines option: %t\n", opts.Has("l:lines"))
	fmt.Printf("Has password option: %t\n", opts.Has("p:password"))
	// Output:
	// Arguments: [file.txt]
	// Has user option: true
	// Has lines option: false
	// Has password option: false
}

func ExampleOptions_Parse() {
	opts := NewOptions()

	input := "-u bob -l 12 file.txt"
	args, _ := opts.Parse(
		strings.Split(input, " "),
		Map{
			"u:user":  {Type: STRING, Value: "john"},
			"l:lines": {Type: INT, Min: 1, Max: 100},
		},
	)

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", opts.GetS("u:user"))
	fmt.Printf("Lines: %d\n", opts.GetI("l:lines"))
	// Output:
	// Arguments: [file.txt]
	// User: bob
	// Lines: 12
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleArguments_Has() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Has command: %t\n", args.Has(0))
	// Output:
	// Arguments: [head file.txt 10]
	// Has command: true
}

func ExampleArguments_Get() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Command: %s\n", args.Get(0).String())
	fmt.Printf("File: %s\n", args.Get(1).String())

	lines, _ := args.Get(2).Int()
	fmt.Printf("Lines: %d\n", lines)
	// Output:
	// Arguments: [head file.txt 10]
	// Command: head
	// File: file.txt
	// Lines: 10
}

func ExampleArguments_Last() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Last argument: %s\n", args.Last())
	// Output:
	// Arguments: [head file.txt 10]
	// Last argument: 10
}

func ExampleArguments_Unshift() {
	opts := NewOptions()

	input := "file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	args = args.Unshift("head")

	fmt.Printf("Arguments: %v\n", args.Strings())
	// Output:
	// Arguments: [head file.txt 10]
}

func ExampleArguments_Append() {
	opts := NewOptions()

	input := "head file.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	args = args.Append("10")

	fmt.Printf("Arguments: %v\n", args.Strings())
	// Output:
	// Arguments: [head file.txt 10]
}

func ExampleArguments_Flatten() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args.Flatten())
	// Output:
	// Arguments: head file.txt 10
}

func ExampleArguments_Strings() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args.Strings())
	// Output:
	// Arguments: [head file.txt 10]
}

func ExampleArguments_Filter() {
	opts := NewOptions()

	input := "parse fileA.txt fileB.jpg fileC.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Text files: %s\n", args[1:].Filter("*.txt"))
	// Output:
	// Arguments: [parse fileA.txt fileB.jpg fileC.txt]
	// Text files: [fileA.txt fileC.txt]
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleMap_Set() {
	m := Map{}

	m.Set("u:user", &V{Value: "john"})
	m.Set("i:id", &V{Type: INT})
}

func ExampleMap_SetIf() {
	m := Map{
		"u:user": &V{Value: "john"},
		"i:id":   &V{Type: INT},
	}

	isAdmin := true

	// If isAdmin is true add additional option
	m.SetIf(isAdmin, "g:group", &V{})
}

func ExampleMap_Delete() {
	m := Map{
		"u:user": &V{Value: "abcd"},
	}

	if m.Delete("u:user") {
		fmt.Printf("Option %s deleted\n", F("u:user"))
	} else {
		fmt.Printf("There is no option %s\n", F("u:user"))
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleArgument_String() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Command: %s\n", args.Get(0).String())
	// Output:
	// Arguments: [head file.txt 10]
	// Command: head
}

func ExampleArgument_Is() {
	opts := NewOptions()

	input := "parse fileA.txt fileB.jpg fileC.txt"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Command is \"parse\": %t\n", args.Get(0).Is("parse"))
	fmt.Printf("Command is \"clear\": %t\n", args.Get(0).Is("clear"))
	// Output:
	// Arguments: [parse fileA.txt fileB.jpg fileC.txt]
	// Command is "parse": true
	// Command is "clear": false
}

func ExampleArgument_Int() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	lines, _ := args.Get(2).Int()
	fmt.Printf("Lines: %d\n", lines)
	// Output:
	// Arguments: [head file.txt 10]
	// Lines: 10
}

func ExampleArgument_Int64() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	lines, _ := args.Get(2).Int64()
	fmt.Printf("Lines: %d\n", lines)
	// Output:
	// Arguments: [head file.txt 10]
	// Lines: 10
}

func ExampleArgument_Uint() {
	opts := NewOptions()

	input := "head file.txt 10"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	lines, _ := args.Get(2).Uint()
	fmt.Printf("Lines: %d\n", lines)
	// Output:
	// Arguments: [head file.txt 10]
	// Lines: 10
}

func ExampleArgument_Float() {
	opts := NewOptions()

	input := "ratio 2.37"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	lines, _ := args.Get(1).Float()
	fmt.Printf("Ratio: %g\n", lines)
	// Output:
	// Arguments: [ratio 2.37]
	// Ratio: 2.37
}

func ExampleArgument_Bool() {
	opts := NewOptions()

	input := "release yes"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	force, _ := args.Get(1).Bool()
	fmt.Printf("Force: %t\n", force)
	// Output:
	// Arguments: [release yes]
	// Force: true
}

func ExampleArgument_ToLower() {
	opts := NewOptions()

	input := "add-user John"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", args.Get(1).ToLower().String())
	// Output:
	// Arguments: [add-user John]
	// User: john
}

func ExampleArgument_ToUpper() {
	opts := NewOptions()

	input := "add-user John"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("User: %s\n", args.Get(1).ToUpper().String())
	// Output:
	// Arguments: [add-user John]
	// User: JOHN
}

func ExampleArgument_Clean() {
	opts := NewOptions()

	input := "run /srv/app//conf/myapp.conf"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Clean: %s\n", args.Get(1).Clean())
	fmt.Printf("Base: %s\n", args.Get(1).Base())
	fmt.Printf("Dir: %s\n", args.Get(1).Dir())
	fmt.Printf("Ext: %s\n", args.Get(1).Ext())
	fmt.Printf("IsAbs: %t\n", args.Get(1).IsAbs())
	// Output:
	// Arguments: [run /srv/app//conf/myapp.conf]
	// Clean: /srv/app/conf/myapp.conf
	// Base: myapp.conf
	// Dir: /srv/app/conf
	// Ext: .conf
	// IsAbs: true
}

func ExampleArgument_Base() {
	opts := NewOptions()

	input := "run /srv/app//conf/myapp.conf"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Clean: %s\n", args.Get(1).Clean())
	fmt.Printf("Base: %s\n", args.Get(1).Base())
	fmt.Printf("Dir: %s\n", args.Get(1).Dir())
	fmt.Printf("Ext: %s\n", args.Get(1).Ext())
	fmt.Printf("IsAbs: %t\n", args.Get(1).IsAbs())
	// Output:
	// Arguments: [run /srv/app//conf/myapp.conf]
	// Clean: /srv/app/conf/myapp.conf
	// Base: myapp.conf
	// Dir: /srv/app/conf
	// Ext: .conf
	// IsAbs: true
}

func ExampleArgument_Dir() {
	opts := NewOptions()

	input := "run /srv/app//conf/myapp.conf"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Clean: %s\n", args.Get(1).Clean())
	fmt.Printf("Base: %s\n", args.Get(1).Base())
	fmt.Printf("Dir: %s\n", args.Get(1).Dir())
	fmt.Printf("Ext: %s\n", args.Get(1).Ext())
	fmt.Printf("IsAbs: %t\n", args.Get(1).IsAbs())
	// Output:
	// Arguments: [run /srv/app//conf/myapp.conf]
	// Clean: /srv/app/conf/myapp.conf
	// Base: myapp.conf
	// Dir: /srv/app/conf
	// Ext: .conf
	// IsAbs: true
}

func ExampleArgument_Ext() {
	opts := NewOptions()

	input := "run /srv/app//conf/myapp.conf"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Clean: %s\n", args.Get(1).Clean())
	fmt.Printf("Base: %s\n", args.Get(1).Base())
	fmt.Printf("Dir: %s\n", args.Get(1).Dir())
	fmt.Printf("Ext: %s\n", args.Get(1).Ext())
	fmt.Printf("IsAbs: %t\n", args.Get(1).IsAbs())
	// Output:
	// Arguments: [run /srv/app//conf/myapp.conf]
	// Clean: /srv/app/conf/myapp.conf
	// Base: myapp.conf
	// Dir: /srv/app/conf
	// Ext: .conf
	// IsAbs: true
}

func ExampleArgument_IsAbs() {
	opts := NewOptions()

	input := "run /srv/app//conf/myapp.conf"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("Clean: %s\n", args.Get(1).Clean())
	fmt.Printf("Base: %s\n", args.Get(1).Base())
	fmt.Printf("Dir: %s\n", args.Get(1).Dir())
	fmt.Printf("Ext: %s\n", args.Get(1).Ext())
	fmt.Printf("IsAbs: %t\n", args.Get(1).IsAbs())
	// Output:
	// Arguments: [run /srv/app//conf/myapp.conf]
	// Clean: /srv/app/conf/myapp.conf
	// Base: myapp.conf
	// Dir: /srv/app/conf
	// Ext: .conf
	// IsAbs: true
}

func ExampleArgument_Match() {
	opts := NewOptions()

	input := "parse fileA.txt fileB.jpg"
	args, _ := opts.Parse(strings.Split(input, " "))

	fmt.Printf("Arguments: %v\n", args)
	m1, _ := args.Get(1).Match("*.txt")
	m2, _ := args.Get(2).Match("*.txt")
	fmt.Printf("%s is match: %t\n", args.Get(1), m1)
	fmt.Printf("%s is match: %t\n", args.Get(2), m2)
	// Output:
	// Arguments: [parse fileA.txt fileB.jpg]
	// fileA.txt is match: true
	// fileB.jpg is match: false
}
