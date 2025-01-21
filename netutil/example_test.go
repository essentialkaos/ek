package netutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetIP() {
	ip := GetIP()

	if ip != "" {
		fmt.Printf("Your IPv4 is %s\n", ip)
	}
}

func ExampleGetIP6() {
	ip := GetIP6()

	if ip != "" {
		fmt.Printf("Your IPv6 is %s\n", ip)
	}
}

func ExampleGetAllIP() {
	ips := GetAllIP()

	if len(ips) > 0 {
		fmt.Printf("All IPv4: %v\n", ips)
	}
}

func ExampleGetAllIP6() {
	ips := GetAllIP6()

	if len(ips) > 0 {
		fmt.Printf("All IPv6: %v\n", ips)
	}
}
