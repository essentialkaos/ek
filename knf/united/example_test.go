package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/options"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCombine() {
	// Load KNF config
	config, err := knf.Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	optMap := options.Map{
		"O:option-one": {},
		"k:option-two": {},
	}

	// Parse command-line options
	_, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
		}

		return
	}

	// Combine combines KNF configuration, options and environment variables
	Combine(
		config,
		Mapping{"main:string", "S:main-string", "MAIN_STRING"},
		Mapping{"main:int", "I:main-int", "MAIN_INT"},
		Mapping{"main:float", "F:main-float", "MAIN_FLOAT"},
		Mapping{"main:boolean", "B:main-boolean", "MAIN_BOOLEAN"},
		Mapping{"main:file-mode", "FM:main-string", "MAIN_FILE_MODE"},
		Mapping{"main:duration", "D:main-string", "MAIN_DURATION"},
		Mapping{"main:size", "s:main-string", "MAIN_SIZE"},
		Mapping{"main:time-duration", "td:main-string", "MAIN_TIME_DURATION"},
		Mapping{"main:timestamp", "T:main-string", "MAIN_TIMESTAMP"},
		Mapping{"main:timezone", "tz:main-string", "MAIN_TIMEZONE"},
		Mapping{"main:list", "L:main-string", "MAIN_LIST"},
	)

	// Also, you can set options and environment variables using helpers
	var (
		mainString       = "main:string"
		mainInt          = "main:int"
		mainFloat        = "main:float"
		mainBoolean      = "main:boolean"
		mainFileMode     = "main:file-mode"
		mainDuration     = "main:duration"
		mainSize         = "main:size"
		mainTimeDuration = "main:time-duration"
		mainTimestamp    = "main:timestamp"
		mainTimezone     = "main:timezone"
		mainList         = "main:list"
	)

	Combine(
		config,
		// Create mapping manually
		Mapping{mainString, ToOption(mainString), ToEnvVar(mainString)},
		// Create simple mapping
		Simple(mainInt),
		Simple(mainFloat),
		Simple(mainBoolean),
		Simple(mainFileMode),
		Simple(mainDuration),
		Simple(mainSize),
		Simple(mainTimeDuration),
		Simple(mainTimestamp),
		Simple(mainTimezone),
		Simple(mainList),
	)

	// Read string value
	GetS("main:string")

	// Read integer value
	GetI("main:int")

	// Read float value
	GetF("main:float")

	// Read boolean value
	GetB("main:boolean")

	// Read file mode value
	GetM("main:file-mode")

	// Read duration as seconds
	GetD("main:duration", SECOND)

	// Read duration as minutes
	GetD("main:duration", MINUTE)

	// Read size
	GetSZ("main:size")

	// Read time duration
	GetTD("main:time-duration")

	// Read timestamp
	GetTS("main:timestamp")

	// Read timezone
	GetTZ("main:timezone")

	// Read list
	GetL("main:list")
}

func ExampleCombineSimple() {
	// Load KNF config
	config, err := knf.Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	optMap := options.Map{
		"O:option-one": {},
		"k:option-two": {},
	}

	// Parse command-line options
	_, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
		}

		return
	}

	// Combine simple combines KNF configuration, options and environment variables
	CombineSimple(config, "test:option-one", "test:option-two")
}

func ExampleAddOptions() {
	m := options.Map{}

	AddOptions(m, "test:option-one", "test:option-two")

	fmt.Printf("Map size: %d\n", len(m))

	// Output:
	// Map size: 2
}

func ExampleGetMapping() {
	// Load KNF config
	config, err := knf.Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Combine combines KNF configuration, options and environment variables
	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	// Print mapping for property test:option-one
	fmt.Println(GetMapping("test:option-one"))
}

func ExampleSimple() {
	m := Simple("test:option-one")

	fmt.Printf("%s → --%s + %s\n", m.Property, m.Option, m.Variable)

	// Output:
	// test:option-one → --test-option-one + TEST_OPTION_ONE
}

func ExampleToOption() {
	fmt.Println(ToOption("main:time-duration"))

	// Output: main-time-duration
}

func ExampleToEnvVar() {
	fmt.Println(ToEnvVar("main:time-duration"))

	// Output: MAIN_TIME_DURATION
}

func ExampleO() {
	fmt.Println(ToOption("main:time-duration"))

	// Output: main-time-duration
}

func ExampleE() {
	fmt.Println(ToEnvVar("main:time-duration"))

	// Output: MAIN_TIME_DURATION
}

func ExampleGetS() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %s\n", GetS("user:name"))
}

func ExampleGetB() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %t\n", GetB("user:is-admin"))
}

func ExampleGetI() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %d\n", GetI("user:uid"))
}

func ExampleGetI64() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %d\n", GetI64("user:uid"))
}

func ExampleGetU() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %d\n", GetU("user:uid"))
}

func ExampleGetU64() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %d\n", GetU64("user:uid"))
}

func ExampleGetF() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %g\n", GetF("user:priority"))
}

func ExampleGetM() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetF("user:default-mode"))
}

func ExampleGetD() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetD("user:timeout", MINUTE))
}

func ExampleGetSZ() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetSZ("user:max-size"))
}

func ExampleGetTD() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetTD("user:timeout"))
}

func ExampleGetTS() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetTS("user:created"))
}

func ExampleGetTZ() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetTZ("service:timezone"))
}

func ExampleGetL() {
	config, _ := knf.Read("/path/to/your/config.knf")

	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	fmt.Printf("Value from config: %v\n", GetL("issue:labels"))
}
