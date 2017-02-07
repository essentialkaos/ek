package strutil

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

func ExampleConcat() {
	fmt.Println(Concat("abc", " ", "123", " ", "ABC"))

	// Output:
	// abc 123 ABC
}

func ExampleEllipsis() {
	fmt.Println(Ellipsis("This is too long message to show", 18))

	// Output:
	// This is too lon...
}

func ExampleSubstr() {
	fmt.Println(Substr("This is funny message", 8, 13))

	// Output:
	// funny
}

func ExampleLen() {
	fmt.Println(Len("Пример 例子 例"))

	// Output:
	// 11
}

func ExampleHead() {
	fmt.Println(Head("This is funny message", 7))

	// Output:
	// This is
}

func ExampleTail() {
	fmt.Println(Tail("This is funny message", 7))

	// Output:
	// message
}

func ExamplePrefixSize() {
	fmt.Println(PrefixSize("#### Header 4", '#'))

	// Output:
	// 4
}

func ExampleSuffixSize() {
	fmt.Println(SuffixSize("Message    ", ' '))

	// Output:
	// 4
}

func ExampleReplaceAll() {
	fmt.Println(ReplaceAll("Message", "es", "?"))

	// Output:
	// M???ag?
}

func ExampleFields() {
	fmt.Printf("%#v\n", Fields("Bob  Alice, 'Mary Key', \"John Dow\""))

	// Output:
	// []string{"Bob", "Alice", "Mary Key", "John Dow"}
}
