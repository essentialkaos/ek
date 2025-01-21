package sdnotify

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleConnect() {
	err := Connect()

	if err != nil {
		panic(err.Error())
	}

	Status("Loading data %d%%", 50)
	Ready()
}
