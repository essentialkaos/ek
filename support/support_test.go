package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SupportSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SupportSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

var gomodData = []byte(`module github.com/essentialkaos/ek/v12

go 1.18

require (
  github.com/essentialkaos/check v1.4.0
  github.com/essentialkaos/go-linenoise/v3 v3.4.0
  golang.org/x/crypto v0.21.0
  golang.org/x/sys v0.18.0
)

require (
  github.com/kr/pretty v0.3.1 // indirect
  github.com/kr/text v0.2.0 // indirect
  github.com/rogpeppe/go-internal v1.11.0 // indirect
)
`)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SupportSuite) TestCollect(c *C) {
	i := Collect("test", "1.2.3")

	c.Assert(i, NotNil)
	c.Assert(i.WithDeps([]Dep{Dep{"test", "1.0", ""}}), NotNil)
	c.Assert(i.WithRevision(""), NotNil)
	c.Assert(i.WithRevision("1234567"), NotNil)
	c.Assert(i.WithPackages([]Pkg{Pkg{"test", "test-1.2.3"}}), NotNil)
	c.Assert(i.WithApps(App{"test", "1.2.3"}), NotNil)
	c.Assert(i.WithChecks(Check{CHECK_OK, "Test", "Test message"}), NotNil)
	c.Assert(i.WithEnvVars("", "TEST", "TERM", "CI"), NotNil)
	c.Assert(i.WithNetwork(&NetworkInfo{PublicIP: "192.168.1.1"}), NotNil)
	c.Assert(i.WithFS([]FSInfo{FSInfo{}, FSInfo{"/", "/dev/vda1", "ext4", 1000, 10000}}), NotNil)
}

func (s *SupportSuite) TestNil(c *C) {
	var i *Info

	c.Assert(i.WithDeps(nil), IsNil)
	c.Assert(i.WithRevision(""), IsNil)
	c.Assert(i.WithPackages(nil), IsNil)
	c.Assert(i.WithApps(App{}), IsNil)
	c.Assert(i.WithChecks(Check{}), IsNil)
	c.Assert(i.WithEnvVars(""), IsNil)
	c.Assert(i.WithNetwork(nil), IsNil)
	c.Assert(i.WithFS(nil), IsNil)

	c.Assert(func() { i.Print() }, NotPanics)
}
