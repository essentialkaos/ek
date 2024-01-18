//go:build linux || freebsd
// +build linux freebsd

package initsystem

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

func ExampleSysV() {
	if SysV() {
		fmt.Println("SysV init system is used")
	}
}

func ExampleUpstart() {
	if Upstart() {
		fmt.Println("Upstart init system is used")
	}
}

func ExampleSystemd() {
	if Systemd() {
		fmt.Println("Systemd init system is used")
	}
}

func ExampleIsPresent() {
	serviceName := "crond"

	if IsPresent(serviceName) {
		fmt.Printf("Service %s is present\n", serviceName)
	} else {
		fmt.Printf("Unknown service %s\n", serviceName)
	}
}

func ExampleIsWorks() {
	serviceName := "crond"

	works, err := IsWorks(serviceName)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if works {
		fmt.Printf("Service %s is working\n", serviceName)
	} else {
		fmt.Printf("Service %s is stopped\n", serviceName)
	}
}

func ExampleIsEnabled() {
	serviceName := "crond"

	enabled, err := IsEnabled(serviceName)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if enabled {
		fmt.Printf("Service %s is enabled\n", serviceName)
	} else {
		fmt.Printf("Service %s is not enabled\n", serviceName)
	}
}
