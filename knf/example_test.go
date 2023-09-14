package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGlobal() {
	// Load global config
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

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
	GetD("section:duration", time.Second)

	// Read duration as minutes
	GetD("section:duration", time.Minute)

	// Check section
	if HasSection("section") {
		// Section exist
	}

	// Check property
	if HasProp("section:string") {
		// Property exist
	}

	// Slice of all sections
	Sections()

	// Slice of all properties in section
	Props("section")
}

func ExampleRead() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", cfg.GetS("section:string"))
}

func ExampleReload() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	changes, err := Reload()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print info about changed values
	for prop, changed := range changes {
		fmt.Printf("Property %s changed → %t\n", prop, changed)
	}
}

func ExampleGetS() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("String value from config: %s\n", GetS("section:string"))
}

func ExampleGetI() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (int) value from config: %d\n", GetI("section:int"))
}

func ExampleGetI64() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (int64) value from config: %d\n", GetI64("section:int64"))
}

func ExampleGetU() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (uint) value from config: %d\n", GetU("section:uint"))
}

func ExampleGetU64() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (uint64) value from config: %d\n", GetU64("section:uint64"))
}

func ExampleGetF() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Floating number value from config: %g\n", GetF("section:float"))
}

func ExampleGetM() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("File mode value from config: %v\n", GetM("section:file-mode"))
}

func ExampleGetD() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Duration value from config (as seconds): %v\n", GetD("section:duration", time.Second))
	fmt.Printf("Duration value from config (as minutes): %v\n", GetD("section:duration", time.Minute))
}

func ExampleIs() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("String value from config equals \"test\": %t\n", Is("section:string", "test"))
	fmt.Printf("Integer value from config equals \"123\": %t\n", Is("section:int", 123))
	fmt.Printf("Floating number value from config equals \"12.34\": %t\n", Is("section:float", 12.34))
}

func ExampleHasSection() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is section \"main\" exist: %t\n", HasSection("main"))
}

func ExampleHasProp() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is property \"section:string\" exist: %t\n", HasProp("section:string"))
}

func ExampleSections() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, section := range Sections() {
		fmt.Printf("%d: %s\n", index, section)
	}
}

func ExampleProps() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, prop := range Props("section") {
		fmt.Printf("%d: %s\n", index, prop)
	}
}

func ExampleConfig_Reload() {
	config, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	changes, err := config.Reload()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print info about changed values
	for prop, changed := range changes {
		fmt.Printf("Property %s changed → %t\n", prop, changed)
	}
}

func ExampleConfig_GetS() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("String value from config: %s\n", cfg.GetS("section:string"))
}

func ExampleConfig_GetI() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (int) value from config: %d\n", cfg.GetI("section:int"))
}

func ExampleConfig_GetI64() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (int64) value from config: %d\n", cfg.GetI64("section:int64"))
}

func ExampleConfig_GetU() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (uint) value from config: %d\n", cfg.GetU("section:uint"))
}

func ExampleConfig_GetU64() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Integer (uint64) value from config: %d\n", cfg.GetU64("section:uint64"))
}

func ExampleConfig_GetF() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Floating number value from config: %g\n", cfg.GetF("section:float"))
}

func ExampleConfig_GetM() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("File mode value from config: %v\n", cfg.GetM("section:file-mode"))
}

func ExampleConfig_GetD() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Duration value from config (as seconds): %v\n", cfg.GetD("section:duration", time.Second))
	fmt.Printf("Duration value from config (as minutes): %v\n", cfg.GetD("section:duration", time.Minute))
}

func ExampleConfig_Is() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("String value from config equals \"test\": %t\n", cfg.Is("section:string", "test"))
	fmt.Printf("Integer value from config equals \"123\": %t\n", cfg.Is("section:int", 123))
	fmt.Printf("Floating number value from config equals \"12.34\": %t\n", cfg.Is("section:float", 12.34))
}

func ExampleConfig_HasSection() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is section \"main\" exist: %t\n", cfg.HasSection("main"))
}

func ExampleConfig_HasProp() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is property \"section:string\" exist: %t\n", cfg.HasProp("section:string"))
}

func ExampleConfig_Sections() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, section := range cfg.Sections() {
		fmt.Printf("%d: %s\n", index, section)
	}
}

func ExampleConfig_Props() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, prop := range cfg.Props("section") {
		fmt.Printf("%d: %s\n", index, prop)
	}
}
