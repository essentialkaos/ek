package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"net/http"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetRequestAddr() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)

	fmt.Println(GetRequestAddr(r))

	// Output: http.cat 443
}

func ExampleGetRequestHost() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)

	fmt.Println(GetRequestHost(r))

	// Output: http.cat
}

func ExampleGetRequestPort() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)

	fmt.Println(GetRequestPort(r))

	// Output: 443
}

func ExampleGetRemoteAddr() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)
	r.RemoteAddr = "127.0.0.1:19371"

	fmt.Println(GetRemoteAddr(r))

	// Output: 127.0.0.1 19371
}

func ExampleGetRemoteHost() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)
	r.RemoteAddr = "127.0.0.1:19371"

	fmt.Println(GetRemoteHost(r))

	// Output: 127.0.0.1
}

func ExampleGetRemotePort() {
	r, _ := http.NewRequest("GET", "https://http.cat/200", nil)
	r.RemoteAddr = "127.0.0.1:19371"

	fmt.Println(GetRemotePort(r))

	// Output: 19371
}

func ExampleGetDescByCode() {
	fmt.Println("200:", GetDescByCode(200))
	fmt.Println("404:", GetDescByCode(404))

	// Output:
	// 200: OK
	// 404: Not Found
}

func ExampleIsURL() {
	url1 := "https://domain.com"
	url2 := "httpj://domain.com"

	fmt.Printf("%s: %t\n", url1, IsURL(url1))
	fmt.Printf("%s: %t\n", url2, IsURL(url2))

	// Output:
	// https://domain.com: true
	// httpj://domain.com: false
}

func ExampleIsHTTP() {
	url1 := "https://domain.com"
	url2 := "http://domain.com"

	fmt.Printf("%s: %t\n", url1, IsHTTP(url1))
	fmt.Printf("%s: %t\n", url2, IsHTTP(url2))

	// Output:
	// https://domain.com: false
	// http://domain.com: true
}

func ExampleIsHTTPS() {
	url1 := "https://domain.com"
	url2 := "http://domain.com"

	fmt.Printf("%s: %t\n", url1, IsHTTPS(url1))
	fmt.Printf("%s: %t\n", url2, IsHTTPS(url2))

	// Output:
	// https://domain.com: true
	// http://domain.com: false
}

func ExampleIsFTP() {
	url1 := "ftp://domain.com"
	url2 := "http://domain.com"

	fmt.Printf("%s: %t\n", url1, IsFTP(url1))
	fmt.Printf("%s: %t\n", url2, IsFTP(url2))

	// Output:
	// ftp://domain.com: true
	// http://domain.com: false
}

func ExampleGetPortByScheme() {
	fmt.Println(GetPortByScheme("https"))
	// Output: 443
}
