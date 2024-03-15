package strutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleQ() {
	defaultValue := "john"
	user := ""

	fmt.Println(Q(user, defaultValue))

	// Output:
	// john
}

func ExampleB() {
	isAdmin := true
	user := "bob"

	fmt.Printf(
		B(isAdmin, "User %s is admin\n", "User %s isn't admin\n"),
		user,
	)

	// Output:
	// User bob is admin
}

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
	fmt.Println(Substr("Пример 例子 例 მაგალითად", 7, 2))

	// Output:
	// 例子
}

func ExampleSubstring() {
	fmt.Println(Substring("Пример 例子 例 მაგალითად", 7, 9))

	// Output:
	// 例子
}

func ExampleExtract() {
	fmt.Println(Extract("This is funny message", 8, 13))

	// Output:
	// funny
}

func ExampleExclude() {
	fmt.Println(Exclude("This is funny message", " funny"))

	// Output:
	// This is message
}

func ExampleLen() {
	fmt.Println(Len("Пример 例子 例 მაგალითად"))
	fmt.Println(Len("😚😘🥰"))

	// Output:
	// 21
	// 3
}

func ExampleLenVisual() {
	k := "🥰 Пример 例子 例 მაგალითად"
	l := LenVisual(k)

	fmt.Println(k)
	fmt.Println(strings.Repeat("^", l))

	// Output:
	// 🥰 Пример 例子 例 მაგალითად
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^
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

func ExampleReplaceIgnoreCase() {
	fmt.Println(ReplaceIgnoreCase(
		"User bob has no item. Add items to user Bob?", "bob", "[Bob]",
	))

	// Output:
	// User [Bob] has no item. Add items to user [Bob]?
}

func ExampleFields() {
	fmt.Printf("%#v\n", Fields("Bob  Alice, 'Mary Key', \"John Dow\""))

	// Output:
	// []string{"Bob", "Alice", "Mary Key", "John Dow"}
}

func ExampleReadField() {
	fmt.Println(ReadField("Bob    Alice\tJohn Mary", 2, true, ' ', '\t'))
	fmt.Println(ReadField("Bob:::Mary:John:", 3, false, ':'))

	// Output:
	// John
	// Mary
}

func ExampleCopy() {
	fmt.Println(Copy("abc"))

	// Output:
	// abc
}

func ExampleBefore() {
	fmt.Println(Before("john@domain.com", "@"))

	// Output:
	// john
}

func ExampleAfter() {
	fmt.Println(After("john@domain.com", "@"))

	// Output:
	// domain.com
}

func ExampleHasPrefixAny() {
	fmt.Println(HasPrefixAny("www.domain.com", "dl", "www"))
	fmt.Println(HasPrefixAny("api.domain.com", "dl", "www"))

	// Output:
	// true
	// false
}

func ExampleHasSuffixAny() {
	fmt.Println(HasSuffixAny("www.domain.com", ".com", ".org", ".net"))
	fmt.Println(HasSuffixAny("www.domain.info", ".com", ".org", ".net"))

	// Output:
	// true
	// false
}

func ExampleIndexByteSkip() {
	// Index from left
	fmt.Println(IndexByteSkip("/home/john/projects/test.log", '/', 2))

	// Index from right
	fmt.Println(IndexByteSkip("/home/john/projects/test.log", '/', -1))

	// Output:
	// 10
	// 10
}
