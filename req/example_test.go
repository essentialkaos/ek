package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleRequest_Do() {
	var response = struct {
		User    string  `json:"user"`
		ID      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{}

	// Configure global engine
	SetUserAgent("my-supper-app", "1.0")
	SetDialTimeout(30.0)
	SetRequestTimeout(30.0)
	SetLimit(15.0)

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

	// decode JSON encoded response
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

	// decode JSON encoded response
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

func ExampleRequest_PostFile() {
	extraFields := map[string]string{
		"user": "john",
		"desc": "My photo",
	}

	// send multipart request with image
	resp, err := Request{
		URL: "https://my.domain.com",
	}.PostFile("/tmp/image.jpg", "file", extraFields)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Can't upload file: %v\n", err)
		return
	}

	fmt.Println("File successfully uploaded!")
}

func ExampleNewRetrier() {
	r := NewRetrier()

	resp, err := r.Get(
		Request{URL: "https://my.domain.com"},
		Retry{Num: 5, Status: STATUS_OK, Pause: time.Second},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// print status code
	fmt.Printf("Status code: %d\n", resp.StatusCode)
}
