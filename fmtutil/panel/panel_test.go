package panel

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"github.com/essentialkaos/ek.v13/fmtc"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PanelSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PanelSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PanelSuite) TestBasicErrorPanel(c *C) {
	Error("Test error", "Message")
}

func (s *PanelSuite) TestBasicWarnPanel(c *C) {
	Warn("Test warn", "Message")
}

func (s *PanelSuite) TestBasicInfoPanel(c *C) {
	Info("Test info", "Message")
}

func (s *PanelSuite) TestNoColor(c *C) {
	fmtc.DisableColors = true
	Panel(
		"使用上のヒント", "{m}", "Test with no colors",
		`Lorem ipsum dolor sit amet.`,
		WRAP, INDENT_OUTER, INDENT_INNER, TOP_LINE, BOTTOM_LINE, LABEL_POWERLINE,
	)
	fmtc.DisableColors = false
}

func (s *PanelSuite) TestPanelAllOptions(c *C) {
	Width = 60
	Indent = 2

	Panel(
		"使用上のヒント", "{m}", "Test all options",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		WRAP, INDENT_OUTER, INDENT_INNER, TOP_LINE, BOTTOM_LINE, LABEL_POWERLINE,
	)

	Width = 88
	Indent = 0
}

func (s *PanelSuite) TestPanelWeird(c *C) {
	Width = -10
	Indent = -10

	Panel(
		"", "{g}", "",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		WRAP, INDENT_OUTER, INDENT_INNER, TOP_LINE, BOTTOM_LINE, LABEL_POWERLINE,
	)

	Panel(
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "{#120}",
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		WRAP, INDENT_OUTER, INDENT_INNER, TOP_LINE, BOTTOM_LINE, LABEL_POWERLINE,
	)

	Width = 999
	Indent = 999

	Panel(
		"", "{#222}", "",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`,
		WRAP, INDENT_OUTER, INDENT_INNER, TOP_LINE, BOTTOM_LINE, LABEL_POWERLINE,
	)
}
