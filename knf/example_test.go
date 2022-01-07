package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
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

	// Read duration in seconds
	GetD("section:duration")

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
