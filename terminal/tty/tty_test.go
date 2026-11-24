package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type TTYSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TTYSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TTYSuite) TestIsTTY(c *C) {
	IsTTY()
	GetWidth()
	GetHeight()

	initOnce = sync.Once{}
	os.Setenv("FAKETTY", "1")
	c.Assert(IsFakeTTY(), Equals, true)

	initOnce = sync.Once{}
	os.Setenv("FAKETTY", "")
	c.Assert(IsFakeTTY(), Equals, false)
}

func (s *TTYSuite) TestIsSystemd(c *C) {
	IsSystemd()
}

func (s *TTYSuite) TestIsTMUX(c *C) {
	initOnce = sync.Once{}
	os.Setenv("TMUX", "/tmp/tmux-1000/default,3739,0")

	isTmux, err := IsTMUX()
	c.Assert(isTmux, Equals, true)
	c.Assert(err, IsNil)

	if runtime.GOOS != "linux" {
		return
	}

	initOnce = sync.Once{}
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

	initOnce = sync.Once{}
	isTmux, err = IsTMUX()
	c.Assert(isTmux, Equals, false)
	c.Assert(err, IsNil)
}
