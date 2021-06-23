package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type UsageSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UsageSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UsageSuite) TestAbout(c *C) {
	about := &About{
		App:     "Application",
		Version: "1.0.0",
		Release: ".A45",
		Build:   "37163",
		Desc:    "Test application",
		Year:    2010,
		Owner:   "Some company",
		License: "MIT",
	}

	about.Render()

	testChecker := func(app, version, data string) (string, time.Time, bool) {
		return "1.0.1", time.Now(), true
	}

	about = &About{
		App:           "Application",
		Version:       "1.0.0",
		Release:       ".A45",
		Desc:          "Test application",
		Owner:         "Some company",
		License:       "MIT",
		UpdateChecker: UpdateChecker{"1", testChecker},
	}

	about.Render()
}

func (s *UsageSuite) TestUsage(c *C) {
	info := NewInfo("", "file")

	info.AddSpoiler("This is usage of spoiler with {g}c{c}o{r}l{m}o{b}r{g}s{!} support")

	info.AddCommand() // will be ignored
	info.AddCommand("print", "Print command")

	info.AddGroup("Command group")

	info.AddCommand("read")
	info.AddCommand("read", "Read command")
	info.AddCommand("read1", "Read command with arguments", "arg1", "arg2")
	info.AddCommand("read2", "Read command with optional argument", "?arg")

	info.AddOption("t:test")
	info.AddOption("t:test", "Test option ")
	info.AddOption("test1", "Test option with argument", "arg")
	info.AddOption("test2", "Test option with optional argument", "?arg")

	info.BoundOptions("read", "t:test", "test1")

	info.AddExample() // will be ignored
	info.AddExample("abc")
	info.AddExample("abc", "Example with description")
	info.AddRawExample() // will be ignored
	info.AddRawExample("echo 123 | myapp")
	info.AddRawExample("echo 123 | myapp", "Example with description")

	info.Render()

	info.Breadcrumbs = false
	info.CommandsColorTag = "{m}"
	info.OptionsColorTag = "{b}"

	info.Render()

	c.Assert(info.GetCommand("read"), NotNil)
	c.Assert(info.GetCommand("read999"), IsNil)

	c.Assert(info.GetOption("t:test"), NotNil)
	c.Assert(info.GetOption("test"), NotNil)
	c.Assert(info.GetOption("test999"), IsNil)
}

func (s *UsageSuite) TestVersionInfo(c *C) {
	c.Assert(isNewerVersion("ABC", "1.0.0"), Equals, false)
	c.Assert(isNewerVersion("1.0.0", "ABC"), Equals, false)

	d1 := time.Unix(time.Now().Unix()-3600, 0)
	d2 := time.Unix(time.Now().Unix()-90000, 0)
	d3 := time.Unix(time.Now().Unix()-1296000, 0)

	printNewVersionInfo("ABC", "1.0.0", d1)
	printNewVersionInfo("1.0.0", "ABC", d1)

	printNewVersionInfo("1.0.0", "2.0.0", d1)
	printNewVersionInfo("1.0.0", "1.1.0", d2)
	printNewVersionInfo("1.0.0", "1.0.1", d3)
}
