package i18n

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Bundle struct {
	GREETING Text
	MESSAGE  Text

	ERRORS *Errors

	// You can also store additional information, which will not be merged by the
	// fallback method
	DateFormat string
	TimeFormat string
}

type Errors struct {
	UNKNOWN_USER Text
	UNKNOWN_ID   Text
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleFallback() {
	en := &Bundle{
		DateFormat: "%D",
		TimeFormat: "%l:%M %p",

		GREETING: "Hello!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	ru := &Bundle{
		DateFormat: "%Y/%m/%d",
		TimeFormat: "%H:%M",

		GREETING: "Привет!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Неизвестный пользователь {{.Username}}",
		},
	}

	kz := &Bundle{
		DateFormat: "%Y/%m/%d",
		TimeFormat: "%H:%M",

		GREETING: "Сәлеметсіз бе!",
	}

	loc, err := Fallback(en, ru, kz)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	l := loc.(*Bundle)

	err = ValidateBundle(l)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := Data{
		"Username": "johndoe",
		"ID":       183,
	}

	fmt.Println(l.GREETING)
	fmt.Println(l.ERRORS.UNKNOWN_USER.With(data))
	fmt.Println(l.ERRORS.UNKNOWN_ID.With(data))

	// Output:
	// Сәлеметсіз бе!
	// Неизвестный пользователь johndoe
	// Unknown ID 183
}

func ExampleIsComplete() {
	en := &Bundle{
		GREETING: "Hello!",
		MESSAGE:  "Hi user!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	ru := &Bundle{
		GREETING: "Привет!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Неизвестный пользователь {{.Username}}",
		},
	}

	isComplete, _ := IsComplete(en)

	fmt.Printf("EN is complete: %t\n", isComplete)

	isComplete, fields := IsComplete(ru)

	fmt.Printf("RU is complete: %t (empty: %s)\n", isComplete, strings.Join(fields, ", "))

	// Output:
	// EN is complete: true
	// RU is complete: false (empty: MESSAGE, ERRORS.UNKNOWN_ID)
}

func ExampleText_With() {
	en := &Bundle{
		GREETING: "Hello!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	data := Data{
		"Username": "johndoe",
		"ID":       183,
	}

	fmt.Println(en.ERRORS.UNKNOWN_USER.With(data))
	fmt.Println(en.ERRORS.UNKNOWN_ID.With(data))

	// Also, you can add prefix/suffix
	fmt.Println(en.ERRORS.UNKNOWN_USER.With(data, "[!] ", " (ERR/21.1)"))

	// Output:
	// Unknown user johndoe
	// Unknown ID 183
	// [!] Unknown user johndoe (ERR/21.1)
}

func ExampleText_Write() {
	en := &Bundle{
		GREETING: "Hello!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	data := Data{
		"Username": "johndoe",
		"ID":       183,
	}

	buf := &bytes.Buffer{}

	en.ERRORS.UNKNOWN_USER.Write(buf, data)

	buf.WriteRune('\n')

	// Also, you can add prefix/suffix
	en.ERRORS.UNKNOWN_USER.Write(buf, data, "- ", " (ERR/21.1)\n")

	fmt.Println(buf.String())

	// Output:
	// Unknown user johndoe
	// - Unknown user johndoe (ERR/21.1)
}

func ExampleText_Add() {
	en := &Bundle{
		GREETING: "Hello",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	fmt.Println(en.GREETING.Add("> ", "!"))

	// Output:
	// > Hello!
}

func ExampleText_Start() {
	en := &Bundle{
		GREETING: "Hello",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	fmt.Println(en.ERRORS.UNKNOWN_USER.Start("[!] "))

	// Output:
	// [!] Unknown user {{.Username}}
}

func ExampleText_End() {
	en := &Bundle{
		GREETING: "Hello",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	fmt.Println(en.GREETING.End(" [i]"))

	// Output:
	// Hello [i]
}

func ExampleText_String() {
	en := &Bundle{
		GREETING: "Hello!",
		ERRORS: &Errors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	fmt.Println(en.GREETING.String())

	// Output:
	// Hello!
}

func ExampleText_Format() {
	en := &Bundle{
		GREETING: "Hello %s!",
	}

	fmt.Println(en.GREETING.Format("Bob"))

	// Output:
	// Hello Bob!
}

func ExampleData_Has() {
	data := Data{
		"ServerNum": 12,
	}

	fmt.Println("ServerNum", data.Has("ServerNum"))
	fmt.Println("ServerName", data.Has("ServerName"))

	// Output:
	// ServerNum true
	// ServerName false
}

func ExampleData_Plural() {
	en := &Bundle{
		GREETING: "Hello!",
		MESSAGE:  `You started {{.ServerNum}} {{.Plural "EN" "ServerNum" "server" "servers"}}`,
	}

	data := Data{
		"ServerNum": 12,
	}

	fmt.Println(en.MESSAGE.With(data))

	// Output:
	// You started 12 servers
}

func ExampleData_PrettyNum() {
	en := &Bundle{
		GREETING: "Hello!",
		MESSAGE:  `Your balance is ${{.PrettyNum "Balance"}}`,
	}

	data := Data{
		"Balance": 3193,
	}

	fmt.Println(en.MESSAGE.With(data))

	// Output:
	// Your balance is $3,193
}

func ExampleData_PrettySize() {
	en := &Bundle{
		GREETING: "Hello!",
		MESSAGE:  `Found {{.PrettyNum "Files"}} {{.Plural "EN" "Files" "file" "files"}} (size: {{.PrettySize "Size"}})`,
	}

	data := Data{
		"Files": 731,
		"Size":  103810746,
	}

	fmt.Println(en.MESSAGE.With(data))

	// Output:
	// Found 731 files (size: 99MB)
}

func ExampleData_PrettyPerc() {
	en := &Bundle{
		GREETING: "Hello!",
		MESSAGE:  `Copied {{.PrettyPerc "Progress"}} of files`,
	}

	data := Data{
		"Progress": 45.31,
	}

	fmt.Println(en.MESSAGE.With(data))

	// Output:
	// Copied 45.3% of files
}
