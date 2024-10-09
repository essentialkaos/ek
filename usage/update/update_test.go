package update

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net"
	"net/http"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
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
}

func (s *UpdateSuite) TearDownSuite(c *C) {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *UpdateSuite) TestGitHubChecker(c *C) {
	githubAPI = s.url + "/github"

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

func (s *UpdateSuite) TestGitLabChecker(c *C) {
	gitlabAPI = s.url + "/gitlab"

	newVersion, releaseDate, hasUpdate := GitLabChecker("", "", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitLabChecker("GitLabChecker", "0.9.9", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitLabChecker("GitLabChecker", "0.9.9", "essentialkaos/unknown")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitLabChecker("GitLabChecker", "0.9.9", "essentialkaos/limited")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitLabChecker("GitLabChecker", "0.9.9", "essentialkaos/garbage")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = GitLabChecker("GitLabChecker", "0.9.9", "essentialkaos/project")

	c.Assert(newVersion, Equals, "1.2.3")
	c.Assert(releaseDate.Unix(), Equals, int64(1589810841))
	c.Assert(hasUpdate, Equals, true)
}

func (s *UpdateSuite) TestUpdateChecker(c *C) {
	basicAPI := s.url + "/basic"
	newVersion, releaseDate, hasUpdate := UpdateChecker("", "", "")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("BasicChecker", "0.9.9", basicAPI+"/unknown")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("BasicChecker", "0.9.9", basicAPI+"/garbage")

	c.Assert(newVersion, Equals, "")
	c.Assert(hasUpdate, Equals, false)

	newVersion, releaseDate, hasUpdate = UpdateChecker("BasicChecker", "0.9.9", basicAPI+"/project")

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

	server.Handler.(*http.ServeMux).HandleFunc("/github/repos/essentialkaos/project/releases/latest", githubInfoHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/github/repos/essentialkaos/unknown/releases/latest", githubNotFoundHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/github/repos/essentialkaos/limited/releases/latest", githubLimitedHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/github/repos/essentialkaos/garbage/releases/latest", githubWrongFormatHandler)

	server.Handler.(*http.ServeMux).HandleFunc("/gitlab/projects/essentialkaos%2Fproject/releases/permalink/latest", gitlabInfoHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/gitlab/projects/essentialkaos/unknown/releases/permalink/latest", gitlabNotFoundHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/gitlab/projects/essentialkaos/limited/releases/permalink/latest", gitlabLimitedHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/gitlab/projects/essentialkaos/garbage/releases/permalink/latest", gitlabWrongFormatHandler)

	server.Handler.(*http.ServeMux).HandleFunc("/basic/project/latest.json", basicInfoHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/basic/unknown/latest.json", basicUnknownHandler)
	server.Handler.(*http.ServeMux).HandleFunc("/basic/garbage/latest.json", basicWrongFormatHandler)

	err = server.Serve(listener)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func githubInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-RateLimit-Remaining", "60")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name": "v1.2.3","published_at": "2020-05-18T14:07:21Z"}`))
}

func githubNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func githubLimitedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-RateLimit-Remaining", "0")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name": "v1.2.3","published_at": "2020-05-18T14:07:21Z"}`))
}

func githubWrongFormatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`DEADBEEF`))
}

func gitlabInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("RateLimit-Remaining", "2000")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name":"v1.2.3","released_at":"2020-05-18T14:07:21.814Z"}`))
}

func gitlabNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func gitlabLimitedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("RateLimit-Remaining", "0")
	w.WriteHeader(200)
	w.Write([]byte(`{"tag_name":"v1.2.3","released_at":"2020-05-18T14:07:21.814Z"}`))
}

func gitlabWrongFormatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`DEADBEEF`))
}

func basicInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{"version": "1.2.3", "date": "2020-01-03T15:18:20Z"}`))
}

func basicUnknownHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func basicWrongFormatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`DEADBEEF`))
}
