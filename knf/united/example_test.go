package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/options"
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

	// Combine combine KNF configuration, options and environment variables
	Combine(
		config,
		Mapping{"test:option-one", "O:option-one", "TEST_OPTION_ONE"},
		Mapping{"test:option-two", "k:option-two", "TEST_OPTION_TWO"},
	)

	// Also you can set options and environment variables using helpers
	var (
		optOne = "test:option-one"
		optTwo = "test:option-two"
	)

	Combine(
		config,
		// Create mapping manually
		Mapping{optOne, ToOption(optOne), ToEnvVar(optOne)},
		// Create simple mapping
		Simple(optTwo),
	)

	// Read string value
	GetS("section:string")

	// Read integer value
	GetI("section:int")

	// Read float value
	GetF("section:float")

	// Read boolean value
	GetB("section:boolean")

	// Read file mode value
	GetM("section:file-mode")

	// Read duration as seconds
	GetD("section:duration", Second)

	// Read duration as minutes
	GetD("section:duration", Minute)

	// Read time duration
	GetTD("section:time-duration")

	// Read timestamp
	GetTS("section:timestamp")

	// Read timezone
	GetTZ("section:timezone")

	// Read list
	GetL("section:list")
}

func ExampleAddOptions() {
	m := options.Map{}

	AddOptions(m, "test:option-one", "test:option-two")

	fmt.Printf("Map size: %d\n", len(m))

	// Output:
	// Map size: 2
}

func ExampleSimple() {
	m := Simple("test:option-one")

	fmt.Printf("%s → --%s + %s\n", m.Property, m.Option, m.Variable)

	// Output:
	// test:option-one → --test-option-one + TEST_OPTION_ONE
}

func ExampleToOption() {
	fmt.Println(ToOption("section:time-duration"))

	// Output: section-time-duration
}

func ExampleToEnvVar() {
	fmt.Println(ToEnvVar("section:time-duration"))

	// Output: SECTION_TIME_DURATION
}

func ExampleO() {
	fmt.Println(ToOption("section:time-duration"))

	// Output: section-time-duration
}

func ExampleE() {
	fmt.Println(ToEnvVar("section:time-duration"))

	// Output: SECTION_TIME_DURATION
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

	fmt.Printf("Value from config: %v\n", GetD("user:timeout", Minute))
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
