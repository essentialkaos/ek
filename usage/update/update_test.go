package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"net/http"
	"testing"
	"time"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type UpdateSuite struct {
	url      string
	port     string
	listener net.Listener
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UpdateSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UpdateSuite) SetUpSuite(c *C) {
	s.port = "30002"
	s.url = "http://127.0.0.1:" + s.port

	go runHTTPServer(s, c)

	time.Sleep(time.Second)
}

func (s *UpdateSuite) TearDownSuite(c *C) {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *UpdateSuite) TestGitHubChecker(c *C) {
	githubAPI = s.url

	newVersion, releaseDate, hasUpdate := GitHubChecker("", "", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "essentialkaos/unknown")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "essentialkaos/limited")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "essentialkaos/garbage")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitHubChecker("GitHubChecker", "0.9.9", "essentialkaos/project")

	c.Assert(newVersion, Equals, "1.2.3")
	c.Assert(releaseDate.Unix(), Equals, int64(1589810841))
	c.Assert(hasUpdate, Equals, true)
}

func (s *UpdateSuite) TestUpdateChecker(c *C) {
	newVersion, releaseDate, hasUpdate := UpdateChecker("", "", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("GitHubChecker", "0.9.9", s.url+"/unknown")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("GitHubChecker", "0.9.9", s.url+"/garbage")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("GitHubChecker", "0.9.9", s.url+"/project")

	c.Assert(newVersion, Equals, "1.2.3")
	c.Assert(releaseDate.Unix(), Equals, int64(1578064700))
	c.Assert(hasUpdate, Equals, true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func runHTTPServer(s *UpdateSuite, c *C) {
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

	server.Handler.(*http.ServeMux).HandleFunc("/repos/essentialkaos/project/releases/latest", ghInfoHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/repos/essentialkaos/unknown/releases/latest", ghNotFoundHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/repos/essentialkaos/limited/releases/latest", ghLimitedHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/repos/essentialkaos/garbage/releases/latest", ghWrongFormatHandler)

	server.Handler.(*http.ServeMux).HandleFunc("/project/latest.json", updInfoHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/unknown/latest.json", updUnknownHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/garbage/latest.json", updWrongFormatHandler)

	err = server.Serve(listener)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func ghInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-RateLimit-Remaining", "1000")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name": "v1.2.3","published_at": "2020-05-18T14:07:21Z"}`))
}

func ghNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func ghLimitedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-RateLimit-Remaining", "0")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name": "v1.2.3","published_at": "2020-05-18T14:07:21Z"}`))
}

func ghWrongFormatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`DEADBEEF`))
}

func updInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{"version": "1.2.3", "date": "2020-01-03T15:18:20Z"}`))
}

func updUnknownHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func updWrongFormatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`DEADBEEF`))
}
