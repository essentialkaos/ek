package spellcheck

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

func ExampleTrain() {
	name := "Frida"
	names := []string{"Jakalyn", "Freda", "Micayla", "Knightley", "Shun"}

	model := Train(names)

	fmt.Printf("%s → %s\n", name, model.Correct(name))
	// Output: Frida → Freda
}

func ExampleModel_Correct() {
	name := "Frida"
	names := []string{"Jakalyn", "Freda", "Micayla", "Knightley", "Shun"}

	model := Train(names)

	fmt.Printf("%s → %s\n", name, model.Correct(name))
	// Output: Frida → Freda
}

func ExampleModel_Suggest() {
	name := "azi"
	names := []string{"Thorin", "Akhai", "Payden", "Ghazi", "Rey", "Axel", "Sahily", "Azriel"}

	model := Train(names)

	fmt.Printf("%s → %s\n", name, model.Suggest(name, 3))
}
