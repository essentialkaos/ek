//+build linux freebsd

package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

func ExampleHasService() {
	serviceName := "crond"

	if HasService(serviceName) {
		fmt.Printf("Service %s is present\n", serviceName)
	} else {
		fmt.Printf("Unknown service %s\n", serviceName)
	}
}

func ExampleGetServiceState() {
	serviceName := "crond"

	works, err := IsServiceWorks(serviceName)

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
