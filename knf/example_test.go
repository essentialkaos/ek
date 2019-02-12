package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

func ExampleValidate() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	errs := Validate([]*Validator{
		{"section:property1", Empty, nil},                      // Return error if value is empty
		{"section:property1", Less, 10},                        // Return error if value less than 10
		{"section:property1", Greater, 50},                     // Return error if value greater than 50
		{"section:property1", Equals, 33},                      // Return error if value equals 33
		{"section:property2", NotContains, []string{"a", "b"}}, // Return error if value not in given slice
	})

	if len(errs) != 0 {
		for _, err = range errs {
			fmt.Printf("Error: %v\n", err)
		}
	}
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
