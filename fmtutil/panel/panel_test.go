package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PanelSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PanelSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PanelSuite) TestBasicErrorPanel(c *C) {
	ErrorPanel("Test error", "Message")
}

func (s *PanelSuite) TestBasicWarnPanel(c *C) {
	WarnPanel("Test warn", "Message")
}

func (s *PanelSuite) TestBasicInfoPanel(c *C) {
	InfoPanel("Test info", "Message")
}

func (s *PanelSuite) TestPanelAllOptions(c *C) {
	Width = 60

	Panel(
		"YOLO", "{m}", "Test all options",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		WRAP, INDENT_OUTER, INDENT_INNER, BOTTOM_LINE, LABEL_POWERLINE,
	)

	Width = 88
}
