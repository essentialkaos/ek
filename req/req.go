// Package req for working with http request
package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
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
	"runtime"
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

// Content types
const (
	CONTENT_TYPE_ATOM         = "application/atom+xml"
	CONTENT_TYPE_EDI          = "application/EDI-X12"
	CONTENT_TYPE_EDIFACT      = "application/EDIFACT"
	CONTENT_TYPE_JSON         = "application/json"
	CONTENT_TYPE_JAVASCRIPT   = "application/javascript"
	CONTENT_TYPE_OCTET_STREAM = "application/octet-stream"
	CONTENT_TYPE_PDF          = "application/pdf"
	CONTENT_TYPE_POSTSCRIPT   = "application/postscript"
	CONTENT_TYPE_SOAP         = "application/soap+xml"
	CONTENT_TYPE_WOFF         = "application/font-woff"
	CONTENT_TYPE_XHTML        = "application/xhtml+xml"
	CONTENT_TYPE_DTD          = "application/xml-dtd"
	CONTENT_TYPE_XOP          = "application/xop+xml"
	CONTENT_TYPE_ZIP          = "application/zip"
	CONTENT_TYPE_GZIP         = "application/gzip"
	CONTENT_TYPE_BITTORRENT   = "application/x-bittorrent"
	CONTENT_TYPE_TEX          = "application/x-tex"
	CONTENT_TYPE_BASIC        = "audio/basic"
	CONTENT_TYPE_L24          = "audio/L24"
	CONTENT_TYPE_MP4_AUDIO    = "audio/mp4"
	CONTENT_TYPE_AAC          = "audio/aac"
	CONTENT_TYPE_MPEG_AUDIO   = "audio/mpeg"
	CONTENT_TYPE_OGG_AUDIO    = "audio/ogg"
	CONTENT_TYPE_VORBIS       = "audio/vorbis"
	CONTENT_TYPE_WMA          = "audio/x-ms-wma"
	CONTENT_TYPE_WAX          = "audio/x-ms-wax"
	CONTENT_TYPE_REALAUDIO    = "audio/vnd.rn-realaudio"
	CONTENT_TYPE_WAV          = "audio/vnd.wave"
	CONTENT_TYPE_WEBM_AUDIO   = "audio/webm"
	CONTENT_TYPE_GIF          = "image/gif"
	CONTENT_TYPE_JPEG         = "image/jpeg"
	CONTENT_TYPE_PJPEG        = "image/pjpeg"
	CONTENT_TYPE_PNG          = "image/png"
	CONTENT_TYPE_SVG          = "image/svg+xml"
	CONTENT_TYPE_TIFF         = "image/tiff"
	CONTENT_TYPE_ICON         = "image/vnd.microsoft.icon"
	CONTENT_TYPE_WBMP         = "image/vnd.wap.wbmp"
	CONTENT_TYPE_HTTP         = "message/http"
	CONTENT_TYPE_IMDN         = "message/imdn+xml"
	CONTENT_TYPE_PARTIAL      = "message/partial"
	CONTENT_TYPE_RFC822       = "message/rfc822"
	CONTENT_TYPE_EXAMPLE      = "model/example"
	CONTENT_TYPE_IGES         = "model/iges"
	CONTENT_TYPE_MESH         = "model/mesh"
	CONTENT_TYPE_VRML         = "model/vrml"
	CONTENT_TYPE_MIXED        = "multipart/mixed"
	CONTENT_TYPE_ALTERNATIVE  = "multipart/alternative"
	CONTENT_TYPE_RELATED      = "multipart/related"
	CONTENT_TYPE_FORM_DATA    = "multipart/form-data"
	CONTENT_TYPE_SIGNED       = "multipart/signed"
	CONTENT_TYPE_ENCRYPTED    = "multipart/encrypted"
	CONTENT_TYPE_CSS          = "text/css"
	CONTENT_TYPE_CSV          = "text/csv"
	CONTENT_TYPE_HTML         = "text/html"
	CONTENT_TYPE_PLAIN        = "text/plain"
	CONTENT_TYPE_PHP          = "text/php"
	CONTENT_TYPE_XML          = "text/xml"
	CONTENT_TYPE_MPEG_VIDEO   = "video/mpeg"
	CONTENT_TYPE_MP4_VIDEO    = "video/mp4"
	CONTENT_TYPE_OGG_VIDEO    = "video/ogg"
	CONTENT_TYPE_QUICKTIME    = "video/quicktime"
	CONTENT_TYPE_WEBM_VIDEO   = "video/webm"
	CONTENT_TYPE_WMV          = "video/x-ms-wmv"
	CONTENT_TYPE_FLV          = "video/x-flv"
	CONTENT_TYPE_3GPP         = "video/3gpp"
	CONTENT_TYPE_3GPP2        = "video/3gpp2"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Query is a map[string]interface{} used for query
type Query map[string]interface{}

// Headers is a map[string]string used for headers
type Headers map[string]string

// Request is basic struct
type Request struct {
	Method            string      // Request method
	URL               string      // Request url
	Query             Query       // Map with query params
	Body              interface{} // Request body
	Headers           Headers     // Map with headers
	ContentType       string      // Content type header
	Accept            string      // Accept header
	BasicAuthUsername string      // Basic auth username
	BasicAuthPassword string      // Basic auth password
	AutoDiscard       bool        // Automatically discard all responses with status code != 200
	FollowRedirect    bool        // Follow redirect
	Close             bool        // Close indicates whether to close the connection after sending request
}

// Response struct contains response data and properties
type Response struct {
	*http.Response
	URL string
}

// RequestError error struct
type RequestError struct {
	class int
	desc  string
}

// Engine is request engine
type Engine struct {
	UserAgent string // UserAgent is default user-agent used for all requests

	Dialer    *net.Dialer     // Dialer default dialer struct
	Transport *http.Transport // Transport is default transport struct
	Client    *http.Client    // Client default client struct

	dialTimeout    float64 // dialTimeout is dial timeout in seconds
	requestTimeout float64 // requestTimeout is request timeout in seconds

	initialized bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	ErrEngineIsNil       = RequestError{ERROR_CREATE_REQUEST, "Engine is nil"}
	ErrClientIsNil       = RequestError{ERROR_CREATE_REQUEST, "Engine.Client is nil"}
	ErrTransportIsNil    = RequestError{ERROR_CREATE_REQUEST, "Engine.Transport is nil"}
	ErrDialerIsNil       = RequestError{ERROR_CREATE_REQUEST, "Engine.Dialer is nil"}
	ErrEmptyURL          = RequestError{ERROR_CREATE_REQUEST, "URL property can't be empty and must be set"}
	ErrUnsupportedScheme = RequestError{ERROR_CREATE_REQUEST, "Unsupported scheme in URL"}
)

// Global is global engine used by default for Request.Do, Request.Get, Request.Post,
// Request.Put, Request.Patch, Request.Head and Request.Delete methods
var Global *Engine = &Engine{
	dialTimeout: 10.0,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// SetUserAgent set user agent based on app name and version for global engine
func SetUserAgent(app, version string, subs ...string) {
	Global.SetUserAgent(app, version, subs...)
}

// SetDialTimeout set dial timeout for global engine
func SetDialTimeout(timeout float64) {
	Global.SetDialTimeout(timeout)
}

// SetRequestTimeout set request timeout for global engine
func SetRequestTimeout(timeout float64) {
	Global.SetRequestTimeout(timeout)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Do send request and process response
func (e *Engine) Do(r Request) (*Response, error) {
	return e.doRequest(r, "")
}

// Get send GET request and process response
func (e *Engine) Get(r Request) (*Response, error) {
	return e.doRequest(r, GET)
}

// Post send POST request and process response
func (e *Engine) Post(r Request) (*Response, error) {
	return e.doRequest(r, POST)
}

// Put send PUT request and process response
func (e *Engine) Put(r Request) (*Response, error) {
	return e.doRequest(r, PUT)
}

// Head send HEAD request and process response
func (e *Engine) Head(r Request) (*Response, error) {
	return e.doRequest(r, HEAD)
}

// Patch send PATCH request and process response
func (e *Engine) Patch(r Request) (*Response, error) {
	return e.doRequest(r, PATCH)
}

// Delete send DELETE request and process response
func (e *Engine) Delete(r Request) (*Response, error) {
	return e.doRequest(r, DELETE)
}

// SetUserAgent set user agent based on app name and version
func (e *Engine) SetUserAgent(app, version string, subs ...string) {
	if e != nil {
		e.UserAgent = fmt.Sprintf(
			"%s/%s (go; %s; %s-%s)",
			app, version, runtime.Version(),
			runtime.GOARCH, runtime.GOOS,
		)

		if len(subs) != 0 {
			e.UserAgent += " " + strings.Join(subs, " ")
		}
	}
}

// SetDialTimeout set dial timeout
func (e *Engine) SetDialTimeout(timeout float64) {
	if e != nil && timeout > 0 {
		if e.Dialer == nil {
			e.dialTimeout = timeout
		} else {
			e.Dialer.Timeout = time.Duration(timeout * float64(time.Second))
		}
	}
}

// SetRequestTimeout set request timeout
func (e *Engine) SetRequestTimeout(timeout float64) {
	if e != nil && timeout > 0 {
		if e.Dialer == nil {
			e.requestTimeout = timeout
		} else {
			e.Client.Timeout = time.Duration(timeout * float64(time.Second))
		}
	}
}

// Do send request and process response
func (r Request) Do() (*Response, error) {
	return Global.doRequest(r, "")
}

// Get send GET request and process response
func (r Request) Get() (*Response, error) {
	return Global.doRequest(r, GET)
}

// Post send POST request and process response
func (r Request) Post() (*Response, error) {
	return Global.doRequest(r, POST)
}

// Put send PUT request and process response
func (r Request) Put() (*Response, error) {
	return Global.doRequest(r, PUT)
}

// Head send HEAD request and process response
func (r Request) Head() (*Response, error) {
	return Global.doRequest(r, HEAD)
}

// Patch send PATCH request and process response
func (r Request) Patch() (*Response, error) {
	return Global.doRequest(r, PATCH)
}

// Delete send DELETE request and process response
func (r Request) Delete() (*Response, error) {
	return Global.doRequest(r, DELETE)
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
	switch e.class {
	case ERROR_BODY_ENCODE:
		return fmt.Sprintf("Can't encode request body (%s)", e.desc)
	case ERROR_SEND_REQUEST:
		return fmt.Sprintf("Can't send request (%s)", e.desc)
	default:
		return fmt.Sprintf("Can't create request struct (%s)", e.desc)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (e *Engine) doRequest(r Request, method string) (*Response, error) {
	// Lazy engine initialization
	if e != nil && !e.initialized {
		initEngine(e)
	}

	err := checkRequest(r)

	if err != nil {
		return nil, err
	}

	err = checkEngine(e)

	if err != nil {
		return nil, err
	}

	if method != "" {
		r.Method = method
	}

	if r.Method == "" {
		r.Method = GET
	}

	if r.Query != nil && len(r.Query) != 0 {
		query, err := encodeQuery(r.Query)

		if err != nil {
			return nil, err
		}

		r.URL += "?" + query
	}

	bodyReader, err := getBodyReader(r.Body)

	if err != nil {
		return nil, RequestError{ERROR_BODY_ENCODE, err.Error()}
	}

	req, err := createRequest(e, r, bodyReader)

	if err != nil {
		return nil, err
	}

	resp, err := e.Client.Do(req)

	if err != nil {
		return nil, RequestError{ERROR_SEND_REQUEST, err.Error()}
	}

	result := &Response{resp, r.URL}

	if resp.StatusCode != 200 && r.AutoDiscard {
		result.Discard()
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func initEngine(e *Engine) {
	if e.Dialer == nil {
		e.Dialer = &net.Dialer{}
	}

	if e.Transport == nil {
		e.Transport = &http.Transport{
			Dial:  e.Dialer.Dial,
			Proxy: http.ProxyFromEnvironment,
		}
	} else {
		e.Transport.Dial = e.Dialer.Dial
	}

	if e.Client == nil {
		e.Client = &http.Client{
			Transport: e.Transport,
		}
	}

	if e.dialTimeout > 0 {
		e.SetDialTimeout(e.dialTimeout)
	}

	if e.requestTimeout > 0 {
		e.SetRequestTimeout(e.requestTimeout)
	}

	if e.UserAgent == "" {
		e.SetUserAgent("goek-http-client", "5.x")
	}

	e.dialTimeout = 0
	e.requestTimeout = 0

	e.initialized = true
}

func checkRequest(r Request) error {
	if r.URL == "" {
		return ErrEmptyURL
	}

	if !isURL(r.URL) {
		return ErrUnsupportedScheme
	}

	return nil
}

func checkEngine(e *Engine) error {
	if e == nil {
		return ErrEngineIsNil
	}

	if e.Dialer == nil {
		return ErrDialerIsNil
	}

	if e.Transport == nil {
		return ErrTransportIsNil
	}

	if e.Client == nil {
		return ErrClientIsNil
	}

	return nil
}

func createRequest(e *Engine, r Request, bodyReader io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.URL, bodyReader)

	if err != nil {
		return nil, RequestError{ERROR_CREATE_REQUEST, err.Error()}
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
		req.Header.Add("Accept", r.Accept)
	}

	if e.UserAgent != "" {
		req.Header.Add("User-Agent", e.UserAgent)
	}

	if r.BasicAuthUsername != "" && r.BasicAuthPassword != "" {
		req.SetBasicAuth(r.BasicAuthUsername, r.BasicAuthPassword)
	}

	if r.Close {
		req.Close = true
	}

	return req, nil
}

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

func encodeQuery(query Query) (string, error) {
	var result string

	for k, v := range query {
		switch v.(type) {
		case string:
			if v == "" {
				result += k + "&"
			} else {
				result += k + "=" + url.QueryEscape(v.(string)) + "&"
			}
		case nil:
			result += k + "&"
		case bool:
			result += k + "=" + fmt.Sprintf("%t", v.(bool)) + "&"
		case int:
			result += k + "=" + fmt.Sprintf("%d", v.(int)) + "&"
		case int8:
			result += k + "=" + fmt.Sprintf("%d", v.(int8)) + "&"
		case int16:
			result += k + "=" + fmt.Sprintf("%d", v.(int16)) + "&"
		case int32:
			result += k + "=" + fmt.Sprintf("%d", v.(int32)) + "&"
		case int64:
			result += k + "=" + fmt.Sprintf("%d", v.(int64)) + "&"
		case uint:
			result += k + "=" + fmt.Sprintf("%d", v.(uint)) + "&"
		case uint8:
			result += k + "=" + fmt.Sprintf("%d", v.(uint8)) + "&"
		case uint16:
			result += k + "=" + fmt.Sprintf("%d", v.(uint16)) + "&"
		case uint32:
			result += k + "=" + fmt.Sprintf("%d", v.(uint32)) + "&"
		case uint64:
			result += k + "=" + fmt.Sprintf("%d", v.(uint64)) + "&"
		case float32:
			result += k + "=" + fmt.Sprintf("%g", v.(float32)) + "&"
		case float64:
			result += k + "=" + fmt.Sprintf("%g", v.(float64)) + "&"
		default:
			return "", fmt.Errorf("Can't encode query - unsupported value type")
		}
	}

	return result[:len(result)-1], nil
}

func isURL(url string) bool {
	switch {
	case len(url) < 10:
		return false
	case url[0:7] == "http://":
		return true
	case url[0:8] == "https://":
		return true
	case url[0:6] == "ftp://":
		return true
	}

	return false
}
