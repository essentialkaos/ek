package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetEngine() {
	fmt.Printf("Container engine: %s\n", GetEngine())
}

func ExampleIsContainer() {
	if !IsContainer() {
		fmt.Println("It's not containerized")
	} else {
		fmt.Printf("It's containerized using %s engine\n", GetEngine())
	}
}

func ExampleGetCPUQuota() {
	quota, isMax, err := GetCPUQuota()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Quota: %g (max: %t)\n", quota, isMax)
}

func ExampleGetMemLimit() {
	limit, isMax, err := GetMemLimit()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Limit: %d (max: %t)\n", limit, isMax)
}

func ExampleGetCPUSet() {
	cpuSet, err := GetCPUSet()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("CPU Set: %v\n", cpuSet)
}
