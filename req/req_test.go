package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
	"github.com/essentialkaos/ek/v13/hashutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_URL_GET          = "/get"
	_URL_POST         = "/post"
	_URL_POST_MULTI   = "/post-multi"
	_URL_PUT          = "/put"
	_URL_HEAD         = "/head"
	_URL_PATCH        = "/patch"
	_URL_DELETE       = "/delete"
	_URL_QUERY        = "/query"
	_URL_HEADERS      = "/headers"
	_URL_CONTENT_TYPE = "/content-type"
	_URL_ACCEPT       = "/accept"
	_URL_USER_AGENT   = "/user-agent"
	_URL_BASIC_AUTH   = "/basic-auth"
	_URL_STRING_RESP  = "/string-response"
	_URL_JSON_RESP    = "/json-response"
	_URL_DISCARD      = "/discard"
	_URL_TIMEOUT      = "/timeout"
	_URL_SAVE         = "/save"
)

const (
	_TEST_CONTENT_TYPE    = "application/json"
	_TEST_ACCEPT          = "application/vnd.example.api+json;version=2"
	_TEST_BASIC_AUTH_USER = "admin"
	_TEST_BASIC_AUTH_PASS = "password"
	_TEST_STRING_RESP     = "Test String Response"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type TestStringer struct{}
type TestPayload struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ReqSuite struct {
	url      string
	port     string
	listener net.Listener
}

type TestStruct struct {
	String  string `json:"string"`
	Integer int    `json:"integer"`
	Boolean bool   `json:"boolean"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ReqSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ReqSuite) SetUpSuite(c *C) {
	s.port = "30001"
	s.url = "http://127.0.0.1:" + s.port

	SetDialTimeout(60.0)
	SetRequestTimeout(60.0)
	SetLimit(1000.0)
	SetUserAgent("req-test", "5", "Test/5.1.1", "Magic/4.2.1")

	go runHTTPServer(s, c)

	time.Sleep(time.Second)
}

func (s *ReqSuite) TearDownSuite(c *C) {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *ReqSuite) TestMethodGet(c *C) {
	getResp, err := Request{URL: s.url + _URL_GET, Method: GET}.Do()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Global.Do(Request{URL: s.url + _URL_GET, Method: GET})

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Request{URL: s.url + _URL_GET}.Do()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Request{URL: s.url + _URL_GET}.Get()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Global.Get(Request{URL: s.url + _URL_GET})

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPost(c *C) {
	postResp, err := Request{URL: s.url + _URL_POST, Method: POST}.Do()

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)

	postResp, err = Global.Do(Request{URL: s.url + _URL_POST, Method: POST})

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)

	postResp, err = Request{URL: s.url + _URL_POST}.Post()

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)

	postResp, err = Global.Post(Request{URL: s.url + _URL_POST})

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPostFile(c *C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/testMultipart.bin"

	err := os.WriteFile(tmpFile, []byte(`DATA8913FIN`), 0644)
	c.Assert(err, IsNil)

	r := Request{URL: s.url + _URL_POST_MULTI, Method: POST}
	postResp, err := r.PostFile(tmpFile, "file", map[string]string{"abc": "123"})

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)

	_, err = r.PostFile(tmpDir+"/unknown", "file", map[string]string{"abc": "123"})

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open .*/unknown: no such file or directory`)

	useFakeFormGenerator = true
	_, err = r.PostFile(tmpFile, "file", map[string]string{"abc": "123"})

	c.Assert(err, NotNil)
	useFakeFormGenerator = false

	ioCopyFunc = func(dst io.Writer, src io.Reader) (int64, error) { return 0, errors.New("") }
	_, err = r.PostFile(tmpFile, "file", map[string]string{"abc": "123"})

	c.Assert(err, NotNil)
}

func (s *ReqSuite) TestMethodPut(c *C) {
	putResp, err := Request{URL: s.url + _URL_PUT, Method: PUT}.Do()

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)

	putResp, err = Global.Do(Request{URL: s.url + _URL_PUT, Method: PUT})

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)

	putResp, err = Request{URL: s.url + _URL_PUT}.Put()

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)

	putResp, err = Global.Put(Request{URL: s.url + _URL_PUT})

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodHead(c *C) {
	headResp, err := Request{URL: s.url + _URL_HEAD, Method: HEAD}.Do()

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)

	headResp, err = Global.Do(Request{URL: s.url + _URL_HEAD, Method: HEAD})

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)

	headResp, err = Request{URL: s.url + _URL_HEAD}.Head()

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)

	headResp, err = Global.Head(Request{URL: s.url + _URL_HEAD})

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPatch(c *C) {
	patchResp, err := Request{URL: s.url + _URL_PATCH, Method: PATCH}.Do()

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)

	patchResp, err = Global.Do(Request{URL: s.url + _URL_PATCH, Method: PATCH})

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)

	patchResp, err = Request{URL: s.url + _URL_PATCH}.Patch()

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)

	patchResp, err = Global.Patch(Request{URL: s.url + _URL_PATCH})

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodDelete(c *C) {
	deleteResp, err := Request{URL: s.url + _URL_DELETE, Method: DELETE}.Do()

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)

	deleteResp, err = Global.Do(Request{URL: s.url + _URL_DELETE, Method: DELETE})

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)

	deleteResp, err = Request{URL: s.url + _URL_DELETE}.Delete()

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)

	deleteResp, err = Global.Delete(Request{URL: s.url + _URL_DELETE})

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestHeaders(c *C) {
	resp, err := Request{
		URL: s.url + _URL_HEADERS,
		Headers: Headers{
			"Header1": "Value1",
			"Header2": "Value2",
		},
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestContentType(c *C) {
	resp, err := Request{
		URL:         s.url + _URL_CONTENT_TYPE,
		ContentType: _TEST_CONTENT_TYPE,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestAccept(c *C) {
	resp, err := Request{
		URL:    s.url + _URL_ACCEPT,
		Accept: _TEST_ACCEPT,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestClose(c *C) {
	getResp, err := Request{
		URL:   s.url + _URL_GET,
		Close: true,
	}.Get()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestUserAgent(c *C) {
	resp, err := Request{URL: s.url + _URL_USER_AGENT}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestBasicAuth(c *C) {
	resp, err := Request{
		URL:       s.url + _URL_BASIC_AUTH,
		Auth:      AuthBasic{_TEST_BASIC_AUTH_USER, _TEST_BASIC_AUTH_PASS},
		ProxyAuth: AuthBasic{_TEST_BASIC_AUTH_USER, _TEST_BASIC_AUTH_PASS},
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestStringResp(c *C) {
	resp, err := Request{URL: s.url + _URL_STRING_RESP}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
	c.Assert(resp.String(), Equals, _TEST_STRING_RESP)
}

func (s *ReqSuite) TestBytesResp(c *C) {
	resp, err := Request{URL: s.url + _URL_STRING_RESP}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	data, err := resp.Bytes()

	c.Assert(err, IsNil)
	c.Assert(data, DeepEquals, []byte(_TEST_STRING_RESP))
}

func (s *ReqSuite) TestJSONResp(c *C) {
	resp, err := Request{URL: s.url + _URL_JSON_RESP}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	testStruct := &TestStruct{}

	err = resp.JSON(testStruct)

	c.Assert(err, IsNil)
	c.Assert(testStruct.String, Equals, "test")
	c.Assert(testStruct.Integer, Equals, 912)
	c.Assert(testStruct.Boolean, Equals, true)
}

func (s *ReqSuite) TestJSONWithHashResp(c *C) {
	resp, err := Request{URL: s.url + _URL_JSON_RESP}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	testStruct := &TestStruct{}

	_, err = resp.JSONWithHash(testStruct, nil)
	c.Assert(err, Equals, hashutil.ErrNilHasher)

	respHash, err := resp.JSONWithHash(testStruct, sha256.New())

	c.Assert(err, IsNil)
	c.Assert(testStruct.String, Equals, "test")
	c.Assert(respHash.String(), Equals, "e03b1cde95e6e4fa740de100773bc41804d03b5297b7b0222888fff3e72d2b7c")

	resp, err = Request{URL: s.url + _URL_GET}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	_, err = resp.JSONWithHash(testStruct, sha256.New())
	c.Assert(err, NotNil)
}

func (s *ReqSuite) TestSaveResp(c *C) {
	resp, err := Request{URL: s.url + _URL_SAVE}.Get()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	err = resp.Save("/test", 0644)
	c.Assert(err, NotNil)

	testDir := c.MkDir()

	err = resp.Save(testDir+"/output.test", 0644)
	c.Assert(err, IsNil)
}

func (s *ReqSuite) TestSaveWithHashResp(c *C) {
	resp, err := Request{URL: s.url + _URL_SAVE}.Get()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	testDir := c.MkDir()

	_, err = resp.SaveWithHash("/test", 0644, sha256.New())
	c.Assert(err, NotNil)
	_, err = resp.SaveWithHash(testDir+"/output.test", 0644, nil)
	c.Assert(err, NotNil)

	resp, _ = Request{URL: s.url + _URL_SAVE}.Get()
	fileHash, err := resp.SaveWithHash(testDir+"/output.test", 0644, sha256.New())

	c.Assert(err, IsNil)
	c.Assert(fileHash.String(), Equals, "9546c567ac10e0d47034582eb9f5e5cfabf1c242c5714cf38fecb0a135f99a75")
}

func (s *ReqSuite) TestDiscard(c *C) {
	resp, err := Request{URL: s.url + _URL_JSON_RESP}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	resp.Discard()

	resp, err = Request{
		URL:         s.url + _URL_DISCARD,
		AutoDiscard: true,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 500)
}

func (s *ReqSuite) TestTimeout(c *C) {
	_, err := Request{
		URL:     s.url + _URL_TIMEOUT,
		Timeout: 10 * time.Millisecond,
	}.Do()

	c.Assert(err, NotNil)

	_, err = Request{
		URL:         s.url + _URL_HEAD,
		Timeout:     10 * time.Millisecond,
		AutoDiscard: true,
	}.Get()

	c.Assert(err, IsNil)
}

func (s *ReqSuite) TestEncoding(c *C) {
	resp, err := Request{
		URL:  s.url + "/404",
		Body: "DEADBEAF",
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)

	resp, err = Request{
		URL:  s.url + "/404",
		Body: []byte("DEADBEAF"),
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)

	r := bytes.NewReader([]byte("DEADBEAF"))

	resp, err = Request{
		URL:  s.url + "/404",
		Body: r,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)

	k := struct{ t string }{"DEADBEAF"}

	resp, err = Request{
		URL:  s.url + "/404",
		Body: k,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)

	resp, err = Request{
		URL:  s.url + "/404",
		Body: func() {},
	}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Can't encode request body: json: unsupported type: func()")
	c.Assert(resp, IsNil)
}

func (s *ReqSuite) TestRequestErrors(c *C) {
	resp, err := Request{}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "URL property can't be empty and must be set")
	c.Assert(resp, IsNil)

	resp, err = Request{URL: "ABCD"}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Unsupported scheme in URL")
	c.Assert(resp, IsNil)

	resp, err = Request{URL: "http://127.0.0.1:60000"}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, `Can't send request: Get "http://127.0.0.1:60000": dial tcp 127.0.0.1:60000: connect: connection refused`)
	c.Assert(resp, IsNil)

	resp, err = Request{URL: "%gh&%ij"}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Unsupported scheme in URL")
	c.Assert(resp, IsNil)

	resp, err = Request{Method: "ЩУП", URL: "http://127.0.0.1"}.Do()

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, `Can't create request: net/http: invalid method "ЩУП"`)
	c.Assert(resp, IsNil)

	eng := &Engine{}
	eng.Init()

	_, _, err = createRequest(eng, Request{Method: "ЩУП", Timeout: 1}, nil)
	c.Assert(err, NotNil)
}

func (s *ReqSuite) TestEngineInit(c *C) {
	var eng *Engine

	eng = &Engine{}
	eng.Init()

	eng = &Engine{Transport: &http.Transport{}}
	eng.Init()
}

func (s *ReqSuite) TestEngineErrors(c *C) {
	var eng *Engine

	resp, err := eng.Do(Request{URL: "https://essentialkaos.com"})

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrNilEngine)
	c.Assert(resp, IsNil)

	eng = &Engine{}
	eng.Init()

	eng.Dialer = nil

	resp, err = eng.Do(Request{URL: "https://essentialkaos.com"})

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrNilDialer)
	c.Assert(resp, IsNil)

	eng = &Engine{}
	eng.Init()
	eng.Transport = nil

	resp, err = eng.Do(Request{URL: "https://essentialkaos.com"})

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrNilTransport)
	c.Assert(resp, IsNil)

	eng = &Engine{}
	eng.Init()
	eng.Init()
	eng.Client = nil

	resp, err = eng.Do(Request{URL: "https://essentialkaos.com"})

	c.Assert(err, NotNil)
	c.Assert(err, Equals, ErrNilClient)
	c.Assert(resp, IsNil)
}

func (s *ReqSuite) TestIsURL(c *C) {
	c.Assert(isURL(""), Equals, false)
	c.Assert(isURL("http://domain.com"), Equals, true)
	c.Assert(isURL("https://domain.com"), Equals, true)
	c.Assert(isURL("ftp://domain.com"), Equals, true)
	c.Assert(isURL("test://domain.com"), Equals, false)
}

func (s *ReqSuite) TestQueryEncoding(c *C) {
	q := Query{}
	c.Assert(q.Encode(), Equals, "")

	q = Query{"a": 1, "b": "abcd", "c": "", "d": nil}

	qr := strings.Split(q.Encode(), "&")
	sort.Strings(qr)
	qrs := strings.Join(qr, "&")

	c.Assert(qrs, Equals, "a=1&b=abcd&c&d")
}

func (s *ReqSuite) TestLimiter(c *C) {
	var l *Limiter

	c.Assert(NewLimiter(0.0), IsNil)

	l.Wait()
}

func (s *ReqSuite) TestRetrier(c *C) {
	r := NewRetrier(Global)

	resp, err := r.Get(
		Request{URL: s.url + _URL_GET},
		Retry{Num: 3},
	)

	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)

	_, err = r.Get(
		Request{URL: "http://127.0.0.1:1"},
		Retry{Num: 3},
	)

	c.Assert(err, NotNil)

	_, err = r.Get(
		Request{URL: s.url + "/unknown"},
		Retry{Num: 3, Status: STATUS_OK, Pause: time.Millisecond},
	)

	c.Assert(err, NotNil)

	_, err = r.Get(
		Request{URL: s.url + "/unknown"},
		Retry{Num: 3, MinStatus: 299},
	)

	c.Assert(err, NotNil)
}

func (s *ReqSuite) TestRetrierErrors(c *C) {
	var r *Retrier

	_, err := r.Do(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Get(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Post(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Put(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Head(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Patch(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)
	_, err = r.Delete(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilRetrier)

	r = &Retrier{}
	_, err = r.Do(Request{}, Retry{Num: 10})
	c.Assert(err, Equals, ErrNilEngine)

	c.Assert(Retry{Num: 3}.Validate(), IsNil)
	c.Assert(Retry{Num: -1}.Validate(), NotNil)
	c.Assert(Retry{Num: 3, Status: 20}.Validate(), NotNil)
	c.Assert(Retry{Num: 3, Status: 2000}.Validate(), NotNil)
	c.Assert(Retry{Num: 3, MinStatus: 20}.Validate(), NotNil)
	c.Assert(Retry{Num: 3, MinStatus: 2000}.Validate(), NotNil)
}

func (s *ReqSuite) TestNil(c *C) {
	var e *Engine

	_, err := e.doRequest(Request{}, GET)
	c.Assert(err, DeepEquals, ErrNilEngine)

	c.Assert(func() { e.SetUserAgent("APP", "1") }, NotPanics)
	c.Assert(func() { e.SetDialTimeout(1) }, NotPanics)
	c.Assert(func() { e.SetRequestTimeout(1) }, NotPanics)
	c.Assert(func() { e.SetLimit(1.0) }, NotPanics)

	var r *Response

	c.Assert(func() { r.Discard() }, NotPanics)

	c.Assert(r.JSON(nil), DeepEquals, ErrNilResponse)
	c.Assert(r.Save("/test", 0644), DeepEquals, ErrNilResponse)
	c.Assert(r.String(), Equals, "")

	_, err = r.Bytes()
	c.Assert(err, DeepEquals, ErrNilResponse)

	_, err = r.JSONWithHash(nil, sha256.New())
	c.Assert(err, DeepEquals, ErrNilResponse)

	_, err = r.SaveWithHash("/test", 0644, sha256.New())
	c.Assert(err, DeepEquals, ErrNilResponse)
}

func (s *ReqSuite) TestEmptyBody(c *C) {
	r := &Response{Response: &http.Response{}}

	c.Assert(r.JSON(nil), DeepEquals, ErrEmptyBody)
	c.Assert(r.Save("/test", 0644), DeepEquals, ErrEmptyBody)
	c.Assert(r.String(), Equals, "")

	_, err := r.Bytes()
	c.Assert(err, DeepEquals, ErrEmptyBody)

	_, err = r.JSONWithHash(nil, sha256.New())
	c.Assert(err, DeepEquals, ErrEmptyBody)

	_, err = r.SaveWithHash("/test", 0644, sha256.New())
	c.Assert(err, DeepEquals, ErrEmptyBody)
}

func (s *ReqSuite) TestAuth(c *C) {
	var a Auth

	r, _ := http.NewRequest("GET", "http://127.0.0.1", nil)

	a = AuthBasic{"John", "Test1234"}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("Authorization"), Equals, "Basic Sm9objpUZXN0MTIzNA==")

	a = AuthBearer{"acbd1234"}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("Authorization"), Equals, "Bearer acbd1234")

	a = AuthOAuth{
		Realm:           "Example",
		ConsumerKey:     "0685bd9184jfhq22",
		Token:           "ad180jjd733klru7",
		SignatureMethod: "HMAC-SHA1",
		Signature:       "wOJIO9A2W5mFwDgiDvZbTSMK",
		Timestamp:       137131200,
		Nonce:           "4572616e48616d6d65724c61686176",
		Version:         "1.0",
	}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("Authorization"), Equals, `OAuth realm="Example", oauth_consumer_key="0685bd9184jfhq22", oauth_token="ad180jjd733klru7", oauth_signature_method="HMAC-SHA1", oauth_signature="wOJIO9A2W5mFwDgiDvZbTSMK", oauth_timestamp="137131200", oauth_nonce="4572616e48616d6d65724c61686176", oauth_version="1.0"`)

	a = AuthDigest{
		Username:  "Mufasa",
		Realm:     "http-auth@example.org",
		URI:       "/dir/index.html",
		Algorithm: "SHA-256",
		Nonce:     "7ypf/xlj9XXwfDPEoM4URrv/xwf94BcCAzFZH4GiTo0v",
		CNonce:    "f2/wE4q74E6zIJEtWaHKaf5wv/H5QzzpXusqGemxURZJ",
		NC:        1,
		QOP:       "auth",
		Response:  "8ca523f5e9506fed4657c9700eebdbec",
		Opaque:    "FQhe/qaU925kfnzjCev0ciny7QMkPqMAFRtzCUYo5tdS",
		UserHash:  true,
	}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("Authorization"), Equals, `Digest username="Mufasa", realm="http-auth@example.org", uri="/dir/index.html", algorithm=SHA-256, nonce="7ypf/xlj9XXwfDPEoM4URrv/xwf94BcCAzFZH4GiTo0v", cnonce="f2/wE4q74E6zIJEtWaHKaf5wv/H5QzzpXusqGemxURZJ", nc=00000001, qop=auth, response="8ca523f5e9506fed4657c9700eebdbec", opaque="FQhe/qaU925kfnzjCev0ciny7QMkPqMAFRtzCUYo5tdS", userhash=true`)

	a = AuthAWS4{
		Credential:    "AKIAIOSFODNN7EXAMPLE/20130524/us-east-1/s3/aws4_request",
		SignedHeaders: "host;range;x-amz-date",
		Signature:     "fe5f80f77d5fa3beca038a248ff027d0445342fe2855ddc963176630326f1024",
	}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("Authorization"), Equals, `AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20130524/us-east-1/s3/aws4_request,SignedHeaders=host;range;x-amz-date,Signature=fe5f80f77d5fa3beca038a248ff027d0445342fe2855ddc963176630326f1024`)

	r, _ = http.NewRequest("GET", "http://127.0.0.1", nil)
	a = AuthAPIKey{Key: "fe5f80f77d5fa3beca038a248ff027d0"}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("X-API-Key"), Equals, `fe5f80f77d5fa3beca038a248ff027d0`)
	c.Assert(r.Header.Get("API-Key"), Equals, `fe5f80f77d5fa3beca038a248ff027d0`)

	r, _ = http.NewRequest("GET", "http://127.0.0.1", nil)
	a = AuthHeader{"X-Auth-Token", "fe5f80f77d5fa3beca038a248ff027d0"}
	a.Apply(r, "Authorization")
	c.Assert(r.Header.Get("X-Auth-Token"), Equals, `fe5f80f77d5fa3beca038a248ff027d0`)
}

func (s *ReqSuite) TestQuery(c *C) {
	resp, err := Request{
		URL: s.url + _URL_QUERY,
		Query: Query{
			"user": "john",
			"id":   1000,
		},
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestQueryParsing(c *C) {
	var q Query
	c.Assert(q.Encode(), Equals, "")

	c.Assert(q.Set("test", true), Equals, false)
	c.Assert(q.SetIf(true, "test", true), Equals, false)
	c.Assert(q.Get("test"), IsNil)
	c.Assert(q.Delete("test"), Equals, false)
	c.Assert(q.DeleteIf(true, "test"), Equals, false)

	q = nil
	c.Assert(q.Encode(), Equals, "")

	ts := &TestStringer{}
	tp := &TestPayload{}

	q = Query{
		"":       "test",
		"test01": nil,
		"test02": "Test 1234",
		"test03": true,
		"test04": false,
		"test05": int(1),
		"test06": int8(2),
		"test07": int16(3),
		"test08": int32(4),
		"test09": int64(5),
		"test10": uint(6),
		"test11": uint8(7),
		"test12": uint16(8),
		"test13": uint32(9),
		"test14": uint64(10),
		"test15": float32(12.35),
		"test16": float64(56.7895),
		"test17": "",
		"test18": ts,
		"test19": tp,
		"test20": []string{"abcd", "1234"},
		"test21": []fmt.Stringer{ts, ts},
		"test22": []int{0, 1, 2},
		"test23": []int8{0, 1, 2},
		"test24": []int16{0, 1, 2},
		"test25": []int32{0, 1, 2},
		"test26": []int64{0, 1, 2},
		"test27": []uint{0, 1, 2},
		"test28": []uint8{0, 1, 2},
		"test29": []uint16{0, 1, 2},
		"test30": []uint32{0, 1, 2},
		"test31": []uint64{0, 1, 2},
		"test32": []float32{0.01, 1.0, 2.231213},
		"test33": []float64{0.01, 1.0, 2.231213},
	}

	q.Set("test40", true)
	q.SetIf(true, "test41", "TEST")
	q.SetIf(false, "test41", "ABDC")

	nq, err := url.ParseQuery(q.Encode())
	c.Assert(err, IsNil)
	c.Assert(nq, NotNil)

	c.Assert(nq.Get(""), Equals, "")
	c.Assert(nq.Has("test01"), Equals, true)
	c.Assert(nq.Get("test01"), Equals, "")
	c.Assert(nq.Get("test02"), Equals, "Test 1234")
	c.Assert(nq.Get("test03"), Equals, "true")
	c.Assert(nq.Get("test04"), Equals, "false")
	c.Assert(nq.Get("test05"), Equals, "1")
	c.Assert(nq.Get("test06"), Equals, "2")
	c.Assert(nq.Get("test07"), Equals, "3")
	c.Assert(nq.Get("test08"), Equals, "4")
	c.Assert(nq.Get("test09"), Equals, "5")
	c.Assert(nq.Get("test10"), Equals, "6")
	c.Assert(nq.Get("test11"), Equals, "7")
	c.Assert(nq.Get("test12"), Equals, "8")
	c.Assert(nq.Get("test13"), Equals, "9")
	c.Assert(nq.Get("test14"), Equals, "10")
	c.Assert(nq.Get("test15"), Equals, "12.35")
	c.Assert(nq.Get("test16"), Equals, "56.7895")
	c.Assert(nq.Has("test17"), Equals, true)
	c.Assert(nq.Get("test17"), Equals, "")
	c.Assert(nq.Get("test18"), Equals, "test")
	c.Assert(nq.Get("test19"), Equals, "test")
	c.Assert(nq.Get("test20"), Equals, "abcd,1234")
	c.Assert(nq.Get("test21"), Equals, "test,test")
	c.Assert(nq.Get("test22"), Equals, "0,1,2")
	c.Assert(nq.Get("test23"), Equals, "0,1,2")
	c.Assert(nq.Get("test24"), Equals, "0,1,2")
	c.Assert(nq.Get("test25"), Equals, "0,1,2")
	c.Assert(nq.Get("test26"), Equals, "0,1,2")
	c.Assert(nq.Get("test27"), Equals, "0,1,2")
	c.Assert(nq.Get("test28"), Equals, "0,1,2")
	c.Assert(nq.Get("test29"), Equals, "0,1,2")
	c.Assert(nq.Get("test30"), Equals, "0,1,2")
	c.Assert(nq.Get("test31"), Equals, "0,1,2")
	c.Assert(nq.Get("test32"), Equals, "0.01,1,2.231213")
	c.Assert(nq.Get("test33"), Equals, "0.01,1,2.231213")

	c.Assert(q.Get("test40"), Equals, true)
	c.Assert(q.Get("test41"), Equals, "TEST")
	c.Assert(q.DeleteIf(false, "test41"), Equals, false)
	c.Assert(q.DeleteIf(true, "test41"), Equals, true)
	c.Assert(q.Get("test41"), Equals, nil)
	c.Assert(q.Delete("test40"), Equals, true)

	c.Assert(
		Query{"test[]": []string{"abc", "def", "123"}}.Encode(),
		Equals, `test%5B%5D=abc&test%5B%5D=def&test%5B%5D=123`,
	)

	c.Assert(
		Query{"test[]": []fmt.Stringer{ts, ts}}.Encode(),
		Equals, `test%5B%5D=test&test%5B%5D=test`,
	)

	c.Assert(
		Query{"test[]": []int{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []int8{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []int16{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []int32{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []int64{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []uint{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []uint8{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []uint16{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []uint32{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []uint64{1, 2, 3}}.Encode(),
		Equals, `test%5B%5D=1&test%5B%5D=2&test%5B%5D=3`,
	)

	c.Assert(
		Query{"test[]": []float32{1.2, 2.5, 3.75}}.Encode(),
		Equals, `test%5B%5D=1.2&test%5B%5D=2.5&test%5B%5D=3.75`,
	)

	c.Assert(
		Query{"test[]": []float64{1.2, 2.5, 3.75}}.Encode(),
		Equals, `test%5B%5D=1.2&test%5B%5D=2.5&test%5B%5D=3.75`,
	)
}

func (s *ReqSuite) TestHeadersHelpers(c *C) {
	var h Headers

	c.Assert(h.Set("X-Test-Header", "test"), Equals, false)
	c.Assert(h.SetIf(true, "X-Test-Header", "test"), Equals, false)
	c.Assert(h.Get("X-Test-Header"), Equals, "")
	c.Assert(h.Delete("X-Test-Header"), Equals, false)
	c.Assert(h.DeleteIf(true, "X-Test-Header"), Equals, false)

	h = Headers{}

	c.Assert(h.Set("X-Test-Header", ""), Equals, false)
	c.Assert(h.Set("", "test"), Equals, false)
	c.Assert(h.SetIf(true, "", "test"), Equals, false)
	c.Assert(h.SetIf(false, "X-Test-Header", "test"), Equals, false)
	c.Assert(h.Get(""), Equals, "")
	c.Assert(h.Delete(""), Equals, false)
	c.Assert(h.DeleteIf(true, ""), Equals, false)
	c.Assert(h.DeleteIf(false, "X-Test-Header"), Equals, false)

	c.Assert(h.Set("X-Test-Header", "test"), Equals, true)
	c.Assert(h.SetIf(true, "X-Test-Header", "test2"), Equals, true)
	c.Assert(h.Get("X-Test-Header"), Equals, "test2")
	c.Assert(h.Delete("X-Test-Header"), Equals, true)
	c.Assert(h.Get("X-Test-Header"), Equals, "")
	c.Assert(h.Set("X-Test-Header", "test"), Equals, true)
	c.Assert(h.DeleteIf(true, "X-Test-Header"), Equals, true)
	c.Assert(h.Get("X-Test-Header"), Equals, "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ReqSuite) BenchmarkQueryEncoding(c *C) {
	q := Query{"a": 1, "b": "abcd", "c": "", "d": nil}

	for range c.N {
		q.Encode()
	}
}

func (s *ReqSuite) BenchmarkGetOk(c *C) {
	for range c.N {
		getResp, err := Request{URL: s.url + _URL_GET, Method: GET}.Do()

		c.Assert(err, IsNil)
		c.Assert(getResp.StatusCode, Equals, 200)
	}
}

func (s *ReqSuite) BenchmarkGetErr(c *C) {
	for range c.N {
		Request{URL: "--", Method: GET}.Do()
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func runHTTPServer(s *ReqSuite, c *C) {
	server := &http.Server{
		Handler:        http.NewServeMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	listener, err := net.Listen("tcp", ":"+s.port)

	if err != nil {
		c.Fatal(err.Error())
	}

	s.listener = listener

	server.Handler.(*http.ServeMux).HandleFunc(_URL_GET, getRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_POST, postRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_POST_MULTI, postMultiRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_PUT, putRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_HEAD, headRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_PATCH, patchRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_DELETE, deleteRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_QUERY, queryRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_HEADERS, headersRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_CONTENT_TYPE, contentTypeRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_ACCEPT, acceptRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_USER_AGENT, uaRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_BASIC_AUTH, basicAuthRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_STRING_RESP, stringRespRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_JSON_RESP, jsonRespRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_DISCARD, discardRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_TIMEOUT, timeoutRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_SAVE, saveRequestHandler)

	err = server.Serve(listener)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != GET {
		w.WriteHeader(801)
		return
	}

	w.WriteHeader(200)
}

func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		w.WriteHeader(802)
		return
	}

	w.WriteHeader(200)
}

func postMultiRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		w.WriteHeader(802)
		return
	}

	f, h, err := r.FormFile("file")

	if f == nil {
		w.WriteHeader(851)
		return
	}

	if h == nil {
		w.WriteHeader(852)
		return
	}

	if err != nil {
		w.WriteHeader(853)
		return
	}

	w.WriteHeader(200)
}

func putRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != PUT {
		w.WriteHeader(803)
		return
	}

	w.WriteHeader(200)
}
func headRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != HEAD {
		w.WriteHeader(804)
		return
	}

	w.WriteHeader(200)
}

func patchRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != PATCH {
		w.WriteHeader(805)
		return
	}

	w.WriteHeader(200)
}

func deleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != DELETE {
		w.WriteHeader(806)
		return
	}

	w.WriteHeader(200)
}

func queryRequestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	switch {
	case query.Get("user") != "john":
		w.WriteHeader(901)
		return
	case query.Get("id") != "1000":
		w.WriteHeader(902)
		return
	}

	w.WriteHeader(200)
}

func headersRequestHandler(w http.ResponseWriter, r *http.Request) {
	header1 := r.Header["Header1"]

	if len(header1) != 1 {
		w.WriteHeader(910)
		return
	}

	if header1[0] != "Value1" {
		w.WriteHeader(911)
		return
	}

	header2 := r.Header["Header2"]

	if len(header2) != 1 {
		w.WriteHeader(912)
		return
	}

	if header2[0] != "Value2" {
		w.WriteHeader(913)
		return
	}

	w.WriteHeader(200)
}

func contentTypeRequestHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header["Content-Type"]

	if len(header) != 1 {
		w.WriteHeader(920)
		return
	}

	if header[0] != _TEST_CONTENT_TYPE {
		w.WriteHeader(921)
		return
	}

	w.WriteHeader(200)
}

func acceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header["Accept"]

	if len(header) != 1 {
		w.WriteHeader(930)
		return
	}

	if header[0] != _TEST_ACCEPT {
		w.WriteHeader(931)
		return
	}

	w.WriteHeader(200)
}

func uaRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() != Global.UserAgent {
		w.WriteHeader(940)
		return
	}

	w.WriteHeader(200)
}

func basicAuthRequestHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(950)
		return
	}

	if user != _TEST_BASIC_AUTH_USER {
		w.WriteHeader(951)
		return
	}

	if pass != _TEST_BASIC_AUTH_PASS {
		w.WriteHeader(952)
		return
	}

	w.WriteHeader(200)
}

func stringRespRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(_TEST_STRING_RESP))
}

func jsonRespRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
  "string": "test",
  "integer": 912,
  "boolean": true }`,
	))
}

func discardRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)

	w.Write([]byte(`{
  "string": "test",
  "integer": 912,
  "boolean": true }`,
	))
}

func timeoutRequestHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	w.WriteHeader(200)
	w.Write([]byte(`{}`))
}

func saveRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`TEST-DATA`))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (t *TestStringer) String() string {
	return "test"
}

func (t *TestPayload) ToQuery(name string) string {
	return "test"
}
