package initsystem

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
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

	switch GetServiceState(serviceName) {
	case STATE_STOPPED:
		fmt.Printf("Service %s is stopped\n", serviceName)
	case STATE_WORKS:
		fmt.Printf("Service %s is working\n", serviceName)
	default:
		fmt.Printf("Unknown service %s\n", serviceName)
	}
}
