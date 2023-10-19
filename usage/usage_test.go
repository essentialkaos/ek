package usage

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type UsageSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UsageSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UsageSuite) TestAbout(c *C) {
	about := &About{
		App:         "Application",
		Version:     "1.8.13",
		Release:     "β4",
		Build:       "ce9d5c6",
		Desc:        "Test application",
		Year:        2010,
		Owner:       "Some company",
		License:     "MIT",
		Environment: Environment{{"A", "1"}, {"B", "2"}},

		AppNameColorTag: "{#99}",
		VersionColorTag: "{#125}",
	}

	about.Render()

	testChecker := func(app, version, data string) (string, time.Time, bool) {
		return "1.0.1", time.Now(), true
	}

	about = &About{
		App:           "Application",
		Version:       "1.0.0",
		Release:       "β4",
		Desc:          "Test application",
		Owner:         "Some company",
		License:       "MIT",
		UpdateChecker: UpdateChecker{"1", testChecker},

		AppNameColorTag: "{ABCD}",
		VersionColorTag: "{ABCD}",
	}

	about.Render()
}

func (s *UsageSuite) TestRawVersion(c *C) {
	about := &About{
		App:         "Application",
		Version:     "1.8.13",
		Release:     "β4",
		Build:       "ce9d5c6",
		Desc:        "Test application",
		Year:        2010,
		Owner:       "Some company",
		License:     "MIT",
		Environment: Environment{{"A", "1"}, {"B", "2"}},
	}

	about.Print(VERSION_SIMPLE)

	c.Assert(getRawVersion(about, VERSION_FULL), Equals, "1.8.13-β4+ce9d5c6")
	c.Assert(getRawVersion(about, VERSION_SIMPLE), Equals, "1.8.13")
	c.Assert(getRawVersion(about, VERSION_MAJOR), Equals, "1")
	c.Assert(getRawVersion(about, VERSION_MINOR), Equals, "8")
	c.Assert(getRawVersion(about, VERSION_PATCH), Equals, "13")
	c.Assert(getRawVersion(about, VERSION_RELEASE), Equals, "β4")
	c.Assert(getRawVersion(about, VERSION_BUILD), Equals, "ce9d5c6")

	c.Assert(getRawVersion(about, "unknown"), Equals, "")

	about.Version = "UnKnOwN"
	c.Assert(getRawVersion(about, VERSION_MAJOR), Equals, "")
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

	info.GetOption("t:test").ColorTag = "{r}"
	info.GetCommand("read").ColorTag = "{r}"

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

	c.Assert(info.GetCommand("read").String(), Equals, "read")
	c.Assert(info.GetCommand("unknown").String(), Equals, "")
	c.Assert(info.GetOption("t:test").String(), Equals, "--test")
	c.Assert(info.GetOption("u:unknown").String(), Equals, "")
}

func (s *UsageSuite) TestDeattachedPrint(c *C) {
	cmd := &Command{Name: "test", Desc: "Test command", ColorTag: "{#99}"}
	opt := &Option{Long: "test", Short: "T", Desc: "Test option", ColorTag: "{#99}"}

	cmd.Print()
	opt.Print()

	cmd.ColorTag = ""
	opt.ColorTag = ""

	cmd.Print()
	opt.Print()
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

func (s *UsageSuite) TestDeprecated(c *C) {
	i := &Info{}
	i.Render()

	m := &Command{}
	m.Render()

	o := &Option{}
	o.Render()

	e := &Example{}
	e.Render()

	a := &About{}
	a.Render()
}

func (s *UsageSuite) TestNils(c *C) {
	var i *Info
	c.Assert(func() { i.AddGroup("test") }, NotPanics)
	c.Assert(func() { i.AddCommand("test") }, NotPanics)
	c.Assert(func() { i.AddOption("test") }, NotPanics)
	c.Assert(func() { i.AddExample("test") }, NotPanics)
	c.Assert(func() { i.AddRawExample("test") }, NotPanics)
	c.Assert(func() { i.AddSpoiler("test") }, NotPanics)
	c.Assert(func() { i.BoundOptions("test", "test") }, NotPanics)
	c.Assert(func() { i.GetCommand("test") }, NotPanics)
	c.Assert(func() { i.GetOption("test") }, NotPanics)
	c.Assert(func() { i.Print() }, NotPanics)

	var m *Command
	c.Assert(func() { m.Print() }, NotPanics)

	var o *Option
	c.Assert(func() { o.Print() }, NotPanics)

	var e *Example
	c.Assert(func() { e.Print() }, NotPanics)

	var a *About
	c.Assert(func() { a.Print() }, NotPanics)
}
