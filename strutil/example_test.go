package strutil

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
	fmt.Println(Substr("This is funny message", 8, 5))

	// Output:
	// funny
}

func ExampleSubstring() {
	fmt.Println(Substr("This is funny message", 8, 13))

	// Output:
	// funny message
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

func ExampleReadField() {
	fmt.Println(ReadField("Bob    Alice\tJohn Mary", 2, true, " ", "\t"))
	fmt.Println(ReadField("Bob:::Mary:John:", 3, false, ":"))

	// Output:
	// John
	// Mary
}

func ExampleCopy() {
	fmt.Println(Copy("abc"))

	// Output:
	// abc
}
