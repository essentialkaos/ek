package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"runtime"
	"strconv"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TTYSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TTYSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TTYSuite) TestIsTTY(c *C) {
	IsTTY()

	os.Setenv("FAKETTY", "1")
	c.Assert(IsFakeTTY(), Equals, true)
	os.Setenv("FAKETTY", "")
}

func (s *TTYSuite) TestIsTMUX(c *C) {
	os.Setenv("TMUX", "/tmp/tmux-1000/default,3739,0")

	isTmux, err := IsTMUX()
	c.Assert(isTmux, Equals, true)
	c.Assert(err, IsNil)

	if runtime.GOOS != "linux" {
		return
	}

	os.Setenv("TMUX", "")

	IsTMUX()

	procFS = "/__unknown__"

	_, err = IsTMUX()
	c.Assert(err, NotNil)

	procFS = c.MkDir()
	ppid := os.Getppid()
	statDir := procFS + "/" + strconv.Itoa(ppid)
	os.MkdirAll(statDir, 0755)
	os.WriteFile(statDir+"/stat", []byte(`1 (systemd) S 0 1 1 0`), 0644)

	isTmux, err = IsTMUX()
	c.Assert(isTmux, Equals, false)
	c.Assert(err, IsNil)
}

func (s *TTYSuite) TestGetSize(c *C) {
	w, h := GetSize()

	c.Assert(w, Not(Equals), -1)
	c.Assert(w, Not(Equals), 0)
	c.Assert(h, Not(Equals), -1)
	c.Assert(h, Not(Equals), 0)
}

func (s *TTYSuite) TestGetWidth(c *C) {
	c.Assert(GetWidth(), Not(Equals), -1)
	c.Assert(GetWidth(), Not(Equals), 0)
}

func (s *TTYSuite) TestGetHeight(c *C) {
	c.Assert(GetHeight(), Not(Equals), -1)
	c.Assert(GetHeight(), Not(Equals), 0)
}

func (s *TTYSuite) TestErrors(c *C) {
	tty = "/non-exist"

	w, h := GetSize()

	c.Assert(w, Equals, -1)
	c.Assert(h, Equals, -1)

	tty = "/dev/tty"
}
