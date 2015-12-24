package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net/http"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type HTTPUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&HTTPUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *HTTPUtilSuite) TestRequestParsing(c *C) {
	req1, _ := http.NewRequest("GET", "http://127.0.0.1/hello", nil)
	req2, _ := http.NewRequest("GET", "http://domain.com/hello", nil)
	req3, _ := http.NewRequest("GET", "https://127.0.0.1:8080", nil)
	req4, _ := http.NewRequest("GET", "https://domain.com:8080", nil)
	req5, _ := http.NewRequest("GET", "", nil)

	c.Assert(req1, NotNil)
	c.Assert(req2, NotNil)
	c.Assert(req3, NotNil)
	c.Assert(req4, NotNil)
	c.Assert(req5, NotNil)

	c.Assert(GetRequestHost(req1), Equals, "127.0.0.1")
	c.Assert(GetRequestHost(req2), Equals, "domain.com")
	c.Assert(GetRequestHost(req3), Equals, "127.0.0.1")
	c.Assert(GetRequestHost(req4), Equals, "domain.com")
	c.Assert(GetRequestHost(req5), Equals, "")

	c.Assert(GetRequestPort(req1), Equals, "80")
	c.Assert(GetRequestPort(req2), Equals, "80")
	c.Assert(GetRequestPort(req3), Equals, "8080")
	c.Assert(GetRequestPort(req4), Equals, "8080")
	c.Assert(GetRequestPort(req5), Equals, "")
}

func (s *HTTPUtilSuite) TestGetDescCode(c *C) {
	c.Assert(GetDescByCode(100), Equals, "Continue")
	c.Assert(GetDescByCode(101), Equals, "Switching Protocols")
	c.Assert(GetDescByCode(200), Equals, "OK")
	c.Assert(GetDescByCode(201), Equals, "Created")
	c.Assert(GetDescByCode(202), Equals, "Accepted")
	c.Assert(GetDescByCode(203), Equals, "Non Authoritative Info")
	c.Assert(GetDescByCode(204), Equals, "No Content")
	c.Assert(GetDescByCode(205), Equals, "Reset Content")
	c.Assert(GetDescByCode(206), Equals, "Partial Content")
	c.Assert(GetDescByCode(300), Equals, "Multiple Choices")
	c.Assert(GetDescByCode(301), Equals, "Moved Permanently ")
	c.Assert(GetDescByCode(302), Equals, "Found")
	c.Assert(GetDescByCode(303), Equals, "See Other")
	c.Assert(GetDescByCode(304), Equals, "Not Modified")
	c.Assert(GetDescByCode(305), Equals, "Use Proxy")
	c.Assert(GetDescByCode(307), Equals, "Temporary Redirect")
	c.Assert(GetDescByCode(400), Equals, "Bad Request")
	c.Assert(GetDescByCode(401), Equals, "Unauthorized")
	c.Assert(GetDescByCode(402), Equals, "Payment Required")
	c.Assert(GetDescByCode(403), Equals, "Forbidden")
	c.Assert(GetDescByCode(404), Equals, "Not Found")
	c.Assert(GetDescByCode(405), Equals, "Method Not Allowed")
	c.Assert(GetDescByCode(406), Equals, "Not Acceptable")
	c.Assert(GetDescByCode(407), Equals, "Proxy Auth Required")
	c.Assert(GetDescByCode(408), Equals, "Request Timeout")
	c.Assert(GetDescByCode(409), Equals, "Conflict")
	c.Assert(GetDescByCode(410), Equals, "Gone")
	c.Assert(GetDescByCode(411), Equals, "Length Required")
	c.Assert(GetDescByCode(412), Equals, "Precondition Failed")
	c.Assert(GetDescByCode(413), Equals, "Request Entity Too Large")
	c.Assert(GetDescByCode(414), Equals, "Request URI TooLong")
	c.Assert(GetDescByCode(415), Equals, "Unsupported Media Type")
	c.Assert(GetDescByCode(416), Equals, "Requested Range Not Satisfiable")
	c.Assert(GetDescByCode(417), Equals, "Expectation Failed")
	c.Assert(GetDescByCode(418), Equals, "Teapot")
	c.Assert(GetDescByCode(500), Equals, "Internal Server Error")
	c.Assert(GetDescByCode(501), Equals, "Not Implemented")
	c.Assert(GetDescByCode(502), Equals, "Bad Gateway")
	c.Assert(GetDescByCode(503), Equals, "Service Unavailable")
	c.Assert(GetDescByCode(504), Equals, "Gateway Timeout")
	c.Assert(GetDescByCode(505), Equals, "HTTP Version Not Supported")
}

func (s *HTTPUtilSuite) TestURLCheck(c *C) {
	c.Assert(IsURL("127.0.0.1"), Equals, false)
	c.Assert(IsURL("127.0.0.1:80"), Equals, false)

	c.Assert(IsURL("ftp://d.pr"), Equals, true)
	c.Assert(IsURL("ftp://domain.com"), Equals, true)
	c.Assert(IsURL("http://domain.com"), Equals, true)
	c.Assert(IsURL("http://domain.com:8080"), Equals, true)
	c.Assert(IsURL("https://domain.com:8080"), Equals, true)
	c.Assert(IsURL("http://domain.com:8080/my/super/method?id=123#TEST"), Equals, true)
}
