package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TerminalSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TerminalSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *TerminalSuite) TestMessage(c *C) {
	Error("Test %s", "Error")
	Warn("Test %s", "Warn")
	Info("Test %s", "Info")

	Error(fmt.Errorf("Error"))
	Error([]string{"Error"})
}

func (s *TerminalSuite) TestMessageWithPrefix(c *C) {
	ErrorPrefix = "▲ "
	WarnPrefix = "▲ "
	InfoPrefix = "▲ "

	Error("Test %s\nMessage\n", "Error")
	Warn("Test %s\nMessage\n", "Warn")
	Info("Test %s\nMessage\n", "Info")

	ErrorPrefix = ""
	WarnPrefix = ""
	InfoPrefix = ""
}

func (s *TerminalSuite) TestAction(c *C) {
	PrintActionMessage("Testing")
	PrintActionStatus(0)

	PrintActionMessage("Testing")
	PrintActionStatus(1)

	PrintActionMessage("Testing")
	PrintActionStatus(2)

	PrintActionMessage("Testing")
	PrintActionStatus(3)
}
