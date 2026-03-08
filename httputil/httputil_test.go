package httputil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net/http"
	"testing"

	. "github.com/essentialkaos/check"
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

	c.Assert(GetRequestHost(nil), Equals, "")

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

	c.Assert(GetPortByScheme("http"), Equals, "80")
	c.Assert(GetPortByScheme("https"), Equals, "443")
	c.Assert(GetPortByScheme("ftp"), Equals, "21")
	c.Assert(GetPortByScheme("unknown"), Equals, "")
}

func (s *HTTPUtilSuite) TestRemoteParsing(c *C) {
	req1, _ := http.NewRequest("GET", "http://127.0.0.1/hello", nil)
	req2, _ := http.NewRequest("GET", "http://127.0.0.1/hello", nil)

	req1.RemoteAddr = ""
	req2.RemoteAddr = "127.0.0.1:12345"

	c.Assert(GetRemoteHost(nil), Equals, "")

	c.Assert(GetRemoteHost(req1), Equals, "")
	c.Assert(GetRemotePort(req1), Equals, "")

	c.Assert(GetRemoteHost(req2), Equals, "127.0.0.1")
	c.Assert(GetRemotePort(req2), Equals, "12345")
}

func (s *HTTPUtilSuite) TestGetDescCode(c *C) {
	c.Assert(GetDescByCode(999), Equals, "")
}

func (s *HTTPUtilSuite) TestURLCheck(c *C) {
	c.Assert(IsURL("127.0.0.1"), Equals, false)
	c.Assert(IsURL("127.0.0.1:80"), Equals, false)
	c.Assert(IsURL("gop://127.0.0.1:80"), Equals, false)

	c.Assert(IsURL("ftp://d.p"), Equals, true)
	c.Assert(IsURL("ftp://domain.com"), Equals, true)
	c.Assert(IsURL("http://domain.com"), Equals, true)
	c.Assert(IsURL("http://domain.com:8080"), Equals, true)
	c.Assert(IsURL("https://domain.com:8080"), Equals, true)
	c.Assert(IsURL("http://domain.com:8080/my/super/method?id=123#TEST"), Equals, true)
}
