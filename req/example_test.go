package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRequest_Do() {

	// Simple request
	resp, err := Request{
		Method: GET,
		URL:    "https://my.domain.com",
		Query: Query{
			"name":     "Bob",
			"id":       120,
			"progress": 12.34,
		},
		Headers: Headers{
			"My-Suppa-Header": "Test",
		},
		ContentType: CONTENT_TYPE_JSON,
		UserAgent:   "My Client 1.0.0",
	}.Do()

	if err != nil {
		return
	}

	// print status code
	fmt.Printf("Status code: %d\n", resp.StatusCode)
}
