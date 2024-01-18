package lock

import "time"

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCreate() {
	err := Create("my_app")

	if err != nil {
		panic(err.Error())
	}

	// do some stuff

	err = Remove("my_app")

	if err != nil {
		panic(err.Error())
	}
}

func ExampleRemove() {
	err := Create("my_app")

	if err != nil {
		panic(err.Error())
	}

	// do your stuff

	err = Remove("my_app")

	if err != nil {
		panic(err.Error())
	}
}

func ExampleHas() {
	if Has("my_app") {
		panic("Can't start application due to lock file")
	}
}

func ExampleWait() {
	ok := Wait("my_app", time.Now().Add(time.Minute))

	if !ok {
		panic("Can't start application due to lock file")
	}

	// do your stuff
}

func ExampleIsExpired() {
	if Has("my_app") {
		if IsExpired("my_app", time.Hour) {
			// looks like lock file was created long time ago, so delete it
			Remove("my_app")
		} else {
			panic("Can't start application due to lock file")
		}
	}
}
