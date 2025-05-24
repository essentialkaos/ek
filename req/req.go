// Package req simplify working with an HTTP requests
package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
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

// Status codes
const (
	STATUS_CONTINUE                        = 100 // RFC 7231, 6.2.1
	STATUS_SWITCHING_PROTOCOLS             = 101 // RFC 7231, 6.2.2
	STATUS_PROCESSING                      = 102 // RFC 2518, 10.1
	STATUS_OK                              = 200 // RFC 7231, 6.3.1
	STATUS_CREATED                         = 201 // RFC 7231, 6.3.2
	STATUS_ACCEPTED                        = 202 // RFC 7231, 6.3.3
	STATUS_NON_AUTHORITATIVE_INFO          = 203 // RFC 7231, 6.3.4
	STATUS_NO_CONTENT                      = 204 // RFC 7231, 6.3.5
	STATUS_RESET_CONTENT                   = 205 // RFC 7231, 6.3.6
	STATUS_PARTIAL_CONTENT                 = 206 // RFC 7233, 4.1
	STATUS_MULTI_STATUS                    = 207 // RFC 4918, 11.1
	STATUS_ALREADY_REPORTED                = 208 // RFC 5842, 7.1
	STATUS_IMUSED                          = 226 // RFC 3229, 10.4.1
	STATUS_MULTIPLE_CHOICES                = 300 // RFC 7231, 6.4.1
	STATUS_MOVED_PERMANENTLY               = 301 // RFC 7231, 6.4.2
	STATUS_FOUND                           = 302 // RFC 7231, 6.4.3
	STATUS_SEE_OTHER                       = 303 // RFC 7231, 6.4.4
	STATUS_NOT_MODIFIED                    = 304 // RFC 7232, 4.1
	STATUS_USE_PROXY                       = 305 // RFC 7231, 6.4.5
	STATUS_TEMPORARY_REDIRECT              = 307 // RFC 7231, 6.4.7
	STATUS_PERMANENT_REDIRECT              = 308 // RFC 7538, 3
	STATUS_BAD_REQUEST                     = 400 // RFC 7231, 6.5.1
	STATUS_UNAUTHORIZED                    = 401 // RFC 7235, 3.1
	STATUS_PAYMENT_REQUIRED                = 402 // RFC 7231, 6.5.2
	STATUS_FORBIDDEN                       = 403 // RFC 7231, 6.5.3
	STATUS_NOT_FOUND                       = 404 // RFC 7231, 6.5.4
	STATUS_METHOD_NOT_ALLOWED              = 405 // RFC 7231, 6.5.5
	STATUS_NOT_ACCEPTABLE                  = 406 // RFC 7231, 6.5.6
	STATUS_PROXY_AUTH_REQUIRED             = 407 // RFC 7235, 3.2
	STATUS_REQUEST_TIMEOUT                 = 408 // RFC 7231, 6.5.7
	STATUS_CONFLICT                        = 409 // RFC 7231, 6.5.8
	STATUS_GONE                            = 410 // RFC 7231, 6.5.9
	STATUS_LENGTH_REQUIRED                 = 411 // RFC 7231, 6.5.10
	STATUS_PRECONDITION_FAILED             = 412 // RFC 7232, 4.2
	STATUS_REQUEST_ENTITY_TOO_LARGE        = 413 // RFC 7231, 6.5.11
	STATUS_REQUEST_URITOO_LONG             = 414 // RFC 7231, 6.5.12
	STATUS_UNSUPPORTED_MEDIA_TYPE          = 415 // RFC 7231, 6.5.13
	STATUS_REQUESTED_RANGE_NOT_SATISFIABLE = 416 // RFC 7233, 4.4
	STATUS_EXPECTATION_FAILED              = 417 // RFC 7231, 6.5.14
	STATUS_TEAPOT                          = 418 // RFC 7168, 2.3.3
	STATUS_UNPROCESSABLE_ENTITY            = 422 // RFC 4918, 11.2
	STATUS_LOCKED                          = 423 // RFC 4918, 11.3
	STATUS_FAILED_DEPENDENCY               = 424 // RFC 4918, 11.4
	STATUS_UPGRADE_REQUIRED                = 426 // RFC 7231, 6.5.15
	STATUS_PRECONDITION_REQUIRED           = 428 // RFC 6585, 3
	STATUS_TOO_MANY_REQUESTS               = 429 // RFC 6585, 4
	STATUS_REQUEST_HEADER_FIELDS_TOO_LARGE = 431 // RFC 6585, 5
	STATUS_UNAVAILABLE_FOR_LEGAL_REASONS   = 451 // RFC 7725, 3
	STATUS_INTERNAL_SERVER_ERROR           = 500 // RFC 7231, 6.6.1
	STATUS_NOT_IMPLEMENTED                 = 501 // RFC 7231, 6.6.2
	STATUS_BAD_GATEWAY                     = 502 // RFC 7231, 6.6.3
	STATUS_SERVICE_UNAVAILABLE             = 503 // RFC 7231, 6.6.4
	STATUS_GATEWAY_TIMEOUT                 = 504 // RFC 7231, 6.6.5
	STATUS_HTTPVERSION_NOT_SUPPORTED       = 505 // RFC 7231, 6.6.6
	STATUS_VARIANT_ALSO_NEGOTIATES         = 506 // RFC 2295, 8.1
	STATUS_INSUFFICIENT_STORAGE            = 507 // RFC 4918, 11.5
	STATUS_LOOP_DETECTED                   = 508 // RFC 5842, 7.2
	STATUS_NOT_EXTENDED                    = 510 // RFC 2774, 7
	STATUS_NETWORK_AUTHENTICATION_REQUIRED = 511 // RFC 6585, 6
)

// USER_AGENT is default user agent
const USER_AGENT = "ek.go"

// ////////////////////////////////////////////////////////////////////////////////// //

// Request is basic struct
type Request struct {
	Method         string        // Request method
	URL            string        // Request URL
	Query          Query         // Map with query params
	Auth           Auth          // Authentication data
	ProxyAuth      Auth          // Proxy authentication data
	Body           any           // Request body
	Headers        Headers       // Map with headers
	ContentType    string        // Content type header
	Accept         string        // Accept header
	Timeout        time.Duration // Request timeout
	AutoDiscard    bool          // Automatically discard all responses with status code > 299
	FollowRedirect bool          // Follow redirect
	Close          bool          // Close indicates whether to close the connection after sending request
}

// Response is struct contains response data and properties
type Response struct {
	*http.Response
	URL string
}

// Engine is request engine
type Engine struct {
	UserAgent string // UserAgent is default user-agent used for all requests

	Dialer    *net.Dialer     // Dialer is default dialer struct
	Transport *http.Transport // Transport is default transport struct
	Client    *http.Client    // Client is default client struct

	limiter        *Limiter // Request limiter
	dialTimeout    float64  // dialTimeout is dial timeout in seconds
	requestTimeout float64  // requestTimeout is request timeout in seconds

	initialized bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilEngine is returned if engine struct is nil
	ErrNilEngine = fmt.Errorf("Engine is nil")

	// ErrNilClient is returned if client struct is nil
	ErrNilClient = fmt.Errorf("Engine.Client is nil")

	// ErrNilTransport is returned if transport is nil
	ErrNilTransport = fmt.Errorf("Engine.Transport is nil")

	// ErrNilDialer is returned if dialer is nil
	ErrNilDialer = fmt.Errorf("Engine.Dialer is nil")

	// ErrNilResponse is returned if response is nil
	ErrNilResponse = fmt.Errorf("Response is nil")

	// ErrEmptyBody is returned if response body has no data
	ErrEmptyBody = fmt.Errorf("Response body is empty")

	// ErrEmptyURL is returned if given URL is empty
	ErrEmptyURL = fmt.Errorf("URL property can't be empty and must be set")

	// ErrUnsupportedScheme is returned if given URL contains unsupported scheme
	ErrUnsupportedScheme = fmt.Errorf("Unsupported scheme in URL")
)

// Global is global engine used by default for Request.Do, Request.Get, Request.Post,
// Request.Put, Request.Patch, Request.Head and Request.Delete methods
var Global = &Engine{
	dialTimeout: 10.0,
}

// ////////////////////////////////////////////////////////////////////////////////// //

var ioCopyFunc = io.Copy
var useFakeFormGenerator = false

// ////////////////////////////////////////////////////////////////////////////////// //

// SetUserAgent sets user agent based on app name and version for global engine
func SetUserAgent(app, version string, subs ...string) {
	Global.SetUserAgent(app, version, subs...)
}

// SetDialTimeout sets dial timeout for global engine
func SetDialTimeout(timeout float64) {
	Global.SetDialTimeout(timeout)
}

// SetRequestTimeout sets request timeout for global engine
func SetRequestTimeout(timeout float64) {
	Global.SetRequestTimeout(timeout)
}

// SetLimit sets a hard limit on the number of requests per second (useful for
// working with APIs)
func SetLimit(rps float64) {
	Global.SetLimit(rps)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Init initializes engine
func (e *Engine) Init() *Engine {
	if e.initialized {
		return e
	}

	if e.Dialer == nil {
		e.Dialer = &net.Dialer{}
	}

	if e.Transport == nil {
		e.Transport = &http.Transport{
			DialContext: e.Dialer.DialContext,
			Proxy:       http.ProxyFromEnvironment,
		}
	} else {
		e.Transport.DialContext = e.Dialer.DialContext
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
		e.SetUserAgent(USER_AGENT, "13")
	}

	e.dialTimeout = 0
	e.requestTimeout = 0

	e.initialized = true

	return e
}

// Do sends request and process response
func (e *Engine) Do(r Request) (*Response, error) {
	return e.doRequest(r, "")
}

// Get sends GET request and process response
func (e *Engine) Get(r Request) (*Response, error) {
	return e.doRequest(r, GET)
}

// Post sends POST request and process response
func (e *Engine) Post(r Request) (*Response, error) {
	return e.doRequest(r, POST)
}

// Put sends PUT request and process response
func (e *Engine) Put(r Request) (*Response, error) {
	return e.doRequest(r, PUT)
}

// Head sends HEAD request and process response
func (e *Engine) Head(r Request) (*Response, error) {
	return e.doRequest(r, HEAD)
}

// Patch sends PATCH request and process response
func (e *Engine) Patch(r Request) (*Response, error) {
	return e.doRequest(r, PATCH)
}

// Delete sends DELETE request and process response
func (e *Engine) Delete(r Request) (*Response, error) {
	return e.doRequest(r, DELETE)
}

// PostFile sends multipart POST request with file data
func (e *Engine) PostFile(r Request, file, fieldName string, extraFields map[string]string) (*Response, error) {
	err := configureMultipartRequest(&r, file, fieldName, extraFields)

	if err != nil {
		return nil, err
	}

	return e.doRequest(r, POST)
}

// SetUserAgent sets user agent based on app name and version
func (e *Engine) SetUserAgent(app, version string, subs ...string) {
	if e == nil {
		return
	}

	e.UserAgent = fmt.Sprintf(
		"%s/%s (go; %s; %s-%s)",
		app, version, runtime.Version(),
		runtime.GOARCH, runtime.GOOS,
	)

	if len(subs) != 0 {
		e.UserAgent += " " + strings.Join(subs, " ")
	}
}

// SetDialTimeout sets dial timeout
func (e *Engine) SetDialTimeout(timeout float64) {
	if e == nil {
		return
	}

	if timeout > 0 {
		if e.Dialer == nil {
			e.dialTimeout = timeout
		} else {
			e.Dialer.Timeout = time.Duration(timeout * float64(time.Second))
		}
	}
}

// SetRequestTimeout sets request timeout
func (e *Engine) SetRequestTimeout(timeout float64) {
	if e == nil {
		return
	}

	if timeout > 0 {
		if e.Dialer == nil {
			e.requestTimeout = timeout
		} else {
			e.Client.Timeout = time.Duration(timeout * float64(time.Second))
		}
	}
}

// SetLimit sets a hard limit on the number of requests per second (useful for
// working with APIs)
func (e *Engine) SetLimit(rps float64) {
	if e == nil {
		return
	}

	e.limiter = NewLimiter(rps)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Do sends request and process response
func (r Request) Do() (*Response, error) {
	return Global.doRequest(r, "")
}

// Get sends GET request and process response
func (r Request) Get() (*Response, error) {
	return Global.Get(r)
}

// Post sends POST request and process response
func (r Request) Post() (*Response, error) {
	return Global.Post(r)
}

// Put sends PUT request and process response
func (r Request) Put() (*Response, error) {
	return Global.Put(r)
}

// Head sends HEAD request and process response
func (r Request) Head() (*Response, error) {
	return Global.Head(r)
}

// Patch sends PATCH request and process response
func (r Request) Patch() (*Response, error) {
	return Global.Patch(r)
}

// Delete sends DELETE request and process response
func (r Request) Delete() (*Response, error) {
	return Global.Delete(r)
}

// PostFile sends multipart POST request with file data
func (r Request) PostFile(file, fieldName string, extraFields map[string]string) (*Response, error) {
	return Global.PostFile(r, file, fieldName, extraFields)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Discard reads response body for closing connection
func (r *Response) Discard() {
	if r == nil || r.Response == nil || r.Body == nil {
		return
	}

	io.Copy(io.Discard, r.Body)
}

// JSON decodes json encoded body
func (r *Response) JSON(v any) error {
	switch {
	case r == nil || r.Response == nil:
		return ErrNilResponse
	case r.Body == nil:
		return ErrEmptyBody
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(v)
}

// Bytes reads response body as byte slice
func (r *Response) Bytes() ([]byte, error) {
	switch {
	case r == nil || r.Response == nil:
		return nil, ErrNilResponse
	case r.Body == nil:
		return nil, ErrEmptyBody
	}

	defer r.Body.Close()

	return io.ReadAll(r.Body)
}

// Save saves response data into a file
func (r *Response) Save(filename string, mode os.FileMode) error {
	switch {
	case r == nil || r.Response == nil:
		return ErrNilResponse
	case r.Body == nil:
		return ErrEmptyBody
	}

	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)

	if err != nil {
		return fmt.Errorf("Can't open file to write: %v", err)
	}

	defer fd.Close()
	defer r.Body.Close()

	_, err = io.Copy(fd, r.Body)

	return err
}

// String reads response body as string
func (r *Response) String() string {
	if r == nil || r.Response == nil {
		return ""
	}

	data, err := r.Bytes()

	if err != nil {
		return ""
	}

	return string(data)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (e *Engine) doRequest(r Request, method string) (*Response, error) {
	// Lazy engine initialization
	if e != nil && !e.initialized {
		e.Init()
	}

	err := checkEngine(e)

	if err != nil {
		return nil, err
	}

	err = checkRequest(r)

	if err != nil {
		return nil, err
	}

	if method != "" {
		r.Method = method
	}

	if r.Method == "" {
		r.Method = GET
	}

	if len(r.Query) != 0 {
		r.URL += "?" + r.Query.Encode()
	}

	bodyReader, contentType, err := getBodyReader(r.Body)

	if err != nil {
		return nil, fmt.Errorf("Can't encode request body: %w", err)
	}

	if r.ContentType == "" {
		r.ContentType = contentType
	}

	req, cancel, err := createRequest(e, r, bodyReader)

	if err != nil {
		return nil, err
	}

	if e.limiter != nil {
		e.limiter.Wait()
	}

	resp, err := e.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Can't send request: %w", err)
	}

	result := &Response{resp, r.URL}

	if resp.StatusCode > 299 && r.AutoDiscard {
		result.Discard()

		if cancel != nil {
			cancel()
		}
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// checkRequest checks request parameters
func checkRequest(r Request) error {
	if r.URL == "" {
		return ErrEmptyURL
	}

	if !isURL(r.URL) {
		return ErrUnsupportedScheme
	}

	return nil
}

// checkEngine checks request engine
func checkEngine(e *Engine) error {
	if e == nil {
		return ErrNilEngine
	}

	if e.Dialer == nil {
		return ErrNilDialer
	}

	if e.Transport == nil {
		return ErrNilTransport
	}

	if e.Client == nil {
		return ErrNilClient
	}

	return nil
}

// createRequest creates new http.Request from Request
func createRequest(e *Engine, r Request, bodyReader io.Reader) (*http.Request, context.CancelFunc, error) {
	var err error
	var req *http.Request
	var cancel context.CancelFunc

	if r.Timeout != 0 {
		var ctx context.Context
		ctx, cancel = context.WithTimeout(context.TODO(), r.Timeout)
		req, err = http.NewRequestWithContext(ctx, r.Method, r.URL, bodyReader)
	} else {
		req, err = http.NewRequest(r.Method, r.URL, bodyReader)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("Can't create request: %w", err)
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

	if r.Auth != nil {
		r.Auth.Apply(req, "Authorization")
	}

	if r.ProxyAuth != nil {
		r.ProxyAuth.Apply(req, "Proxy-Authorization")
	}

	if r.Close {
		req.Close = true
	}

	return req, cancel, nil
}

// configureMultipartRequest configures request for sending multipart data
func configureMultipartRequest(r *Request, file, fieldName string, extraFields map[string]string) error {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return err
	}

	defer fd.Close()

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	part, err := createFormFile(w, fieldName, file)

	if err != nil {
		return err
	}

	_, err = ioCopyFunc(part, fd)

	if err != nil {
		return err
	}

	for k, v := range extraFields {
		w.WriteField(k, v)
	}

	w.Close()

	r.ContentType = w.FormDataContentType()
	r.Body = buf

	return nil
}

// createFormFile creates request from file
func createFormFile(w *multipart.Writer, fieldName, file string) (io.Writer, error) {
	if useFakeFormGenerator {
		return nil, fmt.Errorf("")
	}

	return w.CreateFormFile(fieldName, filepath.Base(file))
}

// getBodyReader returns reader for request body
func getBodyReader(body any) (io.Reader, string, error) {
	switch u := body.(type) {
	case nil:
		return nil, "", nil
	case string:
		return strings.NewReader(u), CONTENT_TYPE_PLAIN, nil
	case io.Reader:
		return u, CONTENT_TYPE_OCTET_STREAM, nil
	case []byte:
		return bytes.NewReader(u), CONTENT_TYPE_OCTET_STREAM, nil
	}

	jsonBody, err := json.MarshalIndent(body, "", "  ")

	if err == nil {
		return bytes.NewReader(jsonBody), CONTENT_TYPE_JSON, nil
	}

	return nil, "", err
}

// isURL returns true if URL is valid
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
