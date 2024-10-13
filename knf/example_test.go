package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

	// Read duration as seconds
	GetD("section:duration", SECOND)

	// Read duration as minutes
	GetD("section:duration", MINUTE)

	// Read time duration
	GetTD("section:time-duration")

	// Read timestamp
	GetTS("section:timestamp")

	// Read timezone
	GetTZ("section:timezone")

	// Read list
	GetL("section:list")

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

	fmt.Printf("Value from config: %s\n", cfg.GetS("service:user"))
}

func ExampleParse() {
	cfg, err := Parse([]byte(`
[service]
	user: john
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", cfg.GetS("service:user"))
	// Output:
	// Value from config: john
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

func ExampleAlias() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Add alias for renamed property "user:username"
	Alias("user:username", "user:name")

	fmt.Printf("Value from config: %s\n", GetS("user:name"))
}

func ExampleGetS() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", GetS("user:name"))
}

func ExampleGetB() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %t\n", GetB("user:is-admin"))
}

func ExampleGetI() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", GetI("user:uid"))
}

func ExampleGetI64() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", GetI64("user:uid"))
}

func ExampleGetU() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", GetU("user:uid"))
}

func ExampleGetU64() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", GetU64("user:uid"))
}

func ExampleGetF() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %g\n", GetF("user:priority"))
}

func ExampleGetM() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetF("user:default-mode"))
}

func ExampleGetD() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetD("user:timeout", MINUTE))
}

func ExampleGetTD() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetTD("user:timeout"))
}

func ExampleGetTS() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetTS("user:created"))
}

func ExampleGetTZ() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetTZ("service:timezone"))
}

func ExampleGetL() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", GetL("issue:labels"))
}

func ExampleIs() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("user.name == bob: %t\n", Is("user:name", "bob"))
	fmt.Printf("user.uid == 512: %t\n", Is("user:uid", 512))
}

func ExampleHasSection() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is section \"main\" exist: %t\n", HasSection("main"))
	fmt.Printf("Is section \"user\" exist: %t\n", HasSection("user"))
}

func ExampleHasProp() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is property \"user:name\" exist: %t\n", HasProp("user:name"))
	fmt.Printf("Is property \"user:path\" exist: %t\n", HasProp("user:path"))
}

func ExampleSections() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, section := range Sections() {
		fmt.Printf("%d: %s\n", index+1, section)
	}
}

func ExampleProps() {
	err := Global("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, prop := range Props("user") {
		fmt.Printf("%d: %s\n", index+1, prop)
	}
}

func ExampleConfig_Merge() {
	cfg1, _ := Parse([]byte(`
[service]
	user: john
`))

	cfg2, _ := Parse([]byte(`
[service]
	user: bob
`))

	fmt.Printf("Value from config (before merge): %s\n", cfg1.GetS("service:user"))

	err := cfg1.Merge(cfg2)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config (after merge): %s\n", cfg1.GetS("service:user"))

	// Output:
	// Value from config (before merge): john
	// Value from config (after merge): bob
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

func ExampleConfig_Alias() {
	cfg, err := Parse([]byte(`
[user]
	username: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Add alias for renamed property "user:username"
	cfg.Alias("user:username", "user:name")

	fmt.Printf("Value from config: %s\n", cfg.GetS("user:name"))

	// Output:
	// Value from config: john
}

func ExampleConfig_GetS() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", cfg.GetS("user:name"))

	// Output:
	// Value from config: john
}

func ExampleConfig_GetB() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %t\n", cfg.GetB("user:is-admin"))

	// Output:
	// Value from config: true
}

func ExampleConfig_GetI() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", cfg.GetI("user:uid"))

	// Output:
	// Value from config: 512
}

func ExampleConfig_GetI64() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", cfg.GetI64("user:uid"))

	// Output:
	// Value from config: 512
}

func ExampleConfig_GetU() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", cfg.GetU("user:uid"))

	// Output:
	// Value from config: 512
}

func ExampleConfig_GetU64() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %d\n", cfg.GetU64("user:uid"))

	// Output:
	// Value from config: 512
}

func ExampleConfig_GetF() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %g\n", cfg.GetF("user:priority"))

	// Output:
	// Value from config: 3.7
}

func ExampleConfig_GetM() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", cfg.GetF("user:default-mode"))

	// Output:
	// Value from config: 644
}

func ExampleConfig_GetD() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", cfg.GetD("user:timeout", MINUTE))

	// Output:
	// Value from config: 3m0s
}

func ExampleConfig_GetTD() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", cfg.GetTD("user:timeout"))

	// Output:
	// Value from config: 3m0s
}

func ExampleConfig_GetTS() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %v\n", cfg.GetTS("user:created"))
}

func ExampleConfig_GetTZ() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", cfg.GetTZ("user:timezone"))

	// Output:
	// Value from config: Europe/Madrid
}

func ExampleConfig_GetL() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Value from config: %s\n", cfg.GetL("user:labels"))

	// Output:
	// Value from config: [system admin]
}

func ExampleConfig_Is() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("user.name == bob: %t\n", cfg.Is("user:name", "bob"))
	fmt.Printf("user.uid == 512: %t\n", cfg.Is("user:uid", 512))

	// Output:
	// user.name == bob: false
	// user.uid == 512: true
}

func ExampleConfig_HasSection() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is section \"main\" exist: %t\n", cfg.HasSection("main"))
	fmt.Printf("Is section \"user\" exist: %t\n", cfg.HasSection("user"))

	// Output:
	// Is section "main" exist: false
	// Is section "user" exist: true
}

func ExampleConfig_HasProp() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Is property \"user:name\" exist: %t\n", cfg.HasProp("user:name"))
	fmt.Printf("Is property \"user:path\" exist: %t\n", cfg.HasProp("user:path"))

	// Output:
	// Is property "user:name" exist: true
	// Is property "user:path" exist: false
}

func ExampleConfig_Sections() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin

[log]
	file: /var/log/app/app.log
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, section := range cfg.Sections() {
		fmt.Printf("%d: %s\n", index+1, section)
	}

	// Output:
	// 1: user
	// 2: log
}

func ExampleConfig_Props() {
	cfg, err := Parse([]byte(`
[user]
	name: john
	uid: 512
	is-admin: true
	priority: 3.7
	default-mode: 0644
	timeout: 3m
	created: 1654424130
	timezone: Europe/Madrid
	labels: system admin
`))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for index, prop := range cfg.Props("user") {
		fmt.Printf("%d: %s\n", index+1, prop)
	}

	// Output:
	// 1: name
	// 2: uid
	// 3: is-admin
	// 4: priority
	// 5: default-mode
	// 6: timeout
	// 7: created
	// 8: timezone
	// 9: labels
}

func ExampleConfig_File() {
	cfg, err := Read("/path/to/your/config.knf")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Path to config: %s\n", cfg.File())
}
