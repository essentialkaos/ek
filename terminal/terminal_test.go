package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ErrorInput struct{}

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

	var buf bytes.Buffer
	buf.WriteString("TestStringer")

	c.Assert(formatMessage(&buf, "", nil), Equals, "TestStringer")
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

func (s *TerminalSuite) TestInput(c *C) {
	var buf bytes.Buffer

	NewLine = true
	dataInput = &buf

	buf.WriteString("Test message ")
	input, err := Read("Title")
	c.Assert(err, IsNil)
	c.Assert(input, Equals, "Test message ")

	buf.WriteString("Y")
	ok, _ := ReadAnswer("Title")
	c.Assert(ok, Equals, true)

	buf.WriteString("n")
	ok, _ = ReadAnswer("Title")
	c.Assert(ok, Equals, false)

	AlwaysYes = true
	buf.WriteString("n")
	ok, _ = ReadAnswer("Title")
	c.Assert(ok, Equals, true)
	AlwaysYes = false

	buf.WriteString("f\ny")
	ok, _ = ReadAnswer("Title", "y")
	c.Assert(ok, Equals, true)

	buf.WriteString("f\nn")
	ok, _ = ReadAnswer("Title", "n")
	c.Assert(ok, Equals, false)

	dataInput = &ErrorInput{}

	_, err = Read("Title")
	c.Assert(err, NotNil)

	_, err = ReadAnswer("Title")
	c.Assert(err, NotNil)

	c.Assert(getAnswerTitle("", ""), Equals, "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (e *ErrorInput) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("error")
}
