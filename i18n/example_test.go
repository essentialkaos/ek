package i18n

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////// //

type Bundle struct {
	GREETING String
	MESSAGE  String

	ERRORS *Errors
}

type Errors struct {
	UNKNOWN_USER String
	UNKNOWN_ID   String
}

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleFallback() {
	en := &Bundle{
		GREETING: "Hello!",
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

	kz := &Bundle{
		GREETING: "Сәлеметсіз бе!",
	}

	loc, _ := Fallback(en, ru, kz)
	l := loc.(*Bundle)

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

func ExampleString_With() {
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

	// Output:
	// Unknown user johndoe
	// Unknown ID 183
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
