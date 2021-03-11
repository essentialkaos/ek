package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleP() {
	fmt.Println(P("I have %d %s", 1, "apple", "apples"))
	fmt.Println(P("I have %d %s", 5, "apple", "apples"))
	// Output: I have 1 apple
	// I have 5 apples
}

func ExamplePS() {
	fmt.Println(PS(Ru, "У меня %d %s", 1, "яблоко", "яблока", "яблок"))
	fmt.Println(PS(Ru, "У меня %d %s", 5, "яблоко", "яблока", "яблок"))
	// Output: У меня 1 яблоко
	// У меня 5 яблок
}

func ExamplePluralize() {
	fmt.Println(Pluralize(1, "apple", "apples"))
	fmt.Println(Pluralize(5, "apple", "apples"))
	// Output: apple
	// apples
}
func ExamplePluralizeSpecial() {
	fmt.Println(PluralizeSpecial(Ru, 1, "яблоко", "яблока", "яблок"))
	fmt.Println(PluralizeSpecial(Ru, 5, "яблоко", "яблока", "яблок"))
	// Output: яблоко
	// яблок
}
