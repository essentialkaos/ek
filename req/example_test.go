package req

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

func ExampleRequest_Do() {
	var response = struct {
		User    string  `json:"user"`
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{}

	// Configure global enagine
	SetUserAgent("my-supper-app", "1.0")
	SetDialTimeout(30.0)
	SetRequestTimeout(30.0)

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
	}.Do()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// print status code
	fmt.Printf("Status code: %d\n", resp.StatusCode)

	// decode json encoded response
	err = resp.JSON(response)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// print response data
	fmt.Printf(
		"User: %s ID: %d Balance: %f\n",
		response.User, response.ID, response.Balance,
	)
}

func ExampleRequest_Get() {
	var response = struct {
		User    string  `json:"user"`
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{}

	resp, err := Request{URL: "https://my.domain.com"}.Get()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// decode json encoded response
	err = resp.JSON(response)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// print response data
	fmt.Printf(
		"User: %s ID: %d Balance: %f\n",
		response.User, response.ID, response.Balance,
	)
}

func ExampleRequest_Post() {
	var request = struct {
		UserID int `json:"user_id"`
	}{
		UserID: 1234,
	}

	var response = struct {
		User    string  `json:"user"`
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{}

	// send post request with basic auth
	resp, err := Request{
		URL:               "https://my.domain.com",
		Body:              request,
		Accept:            CONTENT_TYPE_JSON,
		ContentType:       CONTENT_TYPE_JSON,
		BasicAuthUsername: "someuser",
		BasicAuthPassword: "somepass",
		AutoDiscard:       true,
	}.Post()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// decode json encoded response
	err = resp.JSON(response)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// print response data
	fmt.Printf(
		"User: %s ID: %d Balance: %f\n",
		response.User, response.ID, response.Balance,
	)
}
