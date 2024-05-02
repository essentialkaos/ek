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
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type TTYSuite struct{}

type FakeInfo struct {
	IsChar bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TTYSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TTYSuite) TestIsTTY(c *C) {
	stdout = &FakeInfo{true}
	c.Assert(IsTTY(), Equals, true)
	stdout = &FakeInfo{false}
	c.Assert(IsTTY(), Equals, false)

	os.Setenv("FAKETTY", "1")
	c.Assert(IsFakeTTY(), Equals, true)
	os.Setenv("FAKETTY", "")
	c.Assert(IsFakeTTY(), Equals, false)
}

func (s *TTYSuite) TestIsSystemd(c *C) {
	IsSystemd()
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

// ////////////////////////////////////////////////////////////////////////////////// //

func (f *FakeInfo) Name() string       { return "" }
func (f *FakeInfo) Size() int64        { return 0 }
func (f *FakeInfo) ModTime() time.Time { return time.Time{} }
func (f *FakeInfo) IsDir() bool        { return false }
func (f *FakeInfo) Sys() any           { return nil }

func (f *FakeInfo) Mode() os.FileMode {
	if f.IsChar {
		return os.ModeCharDevice
	}

	return os.ModeSymlink
}
