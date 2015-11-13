package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Error types
const (
	ERROR_BODY_ENCODE    = 1
	ERROR_CREATE_REQUEST = 2
	ERROR_SEND_REQUEST   = 3
)

// Request method
const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	HEAD   = "HEAD"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Request is basic struct
type Request struct {
	Method            string            // Request method
	URL               string            // Request url
	Query             map[string]string // Map with query params
	Body              interface{}       // Request body
	Headers           map[string]string // Map with headers
	ContentType       string            // Content type header
	Accept            string            // Accept header
	BasicAuthUsername string            // Basic auth username
	BasicAuthPassword string            // Basic auth password
	UserAgent         string            // User Agent string
	AutoDiscard       bool              // Automatically discard all responses with status code != 200
}

// Response struct contains response data and properties
type Response struct {
	*http.Response
	URL string
}

// RequestError error struct
type RequestError struct {
	desc string
	Type int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// UserAgent is default user-agent used for all requests
	UserAgent = ""

	// Dialer default dialer struct
	Dialer = &net.Dialer{Timeout: 1 * time.Second}

	// Transport is default transport struct
	Transport = &http.Transport{Dial: Dialer.Dial, Proxy: http.ProxyFromEnvironment}

	// Client default client struct
	Client = &http.Client{Transport: Transport}
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Query build query map from args
func Query(args ...string) map[string]string {
	return sliceToMap(args)
}

// Headers build headers map from args
func Headers(args ...string) map[string]string {
	return sliceToMap(args)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Do execute request
func (r Request) Do() (*Response, error) {
	if r.Method == "" {
		r.Method = GET
	}

	if r.Query != nil && len(r.Query) != 0 {
		r.URL += "?" + encodeQuery(r.Query)
	}

	bodyReader, err := getBodyReader(r.Body)

	if err != nil {
		return nil, RequestError{err.Error(), ERROR_BODY_ENCODE}
	}

	req, err := http.NewRequest(r.Method, r.URL, bodyReader)

	if err != nil {
		return nil, RequestError{err.Error(), ERROR_CREATE_REQUEST}
	}

	if r.Headers != nil && len(r.Headers) != 0 {
		for k, v := range r.Headers {
			req.Header.Add(k, v)
		}
	}

	if r.ContentType != "" {
		req.Header.Add("Content-Type", r.ContentType)
	}

	if r.Accept != "" {
		req.Header.Add("Accept", r.ContentType)
	}

	switch {
	case r.UserAgent != "":
		req.Header.Add("User-Agent", r.UserAgent)
	case UserAgent != "":
		req.Header.Add("User-Agent", UserAgent)
	}

	if r.BasicAuthUsername != "" && r.BasicAuthPassword != "" {
		req.SetBasicAuth(r.BasicAuthUsername, r.BasicAuthPassword)
	}

	resp, err := Client.Do(req)

	if err != nil {
		return nil, RequestError{err.Error(), ERROR_SEND_REQUEST}
	}

	result := &Response{resp, r.URL}

	if resp.StatusCode != 200 && r.AutoDiscard {
		result.Discard()
	}

	return result, nil
}

// Get execute as GET request
func (r Request) Get() (*Response, error) {
	r.Method = GET
	return r.Do()
}

// Get execute as POST request
func (r Request) Post() (*Response, error) {
	r.Method = POST
	return r.Do()
}

// Get execute as PUT request
func (r Request) Put() (*Response, error) {
	r.Method = PUT
	return r.Do()
}

// Get execute as HEAD request
func (r Request) Head() (*Response, error) {
	r.Method = HEAD
	return r.Do()
}

// Get execute as PATCH request
func (r Request) Patch() (*Response, error) {
	r.Method = PATCH
	return r.Do()
}

// Get execute as DELETE request
func (r Request) Delete() (*Response, error) {
	r.Method = DELETE
	return r.Do()
}

// Discard reads response body for closing connection
func (r *Response) Discard() {
	io.Copy(ioutil.Discard, r.Body)
}

// JSON decode json encoded body
func (r *Response) JSON(v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// String read response body as string
func (r *Response) String() string {
	result, _ := ioutil.ReadAll(r.Body)
	return string(result)
}

// Error show error message
func (e RequestError) Error() string {
	switch e.Type {
	case ERROR_BODY_ENCODE:
		return fmt.Sprintf("Can't encode request body (%s)", e.desc)
	case ERROR_CREATE_REQUEST:
		return fmt.Sprintf("Can't create request struct (%s)", e.desc)
	case ERROR_SEND_REQUEST:
		return fmt.Sprintf("Can't send request (%s)", e.desc)
	}

	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getBodyReader(body interface{}) (io.Reader, error) {
	switch body.(type) {
	case nil:
		return nil, nil
	case string:
		return strings.NewReader(body.(string)), nil
	case io.Reader:
		return body.(io.Reader), nil
	case []byte:
		return bytes.NewReader(body.([]byte)), nil
	default:
		jsonBody, err := json.MarshalIndent(body, "", "  ")

		if err == nil {
			return bytes.NewReader(jsonBody), nil
		}

		return nil, err
	}
}

func encodeQuery(query map[string]string) string {
	var result string

	for k, v := range query {
		switch v {
		case "":
			result += k + "&"
		default:
			result += k + "=" + url.QueryEscape(v) + "&"
		}
	}

	return result[:len(result)-1]
}

func sliceToMap(s []string) map[string]string {
	var result = make(map[string]string)

	if len(s)%2 != 0 {
		return nil
	}

	for i := 0; i < len(s)-1; i += 2 {
		result[s[i]] = s[i+1]
	}

	return result
}
